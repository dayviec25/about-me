---
title: "Skills: Teaching Agents Without Bloating Their Memory"
date: "2026-05-04"
author: "David Chung"
description: "How skills use progressive disclosure to give your agent deep how-to knowledge without burning tokens on every request."
---

## The problem

When you onboard a coding agent to a project, the obvious move is to write down everything it needs to know in a memory file — `AGENTS.md`, `CLAUDE.md`, whatever your harness reads. Context on the repo, terminology, how the key components work, conventions, and crucially, the procedures: how to run the end-to-end tests, how to validate a change, how to do the fiddly stuff that bites every fresh agent.

This works, and it's the right instinct. But there's a catch that's easy to miss: **everything in that memory file gets loaded into the system prompt of every single agent session.** Every request pays for it, whether the request needs it or not.

Take the end-to-end testing instructions. Those are only relevant when the agent is **making changes**. If you just ask the agent a question — "how does this component work?" — that entire testing section is dead weight. It silently burns tokens on a request that will never touch a test.

A verbose project memory file is the collective learning of every agent session that ever worked in that repo, so it tends to grow. The bigger it gets, the more this tax compounds.

## The fix: skills

The way to fix this is to move **conditionally useful** information out of the memory file and into a **skill**.

The thing that makes skills work is **progressive disclosure**. When your agent starts up, it only loads the skill's tiny `description` field into the system prompt — just enough to know *what the skill is for*. It does **not** load the body. Only when the agent decides it actually needs that skill does it go read the rest of the file.

That's the whole trick. You can store a huge amount of how-to knowledge — testing procedures, deployment steps, debugging playbooks — and none of it costs you anything on requests that don't need it. The system prompt stays lean. The knowledge is there when, and only when, it's relevant.

| | Everything in the memory file | Moved into a skill |
|---|---|---|
| Loaded into system prompt | Always, in full | Only the one-line `description` |
| Full content read | Every request | Only when the agent needs it |
| Token cost on an unrelated question | You pay for all of it | Roughly nothing |

## Doing it live

You don't hand-craft these. Just ask the agent. The move is exactly this — tell the agent:

> "Let's extract the end-to-end testing instructions in our `AGENTS.md` file into a project-level skill."

The agent pulls that big conditional chunk out of the memory file and drops it into a skill file, leaving the memory file smaller and the procedure intact.

One thing to know: **Claude Code already knows what skills are and how to create them out of the box.** Other harnesses may not understand the concept at all, so the same prompt can fall flat depending on what you're running.

## Teaching any agent to author skills

If your harness doesn't natively grok skills, you can teach it. Install the **`skill-creator`** skill, which was written by Anthropic. Once it's installed, your agent can follow those rules to create well-formed skills for you going forward — no matter which harness you're on.

To install and manage skills across harnesses, the handy tool is the **`npx skills`** CLI from Vercel:

```bash
npx skills
```

It's the main tool for installing and managing skills, and it supports pretty much any agent — which matters if you, like a lot of people, bounce between harnesses and don't want your setup locked to one of them. The repo lives at [vercel-labs/skills](https://github.com/vercel-labs/skills).

## The payoff

The point isn't that skills let you write *more* documentation. It's that they decouple **how much your agent can know** from **how much every request costs**.

Keep your memory file for the stuff that's always relevant — the project shape, the conventions, your hard preferences. Push everything conditional into skills, and let progressive disclosure decide when to surface it. Your agent gets smarter without your system prompt getting heavier.

One word of caution before you go installing skills off the internet, though: popular is not the same as good, and some skills are actively harmful. That's its own post — see [When Skills Hurt Your Agent](/blog/when-skills-hurt-your-agent).
