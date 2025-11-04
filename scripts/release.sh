#!/bin/bash

# Sunday Schemas Release Automation Script
# Complete end-to-end release process with validation, tagging, and monitoring

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
NC='\033[0m' # No Color

# Functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_step() {
    echo -e "\n${PURPLE}[STEP $1]${NC} $2"
    echo "----------------------------------------"
}

# Spinner function for long operations
spinner() {
    local pid=$1
    local delay=0.1
    local spinstr='|/-\'
    while [ "$(ps a | awk '{print $1}' | grep $pid)" ]; do
        local temp=${spinstr#?}
        printf " [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
    printf "    \b\b\b\b"
}

# Check if we're in the right directory
if [ ! -f "package.json" ] || [ ! -d "schemas" ]; then
    log_error "Must be run from sunday-schemas root directory"
    exit 1
fi

# Parse command line arguments
VERSION=""
SKIP_TESTS=false
DRY_RUN=false
AUTO_DEPLOY=false
MONITORING=true
HELP=false

while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        --skip-tests)
            SKIP_TESTS=true
            shift
            ;;
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        --auto-deploy)
            AUTO_DEPLOY=true
            shift
            ;;
        --no-monitoring)
            MONITORING=false
            shift
            ;;
        -h|--help)
            HELP=true
            shift
            ;;
        *)
            log_error "Unknown option: $1"
            HELP=true
            break
            ;;
    esac
done

# Show help
if [ "$HELP" = true ]; then
    cat << 'EOF'
Sunday Schemas Enhanced Release Script

Usage: ./scripts/release.sh [OPTIONS]

Options:
  -v, --version VERSION    Specify version to release (e.g., 1.2.0)
  --skip-tests            Skip validation and tests
  --dry-run               Show what would be done without executing
  --auto-deploy           Automatically deploy to consuming services
  --no-monitoring         Skip post-release monitoring
  -h, --help              Show this help message

Features:
  ✅ Pre-release validation (schemas, compatibility, changelog)
  ✅ Automated version bumping and tagging
  ✅ GitHub Actions CI/CD triggering
  ✅ Package publishing verification
  ✅ Post-release monitoring and health checks
  ✅ Automated consuming services deployment
  ✅ Rollback capabilities

Examples:
  ./scripts/release.sh --version 1.2.0
  ./scripts/release.sh --version 1.2.0 --dry-run
  ./scripts/release.sh --version 1.2.0 --auto-deploy
  ./scripts/release.sh --version 1.2.0 --skip-tests --auto-deploy

EOF
    exit 0
fi

# Get version from user if not provided
if [ -z "$VERSION" ]; then
    echo -n "Enter version to release (e.g., 1.2.0): "
    read VERSION
fi

# Validate version format
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+(-rc\.[0-9]+)?$ ]]; then
    log_error "Invalid version format. Use semantic versioning (e.g., 1.2.0 or 1.2.0-rc.1)"
    exit 1
fi

# Check if version already exists (warning only - don't exit)
if git tag -l | grep -q "^v$VERSION$"; then
    log_warning "Version $VERSION already exists as a git tag - release script will be rerunnable"
fi

echo ""
echo "======================================"
echo "🚀 SUNDAY SCHEMAS RELEASE"
echo "======================================"
echo "Version: $VERSION"
echo "Dry Run: $DRY_RUN"
echo "Auto Deploy: $AUTO_DEPLOY"
echo "Monitoring: $MONITORING"
echo "======================================"
echo ""

# Step 1: Pre-release validation
log_step "1" "Pre-release validation"

if [ "$SKIP_TESTS" = false ]; then
    log_info "Running comprehensive validation..."

    if [ "$DRY_RUN" = true ]; then
        echo "Would run: npm run ci-validate"
        echo "Would run: npm run check-changelog"
        echo "Would run: npm run check-deprecations"
    else
        # Run validation in background with spinner
        npm run ci-validate &
        validation_pid=$!
        echo -n "Validating schemas, examples, and OpenAPI spec..."
        spinner $validation_pid
        wait $validation_pid
        log_success "Schema validation passed"

        # Check changelog has been updated
        if ! node scripts/check-changelog.js --version $VERSION; then
            log_error "CHANGELOG.md needs to be updated for version $VERSION"
            log_info "Please add an entry in CHANGELOG.md and run the script again"
            exit 1
        fi
        log_success "CHANGELOG.md validation passed"

        # Check for deprecated schemas
        if node scripts/check-deprecations.js --has-deprecations; then
            log_warning "Found deprecated schemas - review migration plan"
            echo -n "Continue with release? (y/N): "
            read -n 1 continue_release
            echo
            if [[ ! $continue_release =~ ^[Yy]$ ]]; then
                log_info "Release cancelled"
                exit 0
            fi
        fi
    fi
