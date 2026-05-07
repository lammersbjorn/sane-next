# Instruction Surface Standard

This is the strict authoring standard for prompts, skills, AGENTS files, and agent definitions in `sane-next`.

It exists because research consistently showed the same pattern:

- bloated always-on context hurts performance and cost
- progressive disclosure works better
- executable checks beat repeated prose
- instruction surfaces drift when the same rule appears in multiple files

## Core rules

1. **Always-on surfaces stay tiny.**
   - Root `AGENTS.md`: target **<= 50 lines**
   - Copilot repo instructions: tiny supplement only
   - Never put repo summaries, tutorials, file maps, or large doctrine in always-on context

2. **One rule, one owner.**
   - Put each rule in the smallest surface that can own it
   - Do not duplicate the same policy across `AGENTS.md`, Copilot instructions, skills, overlays, and agent files

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

Forbidden content:

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
- add only Copilot-specific notes that do not belong in `AGENTS.md`
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
3. `Don't Use When` or `Use Main Session When`
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

1. be clear and direct
2. use exact success criteria
3. use examples when behavior is easy to misread
4. prefer exact boundaries over broad style doctrine
5. reserve hard prohibitions for destructive or security-sensitive behavior
6. avoid verbosity constraints that could degrade quality
7. keep prompts structured and easy to scan

## Anti-patterns

- giant AGENTS files
- repo summaries in always-on context
- duplicated policy across surfaces
- one skill doing multiple unrelated jobs
- hidden or vague verification
- placeholder-heavy plans
- passive prose where executable checks should exist
