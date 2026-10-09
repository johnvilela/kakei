# Tasks

Backlog of requested work. One file per task, `NNNN-short-slug.md`, numbered in the order they arrive. Nothing else goes in this folder.

## File format

```markdown
---
tags: [tasks, <module>, <module>]
---

# <Task title>

**Status:** NEW | IN PROGRESS | BLOCKED — <date requested>, <where (chat, issue, session)>
**Branch / PR:** `feat/<slug>` (#NN) — once work starts

## Request

What was asked, in the requester's words where it matters.

## Rules

Bullet list of the rules the feature must respect. Each one a testable sentence.

## Open questions

Questions still waiting on the requester. Remove each one when answered and fold the answer into the rules above.

## Done when

Acceptance list. Checked off in the PR.
```

## Lifecycle

- This folder holds only work that is still to be done: requests parked for later. Work that starts right away gets no task file, and nothing here describes what already shipped.
- A task is created when a request arrives and is not picked up immediately. `Status` moves NEW → IN PROGRESS when the branch is cut.
- Delete the file when the PR is merged into `master`. The how-it-works of what shipped belongs in ai-memory (`decisions/…`, `concepts/…`), not here.
- Rules discovered while building go to [`../rules/`](../rules/), not into the task file.
