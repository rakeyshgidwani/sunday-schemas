# Sunday Schemas Governance

This document provides an overview of the governance process for Sunday platform event schemas and API contracts.

## 🏛️ Schema Working Group (SWG)

The Schema Working Group is responsible for maintaining the integrity, evolution, and compatibility of all Sunday platform schemas.

### Quick Links
- 📋 [Full Governance Process](./docs/governance.md)
- 📖 [SWG Member Handbook](./docs/swg-handbook.md)
- 🗑️ [Deprecation Guide](./docs/deprecation-guide.md)
- 🔄 [Change Process](#change-process)

### Core Members
- **Schema Registry Maintainer** (Lead)
- **Platform Architecture Lead**
- **Data Engineering Lead**
- **Backend Services Lead**

### Contributing Members
- Service owners implementing schema-dependent features
- Data consumers (analytics, ML, reporting teams)
- Frontend teams using generated types

## 📝 Change Process

### 1. For Non-Breaking Changes
1. Create Pull Request using the [PR template](./.github/pull_request_template.md)
2. Ensure all CI checks pass
3. Get approval from **2 SWG core members**
4. Update CHANGELOG.md
5. Merge after approval

### 2. For Breaking Changes
1. Create [RFC issue](./.github/ISSUE_TEMPLATE/schema-change-rfc.md) first
2. Allow **3 business days** for community feedback
3. Get approval from **ALL SWG core members**
4. Create detailed migration guide
5. Ensure **minimum 1 minor version overlap** for deprecated schemas
6. Follow standard PR process with additional approvals

### 3. Schema Deprecation
1. Add `x-deprecated` metadata to schema:
```json
{
  "x-deprecated": {
    "deprecated": true,
    "deprecatedInVersion": "v2.3.0",
    "removalPlannedInVersion": "v2.5.0",
    "reason": "Replaced by new.schema.v2",
    "migrationGuide": "https://docs.sunday.com/migration/..."
  }
}
```
2. Update documentation with deprecation warnings
3. Create migration guide in `docs/migrations/`
4. Maintain **minimum 1 minor version overlap**

## 🔍 Validation

All changes are automatically validated:

```bash
# Check deprecation metadata
npm run check-deprecations

# Validate version overlap requirements
npm run check-version-overlap

# Generate governance reports
npm run deprecation-report
```

### CI Pipeline
- ✅ Schema structural validation
- ✅ Example validation against schemas
- ✅ Backward compatibility checking
- ✅ Deprecation metadata validation
- ✅ Version overlap requirements
- ✅ OpenAPI specification linting

## 📊 Governance Commands

| Command | Description |
|---------|-------------|
| `npm run check-deprecations` | Validate deprecation metadata |
| `npm run check-version-overlap` | Check version overlap requirements |
| `npm run deprecation-report` | Generate deprecation usage report |
| `npm run check-compatibility` | Backward compatibility analysis |
| `npm run check-changelog` | Ensure CHANGELOG is updated |

## 🚨 Emergency Procedures

### Security Issues
- Emergency patches can bypass normal review
- Must be reviewed by Security team + SWG lead
- Retrospective review required within 24 hours

### Production Issues
- Hotfixes allowed with SWG lead + service owner approval
- Must include rollback plan
- Normal review process applies post-hotfix

## 📞 Getting Help

### For Schema Changes
- 💬 **Slack**: #schema-working-group
- 🐛 **Issues**: [Create GitHub issue](./issues/new)
- 📧 **Email**: swg-core@sunday.com

### For Service Integration
- 📚 **Documentation**: [Schema Registry Docs](./docs/)
- 🔧 **Migration Guides**: [Migration Documentation](./docs/migrations/)
- 💡 **Examples**: [Schema Examples](./schemas/examples/)

## 📈 Metrics & Quality

### Governance Health
- Schema change approval time: **< 5 business days**
- Breaking change frequency: **< 1 per quarter**
- Example coverage: **100%** (currently 8/8 schemas)
- Generated type compilation: **100%** success rate

### Current Status
- **Total Schemas**: 8 active
- **Deprecated Schemas**: 0
- **Coverage**: 162% (13 examples for 8 schemas)
- **Last Breaking Change**: None yet

## 🔄 Process Updates

This governance process is itself subject to change:
1. Propose changes via GitHub issue
2. Allow 1 week for community feedback
3. SWG review and approval required
4. Update documentation and communicate changes

## 🚀 Quick Start

### New Contributors
1. Read [Governance Overview](./docs/governance.md)
2. Join #schema-working-group Slack
3. Review recent schema changes for examples
4. Start with small, non-breaking changes

### Service Teams
1. Identify SWG liaison for your team
2. Subscribe to schema change notifications
3. Review schemas your services depend on
4. Participate in RFC discussions

### SWG Members
1. Complete [SWG Handbook](./docs/swg-handbook.md) onboarding
2. Setup GitHub notifications for this repository
3. Schedule recurring review time
4. Join weekly SWG sync meetings

---

## 📋 Templates & Resources

- 🔄 [Pull Request Template](./.github/pull_request_template.md)
- 💡 [RFC Issue Template](./.github/ISSUE_TEMPLATE/schema-change-rfc.md)
- 📚 [Full Documentation](./docs/)
- 🏗️ [Implementation Plan](./IMPLEMENTATION_PLAN.md)

*This governance framework ensures Sunday platform schemas evolve in a controlled, collaborative manner while maintaining system stability and developer productivity.*