# Schema Working Group (SWG) Handbook

This handbook provides detailed guidance for Schema Working Group members on their roles, responsibilities, and operational procedures.

## SWG Member Roles

### SWG Lead (Schema Registry Maintainer)
**Responsibilities:**
- Overall governance process oversight
- Final decision authority for contentious issues
- Release planning and coordination
- External stakeholder communication
- Process improvement and evolution

**Time Commitment:** 5-7 hours/week
**Key Skills:** Schema design, technical leadership, project management

### Core Technical Members

#### Platform Architecture Lead
- Reviews architectural alignment of schema changes
- Ensures schemas support platform scalability goals
- Validates cross-service integration patterns
- Approves technical direction for major changes

#### Data Engineering Lead
- Reviews impact on data pipelines and ETL processes
- Validates performance implications of schema changes
- Ensures compatibility with analytics and ML workflows
- Approves changes affecting data storage patterns

#### Backend Services Lead
- Reviews service integration implications
- Validates API contract changes
- Ensures backward compatibility for service communication
- Approves changes affecting service boundaries

**Time Commitment:** 3-4 hours/week per core member
**Key Skills:** Deep technical expertise in respective domains

### Contributing Members
**Responsibilities:**
- Review changes affecting their services/domains
- Provide implementation feedback
- Test changes in their environments
- Communicate impact to their teams

**Time Commitment:** 1-2 hours/week
**Key Skills:** Domain expertise, service ownership

## SWG Meetings

### Weekly SWG Sync (30 minutes)
**When:** Tuesdays 10:00 AM PT
**Attendees:** Core members + active contributors
**Agenda:**
- Active PR reviews
- Blocking issues
- Upcoming changes
- Process updates

### Bi-weekly Schema Review (60 minutes)
**When:** Thursdays 2:00 PM PT (alternating weeks)
**Attendees:** Full SWG + observers
**Agenda:**
- RFC discussions
- Breaking change reviews
- Deprecation planning
- Community feedback review

### Monthly Planning (90 minutes)
**When:** First Monday of each month, 1:00 PM PT
**Attendees:** Core members
**Agenda:**
- Quarterly roadmap review
- Process improvements
- Metrics review
- Strategic planning

### Meeting Guidelines

#### Preparation
- Review agenda 24 hours before meeting
- Read all relevant RFCs and PRs
- Prepare questions and concerns
- Update on assigned action items

#### During Meetings
- Start/end on time
- Stay focused on agenda
- Document decisions clearly
- Assign specific action items with owners

#### Follow-up
- Meeting notes published within 24 hours
- Action items tracked in GitHub projects
- Decisions communicated to broader team

## Review Process

### PR Review Workflow

#### 1. Automated Validation (Required)
All PRs must pass:
- [ ] Schema structural validation
- [ ] Example validation
- [ ] Backward compatibility check
- [ ] OpenAPI linting
- [ ] Generated type compilation
- [ ] Version overlap validation (for deprecations)

#### 2. Technical Review (Required: 2 core members)
Core members evaluate:
- **Technical Quality**: Schema design, patterns, conventions
- **Compatibility**: Backward compatibility analysis
- **Impact**: Effect on downstream consumers
- **Documentation**: Completeness and clarity

#### 3. Domain Review (Required for relevant changes)
Domain experts evaluate:
- **Service Integration**: Effect on service boundaries
- **Data Pipeline**: Impact on ETL and analytics
- **Performance**: Scalability and efficiency implications
- **Security**: Data privacy and security considerations

#### 4. Breaking Change Review (Required: All core members)
For breaking changes, additional review includes:
- **Migration Strategy**: Feasibility and completeness
- **Timeline**: Reasonable deprecation period
- **Stakeholder Impact**: Effect on all consumers
- **Risk Assessment**: Potential issues and mitigation

### Review Criteria

#### ✅ Approval Criteria
- [ ] Follows established schema patterns and conventions
- [ ] Maintains backward compatibility (or proper deprecation)
- [ ] Includes comprehensive examples and documentation
- [ ] Passes all automated validation
- [ ] Has clear business justification
- [ ] Includes adequate testing strategy

