#!/bin/bash

# Sunday Schemas - Consuming Services Auto-Deployment Script
# Updates all consuming services with new schema version

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

# Configuration
SCHEMA_VERSION=""
DRY_RUN=false
SERVICES_ROOT="../"
SKIP_TESTS=false
PARALLEL=false

# Service configurations
declare -A SERVICES=(
    ["normalizer"]="sunday-data/normalizer:npm:@rakeyshgidwani/sunday-schemas"
    ["insights"]="sunday-data/insights:go:github.com/rakeyshgidwani/sunday-schemas/go"
    ["ui-bff"]="sunday-api/ui-bff:npm:@rakeyshgidwani/sunday-schemas"
    ["frontend"]="sunday-frontend/web:npm:@rakeyshgidwani/sunday-schemas"
)

usage() {
    cat << 'EOF'
Sunday Schemas - Consuming Services Deployment

Usage: ./scripts/deploy-consuming-services.sh [OPTIONS]

Options:
  -v, --version VERSION     Schema version to deploy (e.g., 1.2.0)
  -r, --services-root PATH  Root path for consuming services (default: ../)
  --dry-run                 Show what would be done without executing
  --skip-tests              Skip running tests after updates
  --parallel                Update services in parallel (faster but less safe)
  -h, --help                Show this help message

Examples:
  ./scripts/deploy-consuming-services.sh --version 1.2.0
  ./scripts/deploy-consuming-services.sh --version 1.2.0 --dry-run
  ./scripts/deploy-consuming-services.sh --version 1.2.0 --parallel

EOF
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            SCHEMA_VERSION="$2"
            shift 2
            ;;
        -r|--services-root)
            SERVICES_ROOT="$2"
            shift 2
            ;;
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        --skip-tests)
            SKIP_TESTS=true
            shift
            ;;
        --parallel)
            PARALLEL=true
            shift
            ;;
        -h|--help)
            usage
            exit 0
            ;;
        *)
            log_error "Unknown option: $1"
            usage
            exit 1
            ;;
    esac
done

if [ -z "$SCHEMA_VERSION" ]; then
    log_error "Schema version is required"
    usage
    exit 1
fi

# Function to update a single service
update_service() {
    local service_name="$1"
    local service_config="$2"

    IFS=':' read -r service_path package_type package_name <<< "$service_config"
    local full_path="$SERVICES_ROOT$service_path"

    echo ""
    log_info "Updating $service_name ($package_type package)"
    echo "  Path: $full_path"
    echo "  Package: $package_name"

    # Check if service directory exists
    if [ ! -d "$full_path" ]; then
        log_warning "Service directory not found: $full_path"
        return 1
    fi

    cd "$full_path"

    # Update based on package type
    case $package_type in
        "npm")
            if [ "$DRY_RUN" = true ]; then
                echo "  Would run: npm update $package_name"
            else
                log_info "  Updating NPM package..."
                if npm update "$package_name"; then
                    # Check if update was successful
                    current_version=$(npm list "$package_name" --depth=0 2>/dev/null | grep "$package_name" | grep -o '@[^[:space:]]*' | cut -c2-)
                    log_success "  Updated to version: $current_version"
                else
                    log_error "  Failed to update NPM package"
                    return 1
                fi
            fi
            ;;
        "go")
            if [ "$DRY_RUN" = true ]; then
                echo "  Would run: go get -u $package_name@v$SCHEMA_VERSION"
            else
                log_info "  Updating Go module..."
                if go get -u "$package_name@v$SCHEMA_VERSION"; then
                    log_success "  Updated Go module"
                    # Run go mod tidy to clean up
                    go mod tidy
                else
                    log_error "  Failed to update Go module"
                    return 1
                fi
            fi
            ;;
        *)
            log_error "  Unknown package type: $package_type"
            return 1
            ;;
    esac

    # Run tests if not skipped
    if [ "$SKIP_TESTS" = false ]; then
        if [ "$DRY_RUN" = true ]; then
            echo "  Would run tests"
        else
            log_info "  Running tests..."
            if [ -f "package.json" ]; then
                if npm test; then
                    log_success "  Tests passed"
                else
                    log_error "  Tests failed"
                    return 1
                fi
            elif [ -f "go.mod" ]; then
                if go test ./...; then
                    log_success "  Tests passed"
                else
                    log_error "  Tests failed"
                    return 1
                fi
            else
                log_warning "  No test configuration found"
            fi
        fi
    fi

    return 0
}

# Main execution
echo ""
echo "======================================"
echo "🚀 CONSUMING SERVICES DEPLOYMENT"
echo "======================================"
echo "Schema Version: $SCHEMA_VERSION"
echo "Services Root: $SERVICES_ROOT"
echo "Dry Run: $DRY_RUN"
echo "Parallel: $PARALLEL"
echo "Skip Tests: $SKIP_TESTS"
echo "======================================"

