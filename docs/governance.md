# Sunday Schemas Governance

This document defines the governance process for the Sunday platform event schemas and API contracts.

## Schema Working Group (SWG)

### Overview

The Schema Working Group is responsible for maintaining the integrity, evolution, and compatibility of all Sunday platform schemas. The SWG ensures that schema changes are properly reviewed, documented, and coordinated across all platform services.

### Membership

**Core Members:**
- Schema Registry Maintainer (Lead)
- Platform Architecture Lead
- Data Engineering Lead
- Backend Services Lead

**Contributing Members:**
- Service owners implementing schema-dependent features
- Data consumers (analytics, ML, reporting teams)
- Frontend teams using generated types

**Observers:**
- Security team representative
- DevOps/Infrastructure team
- Product management

### Responsibilities

#### Schema Design & Review
- Review all proposed schema changes for technical soundness
- Ensure backward compatibility requirements are met
- Validate schema changes against platform architectural principles
- Approve or reject schema modification proposals

#### Change Management
- Oversee schema versioning strategy
- Coordinate breaking changes across dependent services
- Manage deprecation timelines and migration paths
- Maintain schema evolution roadmap

#### Quality Assurance
- Ensure comprehensive example coverage for all schemas
- Validate generated types work correctly across languages
- Review documentation for completeness and accuracy
- Oversee testing of schema changes in integration environments

#### Communication & Coordination
- Communicate schema changes to all stakeholders
- Coordinate cross-team impact assessment for breaking changes
- Maintain change logs and migration guides
- Facilitate schema-related technical discussions

## Governance Process

### 1. Schema Change Proposal

All schema changes must follow this process:

#### 1.1 RFC (Request for Comments)
- Create GitHub issue using the "Schema Change RFC" template
- Describe the proposed change and rationale
- Include impact analysis on existing services
- Propose migration strategy if breaking changes are involved

#### 1.2 Community Discussion
- Allow minimum 3 business days for community feedback
- Address comments and concerns from stakeholders
- Update proposal based on feedback

#### 1.3 SWG Review
- SWG reviews technical aspects of the proposal
- Evaluate backward compatibility implications
- Assess impact on downstream consumers
- Make recommendation: Approve, Reject, or Request Changes

### 2. Implementation Process

#### 2.1 Pull Request Requirements
All schema changes must be submitted via Pull Request with:
- **Required Reviews**: Minimum 2 SWG core members must approve
- **Automated Validation**: All CI checks must pass
- **Backward Compatibility**: Compatibility check must pass or explicit breaking change approval
- **Documentation**: Updated docs and examples must be included
- **CHANGELOG**: Entry describing the change must be added

#### 2.2 Breaking Changes
Breaking changes require additional approval:
- **SWG Consensus**: All SWG core members must approve
- **Migration Plan**: Detailed migration guide must be provided
- **Overlap Period**: Old and new versions must coexist for minimum 1 minor version
- **Stakeholder Sign-off**: All affected service owners must acknowledge

#### 2.3 Review Criteria
Pull requests are evaluated on:
- **Technical Quality**: Schema follows best practices and conventions
- **Backward Compatibility**: Changes don't break existing consumers
- **Documentation**: Comprehensive docs and examples included
- **Testing**: Adequate test coverage and validation
- **Performance**: Changes don't negatively impact system performance

### 3. Schema Lifecycle

#### 3.1 Schema States
- **Draft**: Under development, not yet released
- **Active**: In production use, fully supported
- **Deprecated**: Still supported but marked for removal
- **Obsolete**: No longer supported, removed from registry

#### 3.2 Deprecation Process
1. **Deprecation Announcement** (Version N):
   - Add `x-deprecated: true` to schema metadata
   - Update documentation with deprecation notice
   - Communicate timeline to all consumers

2. **Overlap Period** (Version N+1):
   - Continue supporting deprecated schema
   - Provide migration tooling and documentation
   - Monitor usage to ensure successful migration

3. **Removal** (Version N+2):
   - Remove deprecated schema from registry
   - Archive schema documentation
   - Update validation rules

