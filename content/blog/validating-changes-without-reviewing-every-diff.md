---
title: "Stop Reviewing Every Diff: A Validation Pipeline for AI Code"
date: "2026-06-08"
author: "David Chung"
description: "If every AI change needs your manual diff review, you become the bottleneck. Here's a pipeline that validates changes the way a good engineering director shapes quality — through process, not by reading every PR."
---

## The problem

AI writes code fast. Faster than you can review it, which is the whole problem.

If your rule is that every change gets your eyes on the diff before it merges, then **you** become the hard cap on velocity. You can only review so many diffs in a day, so it doesn't matter how many agents you have running — your throughput is pinned to your own reading speed. And let's be honest: nobody became an engineer because they love reviewing diffs all day.

So there's a mindset shift to make here: stop acting like the senior engineer who reviews every PR, and start acting like an **engineering manager or director.** Directors mostly don't read PRs. They shape their team's quality through good process and culture, and trust the team to carry it out. That's exactly the relationship you want with a crew of agents — quality enforced by a pipeline, not by you reading every line.

## The pipeline: no-mistakes

I built a pipeline for this called **no-mistakes** — [github.com/dayviec25/no-mistakes](https://github.com/dayviec25/no-mistakes), free and open source, and runnable as a skill (just say "no mistakes" and it runs).

When the agent says a change is done, you don't open your editor. You hand the change to the pipeline, and it walks it from rough first-pass code all the way to a clean, merged PR. In order:

1. **Branch and commit.** Creates a branch if one doesn't exist yet, and commits the work.
2. **Isolated worktree.** Everything runs in a separate git worktree, so nothing the validation does can touch your working repo.
3. **Infers your real intent.** It reads the agent session to understand what you were actually trying to accomplish — not just what the code literally does.
4. **Rebases up front.** Rebases onto the latest `origin/main` and resolves merge conflicts early, before they can ambush you later.
5. **Adversarial review in a fresh context.** A clean context window reviews the change adversarially. This is where most problems get caught. Obvious ones it **self-corrects**; ambiguous ones with product implications get **escalated to you** for a decision.
6. **End-to-end test, with evidence.** It tests the change against the original intent and **records evidence** — a screenshot, a video, a log — that you can actually inspect to confirm it does what you asked.
7. **Documentation pass.** Updates the relevant docs to match the change.
8. **Lint check.** Makes sure there are no linting problems before it pushes.
9. **Push and open a PR.**
10. **Babysits the PR to merge.** It keeps watching — new merge conflicts, CI failures — and works them until the thing is actually merged, instead of leaving you to nurse a red pipeline.

The slow part is real, but you never sit and watch it. You kick it off and go spin up the next task. You come back when it says everything passed.

## The PR is built for a fast decision

What lands at the end isn't a naked diff. The PR summarizes:

- the **original intent** behind the change,
- **what changed**,
- **how it was tested** (with that recorded evidence to click into),
- what the pipeline **found and fixed** along the way,
- and a **risk assessment**.

That risk assessment is the load-bearing piece. You use it to decide how much of your attention the change actually deserves. For **low-risk changes, you often don't read the diff at all** — you've validated over and over that anything you'd have caught, the pipeline already caught. You spend your scarce review time only on the **riskier** changes, where your judgment genuinely adds something.

That's the whole trick. You're not abdicating quality control — you're spending it where it counts instead of spreading it thin across every trivial change.

## The payoff

This is how you scale the *volume* of code you ship through a crew of agents without losing your grip on quality. The bottleneck stops being your eyeballs and becomes the process, and the process doesn't get tired at 5pm.

You still hold the bar. You just hold it at the end, on the changes that matter, instead of standing in the doorway inspecting every single thing that walks through. Be the director, not the gate.
