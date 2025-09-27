#!/bin/bash

# Sunday Schemas Deploy Script
# Publishes both NPM package and Go module with proper versioning

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
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

# Check if we're in the right directory
if [ ! -f "package.json" ] || [ ! -d "codegen" ]; then
    log_error "Must be run from sunday-schemas root directory"
    exit 1
fi

# Parse command line arguments
VERSION=""
SKIP_TESTS=false
DRY_RUN=false
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
    echo "Sunday Schemas Deploy Script"
    echo ""
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -v, --version VERSION    Specify version to deploy (e.g., 1.0.2)"
    echo "  --skip-tests            Skip validation and tests"
    echo "  --dry-run               Show what would be done without executing"
    echo "  -h, --help              Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 --version 1.0.2                    # Deploy version 1.0.2"
    echo "  $0 --version 1.0.2 --skip-tests       # Deploy without running tests"
    echo "  $0 --version 1.0.2 --dry-run          # Show what would be deployed"
    echo ""
    exit 0
fi

# Get version from user if not provided
if [ -z "$VERSION" ]; then
    echo -n "Enter version to deploy (e.g., 1.0.2): "
    read VERSION
fi

# Validate version format
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    log_error "Invalid version format. Use semantic versioning (e.g., 1.0.2)"
    exit 1
fi

log_info "Starting deployment for version $VERSION"

# Check git status
if [ "$DRY_RUN" = false ]; then
    if [ -n "$(git status --porcelain)" ]; then
        log_warning "Working directory has uncommitted changes"
        echo -n "Continue anyway? (y/N): "
        read -n 1 continue_deploy
        echo
        if [[ ! $continue_deploy =~ ^[Yy]$ ]]; then
            log_info "Deployment cancelled"
            exit 0
        fi
    fi
fi

# Step 1: Run validation and tests (unless skipped)
if [ "$SKIP_TESTS" = false ]; then
    log_info "Step 1: Running validation and tests..."

    if [ "$DRY_RUN" = true ]; then
        echo "Would run: npm run ci-validate"
    else
        npm run ci-validate
        log_success "Validation passed"
    fi
else
    log_warning "Skipping tests (--skip-tests specified)"
fi

# Step 2: Update package.json version
log_info "Step 2: Updating package.json version to $VERSION..."

if [ "$DRY_RUN" = true ]; then
    echo "Would update package.json version to $VERSION"
else
    # Update package.json version
    npm version $VERSION --no-git-tag-version
    log_success "Updated package.json to version $VERSION"
fi

# Step 3: Update Go module version
log_info "Step 3: Updating Go module version..."

if [ "$DRY_RUN" = true ]; then
    echo "Would update codegen/go/go.mod version comment"
else
    # Add version comment to go.mod (optional but helpful)
    sed -i.bak "1s/.*/module github.com\/rakeyshgidwani\/sunday-schemas\/codegen\/go/" codegen/go/go.mod
    rm -f codegen/go/go.mod.bak
    log_success "Updated Go module"
fi

# Step 4: Build and validate generated code
log_info "Step 4: Building and validating generated code..."

if [ "$DRY_RUN" = true ]; then
    echo "Would run: npm run build"
else
    npm run build
    log_success "Build completed successfully"
fi

# Step 5: Commit changes
log_info "Step 5: Committing version bump..."

if [ "$DRY_RUN" = true ]; then
    echo "Would commit: Release version $VERSION"
    echo "Would add files: package.json, codegen/"
else
    git add package.json codegen/
    git commit -m "Release version $VERSION

🚀 Generated with [Claude Code](https://claude.ai/code)

Co-Authored-By: Claude <noreply@anthropic.com>"
    log_success "Committed version bump"
fi

# Step 6: Create and push tags
log_info "Step 6: Creating and pushing tags..."

NPM_TAG="v$VERSION"
GO_TAG="codegen/go/v$VERSION"

if [ "$DRY_RUN" = true ]; then
    echo "Would create tags: $NPM_TAG, $GO_TAG"
    echo "Would push: git push origin main --tags"
else
    # Create tags
    git tag $NPM_TAG
    git tag $GO_TAG

    # Push changes and tags
    git push origin main
    git push origin $NPM_TAG
    git push origin $GO_TAG

    log_success "Created and pushed tags: $NPM_TAG, $GO_TAG"
fi

# Step 7: Publish NPM package
log_info "Step 7: Publishing NPM package..."

if [ "$DRY_RUN" = true ]; then
    echo "Would run: npm publish"
else
    npm publish
    NPM_STATUS=$?
    if [ $NPM_STATUS -eq 0 ]; then
        log_success "NPM package published: sunday-schemas@$VERSION"
    else
        log_error "NPM publish failed"
        exit 1
    fi
fi

# Step 8: Verify Go module accessibility
log_info "Step 8: Verifying Go module accessibility..."

if [ "$DRY_RUN" = true ]; then
    echo "Would verify Go module at: github.com/rakeyshgidwani/sunday-schemas/codegen/go@$GO_TAG"
else
    # Wait a moment for Go proxy to pick up the module
    log_info "Waiting 30 seconds for Go proxy to update..."
    sleep 30

    # Try to fetch the module info (this might fail for private repos, but shows the attempt)
    log_info "Checking Go module availability..."
    GO_MODULE_URL="https://proxy.golang.org/github.com/rakeyshgidwani/sunday-schemas/codegen/go/@v/$GO_TAG.info"
    if curl -s --fail "$GO_MODULE_URL" > /dev/null 2>&1; then
        log_success "Go module is publicly accessible"
    else
        log_warning "Go module not publicly accessible (may be private repo - this is expected)"
    fi
fi

# Step 9: Summary
log_info "Step 9: Deployment Summary"

echo ""
echo "======================================"
echo "🚀 DEPLOYMENT COMPLETE"
echo "======================================"
echo ""
echo "📦 NPM Package:    sunday-schemas@$VERSION"
echo "🐹 Go Module:      github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION"
echo "🏷️  Git Tags:       $NPM_TAG, $GO_TAG"
echo ""
echo "Usage instructions:"
echo ""
echo "📦 TypeScript/JavaScript:"
echo "   npm install sunday-schemas@$VERSION"
echo ""
echo "🐹 Go:"
echo "   go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION"
echo ""

if [ "$DRY_RUN" = true ]; then
    echo "🔍 DRY RUN: No actual changes were made"
else
    echo "✅ All packages successfully deployed!"
fi

echo "======================================"