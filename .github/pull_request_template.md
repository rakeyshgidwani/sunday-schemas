# Schema Change Pull Request

## Change Summary
<!-- Briefly describe what this PR changes -->

## Change Type
- [ ] **New Schema** - Adding a new event schema
- [ ] **Schema Modification** - Modifying an existing schema
- [ ] **Schema Deprecation** - Marking schema as deprecated
- [ ] **API Change** - OpenAPI specification changes
- [ ] **Breaking Change** - ⚠️ Changes that break backward compatibility
- [ ] **Documentation** - Documentation-only changes
- [ ] **Tooling** - Changes to validation, generation, or CI/CD

## Impact Assessment

### Affected Schemas
<!-- List all schemas modified by this PR -->
- [ ] `raw.v0`
- [ ] `md.orderbook.delta.v1`
- [ ] `md.trade.v1`
- [ ] `insights.arb.lite.v1`
- [ ] `insights.movers.v1`
- [ ] `insights.whales.lite.v1`
- [ ] `insights.unusual.v1`
- [ ] `infra.venue_health.v1`
- [ ] **Other:** _____

### Downstream Impact
<!-- Check all that apply -->
- [ ] **Data Pipeline**: May affect event processing or ETL
- [ ] **Analytics**: May impact reporting or ML features
- [ ] **UI/Frontend**: May require frontend type updates
- [ ] **Service Integration**: May affect service-to-service communication
- [ ] **External APIs**: May impact public API contracts
- [ ] **Storage**: May affect data storage schemas

### Services Using These Schemas
<!-- List known services that consume the affected schemas -->
- [ ] `sunday-data/normalizer`
- [ ] `sunday-data/insights`
- [ ] `sunday-api/ui-bff`
- [ ] `sunday-frontend/web`
- [ ] **Other services:** _____

## Schema Working Group Review

### Required Approvals
- [ ] **SWG Core Member 1**: @username
- [ ] **SWG Core Member 2**: @username
- [ ] **Breaking Change Approval** (if applicable): All SWG core members
- [ ] **Security Review** (if sensitive data): @security-team

### Review Checklist
- [ ] **RFC Created**: GitHub issue with RFC label (if major change)
- [ ] **Community Feedback**: Minimum 3 business days for feedback (if breaking change)
- [ ] **Backward Compatibility**: Verified with automated compatibility check
- [ ] **Migration Guide**: Provided for any breaking changes
- [ ] **Examples Updated**: All affected examples updated and validated
- [ ] **Documentation**: Schema docs updated with changes
- [ ] **Generated Types**: TypeScript and Go types compile successfully
- [ ] **CHANGELOG**: Entry added describing changes

## Breaking Changes ⚠️

**If this PR contains breaking changes, complete this section:**

### What breaks?
<!-- Describe exactly what changes in a breaking way -->

### Migration required?
- [ ] **Yes** - Services must update to continue working
- [ ] **No** - Changes are additive/non-breaking despite marking

### Migration strategy
<!-- Describe how services should migrate -->

### Deprecation timeline
<!-- If deprecating old schema -->
- **Deprecation Version**: v___
- **Removal Version**: v___ (minimum +2 minor versions)
- **Support Period**: ___ months

### Affected stakeholders notified?
- [ ] All service teams using affected schemas
- [ ] Product teams for any feature impact
- [ ] Data teams for analytics impact

## Technical Details

### Schema Changes
<!-- Describe technical changes made to schemas -->

### Validation Updates
<!-- Any changes to validation rules or examples -->

### Generated Code Impact
<!-- How this affects TypeScript/Go generated types -->

## Testing

- [ ] **Schema Validation**: All schemas pass structural validation
- [ ] **Example Validation**: All examples validate against schemas
- [ ] **Type Generation**: TypeScript and Go types generate successfully
- [ ] **Compatibility Check**: Backward compatibility verified
- [ ] **Integration Testing**: Tested with dependent services (if applicable)

## Deployment Plan

### Rollout Strategy
- [ ] **Standard Release**: Regular release cycle
- [ ] **Coordinated Release**: Requires coordination with service deployments
- [ ] **Emergency Release**: Critical fix requiring immediate deployment

### Rollback Plan
<!-- Describe how to rollback if issues arise -->

### Monitoring
<!-- How to monitor the impact of this change -->

---

## Pre-merge Checklist

**Before requesting review:**
- [ ] All CI checks pass
- [ ] Examples are comprehensive and realistic
- [ ] Documentation is complete and accurate
- [ ] CHANGELOG entry added
- [ ] Breaking changes properly documented

**For SWG reviewers:**
- [ ] Technical accuracy verified
- [ ] Backward compatibility assessed
- [ ] Impact on downstream consumers evaluated
- [ ] Documentation quality confirmed

**For breaking changes:**
- [ ] All affected service owners acknowledged
- [ ] Migration guide tested and validated
- [ ] Deprecation timeline agreed upon
- [ ] Communication plan executed

---

## Additional Context
<!-- Add any other context about the PR here -->