#### ❌ Rejection Criteria
- [ ] Breaks backward compatibility without proper process
- [ ] Introduces significant performance regression
- [ ] Lacks adequate documentation or examples
- [ ] Conflicts with platform architectural principles
- [ ] Insufficient stakeholder review for breaking changes
- [ ] Security or privacy concerns not addressed

#### 🔄 Revision Request Criteria
- [ ] Technical implementation needs refinement
- [ ] Documentation incomplete or unclear
- [ ] Migration strategy needs more detail
- [ ] Additional stakeholder review required
- [ ] Test coverage insufficient

## Decision Making

### Consensus Building
1. **Technical Discussion**: Open technical debate with data-driven arguments
2. **Stakeholder Input**: Gather feedback from affected service teams
3. **Risk Assessment**: Evaluate potential negative impacts
4. **Alternative Evaluation**: Consider multiple approaches
5. **Consensus Check**: Seek agreement from core members

### Conflict Resolution

#### Level 1: Technical Disagreement
- Schedule dedicated discussion session
- Bring in subject matter experts
- Review similar patterns from other systems
- Document pros/cons of each approach
- Seek compromise solution

#### Level 2: Strategic Differences
- Escalate to SWG Lead + Platform Architecture Lead
- Include Product Management if product impact
- Schedule longer resolution session
- May require executive input
- Document decision rationale

#### Level 3: Deadlock
- Escalate to Engineering Leadership
- Present all perspectives clearly
- Accept final decision from leadership
- Update process to prevent future deadlock

### Voting (Last Resort)
When consensus cannot be reached:
- **Quorum**: Minimum 3 core members must be present
- **Majority**: Simple majority of present core members
- **Tie Breaking**: SWG Lead has deciding vote
- **Documentation**: Record all votes and reasoning

## Quality Assurance

### Pre-Review Checklist
**For PR Authors:**
- [ ] All automated checks pass
- [ ] Examples comprehensive and realistic
- [ ] Documentation complete and accurate
- [ ] Migration guide provided (if breaking change)
- [ ] Stakeholders notified (if breaking change)

**For Reviewers:**
- [ ] Technical accuracy verified
- [ ] Compatibility implications understood
- [ ] Documentation quality confirmed
- [ ] Test strategy validated

### Review Quality Standards

#### Technical Review
- **Depth**: Thorough analysis of schema design
- **Breadth**: Consider all downstream implications
- **Future-proofing**: Evaluate long-term maintainability
- **Performance**: Assess efficiency implications

#### Documentation Review
- **Completeness**: All required documentation present
- **Clarity**: Easy to understand for target audience
- **Accuracy**: Technical details are correct
- **Discoverability**: Well-organized and searchable

### Metrics and Feedback

#### Review Metrics (Tracked Monthly)
- Average review time (target: <5 business days)
- Review quality scores (post-implementation feedback)
- Breaking change frequency (target: <1 per quarter)
- Community satisfaction scores

#### Continuous Improvement
- Quarterly process retrospectives
- Annual SWG effectiveness survey
- Process updates based on feedback
- Training and skill development

## Communication Guidelines

### Internal Communication

#### Slack Channels
- **#schema-working-group**: SWG member coordination
- **#schema-announcements**: Broader team notifications
- **#schema-help**: Community questions and support

#### Email Lists
- **swg-core@sunday.com**: Core member discussions
- **schema-stakeholders@sunday.com**: Stakeholder announcements
- **engineering-all@sunday.com**: Major breaking changes

#### GitHub
- **Issues**: RFC discussions, bug reports
- **Discussions**: General schema-related conversations
- **PRs**: Code review and technical discussion

### External Communication

#### Documentation
- Keep public documentation up to date
- Ensure examples are current and accurate
- Provide clear migration guides
- Maintain comprehensive changelog

#### Community Engagement
- Respond to questions promptly
- Participate in industry schema discussions
- Share best practices and learnings
- Gather feedback from schema consumers

