# ?? Deprecation Notice

## Versioned IPK Filenames (Effective v0.3.1)

### What's Changing?

Starting with this release, we are deprecating versioned IPK filenames in favor of simplified, user-friendly names.

### Timeline

- **v0.3.x releases**: Both filename formats available (backward compatible)
- **v0.4.0 and later**: Only simplified names will be available

### Migration Required

If you are downloading IPK files programmatically, please update your scripts:

| ? Deprecated (v0.4.0) | ? Use Instead |
|------------------------|----------------|
| `lucicodex_0.3.0_mips_24kc.ipk` | `lucicodex-mips.ipk` |
| `lucicodex_0.3.0_mipsel_24kc.ipk` | `lucicodex-mipsle.ipk` |
| `lucicodex_0.3.0_arm_cortex-a7.ipk` | `lucicodex-arm.ipk` |
| `lucicodex_0.3.0_aarch64.ipk` | `lucicodex-arm64.ipk` |
| `lucicodex_0.3.0_x86_64.ipk` | `lucicodex-amd64.ipk` |
| `luci-app-lucicodex_0.3.0_all.ipk` | `luci-app-lucicodex.ipk` |

### Download URLs

Update your download URLs to use simplified names:

```bash
# ? Recommended (stable URLs across versions)
wget https://github.com/aezizhu/LuciCodex/releases/latest/download/lucicodex-mips.ipk

# ? Deprecated (will break in v0.4.0)
wget https://github.com/aezizhu/LuciCodex/releases/latest/download/lucicodex_0.3.0_mips_24kc.ipk
```

### Why This Change?

1. **Simpler URLs**: Easier to remember and document
2. **Stable links**: `/latest/download/lucicodex-mips.ipk` works across versions
3. **Less confusion**: Fewer files to choose from (7 instead of 15)
4. **Better UX**: Matches what's documented in README

### How to Check Version?

After installation, you can verify the version:

```bash
lucicodex -version
# Output: LuciCodex version 0.3.0
```

### Questions?

- GitHub Issues: https://github.com/aezizhu/LuciCodex/issues
- See README for complete installation instructions
