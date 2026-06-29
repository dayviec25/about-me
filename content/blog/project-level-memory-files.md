---
title: "Project Memory Files Are Grown, Not Written"
date: "2026-04-27"
author: "David Chung"
description: "Why the best project-level CLAUDE.md isn't authored up front — it's accumulated, one correction at a time."
---

## The problem

Your global memory file should be tiny — 27 lines or so, personal preferences and universal rules only, because it's taxed on every request (there's a sibling post on that). The **project-level** memory file is the opposite. It's allowed to be verbose, and it should be, because it's doing a completely different job.

This is the per-project `CLAUDE.md` / `AGENTS.md` — and yes, you'll want to symlink those two names together here too, so every harness reads the same file. But the real question isn't where it lives. It's *how it gets written*. Most people sit down and try to author the perfect onboarding doc up front. That's the wrong instinct.

## What goes in it

A good project memory file captures the **collective learning of every agent session in the project**. The kind of content that earns its place:

- **Project overview** — what this thing actually is.
- **Repo layout** — where the important stuff lives.
- **Terminology** — the project-specific words your agent needs to not misuse.
- **How key components work** — the load-bearing parts, explained.
- **End-to-end testing steps** — how to actually verify a change against real behavior.
- **Conventions** — the house style and patterns to follow.

This is genuinely more than 27 lines, and that's fine. Unlike the global file, this only loads when you're working in this project, so the verbosity is paying for itself.

## The method: grow it, don't write it

Here's the part that matters. **Don't write this file by hand up front.** You'll guess wrong about what the agent needs, you'll over-document the obvious, and you'll miss the actual sharp edges.

Instead, grow it through correction. The loop is dead simple:

1. The agent does something wrong.
2. You correct it.
3. You tell it to **record the lesson in the memory file** so it doesn't repeat the mistake.

Do that consistently and the file becomes a living record of every trap your crew has stepped in. Each session makes the next one smarter. You don't need a fancy memory system or a vector database for this — a plain markdown file is the entire mechanism. The crew gets more experienced over time, for free, as a byproduct of you just doing your normal work.

## The failure mode

There's a catch, and you'll hit it eventually: the file gets **bloated**. The same accumulation that makes it valuable also makes it grow without bound, and a lot of what's in there is only *conditionally* useful.

Take the end-to-end testing instructions. They're essential when the agent is making changes — and totally dead weight when you're just asking it a question about the code. But because it's all in the memory file, those instructions load every time regardless, burning tokens on requests that don't need them.

The fix is to move that conditional knowledge into **skills**, which load only when the agent actually decides it needs them. That's a topic of its own — see the sibling post on skills for how to extract sections out of an overgrown memory file. For now, the thing to internalize is the shape of the practice: grow the file through correction, then prune the conditional parts into skills when it gets heavy.

## Was it worth it?

The project memory file is the closest thing you have to institutional knowledge for a crew that has no long-term memory of its own. Authored up front, it's a stale guess. Grown one correction at a time, it's the difference between an agent that keeps relearning your codebase and one that shows up already knowing where the bodies are buried. Correct, record, repeat — and let the file get smarter than you remembered to make it.