## Emergency Procedures

### Security Issues
**Process:**
1. **Immediate Assessment**: Security team + SWG Lead review
2. **Emergency Fix**: Bypass normal review if critical
3. **Communication**: Limited disclosure until fix deployed
4. **Retrospective**: Full review within 48 hours
5. **Process Update**: Update procedures if needed

### Production Outages
**Process:**
1. **Incident Response**: Follow standard incident procedures
2. **Schema Review**: SWG Lead + affected service owner
3. **Hotfix Authorization**: Can bypass review for critical fixes
4. **Post-Incident**: Normal review process applies after resolution
5. **Learning**: Incorporate lessons into future reviews

### Escalation Paths
- **Technical Issues**: SWG Lead → Platform Architecture Lead
- **Timeline Issues**: SWG Lead → Engineering Management
- **Resource Issues**: SWG Lead → Resource allocation team
- **External Issues**: SWG Lead → External partnerships team

## Onboarding New Members

### Core Member Onboarding (2 weeks)

#### Week 1: Foundation
- [ ] Read all governance documentation
- [ ] Review recent schema changes (last 3 months)
- [ ] Attend SWG meetings as observer
- [ ] Complete schema design training
- [ ] Setup development environment

#### Week 2: Practice
- [ ] Review 2-3 non-critical PRs with mentor
- [ ] Participate in RFC discussion
- [ ] Shadow experienced reviewer
- [ ] Present to SWG on area of expertise
- [ ] Complete onboarding feedback survey

### Contributing Member Onboarding (1 week)
- [ ] Read governance overview
- [ ] Join relevant Slack channels
- [ ] Review schemas affecting your domain
- [ ] Attend one SWG meeting
- [ ] Connect with SWG liaison for your team

### Ongoing Development

#### Training Opportunities
- **Schema Design Patterns**: Monthly workshops
- **Backward Compatibility**: Quarterly deep dives
- **Industry Best Practices**: External conference attendance
- **Tool Training**: Hands-on sessions with new tools

#### Knowledge Sharing
- **Lightning Talks**: 10-minute technical presentations
- **Case Studies**: Deep dives on interesting challenges
- **Tool Demos**: Show and tell for new tools/techniques
- **External Learning**: Share insights from conferences/reading

## Performance Standards

### Individual Performance

#### Core Members
- **Review Timeliness**: 95% of reviews completed within SLA
- **Review Quality**: High post-implementation satisfaction scores
- **Meeting Participation**: 90% attendance at required meetings
- **Communication**: Prompt response to requests and questions

#### Contributing Members
- **Domain Expertise**: Provide valuable domain-specific insights
- **Responsiveness**: Participate in relevant reviews promptly
- **Team Communication**: Keep team informed of schema changes
- **Feedback Quality**: Provide constructive, actionable feedback

### Team Performance

#### SWG Effectiveness
- **Decision Speed**: Average time from RFC to decision
- **Quality Outcomes**: Post-implementation issue rate
- **Stakeholder Satisfaction**: Regular feedback surveys
- **Process Efficiency**: Time spent on governance vs. development

#### Community Health
- **Participation**: Number of active contributors
- **Question Response**: Time to answer community questions
- **Documentation Quality**: User feedback on documentation
- **Adoption Success**: Migration success rates

## Offboarding

### Planned Departure
1. **Transition Planning**: 4-week notice preferred
2. **Knowledge Transfer**: Document specialized knowledge
3. **Review Handoff**: Transfer ongoing reviews to other members
4. **Replacement**: Identify and train replacement
5. **Access Removal**: Revoke access after transition complete

### Emergency Departure
1. **Immediate Coverage**: Redistribute responsibilities immediately
2. **Access Review**: Immediate access audit and revocation
3. **Knowledge Recovery**: Gather knowledge from documentation/colleagues
4. **Accelerated Replacement**: Fast-track replacement process

---

*This handbook ensures SWG members have clear guidance on their roles and responsibilities while maintaining high standards for schema governance.*