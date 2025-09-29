# Documentation Automation

This document explains the automated documentation system that keeps usage guides in sync with generated code.

## Problem Solved

**Before Automation:**
- Documentation was manually written and maintained
- High risk of documentation drift from actual code
- Time-consuming to keep docs accurate after schema changes
- Teams received inaccurate examples and function references

**After Automation:**
- Documentation is generated from actual codegen artifacts
- Always accurate and up-to-date
- Automatic updates on every code generation
- Zero maintenance overhead for documentation

## How It Works

### 1. Documentation Generator (`scripts/generate-docs.mjs`)

**What it does:**
- Introspects `codegen/ts/index.ts` to extract actual TypeScript types and constants
- Introspects `codegen/go/*.go` files to extract actual Go functions and types
- Generates comprehensive usage documentation with real examples
- Includes timestamps for tracking freshness

**Features:**
- **Type-safe**: Only documents what actually exists in generated code
- **Complete**: Extracts all types, constants, and functions automatically
- **Structured**: Generates consistent, well-formatted documentation
- **Accurate**: Uses actual codegen output as source of truth

### 2. Documentation Sync Checker (`scripts/check-docs-sync.mjs`)

**What it does:**
- Compares current documentation with freshly generated docs
- Fails CI/CD if documentation is out of sync
- Ignores minor differences like timestamps
- Provides clear instructions for fixing sync issues

**Used in:**
- CI/CD pipeline (`npm run ci-validate`)
- Pre-commit hooks (optional)
- Local development verification

### 3. GitHub Actions Auto-Update (`.github/workflows/auto-update-docs.yml`)

**Triggers:**
- Changes to `codegen/` directory
- Changes to `schemas/` or `openapi/` directories
- Manual workflow dispatch

**Actions:**
- Automatically regenerates documentation
- Commits updates with descriptive message
- Comments on triggering commits
- Prevents documentation drift

## Usage

### Generate Documentation Manually

```bash
# Generate documentation from current codegen
npm run generate:docs

# Check if documentation is in sync
npm run check:docs-sync

# Full build with documentation generation
npm run generate  # Includes automatic doc generation
```

### Integration Points

**1. Code Generation Pipeline:**
```bash
npm run generate
# Runs: generate-ts → generate-go → generate:docs
```

**2. CI/CD Validation:**
```bash
npm run ci-validate
# Includes documentation sync check
```

**3. Release Pipeline:**
```bash
npm run release -- --version 1.2.0
# Ensures docs are generated and synced before release
```

## Generated Documentation Structure

### TypeScript Documentation (`docs/TYPESCRIPT_USAGE.md`)

**Auto-generated sections:**
- ✅ Package installation and import examples
- ✅ Complete type listings with aliases
- ✅ Runtime constants and validation helpers
- ✅ Type-safe event handling patterns
- ✅ Generic utilities and union types
- ✅ Complete working examples

**Extracted from:**
- `codegen/ts/index.ts` - Type exports and aliases
- `codegen/ts/index.ts` - Runtime constants
- Package metadata for installation commands

### Go Documentation (`docs/GO_MODULE_USAGE.md`)

**Auto-generated sections:**
- ✅ Module installation and import examples
- ✅ All available types and structs
- ✅ Schema and venue constants with values
- ✅ Exported function signatures
- ✅ JSON marshal/unmarshal functions
- ✅ Complete working examples

**Extracted from:**
- `codegen/go/constants.go` - Constants and validation functions
- `codegen/go/schemas.go` - Type definitions and marshal functions
- `codegen/go/compat.go` - Helper and compatibility functions

## Customization

### Adding New Documentation Sections

Edit `scripts/generate-docs.mjs`:

```javascript
// Add new introspection for TypeScript
async function introspectTypeScript() {
    // ... existing code ...

    // Add new extraction logic
    const newFeature = extractNewFeature(content);

    return {
        // ... existing exports ...
        newFeature
    };
}

// Update documentation template
async function generateTypeScriptDocs() {
    const tsData = await introspectTypeScript();

    return `
    // ... existing template ...

    ## New Feature Section
    ${tsData.newFeature.map(item => `- ${item}`).join('\n')}
    `;
}
```

### Customizing Documentation Format

The generator uses template literals for formatting. Modify the template strings in:
- `generateTypeScriptDocs()` for TypeScript documentation
- `generateGoDocs()` for Go documentation

### Adding Documentation Validation

Extend `check-docs-sync.mjs` to validate specific content:

```javascript
function validateDocumentationContent(content) {
    // Check for required sections
    if (!content.includes('## Quick Start')) {
        throw new Error('Missing Quick Start section');
    }

    // Check for minimum examples
    if ((content.match(/```typescript/g) || []).length < 3) {
        throw new Error('Insufficient TypeScript examples');
    }
}
```

## Benefits

### For Development Teams

1. **Always Accurate**: Documentation matches actual available functions
2. **Complete Examples**: Working code examples that actually compile
3. **Up-to-Date**: Automatically reflects latest schema changes
4. **Type Safety**: TypeScript examples include proper type annotations

### For Maintainers

1. **Zero Maintenance**: Documentation updates automatically
2. **No Drift**: CI/CD prevents documentation from becoming stale
3. **Consistent Quality**: Generated docs follow consistent patterns
4. **Time Savings**: No manual documentation writing or updating

### For CI/CD

1. **Quality Gates**: Prevents releases with outdated documentation
2. **Automatic Updates**: GitHub Actions maintain documentation freshness
3. **Change Tracking**: Clear commit history of documentation updates
4. **Validation**: Fails builds if documentation is out of sync

## Troubleshooting

### Documentation Sync Failures

**Error**: "Documentation is out of sync with codegen"

**Solution**:
```bash
npm run generate:docs
git add docs/
git commit -m "Update documentation"
```

### Generation Failures

**Error**: "Cannot find codegen artifacts"

**Causes**:
- Codegen hasn't been run
- Missing `codegen/ts/index.ts` or `codegen/go/*.go` files

**Solution**:
```bash
npm run generate  # Regenerate all artifacts
npm run generate:docs  # Then regenerate docs
```

### GitHub Actions Issues

**Error**: "Auto-update documentation workflow failed"

**Common causes**:
- Missing Node.js dependencies
- Permission issues with repository writes
- Codegen pipeline failures

**Check**:
1. Workflow logs in GitHub Actions tab
2. Dependencies in `package.json`
3. Repository permissions for GitHub Actions

## Future Enhancements

### Planned Features

1. **Interactive Examples**: Generate runnable CodeSandbox/StackBlitz examples
2. **API Documentation**: Auto-generate OpenAPI documentation
3. **Schema Visualization**: Generate diagrams from schema definitions
4. **Multilingual Support**: Generate documentation in multiple languages
5. **Dependency Tracking**: Show which schemas depend on others

### Integration Opportunities

1. **IDE Integration**: Generate IDE-specific documentation
2. **Documentation Website**: Auto-publish to documentation sites
3. **Package Metadata**: Include documentation in published packages
4. **Testing**: Generate test cases from documentation examples

---

*This system ensures that documentation is never a liability and always an asset that teams can trust.*