#### 3.3 Version Management
- **Major Versions**: Breaking changes, full SWG approval required
- **Minor Versions**: Backward-compatible additions, standard SWG review
- **Patch Versions**: Bug fixes and documentation updates, expedited review

### 4. Emergency Procedures

#### 4.1 Critical Security Issues
- Emergency patches can bypass normal review process
- Must be reviewed by Security team + SWG lead
- Retrospective review required within 24 hours
- Full documentation must follow within 3 business days

#### 4.2 Production Issues
- Hotfixes for production-breaking changes allowed
- Requires approval from SWG lead + affected service owner
- Must include rollback plan
- Normal review process applies post-hotfix

### 5. Communication Channels

#### 5.1 Regular Meetings
- **SWG Sync**: Weekly 30-minute meeting for active changes
- **Schema Review**: Bi-weekly review of pending proposals
- **Quarterly Planning**: Roadmap review and strategic planning

#### 5.2 Communication Tools
- **GitHub Issues**: Schema change proposals and discussions
- **GitHub Discussions**: General schema-related conversations
- **Slack #schema-working-group**: Real-time coordination
- **Email Lists**: Announcements to broader stakeholder groups

#### 5.3 Documentation
- **Public**: All governance docs, schemas, and examples on GitHub
- **Internal**: Meeting notes and private discussions in company wiki
- **Change Logs**: Comprehensive history of all schema changes

### 6. Quality Gates

#### 6.1 Automated Validation
All schema changes must pass:
- JSON Schema structural validation
- Example validation against schemas
- Backward compatibility checking
- OpenAPI specification linting
- Generated type compilation (TypeScript + Go)

#### 6.2 Manual Review Gates
- Technical accuracy review by subject matter expert
- Impact assessment by affected service owners
- Documentation review for completeness and clarity
- Security review for sensitive data schemas

#### 6.3 Release Validation
Before release to production:
- Integration testing in staging environment
- Performance regression testing
- Generated client library validation
- End-to-end schema pipeline testing

### 7. Metrics & Monitoring

#### 7.1 Governance Metrics
- Schema change approval time (target: <5 business days)
- Breaking change frequency (target: <1 per quarter)
- Community participation in RFC process
- SWG meeting attendance and participation

#### 7.2 Quality Metrics
- Schema validation pass rate (target: 100%)
- Example coverage percentage (target: 100%)
- Generated type compilation success rate (target: 100%)
- Breaking change rollback rate (target: <5%)

#### 7.3 Usage Metrics
- Schema adoption rate across services
- Deprecated schema usage trends
- Generated package download/usage statistics
- Schema evolution velocity (changes per month)

## Escalation Process

### Level 1: Technical Disputes
- Discuss in SWG meeting
- Seek consensus through technical discussion
- Document decision rationale

### Level 2: Strategic Disagreements
- Escalate to Platform Architecture Lead
- Include Product Management if product impact
- Schedule dedicated resolution meeting

### Level 3: Cross-Team Conflicts
- Escalate to Engineering Leadership
- May require executive decision
- Document resolution and update process if needed

## Process Updates

This governance process is itself subject to change through the SWG review process:

1. Propose changes via GitHub issue
2. Allow 1 week for community feedback
3. SWG review and approval required
4. Update documentation and communicate changes

## Getting Started

### For New Schema Contributors
1. Read this governance document thoroughly
2. Join #schema-working-group Slack channel
3. Attend next SWG meeting as observer
4. Review recent schema changes for examples
5. Start with small, non-breaking changes

### For Service Teams
1. Identify SWG liaison for your team
2. Subscribe to schema change notifications
3. Review schemas your services depend on
4. Participate in RFC discussions for relevant changes

### For SWG Members
1. Review SWG member handbook (internal)
2. Setup GitHub notifications for schema repository
3. Schedule recurring review time for proposals
4. Complete SWG onboarding checklist

---

*This governance framework ensures Sunday platform schemas evolve in a controlled, collaborative manner while maintaining system stability and developer productivity.*