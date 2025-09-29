# Release Automation Guide

This guide explains how to use the automated release and deployment scripts for Sunday Schemas.

## Quick Start

### 1. Standard Release (Recommended)
```bash
# Complete automated release process
npm run release -- --version 1.2.0

# Or with the script directly
./scripts/release.sh --version 1.2.0
```

### 2. Preview Release (Dry Run)
```bash
# See what would happen without making changes
npm run release:dry-run -- --version 1.2.0
```

### 3. Fast Release (Skip Tests)
```bash
# Skip validation if you've already run tests
./scripts/release.sh --version 1.2.0 --skip-tests
```

## Available Scripts

### Core Release Script: `./scripts/release.sh`

**Features:**
- ✅ Complete end-to-end release automation
- ✅ Pre-release validation (schemas, compatibility, changelog)
- ✅ Automated version bumping and git tagging
- ✅ GitHub Actions CI/CD triggering
- ✅ Package publication monitoring
- ✅ Post-release validation and health checks
- ✅ Creates post-release checklist

**Usage:**
```bash
./scripts/release.sh [OPTIONS]

Options:
  -v, --version VERSION    Version to release (e.g., 1.2.0)
  --skip-tests            Skip validation and tests
  --dry-run               Preview without making changes
  --auto-deploy           Auto-deploy to consuming services (future)
  --no-monitoring         Skip post-release monitoring
  -h, --help              Show help
```

**Examples:**
```bash
# Standard release
./scripts/release.sh --version 1.2.0

# Preview release
./scripts/release.sh --version 1.2.0 --dry-run

# Quick release (skip tests)
./scripts/release.sh --version 1.2.0 --skip-tests

# Release candidate
./scripts/release.sh --version 1.2.0-rc.1
```

### Consuming Services Deployment: `./scripts/deploy-consuming-services.sh`

**Features:**
- ✅ Updates all consuming services with new schema version
- ✅ Supports both NPM and Go module updates
- ✅ Sequential or parallel deployment modes
- ✅ Built-in testing after updates
- ✅ Rollback-friendly (stops on first failure)

**Usage:**
```bash
./scripts/deploy-consuming-services.sh [OPTIONS]

Options:
  -v, --version VERSION     Schema version to deploy
  -r, --services-root PATH  Root path for services (default: ../)
  --dry-run                 Preview without making changes
  --skip-tests              Skip running tests after updates
  --parallel                Update services in parallel
  -h, --help                Show help
```

**Examples:**
```bash
# Update all services sequentially
./scripts/deploy-consuming-services.sh --version 1.2.0

# Preview updates
./scripts/deploy-consuming-services.sh --version 1.2.0 --dry-run

# Fast parallel updates
./scripts/deploy-consuming-services.sh --version 1.2.0 --parallel
```

## NPM Script Shortcuts

For convenience, these are available as npm scripts:

```bash
# Release scripts
npm run release -- --version 1.2.0
npm run release:dry-run -- --version 1.2.0
npm run release:auto -- --version 1.2.0

# Service deployment scripts
npm run deploy:services -- --version 1.2.0
npm run deploy:services:dry-run -- --version 1.2.0

# Legacy deployment (still available)
npm run deploy -- --version 1.2.0
npm run deploy:dry-run -- --version 1.2.0
```

## Automation Levels

### Level 1: Basic Automation (Current)
```bash
# One command releases packages
npm run release -- --version 1.2.0

# Manual service updates
npm run deploy:services -- --version 1.2.0
```

### Level 2: Full Automation (Future)
```bash
# One command does everything
npm run release:auto -- --version 1.2.0
# ^ This would release + deploy + monitor + notify
```

## What Gets Automated

### ✅ Fully Automated
- Schema validation and compatibility checks
- Code generation (TypeScript & Go)
- Version bumping and changelog updates
- Git tagging and pushing
- GitHub Actions CI/CD triggering
- Package publication to npm and Go modules
- Release asset creation
- Post-release validation

### 🔄 Semi-Automated
- Consuming service updates (script available, manual trigger)
- Post-deployment health checks (script available)
- Team notifications (manual)

### ❌ Still Manual
- CHANGELOG.md content writing (automation verifies format)
- Deployment to production environments
- End-to-end testing coordination
- Rollback decisions

## Pre-Release Checklist

Before running `npm run release`, ensure:

