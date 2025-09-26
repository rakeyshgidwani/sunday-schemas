# Schema Deprecation Guide

This document describes the process for deprecating and removing schemas in the Sunday platform.

## Overview

Schema deprecation allows us to phase out old schemas in a controlled manner while giving consumers time to migrate. The process ensures backward compatibility during a transition period and provides clear migration paths.

## Deprecation Process

### 1. Deprecation Marking

When a schema is marked for deprecation:

#### 1.1 Add Deprecation Metadata
Add the `x-deprecated` extension to the schema root:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://schemas.sunday.com/old.schema.v1.json",
  "title": "Old Schema V1",
  "x-deprecated": {
    "deprecated": true,
    "deprecatedInVersion": "v2.3.0",
    "removalPlannedInVersion": "v2.5.0",
    "reason": "Replaced by new.schema.v2 for improved performance and additional fields",
    "replacedBy": "new.schema.v2",
    "migrationGuide": "https://docs.sunday.com/schemas/migration/old-to-new-v2",
    "deprecationDate": "2025-03-15",
    "plannedRemovalDate": "2025-06-15"
  },
  "description": "⚠️ DEPRECATED: This schema will be removed in v2.5.0. Use new.schema.v2 instead.",
  // ... rest of schema
}
```

#### 1.2 Update Schema Description
- Add clear deprecation warning at the start of the description
- Include replacement schema and migration timeline
- Use warning emoji (⚠️) to make deprecation visible

#### 1.3 Update Documentation
```markdown
# Old Schema V1 ⚠️ DEPRECATED

> **Warning**: This schema is deprecated as of v2.3.0 and will be removed in v2.5.0.
> Please migrate to [`new.schema.v2`](./new.schema.v2.md) before the removal date.
>
> **Migration Guide**: [Old to New V2 Migration](../migration/old-to-new-v2.md)
> **Support Ends**: June 15, 2025

## Overview
<!-- existing documentation with deprecation warnings -->
```

### 2. Migration Support

#### 2.1 Create Migration Guide
Create detailed migration documentation in `docs/migrations/`:

```markdown
# Migrating from old.schema.v1 to new.schema.v2

## Overview
This guide helps you migrate from the deprecated `old.schema.v1` to `new.schema.v2`.

## Field Mapping
| Old Field | New Field | Notes |
|-----------|-----------|-------|
| `oldField` | `newField` | Renamed for clarity |
| `removed` | N/A | Field removed, use `replacement` instead |
| N/A | `added` | New required field, use default value `"default"` |

## Code Examples

### Before (old.schema.v1)
```json
{
  "schema": "old.schema.v1",
  "oldField": "value",
  "removed": "deprecated"
}
```

### After (new.schema.v2)
```json
{
  "schema": "new.schema.v2",
  "newField": "value",
  "replacement": "updated",
  "added": "default"
}
```

### Migration Script
```javascript
function migrateOldToNew(oldEvent) {
  return {
    schema: "new.schema.v2",
    newField: oldEvent.oldField,
    replacement: oldEvent.removed || "updated",
    added: "default",
    // ... other mappings
  };
}
```
```

#### 2.2 Update Examples
- Keep existing examples but add deprecation warnings
- Add migration examples showing old→new transformation
- Create examples for the replacement schema

#### 2.3 Automated Migration Tools
Consider creating migration utilities:
```javascript
// Migration utility
const SchemaMigrator = {
  'old.schema.v1': {
    target: 'new.schema.v2',
    migrate: (event) => ({
      schema: 'new.schema.v2',
      newField: event.oldField,
      replacement: event.removed || 'updated',
      added: 'default'
    })
  }
};
```

### 3. Communication Plan

#### 3.1 Deprecation Announcement
1. **GitHub Issue**: Create deprecation announcement issue
2. **Team Notifications**: Notify all affected service teams
3. **Documentation**: Update all relevant documentation
4. **Generated Types**: Ensure deprecation is visible in generated code

#### 3.2 Timeline Communication
- **Initial Notice**: 3 months before removal
- **Migration Reminders**: Monthly reminders to affected teams
- **Final Warning**: 2 weeks before removal
- **Removal Notice**: When schema is actually removed

### 4. Validation During Deprecation

#### 4.1 Deprecation Warnings
The validation system shows warnings for deprecated schemas:

```bash
⚠️  WARNING: old.schema.v1 is deprecated (since v2.3.0)
   → Will be removed in: v2.5.0 (planned: 2025-06-15)
   → Migrate to: new.schema.v2
   → Guide: https://docs.sunday.com/schemas/migration/old-to-new-v2
