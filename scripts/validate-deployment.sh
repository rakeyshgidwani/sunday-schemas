#!/bin/bash

# Sunday Schemas Deployment Validation Script
# Validates that both NPM package and Go module are accessible and working

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

# Parse command line arguments
VERSION=""
HELP=false
VERBOSE=false
QUICK=false

while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        --verbose)
            VERBOSE=true
            shift
            ;;
        --quick)
            QUICK=true
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
    echo "Sunday Schemas Deployment Validation Script"
    echo ""
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -v, --version VERSION    Validate specific version (e.g., 1.0.2)"
    echo "  --verbose               Show detailed output"
    echo "  --quick                 Skip slow checks (Go proxy, integration tests)"
    echo "  -h, --help              Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 --version 1.0.2                    # Validate version 1.0.2"
    echo "  $0 --version 1.0.2 --verbose          # Validate with detailed output"
    echo "  $0 --version 1.0.2 --quick            # Quick validation (skip slow checks)"
    echo ""
    exit 0
fi

# Get version from user if not provided
if [ -z "$VERSION" ]; then
    echo -n "Enter version to validate (e.g., 1.0.2): "
    read VERSION
fi

# Validate version format
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    log_error "Invalid version format. Use semantic versioning (e.g., 1.0.2)"
    exit 1
fi

log_info "Validating deployment for version $VERSION"
echo ""

# Track validation results
NPM_SUCCESS=false
GO_SUCCESS=false
NPM_IMPORT_SUCCESS=false
GO_IMPORT_SUCCESS=false
TOTAL_CHECKS=0
PASSED_CHECKS=0

# Function to run a check
run_check() {
    local description="$1"
    local command="$2"
    local success_var="$3"

    TOTAL_CHECKS=$((TOTAL_CHECKS + 1))

    if [ "$VERBOSE" = true ]; then
        log_info "Running: $description"
        echo "Command: $command"
    else
        printf "%-50s" "$description..."
    fi

    if eval "$command" > /tmp/validation_output 2>&1; then
        if [ "$VERBOSE" = true ]; then
            log_success "$description"
            if [ -s /tmp/validation_output ]; then
                echo "Output:"
                cat /tmp/validation_output | head -10
                echo ""
            fi
        else
            echo -e " ${GREEN}✅${NC}"
        fi
        PASSED_CHECKS=$((PASSED_CHECKS + 1))
        eval "$success_var=true"
    else
        if [ "$VERBOSE" = true ]; then
            log_error "$description"
            echo "Error output:"
            cat /tmp/validation_output
            echo ""
        else
            echo -e " ${RED}❌${NC}"
        fi
        if [ "$VERBOSE" = true ]; then
            echo "Command that failed: $command"
        fi
    fi
}

# Create temporary directory for tests
TEMP_DIR=$(mktemp -d)
cleanup() {
    rm -rf "$TEMP_DIR"
    rm -f /tmp/validation_output
}
trap cleanup EXIT

echo "🔍 VALIDATION CHECKS"
echo "===================="

# 1. Check NPM package exists
run_check "NPM package exists" \
    "npm view sunday-schemas@$VERSION --json" \
    "NPM_SUCCESS"

# 2. Check NPM package download
if [ "$NPM_SUCCESS" = true ]; then
    run_check "NPM package downloadable" \
        "cd '$TEMP_DIR' && npm pack sunday-schemas@$VERSION" \
        "NPM_DOWNLOAD_SUCCESS"
fi

# 3. Check Go module tag exists
# Try local tags first, then remote if available
if git tag | grep -q "codegen/go/v$VERSION"; then
    run_check "Go module tag exists (local)" \
        "git tag | grep 'codegen/go/v$VERSION'" \
        "GO_TAG_SUCCESS"
elif git remote get-url origin > /dev/null 2>&1; then
    run_check "Go module tag exists (remote)" \
        "git ls-remote --tags origin | grep 'codegen/go/v$VERSION'" \
        "GO_TAG_SUCCESS"