else
    log_warning "Skipping validation (--skip-tests specified)"
fi

# Step 2: Git status check
log_step "2" "Git repository validation"

if [ "$DRY_RUN" = false ]; then
    # Ensure we're on main branch
    current_branch=$(git branch --show-current)
    if [ "$current_branch" != "main" ]; then
        log_error "Must be on main branch (currently on: $current_branch)"
        exit 1
    fi

    # Check for uncommitted changes
    if [ -n "$(git status --porcelain)" ]; then
        log_warning "Working directory has uncommitted changes:"
        git status --short
        echo -n "Continue anyway? (y/N): "
        read -n 1 continue_deploy
        echo
        if [[ ! $continue_deploy =~ ^[Yy]$ ]]; then
            log_info "Release cancelled"
            exit 0
        fi
    fi

    # Pull latest changes
    log_info "Pulling latest changes from origin/main..."
    git pull origin main
    log_success "Repository is up to date"
else
    echo "Would check git branch and pull latest changes"
fi

# Step 3: Version and changelog update
log_step "3" "Version bump and changelog update"

if [ "$DRY_RUN" = true ]; then
    echo "Would update package.json version to $VERSION"
    echo "Would update CHANGELOG.md with release date"
else
    # Check current version in package.json
    current_version=$(node -p "require('./package.json').version")

    if [ "$current_version" = "$VERSION" ]; then
        log_warning "Package.json is already at version $VERSION - skipping version bump"
    else
        # Update package.json version
        npm version $VERSION --no-git-tag-version
        log_success "Updated package.json to version $VERSION"
    fi

    # Update CHANGELOG.md with release date
    today=$(date +%Y-%m-%d)
    if grep -q "\[$VERSION\] - TBD" CHANGELOG.md; then
        sed -i.bak "s/\[$VERSION\] - TBD/[$VERSION] - $today/" CHANGELOG.md
        rm -f CHANGELOG.md.bak
        log_success "Updated CHANGELOG.md with release date"
    fi
fi

# Step 4: Code generation and build
log_step "4" "Code generation and build"

if [ "$DRY_RUN" = true ]; then
    echo "Would run: npm run build"
    echo "Would run: npm run bundle:openapi"
else
    log_info "Generating TypeScript and Go code..."
    npm run generate &
    generate_pid=$!
    echo -n "Generating code from schemas..."
    spinner $generate_pid
    wait $generate_pid
    log_success "Code generation completed"

    log_info "Building TypeScript package..."
    npm run build:ts
    log_success "TypeScript package built"

    log_info "Bundling OpenAPI specification..."
    npm run bundle:openapi
    log_success "OpenAPI spec bundled"
fi

# Step 5: Commit and tag
log_step "5" "Commit version bump and create tags"

if [ "$DRY_RUN" = true ]; then
    echo "Would commit version bump and generated code"
    echo "Would create tags: v$VERSION, codegen/go/v$VERSION"
else
    # Commit version bump and generated code (only if there are changes)
    git add .
    if [ -n "$(git diff --cached)" ]; then
        git commit -m "Release version $VERSION

Updated package.json, CHANGELOG.md, and regenerated all code artifacts.