```

#### 4.2 Usage Tracking
- Monitor usage of deprecated schemas in production
- Track migration progress across services
- Generate deprecation usage reports

### 5. Removal Process

#### 5.1 Pre-removal Validation
Before removing a deprecated schema:

1. **Usage Check**: Verify no production usage remains
2. **Service Confirmation**: Get confirmation from all service teams
3. **Monitoring**: Ensure no events with deprecated schema in logs
4. **Rollback Plan**: Prepare rollback if issues discovered

#### 5.2 Schema Removal
1. **Remove Schema File**: Delete the schema from `schemas/json/`
2. **Remove Examples**: Delete associated example files
3. **Update Validation**: Remove from validation scripts
4. **Update Generation**: Remove from type generation
5. **Archive Documentation**: Move docs to `docs/archived/`

#### 5.3 Post-removal Cleanup
1. **Update Dependencies**: Remove from generated types
2. **Clean CI/CD**: Update validation pipelines
3. **Documentation**: Update schema registry documentation
4. **Communication**: Announce successful removal

## Deprecation Metadata Specification

### Required Fields
- `deprecated: true` - Marks schema as deprecated
- `deprecatedInVersion` - Version when deprecation was introduced
- `reason` - Human-readable reason for deprecation

### Recommended Fields
- `removalPlannedInVersion` - Version when removal is planned
- `replacedBy` - Recommended replacement schema
- `migrationGuide` - URL to migration documentation
- `deprecationDate` - ISO date when deprecated
- `plannedRemovalDate` - ISO date for planned removal

### Optional Fields
- `contact` - Team or person responsible for migration support
- `urgency` - "low", "medium", "high" based on timeline
- `impactLevel` - "low", "medium", "high" based on usage
- `automatedMigration` - Whether automated migration is available

## Version Overlap Requirements

### Minimum Overlap Period
- **Minor Version Changes**: 1 minor version overlap
- **Major Version Changes**: 2 minor versions overlap
- **Critical Security**: May be expedited with SWG approval

### Overlap Timeline Example
```
v2.3.0: Schema deprecated, replacement available
v2.4.0: Both schemas supported, migration encouraged
v2.5.0: Deprecated schema removed, only new schema supported
```

### Support Commitments
During overlap period:
- Both schemas are fully supported
- Examples maintained for both versions
- Generated types include both versions
- Documentation available for both versions
- CI/CD validates both schemas

## Monitoring and Metrics

### Deprecation Tracking
- Number of deprecated schemas
- Age of deprecated schemas (time since deprecation)
- Migration progress by service team
- Usage trends for deprecated schemas

### Quality Metrics
- Time from deprecation to removal
- Number of rollbacks due to premature removal
- Migration success rate
- Post-removal incident rate

## Best Practices

### For Schema Designers
1. **Plan Ahead**: Consider deprecation strategy during initial design
2. **Minimize Breaking Changes**: Use additive changes when possible
3. **Clear Communication**: Provide detailed migration guides
4. **Gradual Migration**: Phase complex migrations into multiple steps

### For Service Teams
1. **Monitor Deprecations**: Subscribe to schema change notifications
2. **Migrate Early**: Don't wait until the last minute
3. **Test Thoroughly**: Validate migrations in staging environments
4. **Communicate Issues**: Report migration problems to SWG immediately

### For SWG Members
1. **Consistent Process**: Follow deprecation process for all changes
2. **Reasonable Timelines**: Allow adequate time for migration
3. **Support Teams**: Provide migration assistance when needed
4. **Monitor Progress**: Track migration progress and offer help

## Emergency Deprecation

For security or critical issues requiring immediate deprecation:

1. **Emergency SWG Review**: Can bypass normal RFC process
2. **Shortened Timeline**: Minimum 2 weeks notice instead of 3 months
3. **Active Support**: SWG provides direct migration assistance
4. **Extended Monitoring**: Additional monitoring during transition
5. **Rollback Ready**: Keep rollback plan available longer than usual

## Automated Tooling

### Deprecation Detection
```bash
npm run check-deprecations  # Check for deprecated schema usage
npm run deprecation-report  # Generate deprecation usage report
npm run migration-status    # Show migration progress across services
```

### Integration with CI/CD
- Automated warnings in PR reviews for deprecated schema usage
- Build warnings when using deprecated schemas
- Regular reports on deprecation progress
- Automated removal when usage drops to zero

---

*This deprecation process ensures smooth transitions while maintaining system stability and giving teams adequate time to migrate.*