1. **CHANGELOG.md is updated**
   ```bash
   # Add an entry for your version:
   ## [1.2.0] - TBD
   ### Added
   - New feature X
   ### Fixed
   - Bug Y
   ```

2. **All changes are committed**
   ```bash
   git status  # Should show clean working directory
   ```

3. **You're on main branch**
   ```bash
   git branch --show-current  # Should show 'main'
   ```

4. **Tests pass locally**
   ```bash
   npm test  # Or use --skip-tests flag
   ```

## Post-Release Workflow

The release script generates a checklist at `/tmp/sunday-schemas-release-{version}.md`. Follow these steps:

### Immediate (0-30 minutes)
- ✅ Verify packages are published
- ✅ Check GitHub release page
- ✅ Monitor GitHub Actions workflow

### Service Updates (30-60 minutes)
```bash
# Use the automation script
npm run deploy:services -- --version 1.2.0

# Or manually update each service:
# cd ../sunday-data/normalizer && npm update @rakeyshgidwani/sunday-schemas
# cd ../sunday-data/insights && go get -u github.com/rakeyshgidwani/sunday-schemas/go@v1.2.0
# cd ../sunday-api/ui-bff && npm update @rakeyshgidwani/sunday-schemas
# cd ../sunday-frontend/web && npm update @rakeyshgidwani/sunday-schemas
```

### Validation (60-90 minutes)
- ✅ Run end-to-end tests
- ✅ Monitor application logs
- ✅ Verify new features work

### Communication
- ✅ Announce in #schema-updates
- ✅ Update team documentation
- ✅ Email stakeholders (major releases)

## Troubleshooting

### Release Script Issues

**Problem: "CHANGELOG.md needs to be updated"**
```bash
# Add version entry to CHANGELOG.md:
## [1.2.0] - TBD
### Added
- Your changes here
```

**Problem: "Working directory has uncommitted changes"**
```bash
git add . && git commit -m "Pre-release cleanup"
# Or use git stash if changes aren't ready
```

**Problem: "Version already exists"**
```bash
git tag -d v1.2.0  # Delete local tag
git push origin :refs/tags/v1.2.0  # Delete remote tag
```

### Package Publication Issues

**Problem: NPM package not found**
```bash
# Check if package exists
npm view @rakeyshgidwani/sunday-schemas@1.2.0

# Clear npm cache and retry
npm cache clean --force
```

**Problem: Go module not available**
```bash
# Check module availability
go list -m github.com/rakeyshgidwani/sunday-schemas/go@v1.2.0

# Force Go proxy refresh
GOPROXY=direct go get github.com/rakeyshgidwani/sunday-schemas/go@v1.2.0
```

### Service Update Issues

**Problem: Service update fails**
```bash
# Check service directory exists
ls -la ../sunday-data/normalizer

# Update services root path
npm run deploy:services -- --version 1.2.0 --services-root /path/to/sunday/services/
```

**Problem: Tests fail after update**
```bash
# Skip tests and update manually
npm run deploy:services -- --version 1.2.0 --skip-tests

# Then run tests manually in each service
cd ../sunday-data/normalizer && npm test
```

## Monitoring and Rollback

### Monitor Release Health
```bash
# Check GitHub Actions
open https://github.com/rakeyshgidwani/sunday-schemas/actions

# Monitor package downloads
npm view @rakeyshgidwani/sunday-schemas --json

# Check Go module usage
curl -s https://api.github.com/repos/rakeyshgidwani/sunday-schemas/releases/latest
```

### Emergency Rollback
```bash
# Unpublish npm package (within 24 hours)
npm unpublish @rakeyshgidwani/sunday-schemas@1.2.0 --force

# Retract Go module version
echo 'retract v1.2.0 // Emergency rollback' >> go/go.mod
git add go/go.mod && git commit -m "Retract v1.2.0" && git push

# Rollback consuming services
kubectl rollout undo deployment/normalizer
kubectl rollout undo deployment/insights
kubectl rollout undo deployment/ui-bff
```

## Future Enhancements

The automation will continue to evolve:

- **Auto-deployment**: Full end-to-end automation with `--auto-deploy`
- **Slack integration**: Automatic team notifications
- **Staging deployment**: Automatic deployment to staging environments
- **Health monitoring**: Automated post-deployment health checks
- **Smart rollbacks**: Automatic rollback on failure detection

This automation significantly reduces the manual effort and human error in releases while maintaining safety and visibility into the process.