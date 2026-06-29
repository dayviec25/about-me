---
title: "Talking to Your Agents: Voice Input for Prompting"
date: "2026-05-18"
author: "David Chung"
description: "Why I've moved almost entirely to voice for prompting coding agents — the research, the local tool that makes it work, and the one case to keep typing."
---

## The problem

Prompting a coding agent well means writing a *lot* — context, intent, the back-and-forth of a real conversation. And typing all of that, all day, is slow. It's the quiet bottleneck in an otherwise fast workflow: the agent can do the work in minutes, but you're still pecking out the request.

The fix is one of those changes that sounds trivial and turns out to be huge: stop typing your prompts. Talk to your agents instead.

## Talking is about 3x faster than typing

This isn't a vibe. A Stanford paper put the two side by side and found that **talking is roughly three times faster than typing**. When prompting is most of what you do, tripling the speed of your primary input is an enormous productivity boost for almost no effort.

There's a fun footnote here, too. Dig into that paper's references and one of them is speech-recognition work from around 2016 co-authored by **Dario Amodei** — now the CEO of Anthropic. So the through-line is almost poetic: we're using speech-recognition technology to talk to Claude, which Dario also helped create. Small world.

## The tool: OpenSuperWhisper

The piece that makes this practical is [OpenSuperWhisper](https://github.com/starmel/OpenSuperWhisper). It's **free and open source** — which is exactly what this kind of foundational tool should be — and it runs the **Whisper model locally** on your machine. Nothing gets shipped off to a cloud service; the transcription happens right there, and the quality is genuinely excellent.

So the loop becomes: hit your shortcut, say what you want, and watch a clean transcription land in front of the agent. No typing, no round-trip to some API.

## The one time you should still type

Voice isn't the answer for *everything*. The clear exception is anything you wouldn't want to say out loud:

- **URLs**
- **File paths**
- Anything fiddly or sensitive

Speaking a URL character by character is miserable whether you're alone or surrounded by people. For those, just type. Voice for the prose, keyboard for the precise strings — that's the split.

## Power tip: teach it your vocabulary

Here's the trick that takes transcription from *good* to *eerily accurate*: OpenSuperWhisper supports a **custom-vocabulary initial prompt** (a system prompt, in the transcription settings).

The problem it solves is that generic speech models mangle the names that matter most to you — your projects, your products, your internal tools. They've never heard of them, so they guess. Drop those names into the initial prompt as common vocabulary, and suddenly it transcribes them correctly every time. If you work on projects with unusual or made-up names, this one setting is the difference between fighting your transcripts and forgetting they're even there.

## The payoff

This is about as close to free productivity as it gets. A local, open-source tool plus a few minutes of vocabulary setup, and your prompting gets roughly three times faster — while your hands barely move. If you're still typing every prompt to your agents, this is the easiest upgrade on the board.
