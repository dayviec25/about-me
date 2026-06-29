---
title: "Free, Serverless Obsidian Sync with Cloudflare R2"
date: "2026-06-28"
author: "David Chung"
description: "How to sync your Obsidian vault across every device for $0 — no server to run, patch, or babysit."
---

## The problem

Obsidian is a local-first notes app. Your vault is just a folder of plain markdown files on disk, which is wonderful — until you want those notes on your phone, your laptop, and your desktop, all staying in sync.

The official answer is **Obsidian Sync**, a paid subscription. It's genuinely good and takes thirty seconds to set up, so if you'd rather not touch any of this, go buy it and move on with your life.

But if you want full control of your data, end-to-end encryption you hold the keys to, and a bill of **zero dollars a month**, you can do it yourself with the community **Self-hosted LiveSync** plugin.

Most guides for LiveSync point you at CouchDB running on a VPS. That works, but it means owning and maintaining an always-on Linux box: a server to patch, a reverse proxy to keep TLS certificates fresh, and a process that dies on reboot if you didn't wire it up carefully. If you don't already have a server — or you just got rid of one — that's a lot of overhead for syncing some text files.

This guide takes the **serverless** route instead: LiveSync now supports S3-compatible object storage, and **Cloudflare R2** has a free tier that's wildly more generous than any markdown vault will ever need. There's nothing to host. You create a bucket, generate an API key, paste it into the plugin, and you're done.

## The one trade-off you should know first

Object storage can't do true real-time **LiveSync** mode — that's reserved for CouchDB and WebRTC peer-to-peer backends. With R2 you sync in **Periodic** mode (every N seconds) or **On-Events** mode (on save, on file open, on startup).

In practice, "sync on save" feels close enough to instant for notes. But if you need keystroke-level real-time mirroring between two devices editing the same file simultaneously, R2 isn't the right backend — go CouchDB. For the other 99% of use cases (write on your laptop, pick it up on your phone a minute later), R2 is perfect.

## What you'll need

