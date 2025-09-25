# Sunday-Schemas Implementation Plan

## Overview

This document outlines the implementation plan for the Sunday-schemas repository, which serves as the single source of truth for event schemas and API contracts across the Sunday platform. The plan is organized into phases with clear dependencies and detailed technical tasks.

## Project Scope

**Deliverables:**
- JSON Schemas for all Kafka event types
- OpenAPI specification for UI BFF endpoints
- Generated TypeScript and Go types
- CI/CD pipeline with compatibility gates
- Documentation and governance processes

**Timeline:** 32 days across 6 sprints
**Team:** Schema Working Group (UI, Platform, Data engineers)

## Phase Index

| Phase | Name | Duration | Key Deliverables |
|-------|------|----------|------------------|
| [Phase 1](#phase-1-foundation-setup-5-days) | Foundation Setup | 5 days | Repository structure, registries, tooling |
| [Phase 2](#phase-2-schema-definitions-8-days) | Schema Definitions | 8 days | All JSON schemas (raw, md, insights, infra) |
| [Phase 3](#phase-3-openapi-specification-3-days) | OpenAPI Specification | 3 days | UI BFF API specification |
| [Phase 4](#phase-4-validation--tooling-6-days) | Validation & Tooling | 6 days | CI validation, compatibility gates |
| [Phase 5](#phase-5-documentation-2-days) | Documentation | 2 days | Mapping docs, example maintenance |
| [Phase 6](#phase-6-code-generation--distribution-6-days) | Code Generation & Distribution | 6 days | TypeScript/Go generation, publishing |
| [Phase 7](#phase-7-governance-process-2-days) | Governance Process | 2 days | Schema Working Group processes |

**Total Duration:** 32 days

---

## Phase 1: Foundation Setup (5 days)

### Phase 1a: Repository Structure Setup (1 day)

**Tasks:**
- Create directory structure per tech spec section 3
- Initialize package.json with required dependencies
- Setup basic README.md and CHANGELOG.md
- Create .gitignore

**Commands:**
```bash
mkdir -p {schemas/{json,examples,registries},openapi,codegen/{ts,go},docs,scripts}
npm init -y
npm install --save-dev ajv spectral @stoplight/spectral-cli jsonlint js-yaml
npm install --save-dev json-schema-to-typescript openapi-typescript
npm install --save-dev json-schema-diff openapi-diff
```

### Phase 1b: Registry Files (2 days) - **Card 113**

**Story:** "Maintain topics.json and registries"

**Technical Tasks:**
- Create `schemas/registries/venues.json` with `["polymarket", "kalshi"]`
- Create `schemas/topics.json` with schema→topic mappings (tech spec 3.1)
- Create placeholder `schemas/registries/instruments.json`
- Add validation script to ensure venue enums match registry

**Acceptance Criteria:**
- topics.json maps schemas → topics
- venues.json lists venues
- instruments.json seeded

### Phase 1c: Initial Tooling Setup (2 days)

**Tasks:**
- Setup basic validation scripts
- Configure linting tools
- Create initial test structure

---

## Phase 2: Schema Definitions (8 days)

*Dependencies: Phase 1 must be complete*

### Phase 2a: Raw Schema (2 days) - **Card 100**

**Story:** "Define raw.v0 envelope schema"
**Persona:** Connector engineer

**Technical Tasks:**
- Implement exact schema from tech spec 4.1 (lines 93-115)
- Add $id: `https://schemas.sunday.dev/raw.v0.envelope.schema.json`
- Enum values must reference `venues.json` registry
- Create 2+ examples with different venues and streams
- Add validation test ensuring examples pass ajv

**Acceptance Criteria:**
- Schema defines venue_id, stream, instrument_native, ts_event_ms, ts_ingest_ms, payload
- Examples validate

### Phase 2b: Market Data Schemas (3 days) - **Cards 101, 102**

#### Card 101: Define md.orderbook.delta schema (2 days)
**Story:** "Define md.orderbook.delta schema"
**Persona:** Data engineer

**Technical Tasks:**
- Implement schema from tech spec 4.2 (lines 146-172)
- Document monotonic seq requirement in $comment
- Add is_snapshot field with default false
- Create example showing snapshot recovery scenario
- Create example showing normal delta update
- Add test for bids/asks array structure [price, size]

**Acceptance Criteria:**
- Schema defines bids, asks, seq, ts_ms, is_snapshot
- Monotonic seq rule documented
- Examples validate

#### Card 102: Define md.trade schema (1 day)
**Story:** "Define md.trade schema"
**Persona:** Data engineer

**Technical Tasks:**
- **FIRST:** Resolve prob vs price_prob naming inconsistency in tech spec
- Implement schema from tech spec 4.3 with consistent field name
- Add prob range validation [0.0, 1.0]
- Create examples for both buy/sell sides
- Add notional_usd as optional field

**Acceptance Criteria:**
- Schema defines side, prob, size, ts_ms
- Examples validate

### Phase 2c: Insights Schemas (2 days) - **Card 103**

**Story:** "Define insights schemas (arb, movers, whales, unusual)"
**Persona:** PM

**Technical Tasks:**
- Implement 4 separate schemas from tech spec 4.4-4.7:
  - `insights.arb.lite.v1.schema.json`
  - `insights.movers.v1.schema.json`
  - `insights.whales.lite.v1.schema.json`
  - `insights.unusual.v1.schema.json`
- Each needs depth_tier, window, or impact enums as specified
- Create realistic examples with proper instrument_id format
- Add edge_bps, delta_bps, zscore numeric validations
- Test persistence_ms, last_seen_ms timestamp fields

**Acceptance Criteria:**
- Schemas for arb, movers, whales, unusual defined with enums + required fields
- Examples validate

### Phase 2d: Infrastructure Schema (1 day) - **Card 104**

**Story:** "Define infra.venue_health schema"
**Persona:** Operator

**Technical Tasks:**
- Implement schema from tech spec 4.8
- Add status enum: `["CONNECTED", "DEGRADED", "STALE"]`
- Create examples for each status type
- Add staleness_seconds and messages_per_second metrics
- Test required vs optional fields

**Acceptance Criteria:**
- Schema defines status, last_event_ts_ms, staleness
- Examples validate

---

## Phase 3: OpenAPI Specification (3 days)

*Dependencies: Phase 2 must be complete (schemas referenced in OpenAPI)*

### Phase 3a: OpenAPI Spec (3 days) - **Card 105**

**Story:** "Publish ui.v1.yaml OpenAPI spec"
**Persona:** UI engineer

**Technical Tasks:**
- Create OpenAPI 3.1 spec with 7 endpoints from tech spec section 5:
  - `GET /markets` → `MarketList`
  - `GET /arb` → array of `ArbLite`
  - `GET /movers` → array of `Mover`
  - `GET /flows` → array of `WhaleLite`
  - `GET /unusual` → array of `Unusual`
  - `GET /calendar` → array of `ResolutionEvent`
  - `GET /venue-health` → array of `VenueHealth`
- Reuse schema components from JSON schemas where possible
- Add query parameters: min_edge_bps, venues, window, limit
- Create MarketList, ResolutionEvent components
- Add realistic response examples
- Validate with spectral linter

**Acceptance Criteria:**
- Endpoints for /markets, /arb, /movers, /flows, /unusual, /calendar, /venue-health defined
- Components reuse schemas

---

## Phase 4: Validation & Tooling (6 days)

*Dependencies: Phases 2-3 must be complete*

### Phase 4a: Validation Setup (3 days) - **Card 109**

**Story:** "Setup validation and linting"
**Persona:** Schema maintainer

**Technical Tasks:**
- Setup package.json scripts for validation:
```json
{
  "validate-schemas": "ajv compile -s 'schemas/json/*.json'",
  "validate-examples": "node scripts/validate-examples.js",
  "lint-openapi": "spectral lint openapi/ui.v1.yaml",
  "lint-json": "jsonlint schemas/**/*.json",
  "test": "npm run validate-schemas && npm run validate-examples && npm run lint-openapi"
}
```
- Configure `ajv` for JSON Schema validation against examples
- Configure `spectral` for OpenAPI linting
- Configure `jsonlint`/`yamllint` for structure checks
- Create validation test suite covering all examples
- Setup pre-commit hooks for validation

**Acceptance Criteria:**
- ajv validates JSON schemas
- spectral lints OpenAPI
- yamllint/jsonlint run in CI

### Phase 4b: Compatibility Gates (3 days) - **Card 110**

**Story:** "Setup compatibility gates"
**Persona:** Schema maintainer

**Technical Tasks:**
- Create `.github/workflows/compatibility.yml`
- Add json-schema-diff check for breaking changes
- Add openapi-diff check for API compatibility
- Create topics.json guard preventing topic remapping
- Create venues.json guard allowing only additions
- Add CHANGELOG.md requirement check
- Block PRs that fail compatibility tests

**Acceptance Criteria:**
- json-schema-diff + oas-diff run in CI
- topics.json and venues.json guarded
- PRs require CHANGELOG entries

---

## Phase 5: Documentation (2 days)

*Dependencies: Phases 2-4 complete*

### Phase 5a: Mapping Documentation (1 day) - **Card 114**

**Story:** "Maintain mapping docs"
**Persona:** Developer

**Technical Tasks:**
- Create `docs/mapping.md` with raw→normalized transformation rules
- Document instrument_id conventions
- Document YES/NO probability conventions

**Acceptance Criteria:**
- Rules for symbol → instrument_id, YES/NO conventions documented

### Phase 5b: Example Maintenance (1 day) - **Card 112**

**Story:** "Maintain examples"
**Persona:** QA

**Technical Tasks:**
- Ensure all examples validate (automated in CI)
- Create comprehensive example coverage
- Document example maintenance process

**Acceptance Criteria:**
- Examples provided for each schema
- CI validates all examples

---

## Phase 6: Code Generation & Distribution (6 days)

*Dependencies: Phases 2-5 must be complete and validated*

### Phase 6a: TypeScript Generation (2 days) - **Card 106**

**Story:** "Generate TypeScript types"
**Persona:** UI engineer

**Technical Tasks:**
- Setup `json-schema-to-typescript` for event schemas
- Setup `openapi-typescript` for API types
- Configure output to `codegen/ts/index.d.ts`
- Create npm package structure for `@sunday/schemas`

**Acceptance Criteria:**
- json-schema-to-typescript generates d.ts files
- Published as @sunday/schemas

### Phase 6b: Go Generation (2 days) - **Card 107**

**Story:** "Generate Go types"
**Persona:** Data engineer

**Technical Tasks:**
- Setup `oapi-codegen` for OpenAPI client/server types
- Setup `quicktype` for JSON Schema structs
- Configure output to `go/` module structure
- Setup Go module `github.com/sunday-xyz/schemas/go`

**Acceptance Criteria:**
- oapi-codegen and quicktype generate Go structs
- Published as github.com/sunday-xyz/schemas/go

### Phase 6c: Publishing Automation (2 days) - **Card 108**

**Story:** "Automate publishing on tag"
**Persona:** Developer

**Technical Tasks:**
- Create GitHub Actions workflow for tag-triggered publishing
- Configure npm publishing to `@sunday/schemas`
- Setup goreleaser for Go module publishing
- Include SBOM and checksum generation

**Acceptance Criteria:**
- GitHub Actions workflow publishes artifacts
- SBOM/checksums included

---

## Phase 7: Governance Process (2 days)

*Dependencies: All previous phases complete*

### Phase 7a: Governance Setup (2 days) - **Card 111**

**Story:** "Define governance process"
**Persona:** Schema working group

**Technical Tasks:**
- Define Schema Working Group process
- Create PR templates requiring WG review
- Document deprecation process with `x-deprecated` markers
- Setup minor version overlap requirements

**Acceptance Criteria:**
- WG review required for schema PRs
- Deprecations marked with x-deprecated/$comment
- Minor version overlap maintained

---

## Sprint Organization

### Sprint 1 (Week 1): Foundation + Raw Schema
- **Goals:** Phase 1 + Card 100
- **Deliverables:** Basic repository structure, raw schema, validation working
- **Risk Mitigation:** Get basic validation pipeline working early

### Sprint 2 (Week 2): Core Schemas
- **Goals:** Cards 101-104 (all event schemas)
- **Deliverables:** Complete schema definition backbone
- **Risk Mitigation:** Resolve prob/price_prob inconsistency early

### Sprint 3 (Week 3): OpenAPI + Validation
- **Goals:** Cards 105, 109
- **Deliverables:** Complete validation suite, OpenAPI spec
- **Risk Mitigation:** Ensure all schemas validate before proceeding

### Sprint 4 (Week 4): CI + Documentation
- **Goals:** Cards 110, 112, 114
- **Deliverables:** Compatibility gates working, documentation complete
- **Risk Mitigation:** Test compatibility tools with realistic schema changes

### Sprint 5 (Week 5): Code Generation
- **Goals:** Cards 106-108
- **Deliverables:** Complete publishing pipeline
- **Risk Mitigation:** Validate generated types compile in downstream repos

### Sprint 6 (3 days): Governance + Polish
- **Goals:** Card 111, final integration testing
- **Deliverables:** Complete governance process
- **Risk Mitigation:** Get Schema Working Group aligned early

---

## Critical Path Dependencies

1. **Venues.json** → All schemas (enum validation)
2. **All schemas** → OpenAPI (component reuse)
3. **Schemas + OpenAPI** → Validation setup
4. **Validation** → Compatibility gates
5. **Everything** → Code generation + publishing

---

## Risk Mitigation

### Technical Risks
- **Issue:** prob/price_prob inconsistency in tech spec
  - **Mitigation:** Resolve in Sprint 1 before implementing trade schema
- **Issue:** Compatibility tools may not work as expected
  - **Mitigation:** Test with realistic schema changes in Sprint 4
- **Issue:** Generated types may not compile in downstream repos
  - **Mitigation:** Test integration early, include in acceptance criteria

### Process Risks
- **Issue:** Schema Working Group alignment on governance
  - **Mitigation:** Get alignment in Phase 1, document decisions
- **Issue:** Publishing pipeline failures
  - **Mitigation:** Test in development environment first
- **Issue:** CI gates blocking legitimate changes
  - **Mitigation:** Test compatibility detection with known good/bad changes

---

## Success Metrics

### Phase 1 Acceptance Criteria (Tech Spec Section 11)
1. ✅ All schemas validate with examples; npm & Go packages publish on tag
2. ✅ OpenAPI `/ui/*` compiles; mock server (prism) can serve sample data
3. ✅ Downstream repos compile against `@sunday/schemas` and Go module at `v1.0.0`
4. ✅ CI gates block incompatible schema changes; CHANGELOG documents every release

### Additional Success Metrics
- Zero schema validation failures in CI
- All 14 Trello stories completed with acceptance criteria met
- Documentation complete and accessible
- Schema Working Group process operational
- Automated publishing working end-to-end

---

## Next Steps

1. **Immediate (Today):**
   - Resolve prob vs price_prob field naming inconsistency in tech spec
   - Set up development environment

2. **Sprint 1 Planning:**
   - Create Sprint 1 backlog with Phase 1 + Card 100 tasks
   - Assign team members to tasks
   - Setup project tracking

3. **Long-term:**
   - Schedule Sprint planning sessions
   - Establish Schema Working Group meeting cadence
   - Plan downstream integration testing