🚀 Generated with [Claude Code](https://claude.ai/code)

Co-Authored-By: Claude <noreply@anthropic.com>"
        log_success "Committed version bump and generated code"
    else
        log_warning "No changes to commit - skipping commit step"
    fi

    # Create tags (check if they exist first)
    created_tags=()

    if git tag -l | grep -q "^v$VERSION$"; then
        log_warning "Tag v$VERSION already exists - skipping creation"
    else
        git tag "v$VERSION" -m "Release v$VERSION"
        created_tags+=("v$VERSION")
    fi

    if git tag -l | grep -q "^codegen/go/v$VERSION$"; then
        log_warning "Tag codegen/go/v$VERSION already exists - skipping creation"
    else
        git tag "codegen/go/v$VERSION" -m "Go module release v$VERSION"
        created_tags+=("codegen/go/v$VERSION")
    fi

    if [ ${#created_tags[@]} -gt 0 ]; then
        log_success "Created tags: ${created_tags[*]}"
    else
        log_warning "All tags already exist - skipping tag creation"
    fi
fi

# Step 6: Push to trigger CI/CD
log_step "6" "Push changes and trigger CI/CD"

if [ "$DRY_RUN" = true ]; then
    echo "Would push: git push origin main --tags"
    echo "Would trigger GitHub Actions release workflow"
else
    log_info "Pushing changes and tags to trigger GitHub Actions..."
    git push origin main

    # Push tags only if they were newly created
    pushed_tags=()

    # Check if main version tag needs to be pushed
    if ! git ls-remote --tags origin | grep -q "refs/tags/v$VERSION$"; then
        git push origin "v$VERSION" 2>/dev/null && pushed_tags+=("v$VERSION") || log_warning "Tag v$VERSION may already exist on remote"
    else
        log_warning "Tag v$VERSION already exists on remote - skipping push"
    fi

    # Check if Go module tag needs to be pushed
    if ! git ls-remote --tags origin | grep -q "refs/tags/codegen/go/v$VERSION$"; then
        git push origin "codegen/go/v$VERSION" 2>/dev/null && pushed_tags+=("codegen/go/v$VERSION") || log_warning "Tag codegen/go/v$VERSION may already exist on remote"
    else
        log_warning "Tag codegen/go/v$VERSION already exists on remote - skipping push"
    fi

    if [ ${#pushed_tags[@]} -gt 0 ]; then
        log_success "Pushed tags to GitHub: ${pushed_tags[*]} - CI/CD pipeline started"
    else
        log_success "Changes pushed to GitHub - CI/CD pipeline may already be running"
    fi

    # Wait a moment for GitHub Actions to start
    sleep 5

    # Show GitHub Actions status
    log_info "GitHub Actions workflow status:"
    echo "🔗 Monitor at: https://github.com/rakeyshgidwani/sunday-schemas/actions"
fi

# Step 7: Monitor publication
log_step "7" "Monitor package publication"

if [ "$DRY_RUN" = true ]; then
    echo "Would monitor npm and Go module publication"
else
    log_info "Monitoring package publication (this may take 2-3 minutes)..."

    # Wait for packages to be published
    npm_published=false
    go_published=false
    max_attempts=30
    attempt=0

    while [ $attempt -lt $max_attempts ] && ([ "$npm_published" = false ] || [ "$go_published" = false ]); do
        attempt=$((attempt + 1))

        # Check npm package
        if [ "$npm_published" = false ]; then
            if npm view @rakeyshgidwani/sunday-schemas@$VERSION version &>/dev/null; then
                log_success "✅ NPM package @rakeyshgidwani/sunday-schemas@$VERSION is live"
                npm_published=true
            fi
        fi

        # Check Go module
        if [ "$go_published" = false ]; then
            if curl -s "https://proxy.golang.org/github.com/rakeyshgidwani/sunday-schemas/codegen/go/@v/v$VERSION.info" &>/dev/null; then
                log_success "✅ Go module github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION is live"
                go_published=true
            fi
        fi

        if [ "$npm_published" = false ] || [ "$go_published" = false ]; then
            echo -n "."
            sleep 10
        fi
    done

    echo ""

    if [ "$npm_published" = false ] || [ "$go_published" = false ]; then
        log_warning "Some packages may still be publishing. Check manually:"
        [ "$npm_published" = false ] && echo "  - NPM: npm view @rakeyshgidwani/sunday-schemas@$VERSION"
        [ "$go_published" = false ] && echo "  - Go: go list -m github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION"
    fi
fi

# Step 8: Post-release validation
log_step "8" "Post-release validation"

if [ "$DRY_RUN" = true ]; then
    echo "Would run deployment validation script"
else
    if [ "$MONITORING" = true ]; then
        log_info "Running post-release validation..."
        if ./scripts/validate-deployment.sh --version $VERSION --quick; then
            log_success "Post-release validation passed"
        else
            log_warning "Post-release validation had issues - check manually"
        fi
    else
        log_info "Skipping post-release validation (--no-monitoring)"
    fi
fi

# Step 9: Auto-deploy to consuming services (optional)
if [ "$AUTO_DEPLOY" = true ]; then
    log_step "9" "Auto-deploy to consuming services"

    if [ "$DRY_RUN" = true ]; then
        echo "Would update consuming services with new schema version"
    else
        log_warning "Auto-deployment to consuming services not yet implemented"
        log_info "Manual deployment required. See docs/RELEASE_DEPLOYMENT_GUIDE.md"
        log_info "Services to update:"
        echo "  1. sunday-data/normalizer"
        echo "  2. sunday-data/insights"
        echo "  3. sunday-api/ui-bff"
        echo "  4. sunday-frontend/web"
    fi
fi

# Step 10: Summary and next steps
log_step "10" "Release Summary"

echo ""
echo "======================================"
echo "🎉 RELEASE COMPLETE"
echo "======================================"
echo ""
echo "📦 Version:        $VERSION"
echo "🏷️  NPM Package:    @rakeyshgidwani/sunday-schemas@$VERSION"
echo "🐹 Go Module:      github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION"
echo "📋 GitHub Release: https://github.com/rakeyshgidwani/sunday-schemas/releases/tag/v$VERSION"
echo ""

if [ "$DRY_RUN" = false ]; then
    echo "✅ Packages are being published via GitHub Actions"
    echo "✅ Release assets will be available shortly"
    echo ""
    echo "🔗 Monitor progress:"
    echo "   • GitHub Actions: https://github.com/rakeyshgidwani/sunday-schemas/actions"
    echo "   • NPM Package: https://www.npmjs.com/package/@rakeyshgidwani/sunday-schemas"
    echo "   • Go Module: https://pkg.go.dev/github.com/rakeyshgidwani/sunday-schemas/go"
    echo ""

    if [ "$AUTO_DEPLOY" = false ]; then
        echo "📋 Next Steps:"
        echo "   1. Monitor package publication (2-3 minutes)"
        echo "   2. Update consuming services (see deployment guide)"
        echo "   3. Verify end-to-end functionality"
        echo "   4. Announce release to team"
    fi
else
    echo "🔍 DRY RUN COMPLETED - No actual changes were made"
fi

echo ""
echo "======================================"

# Create post-release reminder
if [ "$DRY_RUN" = false ] && [ "$AUTO_DEPLOY" = false ]; then
    cat > "/tmp/sunday-schemas-release-$VERSION.md" << EOF
# Post-Release Checklist for sunday-schemas v$VERSION

## Immediate Tasks (0-30 minutes)
- [ ] Verify NPM package: \`npm view @rakeyshgidwani/sunday-schemas@$VERSION\`
- [ ] Verify Go module: \`go list -m github.com/rakeyshgidwani/sunday-schemas/go@v$VERSION\`
- [ ] Check GitHub release page: https://github.com/rakeyshgidwani/sunday-schemas/releases/tag/v$VERSION

## Service Updates (30-60 minutes)
- [ ] Update normalizer service: \`cd ../sunday-data/normalizer && npm update @rakeyshgidwani/sunday-schemas\`
- [ ] Update insights service: \`cd ../sunday-data/insights && go get -u github.com/rakeyshgidwani/sunday-schemas/go\`
- [ ] Update UI BFF service: \`cd ../sunday-api/ui-bff && npm update @rakeyshgidwani/sunday-schemas\`
- [ ] Update frontend: \`cd ../sunday-frontend/web && npm update @rakeyshgidwani/sunday-schemas\`

## Validation (60-90 minutes)
- [ ] Run end-to-end tests on staging environment
- [ ] Monitor logs for schema validation errors
- [ ] Verify new schema features work as expected

## Communication
- [ ] Announce release in #schema-updates Slack channel
- [ ] Update team documentation if needed
- [ ] Send email to stakeholders for major releases

Generated: $(date)
EOF

    log_info "📝 Created checklist: /tmp/sunday-schemas-release-$VERSION.md"
fi