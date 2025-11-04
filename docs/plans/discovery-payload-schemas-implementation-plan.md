# Discovery Payload Schemas Implementation Plan

## Executive Summary

This document outlines the comprehensive implementation plan for adding structured payload schemas to `sunday-schemas` for discovery events, as specified in `docs/04112025/discovery-payload-schemas-specification.md`.

**Objective**: Migrate from untyped `map[string]interface{}` discovery payloads to strongly-typed JSON schemas with generated Go types, ensuring type safety and cross-service consistency.

**Timeline**: 4-5 weeks (sunday-schemas scope only)
**Version Target**: v1.1.0 (additive changes)
**Risk Level**: Low (schema additions only)

## Phase Overview

| Phase | Duration | Key Deliverables | Dependencies |
|-------|----------|------------------|--------------|
| **Phase 1**: Schema Design & Implementation | 2-3 weeks | JSON schemas, Go type generation, validation | Repository access |
| **Phase 2**: Testing & Validation | 1-2 weeks | Test suite, validation framework, examples | Phase 1 complete |

## Phase 1: Schema Design & Implementation (Weeks 1-3)

### 1.1 Schema File Creation

#### Core Schema Files
```
schemas/json/
├── discovery.event-payload.v0.schema.json
├── discovery.series-payload.v0.schema.json
├── discovery.event-metadata.v0.schema.json
├── discovery.series-metadata.v0.schema.json
└── discovery.shared.v0.schema.json
```

#### Deliverables
- [ ] Create `discovery.event-payload.v0.schema.json`
  - EventDiscoveryPayloadV0 schema definition
  - Required fields: event, event_id, event_type, timestamp, venue_id
  - event_type enum: ["discovered", "updated", "expired"]
  - venue_id enum: ["polymarket", "kalshi"] (from venues.json)
  - Optional discovery_meta field

- [ ] Create `discovery.series-payload.v0.schema.json`
  - SeriesDiscoveryPayloadV0 schema definition
  - Same structure as event payload but references SeriesMetadataV0
  - Same enum constraints for event_type and venue_id

- [ ] Create `discovery.event-metadata.v0.schema.json`
  - EventMetadataV0 schema definition
  - kind const: "event" (required)
  - venue_id enum: ["polymarket", "kalshi"] (required)
  - Required fields: kind, venue_id, event_id, title, active, closed, discovered_at, last_seen
  - Parent series relationship fields

- [ ] Create `discovery.series-metadata.v0.schema.json`
  - SeriesMetadataV0 schema definition
  - kind const: "series" (required)
  - venue_id enum: ["polymarket", "kalshi"] (required)
  - Required fields: kind, venue_id, event_id, title, active, closed, discovered_at, last_seen
  - SeriesDataV0 reference for series-specific fields

- [ ] Create `discovery.shared.v0.schema.json`
  - DiscoveryMetaV0 with required fields: batch_id, batch_sequence, batch_total_count, discovery_run_id
  - batch_sequence minimum: 1, batch_total_count minimum: 1
  - RelationshipsV0 (parent/child mappings)
  - FinancialDataV0 with numeric constraints:
    - volume_24h_usd: multipleOf 0.01, minimum 0
    - volume_total_usd: multipleOf 0.01, minimum 0
    - liquidity_total_usd: multipleOf 0.01, minimum 0
    - volume_24h_contracts: integer, minimum 0
    - volume_total_contracts: integer, minimum 0
    - score: minimum 0
    - currency enum: ["USD"]
  - StatusDataV0 (archived, featured, restricted flags)
  - ContractDataV0 (Kalshi contract information)
  - SettlementSourceV0 with required field: name
  - TimestampDataV0 (created_at, updated_at, published_at)
  - CreatorDataV0 (created_by, updated_by)

#### Schema Design Principles
- Use JSON Schema draft/2020-12 specification
- Reference existing venue enums from `schemas/registries/venues.json`
- Maintain backward compatibility with existing envelope structures
- Use `additionalProperties: false` for strict validation where appropriate
- Include comprehensive field descriptions and validation rules

### 1.2 Schema Registry Integration

