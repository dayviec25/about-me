---
title: "When Skills Hurt Your Agent"
date: "2026-05-11"
author: "David Chung"
description: "Two reasons to be skeptical of internet skills — even wildly popular ones — and the one rule that protects you."
---

## The pitch is seductive

Skills are great. They let you give your agent deep, conditional knowledge without bloating its system prompt — if you want the full story on how that works, I wrote about it in [Skills: Teaching Agents Without Bloating Their Memory](/blog/skills-and-progressive-disclosure).

But the same mechanism that makes skills powerful also makes them dangerous. A skill is instructions your agent will follow. The moment you install one, you've handed a stranger a say in what your agent does on your machine. So before you go installing skills off the internet — **even the ones with a ton of GitHub stars** — there are two risks worth sitting with.

## Risk 1: Security

A skill can instruct your agent to run **essentially anything on your machine.**

That's not hyperbole. The agent has access to your shell, your files, your environment. A malicious or careless skill can tell it to do things you'd never approve if you saw them spelled out — and you often won't see them spelled out, because you didn't read the body, you just installed it because it had stars.

The concrete failure mode is exfiltration: a skill that quietly leaks your **API keys**, or even **credentials to your bank account**, to untrusted third parties without you ever knowing. You installed a productivity helper. What you actually installed was a path for someone else's instructions to run with your permissions.

## Risk 2: Popularity is not quality

Set the security problem aside entirely. Even a perfectly safe skill can make your agent **worse**.

Here's a sharp example. There's a skills repo with around **177,000 GitHub stars** — the kind of number that makes you assume it must be excellent. But when the skill is actually run through a benchmark that tests the agent's ability to build programs end to end, the picture flips.

The result:

| Metric | Effect of using the skill |
|---|---|
| Token usage | About **5% more** |
| Result quality | **Worse** |

More tokens, worse output. You'd be paying extra to make your agent dumber.

And here's the kicker — the repo isn't even written by the famous engineer it's associated with. The fame got it viral; rigor had nothing to do with it. **GitHub stars measure popularity, not whether a skill actually helps.** A lot of widely shared skills are just something that worked for one person anecdotally and happened to catch fire. None of that tells you it'll help *you*.

## The rule of thumb

So here's the rule worth internalizing:

> **Don't install any skill that claims to magically improve your agent's performance unless it publishes rigorous evidence backing that claim.**

"Magically make your agent better" with nothing but stars behind it is a red flag, not a recommendation. Evidence — a real benchmark, a reproducible eval, numbers you can check — is the bar. Anything less is someone else's anecdote running with your credentials.

## The takeaway

Skills are a genuinely good tool. Used well, they keep your agent sharp and your system prompt lean. But the install step is a trust decision, and "lots of people installed it" is not trust — it's herd behavior.

Treat every internet skill like code you're about to run as yourself, because that's exactly what it is. Read the body. Demand evidence. And when a skill promises magic, assume it's selling you something until it proves otherwise.
