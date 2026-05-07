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
test -f "$TMP/install/skills/core-workflow/SKILL.md"
test -f "$TMP/install/package.json"
if command -v pi >/dev/null 2>&1; then
  PI_CODING_AGENT_DIR="$TMP/pi-agent" pi install "$TMP/install"
  PI_CODING_AGENT_DIR="$TMP/pi-agent" pi list | grep -q "$TMP/install"
fi
"$GO_BIN" run . doctor --root "$TMP/install"
"$GO_BIN" run . export --config "$REPO/pi-plugin/config-schema.toml" --target codex --target-root "$TMP/export-default"

"$PYTHON_BIN" - <<PY
from pathlib import Path
root = Path("$TMP/export-default/.codex/skills")
expected = ["core-workflow", "rtk-routing", "agent-lanes", "sane-router"]
missing = [name for name in expected if not (root / name / "SKILL.md").exists()]
if missing:
    raise SystemExit(f"missing default exports: {missing}")
if (root / "example-user-pack" / "SKILL.md").exists():
    raise SystemExit("disabled user pack exported unexpectedly")
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
expected = ["core-workflow", "rtk-routing", "agent-lanes", "sane-router", "example-user-pack"]
missing = [name for name in expected if not (root / name / "SKILL.md").exists()]
if missing:
    raise SystemExit(f"missing user-pack exports: {missing}")
if not Path("$TMP/home/.codex/skills/core-workflow/SKILL.md").exists():
    raise SystemExit("default Codex-native export did not write under HOME/.codex/skills")
PY

"$NODE_BIN" --test "$REPO/pi-plugin/plugin.test.js"
grep -q 'resources_discover' "$REPO/pi-plugin/index.ts"
grep -q 'skillPaths' "$REPO/pi-plugin/index.ts"

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
"$GO_BIN" run . update --root "$TMP/install"
test -f "$TMP/install/packs/VERSION"

mkdir -p "$TMP/install/user-packs/custom"
printf 'user-owned\n' > "$TMP/install/user-packs/custom/SKILL.md"
"$GO_BIN" run . uninstall --root "$TMP/install"
test -f "$TMP/install/user-packs/custom/SKILL.md"
test ! -e "$TMP/install/packs"
test ! -e "$TMP/install/pi-plugin"
test ! -e "$TMP/install/extensions"
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