#### Update Existing Schemas
- [ ] Modify `raw.events.discovery.v0.schema.json`
  - Update payload property to reference EventDiscoveryPayloadV0
  - Maintain backward compatibility with map[string]interface{}

- [ ] Modify `raw.series.discovery.v0.schema.json`
  - Update payload property to reference SeriesDiscoveryPayloadV0
  - Maintain backward compatibility with map[string]interface{}

#### Schema Subjects Registration
- [ ] Update existing schema subjects (NO new subjects):
  - Keep existing `raw.events.discovery.v0-value` subject
  - Keep existing `raw.series.discovery.v0-value` subject
  - Update envelope schemas to reference new payload definitions
  - Maintain backward compatibility with existing consumers

### 1.3 Go Type Generation

#### Code Generation Setup
- [ ] Configure `json-schema-to-typescript` for TypeScript types
- [ ] Configure `quicktype` or custom tooling for Go type generation
- [ ] Set up proper JSON tags for Go structs
- [ ] Generate enum constants for DiscoveryKind, EventType, VenueID

#### Generated Type Structure
```
codegen/go/
├── discovery/
│   ├── payloads.go          # EventDiscoveryPayloadV0, SeriesDiscoveryPayloadV0
│   ├── metadata.go          # EventMetadataV0, SeriesMetadataV0
│   ├── shared.go            # Supporting data structures
│   └── enums.go             # DiscoveryKind, EventType constants
```

#### Type Generation Requirements
- [ ] Proper time.Time handling for date-time fields
- [ ] Pointer types for optional fields (*string, *bool, *float64)
- [ ] Struct tags for JSON marshaling/unmarshaling
- [ ] Validation tags for runtime validation
- [ ] Embedded documentation from schema descriptions

### 1.4 Validation Framework

#### Runtime Validation
- [ ] JSON Schema validation using `github.com/santhosh-tekuri/jsonschema`
- [ ] Cross-field validation (e.g., kind="series" requires series_data)
- [ ] Venue-specific field population rules
- [ ] Custom validation for enum values

#### Validation Functions
- [ ] `ValidateEventDiscoveryPayload(payload []byte) error`
- [ ] `ValidateSeriesDiscoveryPayload(payload []byte) error`
- [ ] `ValidateEventMetadata(metadata EventMetadataV0) error`
- [ ] `ValidateSeriesMetadata(metadata SeriesMetadataV0) error`

## Phase 2: Testing & Validation (Weeks 4-5)

### 2.1 Schema Validation Tests

#### Test Data Creation
- [ ] Create valid payload examples for both Polymarket and Kalshi
- [ ] Create invalid payload examples for negative testing
- [ ] Generate edge case examples (minimal required fields, maximum data)
- [ ] Create examples directory: `schemas/examples/discovery/`

#### Test Implementation
- [ ] Unit tests for schema validation
- [ ] Property-based testing for schema compliance
- [ ] Cross-validation between JSON Schema and Go types
- [ ] Performance benchmarks for validation overhead

### 2.2 Type Safety Tests

#### Go Type Tests
- [ ] JSON marshaling/unmarshaling round-trip tests
- [ ] Type assertion tests for generated structs
- [ ] Enum value validation tests
- [ ] Optional field handling tests

#### Integration Tests
- [ ] Mock Kafka message serialization/deserialization
- [ ] Backward compatibility with existing untyped payloads
- [ ] Forward compatibility for schema evolution

### 2.3 Field Mapping Validation

#### Venue Mapping Tests
- [ ] Polymarket field mapping validation
- [ ] Kalshi field mapping validation
- [ ] Cross-venue consistency checks
- [ ] Extra metadata handling verification

#### Data Transformation Tests
- [ ] Test transformation from venue APIs to canonical schema
- [ ] Validate relationship mapping (parent/child events)
- [ ] Test venue-specific field population

## Technical Requirements

### Development Environment

#### Required Tools
- JSON Schema validation: `ajv-cli` or `jsonschema`
- Go type generation: `quicktype` or custom tooling
- TypeScript generation: `json-schema-to-typescript`
- Testing framework: Go testing + `testify`
- Schema diff tools: `json-schema-diff`

