#!/bin/sh
set -eu

# Author: AZ <Aezi.zhu@icloud.com>

# Usage: ARCH=mipsle ./scripts/build-openwrt.sh
ARCH="${ARCH:-mipsle}"
OUT="${OUT:-dist}"

mkdir -p "$OUT"
BIN_OUT="$OUT/lucicodex-linux-${ARCH}"
GOOS=linux GOARCH="$ARCH" go build -trimpath -ldflags "-s -w" -o "$BIN_OUT" ./cmd/lucicodex
ln -sf "$(basename "$BIN_OUT")" "$OUT/g-linux-${ARCH}"
echo "Built $BIN_OUT (and legacy symlink g-linux-${ARCH})"


