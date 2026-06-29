---
title: "Agent Ergonomics: Your Tools Decide Performance, Not Just Your Model"
date: "2026-05-25"
author: "David Chung"
description: "The tools you hand your agent — how it talks to GitHub, the browser, anything external — can cost 3x the tokens for the same work. Here's why, and what to do about it."
---

## The problem

When people try to make their coding agent better, they reach for the obvious levers: a smarter model, a longer prompt, more context. What almost nobody audits is the **tools** — the external interfaces the agent uses to talk to GitHub, the browser, the filesystem, whatever your task touches.

That's a mistake, because those tools quietly decide how much mileage you actually get out of the model. The agent only knows the world through them. A bloated, chatty tool drains your token budget and slows everything down before the model gets a chance to be smart.

Once you see the numbers, it's hard to unsee them.

## The benchmark that should change your mind

Take something every developer's agent does constantly: reading GitHub. Pull requests, issues, recent commits.

The popular way to wire this up is the **GitHub MCP server**. It's official, it's easy, you install it and your agent can suddenly talk to GitHub. So far so good.

Then you measure it. Benchmark the different ways of accessing GitHub for the **exact same tasks**, and the MCP server comes out badly:

| Approach | Token cost | Latency |
|---|---|---|
| GitHub MCP server | ~3x | >2x |
| CLI (same task) | baseline | baseline |

Three times the token cost and more than double the latency, for the same result. Put bluntly: if you're using the GitHub MCP, you're pretty much **wasting both time and money for no clear benefit**. The agent isn't getting anything extra — it's just paying a tax on every call.

That's the part that should sting. This isn't a quality trade-off where you spend more to get more. You're spending more to get the same thing, slower.

## AXI: treating the agent as a first-class citizen

So what's the alternative? In that same benchmark there's an option labeled **AXI** that posts the **lowest cost and the highest success rate** at the same time — the corner of the chart you actually want to live in.

[AXI](https://axi.md/) is a set of design standards for treating agents as first-class citizens, built on a simple observation: there's enormous upside in designing tools *for agents* instead of bolting an agent onto tools built for humans. It's **10 principles** for agent ergonomics — making a tool efficient and pleasant for an agent to use, the same way we obsess over ergonomics for human UIs.

A concrete one: **output format**. The reflexive choice for machine-to-machine data is JSON. But JSON is verbose — all those quotes, braces, and repeated keys are tokens the model has to read and pay for. Switching to a token-efficient output format can save **about 40% of tokens** versus JSON for the same payload. Multiply that across thousands of tool calls and it's real money and real latency.

The point isn't "JSON bad." It's that a tool designed with the agent as a first-class citizen makes a hundred small decisions like that one, and they compound.

## It generalizes past GitHub

GitHub is just the example that's easiest to measure. The principle holds anywhere your agent reaches outside itself.

Browser automation is the obvious next target. Browser tools are notoriously expensive for agents — a single page can dump a wall of DOM into the context window. A tool designed around what the agent actually needs to *decide its next move*, rather than echoing everything, finishes the same task in **fewer turns and fewer tokens**. Same outcome, a fraction of the budget.

AXI is free and open source. If you want to use the tools, the catalog and per-tool setup instructions live at [axi.md](https://axi.md/).

## The takeaway

Before you hand a tool to your agent, **do a little research on its efficiency.** Don't just grab the first MCP server that shows up because it was easy to install — measure what it costs you per task in tokens and latency, and compare it to the boring alternative like a plain CLI.

The model gets all the headlines, but the tools around it set the ceiling on how far that model can carry you. Pick them like they matter, because they do.
