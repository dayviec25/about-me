---
title: "Choosing an Agent Harness (and Why You Should Stay Agnostic)"
date: "2026-04-06"
author: "David Chung"
description: "A survey of four coding-agent harnesses, their honest trade-offs, and why the smartest move is to keep your workflow agnostic to all of them."
---

## The problem

There are a lot of coding-agent harnesses now, and people love to argue about which one is *the best*. I think that's the wrong question. The landscape is moving so fast that whichever harness wins this month may not win the next — so the real skill isn't picking a winner, it's **not getting locked in**.

I run four different harnesses regularly, and I'm deliberately strict about keeping my workflow agnostic to all of them. Here's the survey, with the honest trade-offs of each.

## The four harnesses

### Claude Code

If you're on an Anthropic subscription, **Claude Code** is basically the only practical choice — and that's not a complaint, because it's a good harness. It has the most sensible default experience out of the box and a genuinely rich feature set. You don't have to fight it to get going.

The downsides: it can be a little **buggy** at times, and it's **not as customizable** as some alternatives. You take the defaults largely as given.

### Codex CLI

**Codex CLI** is written in Rust, and you can feel it — it's a bit **smoother** than Claude Code in use. It's also **open source**, which has a neat practical upside: when you hit a problem, you can often have Codex inspect its *own* source code and figure out a workaround on its own.

The trade-off is that it's **light on bells and whistles** and, like Claude Code, **not very customizable**.

### pi coding agent

The **pi coding agent** is built on a single philosophy: be **minimal and highly extensible**. There's almost nothing there by default — which is exactly the point. If you don't want any bloat and you enjoy tinkering to make a tool truly your own, pi is built for you.

The flip side is right there in the philosophy: you'll be assembling much of the experience yourself.

### opencode

**opencode** rounds it out, and it's a strong all-rounder. It has a **buttery-smooth TUI** and good integration with pretty much every model you can find. It's also **more complete out of the box** than pi — you don't have to build as much before it's useful.

That combination makes it the standout **grab-and-go, model-agnostic** option: if you want something you can pull off the shelf and just go, without committing to a single provider, opencode is a good pick.

## How they stack up

| Harness | Strength | The catch |
|---|---|---|
| Claude Code | Sensible defaults, rich features | Sometimes buggy, less customizable |
| Codex CLI | Smooth (Rust); can read its own source | Light on features, not very customizable |
| pi coding agent | Minimal, highly extensible | You build most of it yourself |
| opencode | Smooth TUI, model-agnostic, complete | Little to fault; the well-rounded default |

## The real thesis: stay agnostic

Notice what the table *doesn't* do: crown a winner. That's intentional. Who knows which model or harness will be the best performer next month? Betting your whole workflow on one of them is a bet you don't need to make.

So the discipline is to keep your **workflow** agent-agnostic, even while you happen to use one harness day to day. Everything you build around the agent — your habits, your tooling, your memory — should work no matter which harness you're driving.

One concrete way to do that is the **symlinked-memory trick** I cover in a sibling post: keep a single memory file on disk, and symlink each harness's expected filename (`CLAUDE.md`, `AGENTS.md`, and so on) to point at it. One source of truth, every agent reads it, and switching harnesses costs you nothing. That's agnosticism in practice, not just in principle.

## The payoff

Pick whichever harness feels best today — the defaults of Claude Code, the smoothness of Codex, the minimalism of pi, the all-round polish of opencode. Just don't marry it. Build your workflow so the harness is a swappable part, and the next time the landscape shifts, you get to ride it instead of rebuilding around it.
