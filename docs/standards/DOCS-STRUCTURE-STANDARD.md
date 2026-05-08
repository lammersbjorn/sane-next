# Docs Structure Standard

This standard defines how documentation is organized in `docs/` so humans and agents can find the right source without creating parallel planning surfaces.

## Core rules

1. Keep docs source-grounded and current.
2. Put each durable fact or rule in one canonical place.
3. Prefer deleting or rewriting stale docs over adding corrective notes elsewhere.
4. Do not create repo-local TODO, plan, memory, scratch, or research-dump markdown files.
5. Use `README.md` files as short directory maps, not as extra doctrine.
6. Link to source files, commands, ADRs, or standards instead of duplicating their content.

## Directory ownership

| Directory | Owns | Does not own |
| --- | --- | --- |
| `docs/adr/` | Durable decisions, rationale, rejected alternatives, consequences | Task lists, current status, implementation logs |
| `docs/standards/` | Normative rules for repo structure, tracking, instructions, releases, and implementation discipline | One-off plans, historical narratives |
| `docs/reference/` | Stable background, postmortems, external-source summaries that remain useful | Active decisions, current work state |
| `docs/roadmap/` | Living product direction, release readiness, prioritized next work | Completed implementation diaries, chat transcripts |

## README files

Add a `README.md` to a docs subdirectory when it helps with at least one of these:

- explains what belongs in that directory
- defines naming or status conventions
- points to the highest-value files in a directory with several documents
- prevents agents from creating duplicate or misplaced docs

Keep each directory README short. If it grows into policy, move the policy to `docs/standards/` and link to it.

## ADR rules

- Use numbered filenames: `NNNN-short-kebab-title.md`.
- Keep accepted ADRs historically stable.
- If a decision changes, create a new ADR or mark the old ADR as superseded; do not rewrite history to hide the change.
- ADRs may reference standards that now govern day-to-day behavior.

## Roadmap rules

- Keep the roadmap focused on current and future direction.
- Move completed implementation detail out of the main reading path.
- Every roadmap item should name evidence that would prove completion.
- `TRACK.toml` owns the active execution window; the roadmap owns broader prioritization.

## Reference rules

- Reference docs must say why they still matter.
- Old context belongs here only if it remains useful for decisions or verification.
- If a reference becomes obsolete and no longer informs work, delete it.

## Change checklist

Before adding or changing docs:

1. Identify the reader job: decide, execute, look up reference, or understand history.
2. Pick the owning directory from the table above.
3. Check whether an existing doc should be edited instead of adding a new file.
4. Add links from the relevant directory README only when discoverability improves.
5. Run the repo's docs/tracking validation path when available.