else
    run_check "Go module tag exists" \
        "echo 'No git repository or remote found'" \
        "GO_TAG_SUCCESS"
    GO_TAG_SUCCESS=false
fi

# 4. Check Go module accessibility (skip if quick mode or known to be private)
if [ "$QUICK" = false ]; then
    run_check "Go module proxy accessibility" \
        "curl -s --fail --max-time 10 'https://proxy.golang.org/github.com/rakeyshgidwani/sunday-schemas/codegen/go/@v/v$VERSION.info'" \
        "GO_PROXY_SUCCESS"
else
    log_info "Skipping Go proxy check (--quick mode)"
fi

# 5. Test NPM package import
if [ "$NPM_SUCCESS" = true ]; then
    log_info "Testing NPM package import..."

    # Create test package.json
    cat > "$TEMP_DIR/package.json" << EOF
{
  "name": "test-sunday-schemas",
  "version": "1.0.0"
}
EOF

    # Create test file
    cat > "$TEMP_DIR/test.js" << EOF
const schemas = require('sunday-schemas');

// Test basic import
console.log('Schemas imported successfully');

// Test that we can access schema constants
if (schemas.SCHEMA_CONSTANTS && schemas.SCHEMA_CONSTANTS['raw.v0']) {
    console.log('Schema constants available:', schemas.SCHEMA_CONSTANTS['raw.v0']);
} else {
    console.error('Schema constants not found');
    process.exit(1);
}

// Test that we can access venue constants
if (schemas.VENUE_IDS && schemas.VENUE_IDS.includes('polymarket') && schemas.VENUE_IDS.includes('kalshi')) {
    console.log('Venue constants available:', schemas.VENUE_IDS);
} else {
    console.error('Venue constants not found');
    process.exit(1);
}

console.log('NPM package validation successful');
EOF

    run_check "NPM package import test" \
        "cd '$TEMP_DIR' && npm install sunday-schemas@$VERSION && node test.js" \
        "NPM_IMPORT_SUCCESS"
fi

# 6. Test Go module import
if [ "$GO_TAG_SUCCESS" = true ]; then
    log_info "Testing Go module import..."

    # Create test Go module
    mkdir -p "$TEMP_DIR/go-test"
    cd "$TEMP_DIR/go-test"

    cat > go.mod << EOF
module test-sunday-schemas

go 1.21
EOF

    cat > main.go << EOF
package main

