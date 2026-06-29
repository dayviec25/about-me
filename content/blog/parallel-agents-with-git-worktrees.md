---
title: "Running Agents in Parallel With Git Worktrees"
date: "2026-06-22"
author: "David Chung"
description: "Worktrees let two agents work without stepping on each other — but the bookkeeping is its own tax. Here's how to drop it."
---

## The problem

Once one agent is happily running in a directory, the natural next move is to spin up a second one and double your throughput. So you open another terminal tab, point a fresh agent at the **same repo directory**, and turn it loose.

Then everything breaks. Two agents editing the same working directory **step on each other's toes** — one stages files the other didn't expect, edits collide, and you spend more time untangling the mess than you saved by parallelizing. That's the wall everyone hits the moment they try to scale past a single agent.

The default fix is **git worktrees** — and I built a free, open-source tool that smooths out their rough edges.

## What a worktree actually is

A git worktree is, roughly, a clone of your repo directory — a **separate working directory** backed by the same repository. You create one like this:

```bash
git worktree add ../myproject-feature
```

Now you've got a second directory on the filesystem. An agent can do whatever it wants in there and it **won't conflict** with the agent running in your original directory. Problem solved.

Except it isn't, quite.

## Worktrees become debt

Here's the part nobody warns you about: worktrees turn into **mental overhead** fast.

- **Naming them.** You stare at `git worktree add ../___` and try to think of a meaningful name. You waste five minutes, give up, and call it `myproject2`. (We've all done it.)
- **Remembering what each one was for.** Next week you wander into `myproject2` and have no idea what you were doing there.
- **Knowing if it's live.** Is there still an agent running in this worktree, or is it idle and safe to reuse?
- **Cleaning up.** When you're done, you have to remember to run:

```bash
git worktree remove ../myproject2
```

None of these is hard on its own. But stack them across a dozen worktrees and you've built yourself a second job: a bookkeeping tax that grows with exactly the parallelism you were trying to gain.

## Treehouse: worktrees without the bookkeeping

My fix is a tool called **treehouse** ([on GitHub](https://github.com/dayviec25/treehouse)), and it's refreshingly simple. Instead of manually adding, naming, tracking, and removing worktrees, you let treehouse manage the pool for you.

```bash
treehouse          # drop into a fresh worktree
treehouse          # run it again → another fresh worktree
treehouse status   # list worktrees, in-use vs idle
```

The magic is in the lifecycle. When you're done with a worktree, you just **close the tab** — and treehouse notices, frees it, and **reuses idle worktrees** the next time you ask for one instead of endlessly creating new directories.

So the four chores from above mostly evaporate. You don't name them, you don't track them by hand, `treehouse status` tells you what's live, and cleanup is "close the tab."

## Scaling up with tmux

Worktrees solve the *conflict* problem; they don't solve the *switching* problem. For that, pair treehouse with **tmux**.

tmux lets you run each agent session in its own tab (a "window" in tmux terms), and with a **status bar** at the top you can see at a glance which sessions need your attention. Keyboard shortcuts jump you between them — no reaching for the mouse, no losing your place.

The combination is what makes many parallel sessions actually manageable:

- **treehouse** keeps the agents from colliding on disk.
- **tmux tabs + status bar** keep *you* from getting lost across all of them.

## The payoff

The whole point of running agents in parallel is leverage — more work happening at once than you could ever do serially. But that leverage evaporates if managing the worktrees costs you as much attention as the work itself.

treehouse gives you the parallelism and quietly absorbs the bookkeeping. Spin up another session whenever you want one, close the tab when you're done, and let the tool remember the rest.
