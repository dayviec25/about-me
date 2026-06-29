---
title: "Your Agent's Global Memory File"
date: "2026-04-13"
author: "David Chung"
description: "The one file that loads into every agent session you ever run — and why it should stay tiny."
---

## The problem

When you onboard a coding agent — Claude Code, Codex, whatever — you're working with a fresh recruit who has no idea how you like to work. The fix is a **memory file**: a markdown file the agent reads to ramp up. For Claude Code that's `~/.claude/CLAUDE.md`; other harnesses look for `AGENTS.md` in their own standard location.

There are two kinds: a **global** one that applies everywhere, and a **project-level** one (covered in a sibling post). This post is about the global file, and the single most important thing to understand about it is this: **every line gets injected into the system prompt of every single session, across every project.**

That changes the math completely. This isn't documentation you write once and forget. It's a tax you pay on every request. Bloat here doesn't just sit there harmlessly — it silently burns tokens forever. So the rule is simple: **keep it tiny.** My entire global memory file is about **27 lines** — that's the target I'd aim for. Personal preferences and universal rules only; anything project-specific or conditional belongs elsewhere.

So what earns a spot in 27 lines? Here are three rules worth stealing, and — more importantly — the *why* behind each, because the why is what makes them stick.

## Rule 1 — "Never use em-dash"

Models are trained to reach for the em-dash by default. It shows up everywhere, and once you notice it you can't stop noticing it. The problem is that it reads *robotic*. When you ask an agent to write something with your name on it — a PR description, a commit message, a note to a teammate — that telltale punctuation is a giveaway that a machine wrote it.

One line in the global file fixes it for every session: prefer a plain dash. (Yes, I'm aware of the irony given how this blog is written. Pick the rule that fits your own voice — the point is that *you* decide, not the model's training defaults.)

## Rule 2 — "Don't give too much weight to development cost"

This one is less obvious and more valuable. Here's the rule in full: **when making technical decisions, don't give too much weight to development cost.**

Why does it matter? Because models are trained on human-written data, and humans estimate projects in **days, weeks, and months**. The model absorbs that as ground truth, so when it's weighing options it implicitly assumes the "good" solution is expensive to build and the "cheap" one is the pragmatic choice. The result: it biases toward solutions that are low quality, not scalable, or hard to maintain — all to save a build cost that, for an agent, often doesn't exist.

The memorable demonstration: ask a frontier model to **estimate** how long it would take to build a local 3D first-person shooter with AI enemies. It'll tell you weeks or months. Now ask it to **actually build it** — and you get a playable version in a few minutes. The model doesn't yet seem to know it can code far faster than the humans it learned from.

| | Model's instinct | Reality |
|---|---|---|
| "Estimate the FPS game" | weeks to months | — |
| "Build the FPS game" | — | playable in minutes |

Put the rule in your global file and you correct that bias on every decision the agent makes. It stops talking you out of the right architecture because it thinks the right architecture is "too much work."

## Rule 3 — "Reproduce the bug end-to-end first"

The rule: **for bug fixes, always start by reproducing the bug end-to-end, as close to the real user experience as possible.**

Left to their own devices, agents love to write unit tests. Unit tests feel productive, but they often don't guard the actual product behavior you care about — you end up with green checkmarks and a bug that still ships. Reproducing the issue the way a real user hits it, end-to-end, is far more reliable. Once you can reproduce it for real, the fix is honest, and you have a test that actually proves the thing is fixed.

So: lean into E2E, and tell every session to do so by default.

## Was it worth it?

Three lines of guidance — don't sound like a robot, don't undervalue your own speed, prove the bug is real — and they apply to every task you ever hand an agent. That's the whole appeal of the global memory file: small, durable, high-leverage.

The discipline is in what you *leave out*. The moment something is project-specific, or only relevant some of the time, it doesn't belong here — it's quietly taxing every unrelated request you make. Keep this file lean, push everything else down into project memory and skills, and your agents start every session a little smarter without costing you a thing.
