# Old Sane Postmortem

This file captures what the current `sane` repo taught us.

It is **reference only**. Do not copy the old repo into `sane-next`, and do not treat old Sane claims as the source of truth for the new product.

## What was worth preserving

- outcome-first workflow discipline
- minimal always-on guidance
- skill-first extensibility
- explicit verification and recovery
- RTK-style task-aware tool routing when a target repo already uses RTK

## What became bloated

- package-per-concern architecture
- TUI-heavy product direction
- overlapping planning and tracking surfaces
- duplicated prompt and policy prose across multiple files
- premature model-routing complexity
- too much framework ownership where the host runtime already solved the hard parts

## What helped

- the pack idea
- verification discipline
- skill-first reuse
- explicit operational boundaries

## What hurt

- too many shallow boundaries
- too many docs competing to be "current truth"
- product and framework concerns drifting together
- broad prompt surfaces that wasted context

## Transfer candidates

These are worth reusing conceptually, not copying literally:

1. **packs as a reusable unit**, but kept small and skill-first
2. **RTK-style routing patterns** when the target repo already uses RTK
3. **strong acceptance and reversibility checks**
4. **minimal always-on guidance with heavier trigger-loaded detail**

## Do not import directly

- old Sane source code
- TUI architecture
- multi-package starting structure
- multi-surface planning/doc layout
- broad model-routing assumptions

## Implication for implementation

`sane-next` should preserve the vision while discarding the bloat:

- keep the workflow discipline
- keep packs and extensibility
- keep repo truth explicit
- keep the system small
- let Pi own the runtime internals
