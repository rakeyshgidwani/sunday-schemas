# Sunday-schemas — Product Specification (Phase 1)

## 1. Product Overview
Sunday-schemas is the source of truth for event schemas and API contracts.  
It defines versioned Kafka event schemas, the UI BFF OpenAPI spec, and publishes  
generated types for TypeScript and Go. It is an internal developer-facing product.

## 2. Goals & Objectives
- Consistency: ensure all repos use identical schema definitions.  
- Reliability: CI prevents accidental breaking changes.  
- Developer velocity: provide ready-to-use TS/Go types via npm + Go modules.  
- Traceability: versioned schema definitions with explicit changelog and deprecation policy.  

## 3. Target Users
- Connector engineers (produce raw events into Kafka).  
- Data pipeline engineers (consume raw → normalized → insights).  
- Platform engineers (BFF APIs typed against schemas).  
- UI engineers (TypeScript types generated directly from schemas).  

## 4. Key Features
- Event Schemas: raw envelope, normalized market data, insights (lite), and infra health.  
- Registries & Topic Map: topics.json, venues.json, instruments.json.  
- OpenAPI: /ui/* endpoints defined in ui.v1.yaml.  
- Code Generation & Distribution: TypeScript npm package + Go module.  
- Governance & CI Gates: semantic versioning, compatibility enforcement, schema working group review.  

## 5. User Flows
- Schema Consumption (TS): UI engineers import generated types from @sunday/schemas.  
- Schema Consumption (Go): Data engineers use generated structs from Go module.  
- Governance Flow: Engineer proposes schema change → PR with schema + example → CI validation → WG review → tag + publish.  

## 6. Functional Requirements
- All schemas validate against examples in /schemas/examples.  
- CI prevents breaking changes in Phase 1.  
- npm + Go modules publish automatically on tag.  
- README + mapping.md describe schema roles and rules.  

## 7. Acceptance Criteria
- All schemas and examples validate under CI.  
- @sunday/schemas (TS) and Go module publish at v1.0.0.  
- Downstream repos compile successfully using generated types.  
- OpenAPI spec validates and compiles into mock server.  
- Compatibility gates block any backward-incompatible changes.  

## 8. Non-Goals (Phase 1)
- No runtime services (schemas repo is not a service).  
- No custom code generation beyond TS/Go.  
- No persistence or analytics.  

## 9. Roadmap Notes (Beyond Phase 1)
- Add support for more venues (e.g., Manifold) via minor version bumps.  
- Expand to more language bindings (Python, Rust).  
- Support schema evolution with major-version bumps and migration notes.  
- Optional schema registry service (Phase 2+).  

## Appendix A — Repository Mapping
- /schemas/json/ → Event schemas.  
- /schemas/examples/ → Golden sample payloads.  
- /schemas/registries/ → venues.json, instruments.json.  
- /schemas/topics.json → schema ↔ topic map.  
- /openapi/ui.v1.yaml → UI BFF contract.  
- /codegen/ → TS + Go generation.  
- /docs/ → mapping rules, governance.  