#### Repository Setup
- [ ] Configure pre-commit hooks for schema validation
- [ ] Set up CI/CD pipeline for type generation
- [ ] Implement automated testing for all schema changes
- [ ] Configure dependency management for generated types

### Schema Evolution Strategy

#### Versioning Policy
- V0 → V1: Additive changes only (new optional fields)
- V1 → V2: Breaking changes allowed with migration path
- Schema ID format: `discovery.{type}.{version}.schema.json`
- Generated type suffix: `V0`, `V1`, etc.

#### Change Management
- [ ] Schema Working Group review process
- [ ] Backward compatibility validation
- [ ] Impact assessment for downstream consumers
- [ ] Coordinated rollout with affected services

### Performance Considerations

#### Optimization Targets
- Schema validation overhead: < 5ms per message
- Type conversion performance: < 1ms per message
- Memory overhead: < 10% increase for typed payloads
- Serialization size: No significant increase

#### Monitoring Metrics
- [ ] Validation performance benchmarks
- [ ] Memory usage tracking
- [ ] Error rate monitoring
- [ ] Throughput impact assessment

## Risk Management

### Technical Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|---------|------------|
| Schema validation performance | Medium | High | Benchmark early, optimize validation logic |
| Type generation issues | Medium | Medium | Custom tooling, manual verification |
| Schema evolution compatibility | Low | High | Comprehensive testing, backward compatibility validation |

### Schema Registry Risks

- [ ] Backup existing schemas before modifications
- [ ] Test schema evolution in staging environment
- [ ] Monitor schema registry performance impact
- [ ] Plan for schema registry failover scenarios

## Success Metrics

### Technical Metrics

#### Schema Implementation
- [ ] All discovery payload schemas defined with proper constraints
- [ ] Generated Go types match JSON schema exactly
- [ ] Complete validation framework implemented
- [ ] Test coverage > 95% for all schemas

#### Schema Quality
- [ ] All schema constraints properly enforced (enums, bounds, required fields)
- [ ] Backward compatibility with existing envelope structures
- [ ] Comprehensive documentation and examples
- [ ] Performance benchmarks within targets

### Performance Metrics

#### Validation Performance
- [ ] Schema validation latency < 5ms (95th percentile)
- [ ] Type generation completes in < 30 seconds
- [ ] Test suite runs in < 2 minutes
- [ ] No memory leaks in validation functions

## Dependencies & Prerequisites

### Technical Dependencies

#### Repository Access
- [ ] Write access to `sunday-schemas` repository
- [ ] CI/CD pipeline configuration access
- [ ] Package publishing permissions (npm, Go modules)

#### Tool Dependencies
- [ ] JSON Schema validation tools (ajv, jsonschema)
- [ ] Code generation tools (quicktype, json-schema-to-typescript)
- [ ] Testing frameworks (Go testing, Jest for TS)
- [ ] Schema diff and validation utilities

### Organizational Dependencies

#### Team Coordination
- [ ] Schema Working Group approval for changes
- [ ] DevOps team support for CI/CD changes

## Communication Plan

### Stakeholder Communication

#### Weekly Status Updates
- [ ] Progress against phase milestones
- [ ] Risk assessment and mitigation status
- [ ] Blockers and resolution timeline

#### Technical Reviews
- [ ] Schema design review with Working Group
- [ ] Code review process for all changes
- [ ] Performance impact assessment review

### Documentation Strategy

#### Technical Documentation
- [ ] Schema reference with examples
- [ ] Testing and validation procedures
- [ ] Type generation and usage documentation

## Conclusion

This implementation plan provides a structured approach to implementing strongly-typed discovery payload schemas in sunday-schemas. The 2-phase approach focuses on schema creation and thorough testing while maintaining backward compatibility.

The success of this implementation depends on thorough testing, proper validation framework implementation, and clear documentation. The resulting typed schema system will provide a solid foundation for type-safe discovery event processing across the Sunday platform.

---

**Document Version**: 1.0
**Created**: 2025-11-04
**Author**: sunday-schemas implementation team
**Review Date**: Weekly during implementation
**Approval Required**: Schema Working Group