#!/bin/sh
set -eu

# Author: aezizhu

# Usage: ARCH=mipsle ./scripts/build-openwrt.sh
ARCH="${ARCH:-mipsle}"
OUT="${OUT:-dist}"

mkdir -p "$OUT"
GOOS=linux GOARCH="$ARCH" go build -trimpath -ldflags "-s -w" -o "$OUT/g-${ARCH}" ./cmd/g
echo "Built $OUT/g-${ARCH}"


