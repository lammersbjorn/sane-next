#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO="$(cd "$ROOT/.." && pwd)"
TMP="$REPO/.tmp/acceptance"
GO_BIN="${GO_BIN:-go}"
NODE_BIN="${NODE_BIN:-node}"
PYTHON_BIN="${PYTHON_BIN:-python3}"

rm -rf "$TMP" "$ROOT/cli"
mkdir -p "$TMP"

cd "$ROOT"
"$GO_BIN" test ./...
"$GO_BIN" build ./...

"$GO_BIN" run . install --root "$TMP/install"
test -f "$TMP/install/packs/core-workflow/SKILL.md"
test -f "$TMP/install/pi-plugin/index.ts"
test -f "$TMP/install/pi-plugin/config-schema.toml"
test -f "$TMP/install/extensions/sane-next/index.ts"
test -f "$TMP/install/themes/github-dark-pro.json"
test ! -e "$TMP/install/skills"
test -f "$TMP/install/package.json"
"$PYTHON_BIN" - <<PY
from pathlib import Path
import json
package = json.loads(Path("$TMP/install/package.json").read_text())
if "skills" in package.get("pi", {}):
    raise SystemExit("generated Pi package should let the Sane extension discover enabled skills")
if package.get("pi", {}).get("extensions") != ["./extensions/sane-next/index.ts"]:
    raise SystemExit("generated Pi package extension manifest is wrong")
if package.get("pi", {}).get("themes") != ["./themes"]:
    raise SystemExit("generated Pi package theme manifest is wrong")
PY
grep -q 'resources_discover' "$REPO/pi-plugin/index.ts"
grep -q 'before_agent_start' "$REPO/pi-plugin/index.ts"
grep -q 'session_start' "$REPO/pi-plugin/index.ts"
grep -q 'tool_call' "$REPO/pi-plugin/index.ts"
grep -q 'user_bash' "$REPO/pi-plugin/index.ts"
grep -q 'registerCommand("sane-status"' "$REPO/pi-plugin/index.ts"
grep -q 'registerCommand("sane-lanes"' "$REPO/pi-plugin/index.ts"
grep -q 'registerCommand("sane-goal"' "$REPO/pi-plugin/index.ts"
grep -q 'buildSubagentRoutingHint' "$REPO/pi-plugin/index.ts"
grep -q 'appendEntry(LEDGER_ENTRY_TYPE' "$REPO/pi-plugin/index.ts"
grep -q 'buildRelevantLedgerContext' "$REPO/pi-plugin/index.ts"
grep -q 'skillPaths' "$REPO/pi-plugin/index.ts"
if command -v pi >/dev/null 2>&1; then
  PI_CODING_AGENT_DIR="$TMP/pi-agent" pi install "$TMP/install"
  PI_CODING_AGENT_DIR="$TMP/pi-agent" pi list | grep -q "$TMP/install"
elif [[ "${SANE_NEXT_REQUIRE_PI:-0}" == "1" ]]; then
  echo "pi is required for acceptance but was not found on PATH" >&2
  exit 1
else
  echo "pi not found; skipped live pi install (set SANE_NEXT_REQUIRE_PI=1 to require it)"
fi
"$GO_BIN" run . doctor --root "$TMP/install"
"$GO_BIN" run . export --config "$REPO/pi-plugin/config-schema.toml" --target codex --target-root "$TMP/export-default"
"$GO_BIN" run . export --config "$REPO/pi-plugin/config-schema.toml" --target pi --target-root "$TMP/export-pi"

"$PYTHON_BIN" - <<PY
from pathlib import Path
codex_root = Path("$TMP/export-default/.codex/skills")
pi_root = Path("$TMP/export-pi/.pi/skills")
root = codex_root
expected = ["core-workflow", "rtk-routing", "agent-lanes", "sane-router", "craft-router", "frontend-craft", "frontend-review", "frontend-accessibility", "docs-writing", "ux-copy"]
missing = [name for name in expected if not (root / name / "SKILL.md").exists()]
if missing:
    raise SystemExit(f"missing default exports: {missing}")
missing_pi = [name for name in expected if not (pi_root / name / "SKILL.md").exists()]
if missing_pi:
    raise SystemExit(f"missing Pi skill exports: {missing_pi}")
if (root / "example-user-pack" / "SKILL.md").exists():
    raise SystemExit("disabled user pack exported unexpectedly")
