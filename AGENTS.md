Guidance for any AI coding agent (Claude Code, Codex, Cursor, etc.) in this repo. `CLAUDE.md` is `@AGENTS.md`; this file is the single source.

## The `ai/` folder

Two things live there, nothing else:

- [ai/rules/](ai/rules/) — **binding**. Read the relevant page before any code change: `tdd` (red/green for every feature and module, real SQLite file per test case), `commit-messages` (single-line Conventional Commits, explicit staging, no trailers), `git-hooks` (per-clone `gofmt`/`go vet` pre-commit).
- [ai/tasks/](ai/tasks/) — backlog of requested work, one file per task with a status header. See [ai/tasks/README.md](ai/tasks/README.md) for the format.

## Everything else is ai-memory

Concepts (how a domain works), module specs (`concepts/modules/01`–`11`), decisions (`decisions/0001`–`0026`, ADRs), gotchas (traps already hit) and session history live in ai-memory, not in the repo. Query it before non-trivial changes (`memory_query`, `memory_read_page`; pages such as `decisions/0013-data-integrity-fixes-and-known-gaps.md`, `gotchas/huh-form-skips-validators-on-eof.md`). Write there only when the user asks to remember something. Do not recreate a `wiki/`, `concepts/`, `decisions/` or `gotchas/` folder in the repo.

In `ai/rules/`, a `[[rules/<name>]]` link points to a sibling page; any other `[[path]]` link is an ai-memory page.

## Working agreements

- `master` takes pull requests only; cut a `feat/`, `fix/` or `docs/` branch from `master` and open a PR. Bumping `var version` in `cmd/main.go` releases on merge (ai-memory `decisions/0019-pr-only-master-with-ci-and-release-on-merge.md`).
