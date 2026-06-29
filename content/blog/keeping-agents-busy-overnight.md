---
title: "Keeping Agents Busy While You Sleep"
date: "2026-06-15"
author: "David Chung"
description: "How to hand an agent a long-running objective and review the results in the morning — without torching your weekly quota."
---

## The problem

Once you get comfortable working with coding agents, a pattern emerges. You spend real effort at the **start** of a task — clarifying what you actually want — and real effort at the **end**, holding the bar on quality. The whole middle? The agent does that on its own.

Which raises an obvious question: the more of that middle you can hand off, the more work you can run in parallel. So how do you get an agent to keep working in the middle for **longer and longer**?

That depends on giving it bigger, more complex objectives. And the most extreme version of "longer" is the one I find most appealing: I sleep seven or eight hours a night. How do I keep an agent productive that entire time?

I built a tool for exactly this, and it's free and open source.

## Enter agent-loop

The tool is called **agent-loop**, and it's built for long-running tasks. It's [on GitHub](https://github.com/dayviec25/agent-loop), and the usage is about as simple as it gets: you give it an objective, and it keeps going in a loop until it hits a **stop condition you defined**.

That's the whole idea. You're not babysitting a single prompt. You're handing over a goal and letting the agent iterate against it autonomously, overnight, while you're unconscious.

The best example — and the one I keep coming back to — is this:

> Pretend you are a seven-year-old kid using the app end to end. Find the first usability problem that would confuse a kid or stop them from knowing how to proceed. If you find one, stop and fix it. Then rinse and repeat.

Set that loose at midnight and it'll churn through usability issues one at a time, committing fixes as it goes. You wake up to a branch full of small, focused improvements.

## What you can actually watch

While it runs, agent-loop gives you a live view so you're not flying blind:

- **Token usage** — how much you've burned so far.
- **Iterations** — shown as a row of **moons**, one per loop.
- **Commits** — how many changes have actually landed.

Or you ignore all of it and go to bed, secure in the knowledge that the loop won't stop until there's nothing left to find.

## What it's good at — and what it isn't

This is the honest part. A loop running unsupervised for eight hours is only as useful as the objective is **checkable**. agent-loop shines on two kinds of goals:

**Verifiable objectives** — where success is measurable:

| Objective | Why it works |
|---|---|
| Reduce page load time | There's a number that goes down |
| Improve E2E test coverage | Coverage is a metric you can track |
| Experiment to improve a metric | The agent can run, measure, adjust, repeat |

**Trust-the-judgment objectives** — like the seven-year-old example, where you trust the agent to have reasonable taste about what counts as a real problem.

What you should *not* do is point an unsupervised overnight loop at something vague and unmeasurable and expect to love the result. Give it a target it can evaluate itself against, or give it a domain where its judgment is good enough that you'll keep most of what it produces.

## "But doesn't /go already do this?"

Fair question. The newer **`/go` command** in Codex and Claude Code does something similar — kick off a goal and let the agent run.

The difference is control. agent-loop lets you set:

- a **token cap**
- an **iteration cap**
- a **precise stop condition**

That matters more than it sounds. If you fire off an open-ended goal in Codex or Claude Code before bed, there's a real chance you wake up to discover your **entire weekly quota is gone**. The hard caps in agent-loop are exactly what keep an overnight run from quietly emptying your budget while you sleep.

## The payoff

When you wake up, you don't get a finished feature dropped into your lap — you get a **list of commits on a new branch**. You skim them over coffee, keep the ones you like, and drop the rest.

That's the trade I want from an overnight agent: it does the tedious middle while I sleep, and I keep final say in the morning. Hand off the grind, keep the judgment.
