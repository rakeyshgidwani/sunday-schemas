# Package Publishing Setup Guide

This guide explains how to set up automated publishing for the Sunday Schemas packages.

## Prerequisites

### 1. NPM Account Setup
1. Create an NPM account at https://www.npmjs.com/
2. Create an NPM access token:
   ```bash
   npm login
   npm token create --read-and-publish
   ```
3. Add the token as `NPM_TOKEN` secret in GitHub repository settings

### 2. GitHub Repository Setup
1. Ensure repository has proper permissions for GitHub Actions
2. Add the following secrets in repository settings:
   - `NPM_TOKEN`: Your NPM access token
   - `GITHUB_TOKEN`: Automatically provided by GitHub

## Publishing Process

### Automated Publishing (Recommended)

1. **Create a release tag**:
   ```bash
   # Create and push a version tag
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **GitHub Actions will automatically**:
   - Validate schemas and run compatibility checks
   - Generate TypeScript and Go artifacts
   - Publish `@sunday/schemas` to NPM
   - Tag Go module for `go get` compatibility
   - Create GitHub release with usage instructions

### Manual Publishing (Development)

#### NPM Package
```bash
# Navigate to TypeScript package
cd codegen/ts

# Update version
npm version patch  # or minor/major

# Publish
npm publish
```

#### Go Module
```bash
# Create version tag for Go module
git tag codegen/go/v1.0.0
git push origin codegen/go/v1.0.0
```

## Testing the Published Packages

### Test NPM Package
```bash
# Create test project
mkdir test-sunday-schemas && cd test-sunday-schemas
npm init -y
npm install @sunday/schemas@latest

# Test import
node -e "console.log(require('@sunday/schemas'))"
```

### Test Go Module
```bash
# Create test module
mkdir test-go && cd test-go
go mod init test

# Add dependency
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@latest

# Test import
cat > main.go << 'EOF'
package main

import (
    "fmt"
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func main() {
    fmt.Println("Schemas loaded successfully")
    fmt.Printf("Supported venues: %v\n", schemas.AllVenues())
}
EOF

go run main.go
```

## Version Management

### Semantic Versioning
- **Patch (v1.0.1)**: Bug fixes, documentation
- **Minor (v1.1.0)**: New schemas, backward-compatible changes
- **Major (v2.0.0)**: Breaking changes

### Release Process
1. Update schemas/code as needed
2. Run validation: `npm run validate && npm run check-compatibility`
3. Create git tag: `git tag v1.x.y && git push origin v1.x.y`
4. GitHub Actions handles the rest

## Troubleshooting

### NPM Publish Fails
- Check NPM token permissions
- Verify package name isn't taken
- Ensure version number is incremented

### Go Module Not Found
- Wait 5-10 minutes after tagging (proxy cache)
- Use `GOPROXY=direct go get ...` to bypass cache
- Verify tag format: `codegen/go/v1.0.0`

### GitHub Actions Fail
- Check repository secrets are set
- Verify workflow file syntax
- Review Actions logs for specific errors

## Package URLs

Once published, packages will be available at:
- **NPM**: https://www.npmjs.com/package/@sunday/schemas
- **Go**: https://pkg.go.dev/github.com/rakeyshgidwani/sunday-schemas/codegen/go