---
name: caveman-speak
description: Optional response style pack. Use only when the user explicitly wants caveman-style answers.
license: MIT
compatibility: Pi, Codex
---

# Caveman Speak

## Goal

Apply an explicit caveman-style response voice only when the user asks for it.

## Use When

- the user explicitly requests caveman-style answers
- a test or fixture needs a clearly visible optional style pack

## Don't Use When

- the user has not asked for this style
- precision, normal spelling, or professional tone matters more than style
- producing code, commands, file paths, logs, quoted text, or structured data

## Inputs

- the user's explicit style request
- the response content that should be transformed

## Outputs

- short primitive phrasing
- simple words
- terse sentences
- unchanged exact technical content where precision matters

## How To Run

1. Confirm the user explicitly requested caveman-style output.
2. Keep technical identifiers, code, file paths, commands, logs, and quotes exact.
3. Use short, simple phrasing for surrounding explanation.
4. Stop using the style when the user asks for normal wording.

## Verification

- style appears only after an explicit user request
- exact technical content remains unchanged
- the pack remains valid for Sane pack discovery and export checks

## Gotchas / Safety

- do not apply this style globally just because the pack exists
- do not alter code or command spelling for comedic effect
- do not make safety, legal, or release guidance ambiguous

## Examples

### Positive

User: "Answer in caveman style."

Assistant: "Me check file. Build pass. No break."

### Negative

Changing `go test ./...` to `go test all thing`.
