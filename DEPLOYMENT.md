# Deployment Guide

This guide explains how to deploy new versions of sunday-schemas to both NPM and Go module registries.

## 🚀 Quick Deploy

Deploy a new version with one command:

```bash
# Interactive deployment (will prompt for version)
npm run deploy

# Deploy specific version
./scripts/deploy.sh --version 1.0.3

# Dry run to see what would happen
npm run deploy:dry-run --version 1.0.3
```

## 📋 Deploy Script Features

### ✅ Automated Steps
1. **Validation & Tests** - Runs `npm run ci-validate`
2. **Version Bump** - Updates `package.json` version
3. **Code Generation** - Runs `npm run build`
4. **Git Commit** - Commits version changes
5. **Git Tags** - Creates both NPM (`v1.0.3`) and Go (`codegen/go/v1.0.3`) tags
6. **NPM Publish** - Publishes to npmjs.org
7. **Go Module** - Makes module available via Git tags
8. **Verification** - Checks deployment success

### 🎛️ Options
- `--version 1.0.3` - Specify version to deploy
- `--skip-tests` - Skip validation and tests
- `--dry-run` - Show what would be done without executing
- `--help` - Show usage help

## 🛠️ Usage Examples

### Standard Deployment
```bash
# Deploy version 1.0.3 with full validation
./scripts/deploy.sh --version 1.0.3
```

### Quick Development Deploy
```bash
# Skip tests for faster deployment
npm run deploy:skip-tests --version 1.0.3-beta.1
```

### Preview Changes
```bash
# See exactly what would be deployed
npm run deploy:dry-run --version 1.0.3
```

### Interactive Mode
```bash
# Script will prompt for version
npm run deploy
```

## 📦 What Gets Published

### NPM Package (`sunday-schemas`)
- **Registry**: https://npmjs.org/package/sunday-schemas
- **Tag**: `v1.0.3`
- **Usage**: `npm install sunday-schemas@1.0.3`

### Go Module (`github.com/rakeyshgidwani/sunday-schemas/codegen/go`)
- **Registry**: Go module proxy
- **Tag**: `codegen/go/v1.0.3`
- **Usage**: `go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.3`

## 🔍 Deployment Validation

### Automatic Validation
The deploy script automatically runs validation after deployment to ensure both packages are accessible.

### Manual Validation
You can also validate deployments manually:

```bash
# Validate specific version
./scripts/validate-deployment.sh --version 1.0.3

# Quick validation (skip slow checks)
npm run validate-deployment:quick --version 1.0.3

# Verbose output for debugging
./scripts/validate-deployment.sh --version 1.0.3 --verbose
```

### What Gets Validated
✅ **NPM Package Checks:**
- Package exists on npmjs.org
- Package is downloadable
- Package imports correctly
- TypeScript definitions included

✅ **Go Module Checks:**
- Git tags exist (both `v1.0.3` and `codegen/go/v1.0.3`)
- Module imports and compiles
- Schema constants accessible
- Validation functions work

### Sample Validation Output

After deployment, you'll see:

```
======================================
🚀 DEPLOYMENT COMPLETE
======================================

📦 NPM Package:    sunday-schemas@1.0.3
🐹 Go Module:      github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.3
🏷️  Git Tags:       v1.0.3, codegen/go/v1.0.3

Usage instructions:

📦 TypeScript/JavaScript:
   npm install sunday-schemas@1.0.3

🐹 Go:
   go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.3

✅ All packages successfully deployed!
======================================
```

## 🚨 Rollback

If deployment fails or issues are found:

### NPM Rollback
```bash
# Deprecate problematic version
npm deprecate sunday-schemas@1.0.3 "Version deprecated due to issues"

# Users can downgrade
npm install sunday-schemas@1.0.2
```

### Go Module Rollback
```bash
# Delete problematic tags
git tag -d v1.0.3
git tag -d codegen/go/v1.0.3
git push origin :refs/tags/v1.0.3
git push origin :refs/tags/codegen/go/v1.0.3

# Users can downgrade
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.2
```

## 🔧 Manual Deployment

If the script fails, you can deploy manually:

### NPM Only
```bash
npm version 1.0.3 --no-git-tag-version
git add package.json
git commit -m "Release version 1.0.3"
git tag v1.0.3
git push origin main
git push origin v1.0.3
npm publish
```

### Go Module Only
```bash
git tag codegen/go/v1.0.3
git push origin codegen/go/v1.0.3
```

## 📋 Pre-Deployment Checklist

Before deploying, ensure:

- [ ] All tests pass (`npm run ci-validate`)
- [ ] CHANGELOG.md is updated
- [ ] Version follows semantic versioning
- [ ] No uncommitted changes (or use `--force` flag)
- [ ] NPM authentication is set up
- [ ] Git push permissions are available

## 🎯 Best Practices

1. **Use Semantic Versioning**: `1.0.3` for patches, `1.1.0` for features, `2.0.0` for breaking changes
2. **Test First**: Always run `npm run ci-validate` before deploying
3. **Update Changelog**: Document what changed in each version
4. **Dry Run**: Use `--dry-run` to preview deployments
5. **Monitor**: Check that both packages are accessible after deployment

---

**Next**: After deployment, update dependent projects (like sunday-connectors) to use the new version.