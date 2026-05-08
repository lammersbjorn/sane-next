#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO="$(cd "$ROOT/.." && pwd)"
DIST="${DIST:-$REPO/dist}"
VERSION="${VERSION:-}"

if [[ -z "$VERSION" ]]; then
  VERSION="$(cd "$ROOT" && go run . version)"
fi

rm -rf "$DIST"
mkdir -p "$DIST"

package_target() {
  local goos="$1"
  local goarch="$2"
  local exe="sane-next"
  local archive_dir="$DIST/sane-next_${VERSION}_${goos}_${goarch}"
  local binary="$archive_dir/$exe"

  if [[ "$goos" == "windows" ]]; then
    exe="sane-next.exe"
    binary="$archive_dir/$exe"
  fi

  mkdir -p "$archive_dir"
  (cd "$ROOT" && GOOS="$goos" GOARCH="$goarch" go build -o "$binary" .)
  cp "$REPO/README.md" "$archive_dir/README.md"
  if [[ -f "$REPO/LICENSE" ]]; then
    cp "$REPO/LICENSE" "$archive_dir/LICENSE"
  fi

  if [[ "$goos" == "windows" ]]; then
    (cd "$DIST" && zip -qr "sane-next_${VERSION}_${goos}_${goarch}.zip" "$(basename "$archive_dir")")
  else
    (cd "$DIST" && tar -czf "sane-next_${VERSION}_${goos}_${goarch}.tar.gz" "$(basename "$archive_dir")")
  fi
}

package_target linux amd64
package_target linux arm64
package_target darwin amd64
package_target darwin arm64
package_target windows amd64
package_target windows arm64

echo "release archives written to $DIST"