import (
    "fmt"
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func main() {
    // Test basic import
    fmt.Println("Go module imported successfully")

    // Test schema constants
    if schemas.SchemaRAW_V0 != "" {
        fmt.Println("Schema constants available:", schemas.SchemaRAW_V0)
    } else {
        fmt.Println("Error: Schema constants not found")
        return
    }

    // Test venue constants
    if schemas.Polymarket != "" && schemas.Kalshi != "" {
        fmt.Println("Venue constants available:", schemas.Polymarket, schemas.Kalshi)
    } else {
        fmt.Println("Error: Venue constants not found")
        return
    }

    // Test validation functions
    if err := schemas.ValidateSchema(string(schemas.SchemaRAW_V0)); err != nil {
        fmt.Println("Error: Schema validation failed:", err)
        return
    }

    if err := schemas.ValidateVenue(string(schemas.Polymarket)); err != nil {
        fmt.Println("Error: Venue validation failed:", err)
        return
    }

    fmt.Println("Go module validation successful")
}
EOF

    # Add note about private repositories
    if [ "$VERBOSE" = true ]; then
        log_info "Note: Private repositories may require authentication for Go module access"
    fi

    run_check "Go module import test" \
        "cd '$TEMP_DIR/go-test' && go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION && go run main.go" \
        "GO_IMPORT_SUCCESS"
else
    log_warning "Skipping Go module import test (tag not found)"
fi

# 7. Additional checks if not in quick mode
if [ "$QUICK" = false ]; then
    # Check package contents
    if [ "$NPM_SUCCESS" = true ]; then
        run_check "NPM package contains TypeScript types" \
            "cd '$TEMP_DIR' && npm pack sunday-schemas@$VERSION && tar -tf sunday-schemas-$VERSION.tgz | grep -q '\.d\.ts'" \
            "NPM_TYPES_SUCCESS"

        run_check "NPM package contains JavaScript files" \
            "cd '$TEMP_DIR' && tar -tf sunday-schemas-$VERSION.tgz | grep -q '\.js'" \
            "NPM_JS_SUCCESS"
    fi

    # Check Go module contains expected files
    run_check "Go module contains schema types" \
        "curl -s 'https://raw.githubusercontent.com/rakeyshgidwani/sunday-schemas/codegen/go/v$VERSION/schemas.go' | grep -q 'type.*Envelope'" \
        "GO_TYPES_SUCCESS"
fi

# Summary
echo ""
echo "📊 VALIDATION SUMMARY"
echo "===================="
echo ""

# Overall status
if [ "$PASSED_CHECKS" -eq "$TOTAL_CHECKS" ]; then
    log_success "All validation checks passed ($PASSED_CHECKS/$TOTAL_CHECKS)"
    OVERALL_SUCCESS=true
else
    log_error "Some validation checks failed ($PASSED_CHECKS/$TOTAL_CHECKS)"
    OVERALL_SUCCESS=false
fi

echo ""
echo "📦 NPM Package (sunday-schemas@$VERSION):"
if [ "$NPM_SUCCESS" = true ]; then
    echo -e "   Registry: ${GREEN}✅ Available${NC}"
else
    echo -e "   Registry: ${RED}❌ Not found${NC}"
fi

if [ "$NPM_IMPORT_SUCCESS" = true ]; then
    echo -e "   Import:   ${GREEN}✅ Working${NC}"
elif [ "$NPM_SUCCESS" = true ]; then
    echo -e "   Import:   ${RED}❌ Failed${NC}"
else
    echo -e "   Import:   ${YELLOW}⏭️  Skipped${NC}"
fi

echo ""
echo "🐹 Go Module (github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION):"
if [ "$GO_TAG_SUCCESS" = true ]; then
    echo -e "   Git Tag:  ${GREEN}✅ Available${NC}"
else
    echo -e "   Git Tag:  ${RED}❌ Not found${NC}"
fi

if [ "$GO_IMPORT_SUCCESS" = true ]; then
    echo -e "   Import:   ${GREEN}✅ Working${NC}"
elif [ "$GO_TAG_SUCCESS" = true ]; then
    echo -e "   Import:   ${RED}❌ Failed${NC}"
else
    echo -e "   Import:   ${YELLOW}⏭️  Skipped${NC}"
fi

echo ""
echo "🔗 Usage Instructions:"
echo ""
if [ "$NPM_SUCCESS" = true ]; then
    echo "📦 TypeScript/JavaScript:"
    echo "   npm install sunday-schemas@$VERSION"
    echo ""
fi

if [ "$GO_TAG_SUCCESS" = true ]; then
    echo "🐹 Go:"
    echo "   go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v$VERSION"
    echo ""
fi

echo "===================="

# Exit with appropriate code
if [ "$OVERALL_SUCCESS" = true ]; then
    echo -e "${GREEN}🎉 Deployment validation completed successfully!${NC}"
    exit 0
else
    echo -e "${RED}💥 Deployment validation failed!${NC}"
    echo ""
    echo "Common issues:"
    echo "• NPM package not found: Check if 'npm publish' succeeded"
    echo "• Go module not found: Check if git tags were pushed correctly"
    echo "• Import failures: Check for syntax errors in generated code"
    echo "• Proxy issues: May indicate private repository (expected for private repos)"
    exit 1
fi