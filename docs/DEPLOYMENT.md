# Deployment Guide

## Simple Release

```bash
./scripts/simple-release.sh 1.2.3
```

**That's it!** The script handles:
- ✅ Updates both package.json files
- ✅ Commits and tags correctly
- ✅ Pushes to trigger CI/CD

## Advanced Release (with validation)

```bash
./scripts/release.sh --version 1.2.3
```

## Published Packages

- **npm**: `@rakeyshgidwani/sunday-schemas@X.Y.Z`
- **Go**: `github.com/rakeyshgidwani/sunday-schemas/codegen/go@vX.Y.Z`

## Verification

```bash
# Check Go module
go list -m github.com/rakeyshgidwani/sunday-schemas/codegen/go@vX.Y.Z

# Check npm package
npm view @rakeyshgidwani/sunday-schemas@X.Y.Z
```