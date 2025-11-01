# How to Create Your First Release

## The Problem

Your README.md references downloadable IPK files from GitHub Releases, but no releases exist yet because:
1. The release workflow only triggers when you push a git tag
2. No tags have been created yet

## Solution: Create a Release

Follow these steps to create your first release with the IPK files:

### Step 1: Ensure All Changes Are Committed

```bash
git status
# Make sure working tree is clean
```

### Step 2: Create and Push a Version Tag

```bash
# Create a tag for version 0.3.0 (or whatever your current version is)
git tag -a v0.3.0 -m "Release v0.3.0 - Initial public release"

# Push the tag to GitHub (this will trigger the release workflow)
git push origin v0.3.0
```

### Step 3: Monitor the Release Workflow

1. Go to your GitHub repository: https://github.com/aezizhu/LuciCodex
2. Click on the "Actions" tab
3. You should see a "release" workflow running
4. Wait for it to complete (usually 2-5 minutes)

### Step 4: Verify the Release

Once the workflow completes:

1. Go to: https://github.com/aezizhu/LuciCodex/releases
2. You should see a new release "v0.3.0" with the following files:
   - `lucicodex-mips.ipk` (simplified name for README instructions)
   - `lucicodex-arm.ipk` (simplified name for README instructions)
   - `lucicodex-amd64.ipk` (simplified name for README instructions)
   - `lucicodex-arm64.ipk` (bonus: ARM 64-bit)
   - `lucicodex-mipsle.ipk` (bonus: MIPS little-endian)
   - `luci-app-lucicodex.ipk` (web interface)
   - Plus versioned files and SHA256SUMS

### Step 5: Test the Download Links

```bash
# Test that the download links from your README work
wget https://github.com/aezizhu/LuciCodex/releases/latest/download/lucicodex-mips.ipk
# Should download successfully
```

## What I Fixed

I updated the `.github/workflows/release.yml` file to:
1. Build the IPK files using your build script
2. Create user-friendly filenames (like `lucicodex-mips.ipk`) that match your README
3. Keep the original versioned files too (like `lucicodex_0.3.0_mips_24kc.ipk`)
4. Upload all files to the GitHub release

## Future Releases

For future releases, just repeat steps 2-5 with a new version number:

```bash
# Update version in code/scripts first, then:
git tag -a v0.3.1 -m "Release v0.3.1 - Bug fixes"
git push origin v0.3.1
```

The workflow will automatically build and publish the new release.

## Manual Release (Alternative)

If you prefer to test locally first:

```bash
# Build the IPK files locally
VERSION=0.3.0 OUT=dist ./scripts/build-release-assets.sh

# Check what was built
ls -lh dist/

# Create a release manually using GitHub CLI
gh release create v0.3.0 \
  --title "Release v0.3.0" \
  --notes "Initial public release with IPK packages" \
  dist/*.ipk
```

## Troubleshooting

**Q: The workflow failed!**
- Check the Actions tab for error logs
- Make sure you have proper permissions set in repository settings
- Ensure `GITHUB_TOKEN` has write permissions (Settings ? Actions ? General ? Workflow permissions)

**Q: Files have wrong names in the release**
- The workflow now creates both simplified names AND versioned names
- Users can use either `lucicodex-mips.ipk` or `lucicodex_0.3.0_mips_24kc.ipk`

**Q: How do I delete a bad release?**
```bash
gh release delete v0.3.0
git tag -d v0.3.0
git push origin :refs/tags/v0.3.0
```
Then recreate it with the steps above.
