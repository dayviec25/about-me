---
title: "Recruiting a First Mate: From Juggling Agents to Captaining a Crew"
date: "2026-06-29"
author: "David Chung"
description: "Running many agents in parallel is exhausting. The fix isn't more discipline — it's a first mate that manages the crew for you."
---

## The problem with winning

Say you've gotten good at parallel agents. You've got worktrees keeping them from colliding, tmux tabs lined up, three or four sessions all running at once. This feels like the destination.

It isn't. After a while you discover that juggling all those sessions is genuinely **exhausting**. The constant context-switching, the mental effort of reminding yourself what each session was even doing — it wears you down. You're playing whack-a-mole with an ever-growing number of agents, and that's not a satisfying endgame. It's a new bottleneck wearing the costume of productivity.

The fix is the one that finally made the whole thing click for me: you don't need more discipline. You need a **first mate**.

## A first mate, not more crew

The idea is simple. Instead of talking to every individual agent yourself, you talk to **one** thing — a first mate — and *it* manages the crew for you. You stay the captain. It does the juggling.

I built a tool called **firstmate** ([on GitHub](https://github.com/dayviec25/firstmate)), it's free and open source, and it's new. Getting started is just:

```bash
git clone https://github.com/dayviec25/firstmate
# then run an agent inside it and start talking
```

The first time you run it, it does a quick **talk-through setup** — no config files, just a conversation about your preferences. One of the things it asks is **how strict** you want to be with changes. Choosing **"full gate to PR"**, for example, means every change runs through **no-mistakes** to validate it before it ever reaches you as a PR.

## What it does behind the scenes

Here's where it gets good. When you hand the first mate a request, it doesn't just run one agent. It:

1. **Decomposes** the request into parallel tasks.
2. **Spins up tmux tabs** for each — the same thing you'd do by hand.
3. **Calls treehouse** to get an isolated worktree per task.
4. **Runs an agent** in each worktree to do the actual work.
5. **Runs no-mistakes** to get each change validated and the PRs ready for review.

All of that happens out of sight. You don't manage tabs, you don't name worktrees, you don't shepherd anything through review. You just keep feeding it work:

> "Add an `update` command to the CLI for these three repos that bumps each one's npm version to the latest."

The first mate recognizes that's not one task but **three parallel tasks**, fans them out, and gets to work. Or you give it something more open-ended:

> "Look at the three most recent open issues in this repo and let's discuss which ones are actionable."

And it pulls the issues while the background agents keep churning. Watching it context-switch on your behalf is oddly satisfying — precisely because you know that's the work you'd otherwise be doing yourself.

## Everything, composed

What makes firstmate more than a convenience wrapper is that it's **all the other tools coming together as one workflow**:

| Piece | Job |
|---|---|
| voice | how you talk to it |
| lavish | clarifying and planning the work |
| treehouse | isolated worktrees per task |
| no-mistakes | validating changes into clean PRs |
| gnhf | keeping agents running on long objectives |

Individually, each is a sharp tool. Composed behind a single conversation, they become a crew you command by talking — not a pile of sessions you babysit.

## The captain's mindset

Here's the twist that tells you it's working: once a first mate is absorbing all that overhead, you start to **run out of ideas** for what to ask it to do.

That sounds like a problem. It's the opposite. It means the bottleneck has shifted — it's no longer the agents, no longer the tooling, no longer your capacity to juggle sessions. **The bottleneck is now you.**

And that's exactly where you want it. When execution is cheap and parallel, the scarce resource becomes **knowing what's worth building**. The real work moves upstream: talking to your users, studying the competitive landscape, charting a direction worth pointing the crew at.

That's the shift from sailor to captain. You stop pulling on every rope yourself and start deciding where the ship should go. Recruit a first mate, and the question stops being "how do I keep up with my agents?" and becomes "where do I want to take them?"
