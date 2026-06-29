---
title: "Why I Do Agentic Engineering in the Terminal"
date: "2026-03-30"
author: "David Chung"
description: "The case for a keyboard-first, terminal-centric agentic workflow — flow you don't break, and a setup that follows you to your phone."
---

## The problem

When you work with coding agents all day, you spend a lot of time switching between things: the agent, the editor, the shell, a browser. Every one of those switches is a chance to break your flow. And the single biggest flow-breaker isn't the agent at all — it's your hand leaving the keyboard.

I resisted this for a long time, but I now do almost everything in the terminal — and the reasons are more practical than aesthetic.

## Reason one: your hands never leave the keyboard

This sounds minor. It isn't. When your hands stay on the keyboard, you stay in the flow. The moment you reach for the mouse — even for a second, even just to click a tab — your brain has to context switch. Do that every few seconds across a full day and the cost adds up to something real.

Yes, plenty of GUI apps have great keybindings. But keyboard control isn't their **primary interaction paradigm**. They're built mouse-first, and the keyboard is bolted on. That makes it genuinely hard to build the muscle memory of *hands always on home row*, because every so often you still have to grab the mouse anyway. Terminal apps are the opposite: they're designed keyboard-first, so there's simply no reason for your hands to go anywhere else.

## Reason two: the same workflow runs everywhere

The second reason is portability. A terminal-centric setup runs **exactly the same** on every machine you touch — and even on your phone. That's surprisingly hard to replicate any other way, and once you have it, going back feels like a downgrade.

So what's actually in the stack? Three pieces.

### WezTerm — the terminal emulator

[WezTerm](https://wezterm.org) is a high-performance terminal emulator with around **26k GitHub stars**. Two things make it stand out:

- **It's truly cross-platform.** It works the same on Windows, Mac, and Linux. If you've ever been forced onto Windows for work, having an identical terminal there is a genuine lifesaver.
- **It's configured in Lua.** Not static key-value settings — an actual scripting language, so your config can have conditions and logic and be as dynamic as you want. And it **hot-reloads**: change your color scheme and it updates instantly, no restart.

### tmux — the multiplexer

Inside WezTerm I run **tmux**, short for *terminal multiplexer*. The easiest way to understand it is by what it lets you do:

- **Split into panes.** Run an agent in one pane, an editor in another, and keep a third pane just for running commands.
- **Spin up windows (tabs).** Great for running multiple agent sessions in parallel.
- **Persistent sessions.** This is the killer feature. You can *detach* from a session and drop back to a plain shell — then re-attach later and land in the exact state you left. Even better, you can attach to the same session from **another device**, like your laptop or your phone. That continuity is the part that's nearly impossible to get without a terminal-centric workflow.

One caveat: stock tmux doesn't look or behave like the polished setups you see in videos. The status bar, the keybinds, the styling — that's all configuration you'll want to invest some time in.

### Neovim — the editor

The editor is **Neovim**, a modern Vim. Its whole purpose is to keep your hands on the keyboard:

- Move and scroll with keys; drop into insert mode to type.
- `dd` deletes the current line; `u` undoes it.
- **Relative line numbers** mean you can see a target is 11 lines up and just type `11k` to jump there.
- Plugins extend it — `space-s` to grep the codebase, `space-f` to find files by name.

Once the motions are in muscle memory, you navigate faster than any mouse ever could.

## The honest part: there's a learning curve

I won't pretend this is free. Vim especially has a real learning curve — there's a reason "how do I exit Vim" is a running joke. tmux needs configuration before it feels good. WezTerm's Lua config is powerful precisely because it expects you to write some.

But here's the thing worth holding onto: the **concepts** matter more than the specific tools. Keyboard-first interaction, parallel sessions, persistent state you can pick up anywhere — these apply just as well to GUI workflows. If the terminal isn't for you, steal the ideas and leave the mechanics.

## The payoff

Flow and portability. Your hands stay put, so you stay in the zone — and your entire workflow follows you from desktop to laptop to phone without missing a beat. For a job that's increasingly about orchestrating a crew of agents rather than typing every character yourself, that's exactly the foundation you want.
