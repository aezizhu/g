# ? Release v0.3.0 Successfully Created!

## Summary

The first release of LuciCodex has been successfully created and published to GitHub!

## What Was Fixed

### 1. **Build Script Issues**
   - Fixed relative path problem in IPK packaging
   - Added absolute path resolution using `readlink -f`
   - Now correctly creates IPK files in the output directory

### 2. **Workflow Configuration**
   - Added `permissions: contents: write` to allow release creation
   - Added file renaming step to create user-friendly filenames
   - Workflow now creates both simplified AND versioned filenames

### 3. **Release Process**
   - Created and pushed tag `v0.3.0`
   - Workflow automatically built all architectures
   - Release published with all required files

## Release Details

**Release URL:** https://github.com/aezizhu/LuciCodex/releases/tag/v0.3.0

**Available Files:**

### User-Friendly Names (for README instructions):
- ? `lucicodex-mips.ipk` - For MIPS routers (most common)
- ? `lucicodex-arm.ipk` - For ARM routers  
- ? `lucicodex-amd64.ipk` - For x86_64 routers
- ? `lucicodex-arm64.ipk` - For ARM 64-bit routers
- ? `lucicodex-mipsle.ipk` - For MIPS little-endian routers
- ? `luci-app-lucicodex.ipk` - Web interface package

### Versioned Names (for advanced users):
- `lucicodex_0.3.0_mips_24kc.ipk`
- `lucicodex_0.3.0_arm_cortex-a7.ipk`
- `lucicodex_0.3.0_x86_64.ipk`
- `lucicodex_0.3.0_aarch64.ipk`
- `lucicodex_0.3.0_mipsel_24kc.ipk`
- `luci-app-lucicodex_0.3.0_all.ipk`

### Additional Files:
- `SHA256SUMS` - Checksums for verification

## Download Links Verified ?

All download links from the README are working correctly:

```bash
# MIPS (2.1 MB) ?
wget https://github.com/aezizhu/LuciCodex/releases/latest/download/lucicodex-mips.ipk

# ARM (2.1 MB) ?
wget https://github.com/aezizhu/LuciCodex/releases/latest/download/lucicodex-arm.ipk

# AMD64 (2.3 MB) ?
wget https://github.com/aezizhu/LuciCodex/releases/latest/download/lucicodex-amd64.ipk
```

## Changes Made to Repository

### Modified Files:
1. **`.github/workflows/release.yml`**
   - Added `permissions: contents: write`
   - Added file renaming step to create simplified filenames
   - Now uploads from `release/` directory with both simplified and versioned files

2. **`scripts/build-release-assets.sh`**
   - Fixed path resolution in `ipk_pack_lucicodex()` function
   - Fixed path resolution in `ipk_pack_luci()` function
   - Now uses `readlink -f` to get absolute paths before cd'ing into temp directories

### Commits:
- `7eae813` - feat: Add release workflow and instructions
- `2154c04` - fix: Use absolute paths in build script for IPK creation
- `f27f0db` - fix: Add contents write permission to release workflow

### Tag Created:
- `v0.3.0` - Release v0.3.0 - Fix IPK download links with simplified filenames

## Future Releases

To create future releases, simply:

```bash
# 1. Update version in code/scripts
# 2. Commit changes
# 3. Create and push a tag
git tag -a v0.3.1 -m "Release v0.3.1 - Description"
git push origin v0.3.1

# The workflow will automatically:
# - Build IPK files for all architectures
# - Create simplified filenames
# - Upload everything to GitHub Releases
```

## Verification Results

All three main download links were tested and verified:
- ? MIPS IPK downloads correctly (2.1 MB, valid Debian package)
- ? ARM IPK downloads correctly (2.1 MB, valid Debian package)  
- ? AMD64 IPK downloads correctly (2.3 MB, valid Debian package)

File validation confirms all are proper IPK packages:
```
Debian binary package (format 2.0), with control.tar.gz/, data compression gz
```

## Next Steps

Users can now install LuciCodex following the README instructions:

```bash
cd /tmp
wget https://github.com/aezizhu/LuciCodex/releases/latest/download/lucicodex-mips.ipk
opkg update
opkg install /tmp/lucicodex-*.ipk
```

The release is complete and fully functional! ??
