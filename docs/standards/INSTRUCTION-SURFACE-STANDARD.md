# Instruction Surface Standard

This is the strict authoring standard for prompts, skills, AGENTS files, and agent definitions in `sane-next`.

It exists because current prompt-engineering research and vendor guidance consistently show the same pattern:

- bloated always-on context hurts performance and cost
- progressive disclosure works better
- executable checks beat repeated prose
- instruction surfaces drift when the same rule appears in multiple files
- positive behavioral instructions work better than long lists of negative constraints
- negative prompting is unreliable: models often struggle to follow prompts that mainly say what behavior to suppress

Research basis: OpenAI prompt guidance for GPT-5-family models recommends clear goals, success criteria, examples, tool-use policy, and saying what to do instead of only what to suppress. Anthropic guidance similarly favors positive examples and direct instructions. ETH Zurich AI Center / EMNLP 2024 research on negative prompting found that LLMs often fail to reliably suppress prohibited content or behavior when steered mainly by negative constraints.

## Core rules

1. **Always-on surfaces stay tiny.**
   - Root `AGENTS.md`: target **<= 50 lines**
   - Copilot repo instructions: tiny supplement only
   - Never put repo summaries, tutorials, file maps, or large doctrine in always-on context

2. **One rule, one owner.**
   - Put each rule in the smallest surface that can own it
   - Put shared policy in one owning surface and link to it from `AGENTS.md`, Copilot instructions, skills, overlays, and agent files

3. **Progressive disclosure is mandatory.**
   - Always-on: only durable repo truth
   - Trigger-loaded: skill body
   - On-demand: `references/`, `scripts/`, `agents/`

4. **Executable enforcement beats prose repetition.**
   - If a rule exists mainly to prevent a bad action, prefer a hook, validator, test, or typed contract
   - Use prose to point to the check, not to restate the same warning repeatedly

5. **Research before implementation is mandatory.**
   - For ecosystem-dependent choices, research current sources first
   - Label mirrors, summaries, and incomplete sources as lower confidence

## AGENTS.md

Allowed sections:

1. `Product frame`
2. `Startup rules`
3. `Hard boundaries`
4. `Done`

Allowed content:

- product/runtime boundary
- exact current-truth files to read first
- hard repo boundaries
- minimal done criteria

Route elsewhere:

- long process doctrine
- model fandom
- historical narrative
- repo overviews
- duplicated skill content
- speculative future plans

## Copilot repo instructions

Purpose:

- tiny supplement for GitHub Copilot-specific behavior only

Rules:

- point to `AGENTS.md` as the primary repo instruction surface
- add Copilot-specific notes only when `AGENTS.md` is not the right owner
- keep it shorter than `AGENTS.md`

## Skills

### Directory structure

```text
.agents/skills/<skill-name>/
  SKILL.md
  references/
  scripts/
  agents/
```

Only `SKILL.md` is required.

### Naming

- skill directories: `lowercase-hyphens`
- frontmatter `name`: matches directory name
- keep names specific and concrete

### Frontmatter

```yaml
---
name: lowercase-hyphens
description: Use this skill when...
license: MIT
compatibility: Pi, Codex
---
```

Rules:

- `description` is the discovery contract
- write it last
- focus on exact trigger conditions
- include sharp exclusions when useful

### Skill body

Required sections, in order:

1. `Goal`
2. `Use When`
3. `Use Main Session When`, `Route Elsewhere When`, or another positive routing section
4. `Inputs`
5. `Outputs`
6. `How To Run`
7. `Verification`
8. `Gotchas / Safety`
9. `Examples` (recommended)

Rules:

- one job per skill
- short action-oriented body
- keep gotchas in the body, not in references
- move bulky detail to `references/`
- put scripts in `scripts/` and call them explicitly

## Agent definitions

Every agent definition must declare:

1. role and non-goals
2. write boundary
3. expected output shape
4. verification expectation

### Codex-native agent files

Use TOML.

```toml
name = "agent_name"
description = "..."
model = "gpt-5.5"
model_reasoning_effort = "low"
sandbox_mode = "workspace-write"

developer_instructions = """
[role] [boundary] [output] [verification]
"""
```

### Markdown/YAML agent files

Use for Claude-style agent surfaces.

```markdown
---
name: agent-name
description: Use this agent when...
model: inherit
tools: ["Read", "Grep", "Bash"]
---

[role, process, output format, verification]
```

## Prompt writing rules

1. State the desired action path: "When X, do Y, because success means Z."
2. Use exact success criteria.
3. Use examples when behavior is easy to misread.
4. Prefer routing rules and workflow recipes over broad style doctrine.
5. Reserve hard prohibitions for destructive, security-sensitive, or architectural boundary behavior.
6. Convert ordinary negative rules into positive alternatives.
7. Use verbosity guidance that preserves quality.
8. Keep prompts structured and easy to scan.

### Positive-instruction rewrite rule

For each negative instruction, choose one bucket:

| Bucket | Prompt shape |
| --- | --- |
| Safety or irreversible boundary | Keep one short hard-stop rule. |
| Routing choice | Rewrite as `If this task matches X, use Y.` |
| Workflow preference | Rewrite as an ordered action step. |
| Quality bar | Rewrite as success criteria or verification. |
| Stale warning | Delete it. |

Examples:

```md
When writing docs:
1. Check docs/README.md for the owning location.
2. Edit the existing owning doc when one exists.
3. If no approved location fits, ask before creating a new docs path.
```

```md
When implementation appears to require Pi runtime changes:
1. Treat Pi as upstream runtime.
2. Build overlays, packs, adapters, or companion tools in this repo.
3. Stop and ask before choosing a runtime fork path.
```

## Anti-patterns

- giant AGENTS files
- repo summaries in always-on context
- duplicated policy across surfaces
- one skill doing multiple unrelated jobs
- hidden or vague verification
- placeholder-heavy plans
- passive prose where executable checks should exist
