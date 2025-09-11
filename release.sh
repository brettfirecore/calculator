#!/usr/bin/env bash
set -euo pipefail

# -------- Settings --------
# Build all commands under ./cmd/* (common Go layout)
CMD_DIRS=()
if [[ -d cmd ]]; then
  # one-level subdirs under cmd/ that contain Go files
  while IFS= read -r -d '' d; do CMD_DIRS+=("$d"); done < <(find cmd -mindepth 1 -maxdepth 1 -type d -print0)
fi
# Fallback: if there is a main package at repo root, also build it
if [[ ${#CMD_DIRS[@]} -eq 0 && -n "$(go list -f '{{.Name}}' . 2>/dev/null || true)" ]]; then
  if [[ "$(go list -f '{{.Name}}' .)" == "main" ]]; then
    CMD_DIRS+=(".")
  fi
fi
if [[ ${#CMD_DIRS[@]} -eq 0 ]]; then
  echo "No commands found (looked in ./cmd/* and current dir)."
  exit 1
fi

VERSION="${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo v0.0.0-dev)}"
DATE="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
# These target the 'main' package vars in each command
LDFLAGS="-s -w -X 'main.version=${VERSION}' -X 'main.buildDate=${DATE}'"
export CGO_ENABLED=0

TARGETS=(
  "linux/amd64"
  "linux/arm64"
  "darwin/arm64"
  "windows/amd64"
  "windows/arm64"
)

rm -rf dist
mkdir -p dist

for cmd in "${CMD_DIRS[@]}"; do
  APP_NAME="$(basename "$cmd")"
  # If cmd is ".", name the app after the repo folder
  [[ "$cmd" == "." ]] && APP_NAME="$(basename "$PWD")"

  for target in "${TARGETS[@]}"; do
    IFS=/ read -r GOOS GOARCH <<<"${target}"
    OUT="dist/${APP_NAME}-${VERSION}-${GOOS}-${GOARCH}"
    [[ "${GOOS}" == "windows" ]] && OUT="${OUT}.exe"

    echo ">> Building ${OUT} from ./${cmd}"
    GOOS="${GOOS}" GOARCH="${GOARCH}" \
      go build -trimpath -ldflags="${LDFLAGS}" -o "${OUT}" "./${cmd}"

    [[ "${GOOS}" != "windows" ]] && chmod +x "${OUT}"
  done
done

echo ">> Generating checksums"
( cd dist && sha256sum * > SHA256SUMS.txt )

echo "✅ Done. See ./dist"
