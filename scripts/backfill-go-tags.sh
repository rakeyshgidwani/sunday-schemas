#!/bin/bash

# Sunday Schemas - Backfill Go Module Tags
# Creates missing go/ prefixed tags for existing releases

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
log_warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# Check if we're in the right directory
if [ ! -f "package.json" ] || [ ! -d "go" ]; then
    log_error "Must be run from sunday-schemas root directory"
    exit 1
fi

DRY_RUN=false

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        -h|--help)
            cat << 'EOF'
Sunday Schemas - Backfill Go Module Tags

Creates missing go/ prefixed tags for existing releases to make them
available as proper Go modules.

Usage: ./scripts/backfill-go-tags.sh [OPTIONS]

Options:
  --dry-run    Show what would be done without executing
  -h, --help   Show this help message

Background:
Go modules in subdirectories need both root tags (v1.0.9) and prefixed
tags (go/v1.0.9) to be properly discoverable by `go get`.

EOF
            exit 0
            ;;
        *)
            log_error "Unknown option: $1"
            exit 1
            ;;
    esac
done

echo ""
echo "======================================"
echo "🏷️  GO MODULE TAG BACKFILL"
echo "======================================"
echo "Dry Run: $DRY_RUN"
echo "======================================"
echo ""

# Get all existing root-level version tags
root_tags=$(git tag -l | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | sort -V)

if [ -z "$root_tags" ]; then
    log_warning "No version tags found"
    exit 0
fi

log_info "Found root version tags:"
echo "$root_tags"
echo ""

# Check which go/ tags already exist
existing_go_tags=$(git tag -l | grep -E '^go/v' | sort -V)
if [ -n "$existing_go_tags" ]; then
    log_info "Existing go/ tags:"
    echo "$existing_go_tags"
    echo ""
fi

# Process each root tag
created_count=0
skipped_count=0

for tag in $root_tags; do
    go_tag="go/$tag"
    codegen_go_tag="codegen/go/$tag"

    # Check if go/ tag already exists
    if git tag -l | grep -q "^$go_tag$"; then
        log_info "⏭️  $go_tag already exists, skipping"
        skipped_count=$((skipped_count + 1))
    else
        if [ "$DRY_RUN" = true ]; then
            echo "Would create: $go_tag -> $tag"
        else
            # Create the go/ prefixed tag pointing to the same commit as the root tag
            if git tag "$go_tag" "$tag" -m "Go module release $go_tag"; then
                log_success "✅ Created $go_tag"
                created_count=$((created_count + 1))
            else
                log_error "❌ Failed to create $go_tag"
            fi
        fi
    fi

    # Check if codegen/go/ tag already exists
    if git tag -l | grep -q "^$codegen_go_tag$"; then
        log_info "⏭️  $codegen_go_tag already exists, skipping"
        skipped_count=$((skipped_count + 1))
    else
        if [ "$DRY_RUN" = true ]; then
            echo "Would create: $codegen_go_tag -> $tag"
        else
            # Create the codegen/go/ prefixed tag pointing to the same commit as the root tag
            if git tag "$codegen_go_tag" "$tag" -m "Codegen Go module release $codegen_go_tag"; then
                log_success "✅ Created $codegen_go_tag"
                created_count=$((created_count + 1))
            else
                log_error "❌ Failed to create $codegen_go_tag"
            fi
        fi
    fi
done

echo ""
log_info "Summary:"
echo "  Created: $created_count"
echo "  Skipped: $skipped_count"

if [ "$DRY_RUN" = false ] && [ $created_count -gt 0 ]; then
    echo ""
    log_info "Pushing new go/ tags to remote..."

    # Push all new tags
    for tag in $root_tags; do
        go_tag="go/$tag"
        codegen_go_tag="codegen/go/$tag"

        if git tag -l | grep -q "^$go_tag$"; then
            # Only push if this tag was just created (not skipped)
            if ! git ls-remote --tags origin | grep -q "refs/tags/$go_tag$"; then
                if [ "$DRY_RUN" = false ]; then
                    git push origin "$go_tag"
                    log_success "📤 Pushed $go_tag"
                fi
            fi
        fi

        if git tag -l | grep -q "^$codegen_go_tag$"; then
            # Only push if this tag was just created (not skipped)
            if ! git ls-remote --tags origin | grep -q "refs/tags/$codegen_go_tag$"; then
                if [ "$DRY_RUN" = false ]; then
                    git push origin "$codegen_go_tag"
                    log_success "📤 Pushed $codegen_go_tag"
                fi
            fi
        fi
    done

    echo ""
    log_success "🎉 Go module tags backfilled successfully!"

    echo ""
    echo "🔍 Verify the tags are working:"
    echo "  # Modular structure (separate packages):"
    echo "  go list -m -versions github.com/rakeyshgidwani/sunday-schemas/go"
    echo "  go get github.com/rakeyshgidwani/sunday-schemas/go@$(echo "$root_tags" | tail -1)"
    echo ""
    echo "  # Unified package structure (all types in one package):"
    echo "  go list -m -versions github.com/rakeyshgidwani/sunday-schemas/codegen/go"
    echo "  go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@$(echo "$root_tags" | tail -1)"

elif [ "$DRY_RUN" = true ]; then
    echo ""
    echo "🔍 DRY RUN: No actual changes were made"
    echo "Run without --dry-run to create the tags"

elif [ $created_count -eq 0 ]; then
    echo ""
    log_info "ℹ️  All go/ tags already exist - nothing to do"
fi

echo ""
echo "======================================"