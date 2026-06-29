---
title: "Plan With Interactive Artifacts, Not Walls of Text"
date: "2026-06-01"
author: "David Chung"
description: "Markdown plans give you a wall of text you can't point at. An interactive artifact in your own design system lets you annotate, click, and decide — and saves you grief during implementation."
---

## The problem

You ask your agent to plan a feature. Maybe you trigger plan mode, maybe you ask it to write the plan to a markdown file. Either way, what comes back is a **wall of text** — three or four options, each described in a few dense paragraphs, that you now have to read top to bottom.

That format fails you in two specific ways.

First, it's hard to scan. Options written as prose all blur together, and "revamp the home screen with an achievement system" reads fine as a sentence while telling you almost nothing about what it will actually *look* like.

Second — and this is the one that really bites — **you can't point at anything.** If part of the plan is wrong, your only move is to describe, in more text, which part you mean. You can't select a piece of the plan and say "this, this is the bit I don't like." So feedback gets vague, the agent guesses at what you meant, and the loop drags.

## The better flow: an interactive artifact

Instead of printing a wall of text, the agent generates an **interactive HTML artifact** and opens it in the browser. Crucially, it renders that artifact in **your project's own design system** — so the options don't just look like a generic mockup, they look like your actual app.

That changes the planning conversation completely:

- **Options are laid out visually.** You see the proposed screens side by side, in the real styling, instead of decoding a paragraph. Reviewing a concept becomes glancing at it, not parsing it.
- **You can annotate specific parts.** Drop a comment directly on the piece you're reacting to. That's the thing a markdown plan can't do — targeted, located feedback.
- **You click your decisions.** The open questions sit at the bottom as actual options you select, rather than questions you answer in prose.
- **You never leave the artifact.** Annotations and choices get sent straight back to the agent from inside the artifact — no detour back to the terminal to type up what you meant.

The whole thing turns planning from "read my essay and reply with an essay" into "look at this, mark it up, decide." It's just a far more honest way to evaluate something that's ultimately visual.

## The tool: lavish

I built a tool for this called **lavish** — [github.com/dayviec25/lavish-axi](https://github.com/dayviec25/lavish-axi), free and open source. (The name is a play on being "richer than a rich editor.")

The part that makes it stick is that you install it **as a skill**. Skills load lazily — the agent only reads the description up front, and only pulls in the full instructions when it decides it needs them. So once lavish is installed, the agent automatically reaches for it on planning-type questions, without you having to remember to ask. You just say "come up with some options and let's discuss," and it launches the editor on its own.

It's consistent enough that you have to go out of your way to *not* use it. Saying "don't use lavish" becomes the special case.

## Why this pays off later

Here's the real reason interactive planning matters, and it has almost nothing to do with the planning step looking nice.

When you clarify requirements well **up front**, you barely have to interfere during implementation. The agent already knows what you want, in detail, with the ambiguous calls already decided. So it can run the whole middle of the task on its own while you go do something else.

The vague-markdown-plan flow does the opposite. Fuzzy requirements mean the agent guesses, you catch it mid-build, you correct it, it re-guesses — constant interruptions through the part of the work that should be hands-off.

This fits a broader pattern worth internalizing: **you spend your time at the start and the end of a task, not the middle.** At the start, you plan the requirements clearly. At the end, you hold the bar on quality. The middle — the actual implementation — is where the agent works alone, and the more of that middle you can free up, the more tasks you can run in parallel.

An interactive artifact is how you make the start count. Get the requirements right where it's cheap to get them right, and the rest of the task stops fighting you.
