---
title: "One Memory File, Two Names: Symlinking CLAUDE.md and AGENTS.md"
date: "2026-04-20"
author: "David Chung"
description: "How to keep your agent memory file in sync across Claude Code, Codex, and every other harness — with one symlink."
---

## The problem

If you take agentic engineering seriously, you probably don't use just one harness. Claude Code is the obvious pick on an Anthropic subscription. Codex CLI is smooth and open source. There's pi if you like things minimal, and opencode if you want something model-agnostic off the shelf. Each has strengths, and the smart move is to stay flexible.

But here's the friction: **Claude Code reads `CLAUDE.md`, and almost everyone else reads `AGENTS.md`.** Same purpose — the memory file that ramps your agent up on how you work — two different filenames. Maintain both by hand and they *will* drift. You'll fix a rule in one, forget the other, and quietly hand different instructions to different agents working on the same code.

## The fix: one file, two names

You don't need two files. You need one source of truth that answers to both names. That's exactly what a **symbolic link** is for — make one filename a symlink to the other, and both resolve to the same bytes on disk. Edit either name and you've edited the one underlying file.

It's a thirty-second change:

```bash
# global
cd ~/.claude && ln -s ~/.codex/AGENTS.md CLAUDE.md

# project (commit both; the symlink travels with the repo)
cd your-project && ln -s AGENTS.md CLAUDE.md
```

A couple of things worth calling out:

- **Pick a real file and a link.** In the examples above, the actual content lives in `AGENTS.md`, and `CLAUDE.md` is the pointer. It doesn't matter which direction you choose — just be consistent so you always know where the real file is.
- **Commit the symlink in the project case.** Git stores symlinks natively, so the link travels with the repo. Anyone who clones it gets both filenames resolving to your single memory file, no setup required.

## Why bother

The payoff is being **genuinely agent-agnostic**. This isn't a tidiness thing — it's a hedge. The landscape moves fast, and which model or harness is best genuinely changes month to month. If your workflow is welded to one tool's filename, switching means re-plumbing your memory files every time. With the symlink in place, you can swap Claude Code for Codex tomorrow, or run both side by side, and they both read the same instructions automatically.

And it works at **both levels** — your global `~/.claude` memory file and each project's memory file. Set it up once per location and you stop thinking about it.

## Was it worth it?

For a single `ln -s`, you eliminate an entire category of "wait, why is this agent behaving differently" bugs, and you keep yourself free to chase the best tool without paying a tax to switch. One file, two names, zero drift. That's the kind of trade I'll take every time.