craft = ["craft-router", "frontend-craft", "frontend-review", "frontend-accessibility", "docs-writing", "ux-copy"]
required_sections = ["## Goal", "## Use When", "## Inputs", "## Outputs", "## How To Run", "## Verification", "## Gotchas / Safety"]
for name in craft:
    body = (root / name / "SKILL.md").read_text()
    if f"name: {name}" not in body or "compatibility: Pi, Codex" not in body:
        raise SystemExit(f"{name} frontmatter missing expected metadata")
    missing_sections = [section for section in required_sections if section not in body]
    if missing_sections:
        raise SystemExit(f"{name} missing required sections: {missing_sections}")
    upstream = root / name / "references" / "UPSTREAM.md"
    if not upstream.exists():
        raise SystemExit(f"{name} missing upstream provenance")
    if upstream.stat().st_size > 5000:
        raise SystemExit(f"{name} upstream provenance is too large for non-vendored notes")
router = (root / "craft-router" / "SKILL.md").read_text()
if "dispatch-only" not in router or "Choose " + chr(96) + "frontend-craft" + chr(96) not in router:
    raise SystemExit("craft-router does not read like a dispatch-only router")
for doctrine in ["outer radius", "tabular numbers", "WCAG compliance", "Diátaxis"]:
    if doctrine in router:
        raise SystemExit(f"craft-router contains subordinate doctrine: {doctrine}")
PY

"$PYTHON_BIN" - <<PY
from pathlib import Path
config = Path("$REPO/pi-plugin/config-schema.toml").read_text()
config = config.replace('id = "example-user-pack"\nenabled = false', 'id = "example-user-pack"\nenabled = true')
Path("$TMP/user-pack-enabled.toml").write_text(config)
PY

"$GO_BIN" run . export --config "$TMP/user-pack-enabled.toml" --source-root "$REPO/pi-plugin" --target codex --target-root "$TMP/export-user"
HOME="$TMP/home" "$ROOT/cli" export --config "$REPO/pi-plugin/config-schema.toml" --target codex

"$PYTHON_BIN" - <<PY
from pathlib import Path
root = Path("$TMP/export-user/.codex/skills")
expected = ["core-workflow", "rtk-routing", "agent-lanes", "sane-router", "craft-router", "frontend-craft", "frontend-review", "frontend-accessibility", "docs-writing", "ux-copy", "example-user-pack"]
missing = [name for name in expected if not (root / name / "SKILL.md").exists()]
if missing:
    raise SystemExit(f"missing user-pack exports: {missing}")
if not Path("$TMP/home/.codex/skills/core-workflow/SKILL.md").exists():
    raise SystemExit("default Codex-native export did not write under HOME/.codex/skills")
PY

"$NODE_BIN" --test "$REPO/pi-plugin/plugin.test.js"
rm -rf "$TMP/install/exports"
rm -f "$TMP/install/pi-plugin/index.ts"
if "$GO_BIN" run . doctor --root "$TMP/install" >"$TMP/doctor-broken.log" 2>&1; then
  echo "doctor unexpectedly passed with missing plugin" >&2
  exit 1
fi
"$GO_BIN" run . repair --root "$TMP/install"
test -d "$TMP/install/exports"
test -f "$TMP/install/pi-plugin/index.ts"
test -f "$TMP/install/extensions/sane-next/index.ts"
test -f "$TMP/install/themes/github-dark-pro.json"
"$GO_BIN" run . update --root "$TMP/install"
test -f "$TMP/install/packs/VERSION"

mkdir -p "$TMP/install/user-packs/custom"
printf 'user-owned\n' > "$TMP/install/user-packs/custom/SKILL.md"
"$GO_BIN" run . uninstall --root "$TMP/install"
test -f "$TMP/install/user-packs/custom/SKILL.md"
test ! -e "$TMP/install/packs"
test ! -e "$TMP/install/pi-plugin"
test ! -e "$TMP/install/extensions"
test ! -e "$TMP/install/themes"
test ! -e "$TMP/install/skills"
test ! -e "$TMP/install/package.json"
test ! -e "$TMP/install/exports"

"$PYTHON_BIN" - <<PY
from pathlib import Path
import re
roadmap = Path("$REPO/docs/roadmap/ROADMAP.md").read_text()
if "annotated tags" not in roadmap or "Linux-only" not in roadmap:
    raise SystemExit("release discipline text missing")
tags = [line.strip() for line in roadmap.splitlines() if "v0.y.z" in line]
if not tags:
    raise SystemExit("pre-stable tag policy missing")
PY

rm -rf "$TMP" "$ROOT/cli"
echo "acceptance ok"