# Store original directory
ORIGINAL_DIR=$(pwd)

# Track successes and failures
declare -A SERVICE_STATUS
successful_updates=0
failed_updates=0

# Update services
if [ "$PARALLEL" = true ]; then
    log_info "Updating services in parallel..."

    # Run updates in background
    for service_name in "${!SERVICES[@]}"; do
        (
            if update_service "$service_name" "${SERVICES[$service_name]}"; then
                echo "SUCCESS:$service_name" > "/tmp/deploy-$service_name.status"
            else
                echo "FAILED:$service_name" > "/tmp/deploy-$service_name.status"
            fi
        ) &
    done

    # Wait for all background jobs
    wait

    # Collect results
    for service_name in "${!SERVICES[@]}"; do
        if [ -f "/tmp/deploy-$service_name.status" ]; then
            status=$(cat "/tmp/deploy-$service_name.status")
            if [[ $status == SUCCESS:* ]]; then
                SERVICE_STATUS[$service_name]="success"
                successful_updates=$((successful_updates + 1))
            else
                SERVICE_STATUS[$service_name]="failed"
                failed_updates=$((failed_updates + 1))
            fi
            rm -f "/tmp/deploy-$service_name.status"
        fi
    done
else
    log_info "Updating services sequentially..."

    # Dependency order: normalizer -> insights -> ui-bff -> frontend
    service_order=("normalizer" "insights" "ui-bff" "frontend")

    for service_name in "${service_order[@]}"; do
        if [ -n "${SERVICES[$service_name]}" ]; then
            if update_service "$service_name" "${SERVICES[$service_name]}"; then
                SERVICE_STATUS[$service_name]="success"
                successful_updates=$((successful_updates + 1))
            else
                SERVICE_STATUS[$service_name]="failed"
                failed_updates=$((failed_updates + 1))

                # Ask if we should continue
                if [ "$DRY_RUN" = false ]; then
                    echo -n "Continue with remaining services? (y/N): "
                    read -n 1 continue_deploy
                    echo
                    if [[ ! $continue_deploy =~ ^[Yy]$ ]]; then
                        log_info "Deployment stopped"
                        break
                    fi
                fi
            fi
        fi
    done
fi

# Return to original directory
cd "$ORIGINAL_DIR"

# Summary
echo ""
echo "======================================"
echo "📊 DEPLOYMENT SUMMARY"
echo "======================================"
echo "Successful: $successful_updates"
echo "Failed: $failed_updates"
echo ""

for service_name in "${!SERVICE_STATUS[@]}"; do
    status="${SERVICE_STATUS[$service_name]}"
    if [ "$status" = "success" ]; then
        log_success "$service_name: ✅ Updated successfully"
    else
        log_error "$service_name: ❌ Update failed"
    fi
done

echo ""

if [ $failed_updates -eq 0 ]; then
    log_success "🎉 All services updated successfully!"

    if [ "$DRY_RUN" = false ]; then
        echo ""
        echo "🔍 Next steps:"
        echo "  1. Deploy updated services to staging/production"
        echo "  2. Run end-to-end tests"
        echo "  3. Monitor for schema validation errors"
        echo "  4. Verify new functionality works as expected"

        # Generate deployment commands
        cat > "/tmp/sunday-services-deploy-commands.sh" << EOF
#!/bin/bash
# Generated deployment commands for updated services

# Deploy to staging
$(for service_name in "${!SERVICE_STATUS[@]}"; do
    if [ "${SERVICE_STATUS[$service_name]}" = "success" ]; then
        service_config="${SERVICES[$service_name]}"
        service_path=$(echo "$service_config" | cut -d':' -f1)
        echo "kubectl set image deployment/$service_name $service_name=$service_name:$(date +%Y%m%d)-schema-$SCHEMA_VERSION"
    fi
done)

# Health checks
$(for service_name in "${!SERVICE_STATUS[@]}"; do
    if [ "${SERVICE_STATUS[$service_name]}" = "success" ]; then
        echo "kubectl logs -f deployment/$service_name | grep -i 'schema\\|error'"
    fi
done)
EOF

        log_info "📝 Generated deployment commands: /tmp/sunday-services-deploy-commands.sh"
    fi
else
    log_error "⚠️  Some services failed to update"
    echo ""
    echo "🔧 Troubleshooting:"
    echo "  1. Check service logs for specific errors"
    echo "  2. Verify package availability:"
    echo "     - NPM: npm view @rakeyshgidwani/sunday-schemas@$SCHEMA_VERSION"
    echo "     - Go: go list -m github.com/rakeyshgidwani/sunday-schemas/go@v$SCHEMA_VERSION"
    echo "  3. Retry individual services manually"
    echo "  4. Check for breaking changes in CHANGELOG.md"

    exit 1
fi

echo "======================================"