- An [Obsidian](https://obsidian.md) install on each device you want to sync (desktop, iOS, Android — all supported).
- A free [Cloudflare](https://dash.cloudflare.com) account. R2 requires you to add a payment method to enable the service, but the free tier covers this use case completely — you won't be charged for a notes vault.
- About fifteen minutes.

A quick note on cost: R2's free tier includes **10 GB of storage**, plus a monthly allowance of Class A (write) and Class B (read) operations, and **no egress fees**. A text vault is a few megabytes. The only way to brush against the free limits is to sync extremely aggressively across many devices, and even then the fix is simply using a sensible sync interval. More on that at the end.

---

## Step-by-step setup

### Step 1 — Create the R2 bucket

In the Cloudflare dashboard, open **R2** in the sidebar and click **Create bucket**. Give it a name like `obsidian-vault`. Leave the location and jurisdiction on their defaults. That's your storage created.

### Step 2 — Create an R2 API token

From the R2 overview page, click **Manage R2 API Tokens** (top right), then **Create API Token**.

- Set the permission to **Object Read & Write**.
- Optionally scope the token to only the bucket you just made — good hygiene.
- Create it.

Cloudflare will show you three values **exactly once**. Copy all of them somewhere safe right now (a password manager, not inside Obsidian):

- **Access Key ID**
- **Secret Access Key**
- **S3 endpoint**, in the form `https://<ACCOUNT_ID>.r2.cloudflarestorage.com`

If you miss the endpoint, your Account ID is shown on the R2 overview page — just drop it into that URL pattern.

### Step 3 — Handle CORS (the part everyone gets stuck on)

R2's S3 API doesn't return browser-friendly CORS headers by default, and a failed connection test here is the single most common LiveSync + R2 complaint. You have two ways around it.

**Option A — the easy way (recommended).** Skip CORS configuration entirely. Instead, in the plugin (Step 5), enable the **"Use Custom HTTP Handler"** toggle — sometimes labeled *Use Internal API*. This routes requests through Obsidian's own networking layer instead of the browser's fetch, sidestepping CORS completely.

**Option B — configure it properly.** If you'd rather set the policy, open your bucket → **Settings** → **CORS Policy** → **Add CORS policy**, and paste:

```json
[
  {
    "AllowedOrigins": ["app://obsidian.md", "capacitor://localhost", "http://localhost"],
    "AllowedMethods": ["GET", "PUT", "HEAD", "DELETE"],
    "AllowedHeaders": ["*"]
  }
]
```

Those three origins are the ones Obsidian uses on desktop and mobile. Save it. (CORS changes can take up to ~30 seconds to propagate.)

### Step 4 — Install the plugin

Do this on the device that holds your **most complete vault** first — it'll be the source of truth for the initial upload.

In Obsidian: **Settings → Community plugins → Browse →** search **"Self-hosted LiveSync" →** Install → Enable. Open the plugin's settings to start the setup wizard.

### Step 5 — Configure the remote

Set **Remote Type** to **Object Storage (S3/MinIO/R2)**, then fill in:

| Field | Value |
|---|---|
| Endpoint URL | `https://<ACCOUNT_ID>.r2.cloudflarestorage.com` |
| Access Key ID | from Step 2 |
| Secret Access Key | from Step 2 |
| Region | `auto` *(R2-specific — do **not** use a normal AWS region)* |
| Bucket | `obsidian-vault` |
| Use Custom HTTP Handler | **ON** (if you took the easy CORS route in Step 3) |

Then turn on **End-to-End Encryption** and set a strong vault passphrase. With this enabled, R2 only ever stores ciphertext — even Cloudflare can't read your notes without the passphrase. **Write this passphrase down.** If you lose it, your synced data is unrecoverable.

Click **Test Connection**. Green means you're good. A failure here is almost always the endpoint format or the CORS handler toggle being off — see troubleshooting below.

### Step 6 — Choose a sync mode and do the first sync

Pick your sync behavior:

- **On Events** with *sync on save* enabled feels the closest to real-time.
- **Periodic Sync** every 30–60 seconds is a solid, low-maintenance default.

You can combine both. Run the initial synchronization from this device so your full vault uploads to R2.

### Step 7 — Add your other devices

Don't hand-enter all those settings on each device. Instead, on the already-configured device:

1. Open the command palette and run **"Copy settings as a new setup URI"**.
2. Set a passphrase for the URI (this encrypts the settings blob — it's separate from your vault passphrase).
3. Obsidian copies a long `obsidian://setuplivesync?settings=...` string to your clipboard. Save it somewhere outside Obsidian.

Then on each new device: install the plugin, open the command palette, run **"Use the copied setup URI"**, paste the string, and enter the URI passphrase. When prompted, set it up as a **secondary / subsequent device** so it pulls down from R2 rather than overwriting what's already there.

That's it — your vault now syncs across every device through R2.

---

## Troubleshooting

**Connection test fails immediately.**
Check the endpoint format first: it must be `https://<ACCOUNT_ID>.r2.cloudflarestorage.com` with no bucket name appended and no trailing slash. Then confirm **Region is `auto`** — a normal AWS region string will fail against R2.

**Connection works on desktop but not mobile.**
Mobile Obsidian only accepts HTTPS endpoints. The R2 endpoint is HTTPS natively, so this usually points back to CORS — enable the **Use Custom HTTP Handler** toggle, or double-check your CORS policy includes `capacitor://localhost`.

**"Failed to initialise the encryption key."**
Your E2EE passphrase doesn't match between devices. Every device must use the identical passphrase. The setup-URI method in Step 7 avoids this by carrying the encryption settings for you.

**It connects but nothing syncs.**
Remember object storage can't use LiveSync mode — make sure you actually selected **Periodic** or **On Events**, not LiveSync, in Step 6.

---

## Keeping it free and healthy

- **Use a sane sync interval.** A 30–60 second periodic interval (rather than 5s) keeps your monthly operation count comfortably inside the free tier, even across several devices.
- **Never stack sync engines.** Do not run iCloud, Dropbox, Google Drive, or Obsidian Sync on the same vault folder at the same time as LiveSync. Two sync systems fighting over the same markdown files is the fastest way to corrupt a vault. Let R2 be the only sync layer.
- **Back up separately.** Sync is not backup. Consider the companion *Differential ZIP Backup* plugin, or periodic copies of the vault folder, so a bad sync or accidental delete-everywhere doesn't take your notes with it.
- **Wait for the indicators.** When LiveSync shows progress in the status bar, let it finish before closing Obsidian to avoid interrupting a write.

---

## Was it worth it?

For the price of about fifteen minutes, you get private, encrypted, cross-device Obsidian sync with no server to maintain and no monthly bill. You hold the encryption keys, your data lives in storage you control, and there's no always-on machine quietly accruing security patches and reboots.

If you later decide you want true real-time sync — or you end up with a spare always-on box — you can migrate to a CouchDB backend without losing anything. But for most people who just want their notes everywhere, R2 is the path of least resistance, and it's free.

Happy syncing.
