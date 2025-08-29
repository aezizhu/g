#!/bin/bash
set -euo pipefail

# Build multi-arch binaries and package .ipk for g and luci-app-g
# Requires: go, tar, ar (binutils)

VERSION=${VERSION:-"0.1.0"}
OUT=${OUT:-"dist"}
ARCHES=(amd64 arm64 arm mipsle mips)
GOARM_DEFAULT=7

mkdir -p "$OUT"

build_bin() {
  local arch="$1"
  local outbin
  if [[ "$arch" == "arm" ]]; then
    GOOS=linux GOARCH=arm GOARM=${GOARM:-$GOARM_DEFAULT} \
      go build -trimpath -ldflags "-s -w" -o "$OUT/g-linux-${arch}v${GOARM:-$GOARM_DEFAULT}" ./cmd/g
    outbin="$OUT/g-linux-${arch}v${GOARM:-$GOARM_DEFAULT}"
  else
    GOOS=linux GOARCH="$arch" \
      go build -trimpath -ldflags "-s -w" -o "$OUT/g-linux-${arch}" ./cmd/g
    outbin="$OUT/g-linux-${arch}"
  fi
  echo "$outbin"
}

ipk_pack_g() {
  local arch="$1"; shift
  local binpath="$1"; shift
  local arch_ipk="$arch"
  case "$arch" in
    amd64) arch_ipk="x86_64";;
    arm64) arch_ipk="aarch64";;
    arm) arch_ipk="arm_cortex-a7";;
    mipsle) arch_ipk="mipsel_24kc";;
    mips) arch_ipk="mips_24kc";;
  esac
  local work
  work=$(mktemp -d)
  mkdir -p "$work/control" "$work/data/usr/bin"
  install -m0755 "$binpath" "$work/data/usr/bin/g"
  cat > "$work/control/control" <<EOF
Package: g
Version: $VERSION
Architecture: $arch_ipk
Maintainer: aezizhu
Section: utils
Priority: optional
Depends: libc
Description: Natural-language CLI for OpenWrt
EOF
  (cd "$work"; echo 2.0 > debian-binary; tar -czf control.tar.gz -C control .; tar -czf data.tar.gz -C data .; ar -r "$OUT/g_${VERSION}_${arch_ipk}.ipk" debian-binary control.tar.gz data.tar.gz >/dev/null)
  rm -rf "$work"
}

ipk_pack_luci() {
  local work
  work=$(mktemp -d)
  mkdir -p "$work/control" "$work/data/usr/lib/lua/luci/controller" "$work/data/usr/lib/lua/luci/view/g"
  install -m0644 package/luci-app-g/luasrc/controller/g.lua "$work/data/usr/lib/lua/luci/controller/g.lua"
  install -m0644 package/luci-app-g/luasrc/view/g/overview.htm "$work/data/usr/lib/lua/luci/view/g/overview.htm"
  cat > "$work/control/control" <<EOF
Package: luci-app-g
Version: $VERSION
Architecture: all
Maintainer: aezizhu
Section: luci
Priority: optional
Depends: luci-base, g
Description: LuCI web UI for g
EOF
  (cd "$work"; echo 2.0 > debian-binary; tar -czf control.tar.gz -C control .; tar -czf data.tar.gz -C data .; ar -r "$OUT/luci-app-g_${VERSION}_all.ipk" debian-binary control.tar.gz data.tar.gz >/dev/null)
  rm -rf "$work"
}

sha256sum_all() {
  (cd "$OUT" && sha256sum * > SHA256SUMS)
}

main() {
  for arch in "${ARCHES[@]}"; do
    bin=$(build_bin "$arch")
    ipk_pack_g "$arch" "$bin"
  done
  ipk_pack_luci
  sha256sum_all
}

main "$@"


