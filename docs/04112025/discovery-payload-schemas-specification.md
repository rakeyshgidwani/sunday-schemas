# Discovery Payload Schemas Specification for sunday-schemas

## Executive Summary

This specification requests the addition of structured payload schemas to `sunday-schemas` for discovery events. Currently, discovery messages use typed envelopes but untyped `map[string]interface{}` payloads, creating type safety gaps and schema drift risks across services.

**Goal**: Migrate discovery payload definitions from local service types to centralized, versioned schemas in `sunday-schemas` to ensure cross-service consistency and type safety.

## Current State Analysis

### ✅ What Exists in sunday-schemas v1.0.15
- `RawEventsDiscoveryV0` and `RawSeriesDiscoveryV0` envelope structures
- `RawEventsDiscoveryV0Envelope` and `RawSeriesDiscoveryV0Envelope`
- Proper schema constants and stream definitions

### ❌ What's Missing
- **Structured payload types**: Both discovery types use `Payload map[string]interface{}`
- **Payload validation**: No compile-time or runtime validation of discovery data
- **Cross-service consistency**: Each service defines its own payload structures

### 🏠 Current Local Implementation
- `sunday-connectors` defines `EventMetadata` struct locally in `internal/discovery/discovery.go`
- Manual JSON marshaling into untyped payload maps
- Risk of schema drift between services consuming discovery events

## Scope and Requirements

### Discovery Data Types
Based on analysis of Polymarket and Kalshi APIs, we need schemas for:

1. **Event Discovery** - Individual prediction market events
2. **Series Discovery** - Collections/groups of related events
3. **Unified field mapping** from venue-specific APIs to canonical schema

### Key Design Principles
- **Backward Compatibility**: New schemas must not break existing consumers
- **Venue Agnostic**: Schema supports multiple prediction market venues
- **Extensible**: Design accommodates future venue integrations
- **Type Safe**: Strong typing for all common fields
- **Flexible**: `extra_metadata` for venue-specific data

## Proposed Schema Definitions

### Core Payload Structures

#### EventDiscoveryPayloadV0
```json
{
  "type": "object",
  "properties": {
    "event": { "$ref": "#/definitions/EventMetadataV0" },
    "event_id": { "type": "string", "description": "Unique event identifier for this discovery message" },
    "event_type": { "type": "string", "enum": ["discovered", "updated", "expired"] },
    "timestamp": { "type": "string", "format": "date-time" },
    "venue_id": { "type": "string", "enum": ["polymarket", "kalshi"] },
    "discovery_meta": { "$ref": "#/definitions/DiscoveryMetaV0" }
  },
  "required": ["event", "event_id", "event_type", "timestamp", "venue_id"]
}
```

#### SeriesDiscoveryPayloadV0
```json
{
  "type": "object",
  "properties": {
    "event": { "$ref": "#/definitions/SeriesMetadataV0" },
    "event_id": { "type": "string", "description": "Unique event identifier for this discovery message" },
    "event_type": { "type": "string", "enum": ["discovered", "updated", "expired"] },
    "timestamp": { "type": "string", "format": "date-time" },
    "venue_id": { "type": "string", "enum": ["polymarket", "kalshi"] },
    "discovery_meta": { "$ref": "#/definitions/DiscoveryMetaV0" }
  },
  "required": ["event", "event_id", "event_type", "timestamp", "venue_id"]
}
```

### Metadata Structures

#### EventMetadataV0
```json
{
  "type": "object",
  "properties": {
    "kind": { "type": "string", "const": "event" },
    "venue_id": { "type": "string", "enum": ["polymarket", "kalshi"] },
    "event_id": { "type": "string", "description": "Venue-specific event identifier" },
    "title": { "type": "string" },
    "description": { "type": "string" },
    "category": { "type": "string" },
    "active": { "type": "boolean" },
    "closed": { "type": "boolean" },
    "start_date": { "type": "string", "format": "date-time" },
    "end_date": { "type": "string", "format": "date-time" },
    "discovered_at": { "type": "string", "format": "date-time" },
    "last_seen": { "type": "string", "format": "date-time" },
    "parent_series_id": { "type": "string", "description": "Series ID this event belongs to" },
    "parent_series_title": { "type": "string", "description": "Series title for convenience" },
    "tags": {
      "type": "array",
      "items": { "type": "string" },
      "description": "Structured tags for categorization"
    },
    "relationships": { "$ref": "#/definitions/RelationshipsV0" },
    "extra_metadata": {
      "type": "object",
      "additionalProperties": true,
      "description": "Venue-specific fields that don't fit canonical schema"
    }
  },
  "required": ["kind", "venue_id", "event_id", "title", "active", "closed", "discovered_at", "last_seen"]
}
```

#### SeriesMetadataV0
```json
{
  "type": "object",
  "properties": {
    "kind": { "type": "string", "const": "series" },
    "venue_id": { "type": "string", "enum": ["polymarket", "kalshi"] },
    "event_id": { "type": "string", "description": "Series identifier (note: field name kept for compatibility)" },
    "title": { "type": "string" },
    "description": { "type": "string" },
    "category": { "type": "string" },
    "active": { "type": "boolean" },
    "closed": { "type": "boolean" },
    "discovered_at": { "type": "string", "format": "date-time" },
    "last_seen": { "type": "string", "format": "date-time" },
    "tags": {
      "type": "array",
      "items": { "type": "string" },
      "description": "Structured tags for categorization"
    },
    "child_event_ids": {
      "type": "array",
      "items": { "type": "string" },
      "description": "Event IDs that belong to this series"
    },
    "relationships": { "$ref": "#/definitions/RelationshipsV0" },
    "series_data": { "$ref": "#/definitions/SeriesDataV0" },
    "extra_metadata": {
      "type": "object",
      "additionalProperties": true,
      "description": "Venue-specific fields that don't fit canonical schema"
    }
  },
  "required": ["kind", "venue_id", "event_id", "title", "active", "closed", "discovered_at", "last_seen"]
}
```

#### SeriesDataV0 (Series-Specific Fields)
```json
{
  "type": "object",
  "properties": {
    "ticker": { "type": "string", "description": "Series ticker/symbol" },
    "slug": { "type": "string", "description": "URL-friendly series identifier" },
    "subtitle": { "type": "string", "description": "Short series subtitle" },
    "series_type": { "type": "string", "description": "Type/classification of series" },
    "recurrence": { "type": "string", "description": "Series recurrence pattern" },
    "image_url": { "type": "string", "format": "uri" },
    "icon_url": { "type": "string", "format": "uri" },
    "layout": { "type": "string", "description": "UI layout hint" },
    "financial": { "$ref": "#/definitions/FinancialDataV0" },
    "status": { "$ref": "#/definitions/StatusDataV0" },
    "contract": { "$ref": "#/definitions/ContractDataV0" },
    "timestamps": { "$ref": "#/definitions/TimestampDataV0" },
    "creators": { "$ref": "#/definitions/CreatorDataV0" }
  }
}
```

### Supporting Data Structures

#### FinancialDataV0
```json
{
  "type": "object",
  "properties": {
    "volume_24h_usd": {
      "type": "number",
      "multipleOf": 0.01,
      "minimum": 0,
      "description": "24-hour USD volume, decimal precision to 2 places"
    },
    "volume_total_usd": {
      "type": "number",
      "multipleOf": 0.01,
      "minimum": 0,
      "description": "Total USD volume, decimal precision to 2 places"
    },
    "liquidity_total_usd": {
      "type": "number",
      "multipleOf": 0.01,
      "minimum": 0,
      "description": "Total USD liquidity, decimal precision to 2 places"
    },
    "volume_24h_contracts": {
      "type": "integer",
      "minimum": 0,
      "description": "24-hour contract volume count"
    },
    "volume_total_contracts": {
      "type": "integer",
      "minimum": 0,
      "description": "Total contract volume count"
    },
    "score": {
      "type": "number",
      "minimum": 0,
      "description": "Ranking/scoring metric (unitless)"
    },
    "currency": {
      "type": "string",
      "enum": ["USD"],
      "description": "Currency unit for monetary values"
    }
  },
  "additionalProperties": false
}
```

#### StatusDataV0
```json
{
  "type": "object",
  "properties": {
    "archived": { "type": "boolean" },
    "is_new": { "type": "boolean", "description": "Newly featured series" },
    "featured": { "type": "boolean" },
    "restricted": { "type": "boolean", "description": "Access restrictions apply" },
    "is_template": { "type": "boolean", "description": "Template for event generation" },
    "competitive": { "type": "string", "description": "Competitive mode/flag" },
    "comments_enabled": { "type": "boolean" }
  }
}
```

#### ContractDataV0
```json
{
  "type": "object",
  "properties": {
    "contract_url": { "type": "string", "format": "uri" },
    "contract_terms_url": { "type": "string", "format": "uri" },
    "fee_type": { "type": "string", "description": "Fee calculation method" },
    "fee_multiplier": { "type": "number" },
    "additional_prohibitions": {
      "type": "array",
      "items": { "type": "string" }
    },
    "settlement_sources": {
      "type": "array",
      "items": { "$ref": "#/definitions/SettlementSourceV0" }
    }
  }
}
```

#### SettlementSourceV0
```json
{
  "type": "object",
  "properties": {
    "name": { "type": "string" },
    "url": { "type": "string", "format": "uri" }
  },
  "required": ["name"]
}
```

#### TimestampDataV0
```json
{
  "type": "object",
  "properties": {
    "published_at": { "type": "string", "format": "date-time" },
    "created_at": { "type": "string", "format": "date-time" },
    "updated_at": { "type": "string", "format": "date-time" }
  }
}
```

#### CreatorDataV0
```json
{
  "type": "object",
  "properties": {
    "created_by": { "type": "string" },
    "updated_by": { "type": "string" }
  }
}
```

#### RelationshipsV0
```json
{
  "type": "object",
  "properties": {
    "series_id": { "type": "string", "description": "Parent series identifier" },
    "event_ids": {
      "type": "array",
      "items": { "type": "string" },
      "description": "Child event identifiers"
    },
    "instrument_ids": {
      "type": "array",
      "items": { "type": "string" },
      "description": "Related venue instrument identifiers"
    }
  },
  "description": "Parent/child relationship mappings"
}
```

#### DiscoveryMetaV0
```json
{
  "type": "object",
  "properties": {
    "batch_id": { "type": "string", "description": "Unique identifier for this discovery batch" },
    "batch_sequence": { "type": "integer", "minimum": 1, "description": "Position of this item within the batch" },
    "batch_total_count": { "type": "integer", "minimum": 1, "description": "Total number of items in this batch" },
    "discovery_run_id": { "type": "string", "description": "Unique identifier for the entire discovery run" }
  },
  "required": ["batch_id", "batch_sequence", "batch_total_count", "discovery_run_id"],
  "description": "Metadata about the discovery batch/run for monitoring and sequencing"
}
```

## Field Mapping Analysis

### Venue-Specific Field Mappings

| Canonical Field | Polymarket Path | Kalshi Path | Notes |
|----------------|-----------------|-------------|-------|
| venue_id | N/A | N/A | Set by pipeline |
| event_id (series) | [].id | series[].ticker | Polymarket uses numeric ID, Kalshi uses ticker |
| title | [].title | series[].title | Direct mapping |
| ticker | [].ticker | series[].ticker | Available in both |
| slug | [].slug | N/A | Polymarket only |
| subtitle | [].subtitle | N/A | Polymarket only |
| series_type | [].seriesType | N/A | Polymarket only |
| recurrence | [].recurrence | series[].frequency | Different field names |
| category | [].category | series[].category | Direct mapping |
| image_url | [].image | N/A | Polymarket only |
| icon_url | [].icon | N/A | Polymarket only |
| volume_24h_usd | [].volume24hr | N/A | Polymarket only - USD amount |
| volume_total_usd | [].volume | N/A | Polymarket only - USD amount |
| tags | N/A | series[].tags[] | Kalshi only - structured array |
| contract_url | N/A | series[].contract_url | Kalshi only |
| contract_terms_url | N/A | series[].contract_terms_url | Kalshi only |
| fee_type | N/A | series[].fee_type | Kalshi only |
| additional_prohibitions | N/A | series[].additional_prohibitions[] | Kalshi only |
| settlement_sources | N/A | series[].settlement_sources[] | Kalshi only |
| parent_series_id | Derived from context | Derived from context | Relationship mapping |
| child_event_ids | Derived from events API | Derived from events API | Relationship mapping |

## Schema Registry Integration

### Schema Subjects
- `raw.events.discovery.v0-value`
- `raw.series.discovery.v0-value`

### Envelope Schema Updates
Update existing envelope schemas to reference new payload types:

```json
{
  "type": "object",
  "properties": {
    "envelope": { "$ref": "#/definitions/RawEventsDiscoveryV0Envelope" },
    "payload": { "$ref": "#/definitions/EventDiscoveryPayloadV0" }
  }
}
```

## Migration Strategy

### Phase 1: Schema Addition (sunday-schemas)
- Add new payload schema definitions
- Generate Go types with proper JSON tags
- Version as v1.1.0 with new discovery payload types
- Maintain backward compatibility with `map[string]interface{}` payloads

### Phase 2: sunday-connectors Integration
- Update `sunday-connectors` to import and use new typed payloads
- Replace local `EventMetadata` with `sundayschemas.EventMetadataV0`
- Update venue adapters to populate structured types
- Add validation for proper schema compliance

### Phase 3: Cross-Service Adoption
- Other services adopt new typed payloads
- Deprecation notices for untyped payload usage
- Documentation updates for downstream consumers

### Phase 4: Legacy Cleanup
- Remove compatibility shims
- Enforce typed payload validation
- Schema evolution to v2 if needed

## Validation Requirements

### Compile-Time Validation
- Strong typing for all common fields
- Enum validation for venue_id, event_type, kind fields
- Required field enforcement

### Runtime Validation
- JSON schema validation for payloads
- Cross-field validation (e.g., kind="series" requires series_data)
- Venue-specific field population rules

### Backward Compatibility
- New schemas accept existing untyped payloads during transition
- Gradual migration path with deprecation warnings
- Version bumps for breaking changes

## Expected Go Type Generation

```go
// Generated types in sunday-schemas
type EventDiscoveryPayloadV0 struct {
    Event         EventMetadataV0  `json:"event"`
    EventID       string           `json:"event_id"`
    EventType     EventType        `json:"event_type"`
    Timestamp     time.Time        `json:"timestamp"`
    VenueID       VenueID          `json:"venue_id"`
    DiscoveryMeta *DiscoveryMetaV0 `json:"discovery_meta,omitempty"`
}

type SeriesDiscoveryPayloadV0 struct {
    Event         SeriesMetadataV0 `json:"event"`
    EventID       string           `json:"event_id"`
    EventType     EventType        `json:"event_type"`
    Timestamp     time.Time        `json:"timestamp"`
    VenueID       VenueID          `json:"venue_id"`
    DiscoveryMeta *DiscoveryMetaV0 `json:"discovery_meta,omitempty"`
}

type EventMetadataV0 struct {
    Kind               DiscoveryKind    `json:"kind"`
    VenueID            VenueID          `json:"venue_id"`
    EventID            string           `json:"event_id"`
    Title              string           `json:"title"`
    Description        string           `json:"description,omitempty"`
    Category           string           `json:"category,omitempty"`
    Active             bool             `json:"active"`
    Closed             bool             `json:"closed"`
    StartDate          *time.Time       `json:"start_date,omitempty"`
    EndDate            *time.Time       `json:"end_date,omitempty"`
    ParentSeriesID     string           `json:"parent_series_id,omitempty"`
    ParentSeriesTitle  string           `json:"parent_series_title,omitempty"`
    Tags               []string         `json:"tags,omitempty"`
    Relationships      *RelationshipsV0 `json:"relationships,omitempty"`
    DiscoveredAt       time.Time        `json:"discovered_at"`
    LastSeen           time.Time        `json:"last_seen"`
    ExtraMetadata      map[string]any   `json:"extra_metadata,omitempty"`
}

type SeriesMetadataV0 struct {
    Kind            DiscoveryKind    `json:"kind"`
    VenueID         VenueID          `json:"venue_id"`
    EventID         string           `json:"event_id"`
    Title           string           `json:"title"`
    Description     string           `json:"description,omitempty"`
    Category        string           `json:"category,omitempty"`
    Active          bool             `json:"active"`
    Closed          bool             `json:"closed"`
    Tags            []string         `json:"tags,omitempty"`
    ChildEventIDs   []string         `json:"child_event_ids,omitempty"`
    Relationships   *RelationshipsV0 `json:"relationships,omitempty"`
    DiscoveredAt    time.Time        `json:"discovered_at"`
    LastSeen        time.Time        `json:"last_seen"`
    SeriesData      *SeriesDataV0    `json:"series_data,omitempty"`
    ExtraMetadata   map[string]any   `json:"extra_metadata,omitempty"`
}

type SeriesDataV0 struct {
    Ticker           string            `json:"ticker,omitempty"`
    Slug             string            `json:"slug,omitempty"`
    Subtitle         string            `json:"subtitle,omitempty"`
    SeriesType       string            `json:"series_type,omitempty"`
    Recurrence       string            `json:"recurrence,omitempty"`
    ImageURL         string            `json:"image_url,omitempty"`
    IconURL          string            `json:"icon_url,omitempty"`
    Layout           string            `json:"layout,omitempty"`
    Financial        *FinancialDataV0  `json:"financial,omitempty"`
    Status           *StatusDataV0     `json:"status,omitempty"`
    Contract         *ContractDataV0   `json:"contract,omitempty"`
    Timestamps       *TimestampDataV0  `json:"timestamps,omitempty"`
    Creators         *CreatorDataV0    `json:"creators,omitempty"`
}

// Supporting enums
type DiscoveryKind string
const (
    DiscoveryKindEvent  DiscoveryKind = "event"
    DiscoveryKindSeries DiscoveryKind = "series"
)

type EventType string
const (
    EventTypeDiscovered EventType = "discovered"
    EventTypeUpdated    EventType = "updated"
    EventTypeExpired    EventType = "expired"
)

type VenueID string
const (
    VenueIDPolymarket VenueID = "polymarket"
    VenueIDKalshi     VenueID = "kalshi"
)

// Supporting data structures
type FinancialDataV0 struct {
    Volume24hUSD         *float64 `json:"volume_24h_usd,omitempty"`
    VolumeTotalUSD       *float64 `json:"volume_total_usd,omitempty"`
    LiquidityTotalUSD    *float64 `json:"liquidity_total_usd,omitempty"`
    Volume24hContracts   *int     `json:"volume_24h_contracts,omitempty"`
    VolumeTotalContracts *int     `json:"volume_total_contracts,omitempty"`
    Score                *float64 `json:"score,omitempty"`
    Currency             string   `json:"currency,omitempty"`
}

type StatusDataV0 struct {
    Archived        *bool  `json:"archived,omitempty"`
    IsNew           *bool  `json:"is_new,omitempty"`
    Featured        *bool  `json:"featured,omitempty"`
    Restricted      *bool  `json:"restricted,omitempty"`
    IsTemplate      *bool  `json:"is_template,omitempty"`
    Competitive     string `json:"competitive,omitempty"`
    CommentsEnabled *bool  `json:"comments_enabled,omitempty"`
}

type ContractDataV0 struct {
    ContractURL             string                `json:"contract_url,omitempty"`
    ContractTermsURL        string                `json:"contract_terms_url,omitempty"`
    FeeType                 string                `json:"fee_type,omitempty"`
    FeeMultiplier           *float64              `json:"fee_multiplier,omitempty"`
    AdditionalProhibitions  []string              `json:"additional_prohibitions,omitempty"`
    SettlementSources       []SettlementSourceV0  `json:"settlement_sources,omitempty"`
}

type SettlementSourceV0 struct {
    Name string `json:"name"`
    URL  string `json:"url,omitempty"`
}

type TimestampDataV0 struct {
    PublishedAt *time.Time `json:"published_at,omitempty"`
    CreatedAt   *time.Time `json:"created_at,omitempty"`
    UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type CreatorDataV0 struct {
    CreatedBy string `json:"created_by,omitempty"`
    UpdatedBy string `json:"updated_by,omitempty"`
}

type RelationshipsV0 struct {
    SeriesID      string   `json:"series_id,omitempty"`
    EventIDs      []string `json:"event_ids,omitempty"`
    InstrumentIDs []string `json:"instrument_ids,omitempty"`
}

type DiscoveryMetaV0 struct {
    BatchID          string `json:"batch_id"`
    BatchSequence    int    `json:"batch_sequence"`
    BatchTotalCount  int    `json:"batch_total_count"`
    DiscoveryRunID   string `json:"discovery_run_id"`
}
```

## Testing Requirements

### Schema Validation Tests
- Valid payload acceptance
- Invalid payload rejection
- Field type validation
- Required field enforcement

### Cross-Compatibility Tests
- Existing untyped payloads still deserialize
- New typed payloads serialize correctly
- Round-trip JSON marshaling/unmarshaling

### Integration Tests
- sunday-connectors integration with new types
- Downstream consumer compatibility
- Performance impact assessment

## Success Criteria

### ✅ Type Safety
- All discovery payloads use strongly-typed structs
- Compile-time validation of field types and requirements
- IDE auto-completion and validation support

### ✅ Schema Consistency
- Single source of truth for discovery payload structure
- Centralized schema evolution and versioning
- Cross-service compatibility guarantees

### ✅ Backward Compatibility
- Existing consumers continue working during migration
- Gradual adoption path with clear deprecation timeline
- No breaking changes for essential workflows

### ✅ Developer Experience
- Clear documentation and examples
- Auto-generated types with proper Go conventions
- Validation errors provide actionable feedback

## Timeline and Dependencies

### Dependencies
- Access to `sunday-schemas` repository for schema addition
- Coordination with downstream services using discovery data
- Schema Registry access for subject registration

### Estimated Timeline
- **Week 1-2**: Schema design and JSON Schema definition
- **Week 3**: Go type generation and validation
- **Week 4**: Integration testing with sunday-connectors
- **Week 5-6**: Documentation and migration guides
- **Week 7+**: Cross-service adoption and legacy cleanup

## Questions for sunday-schemas Team

1. **Versioning Strategy**: Should this be v1.1.0 or v2.0.0 given the scope of new types?
2. **Naming Conventions**: Any preferences for type names (V0 suffix, DiscoveryPayload vs Discovery)?
3. **Schema Organization**: Should discovery schemas be in separate files or integrated with existing schemas?
4. **Validation Level**: How strict should runtime validation be during the transition period?
5. **Breaking Changes**: What's the policy for schema evolution and backward compatibility?
6. **Documentation**: What format do you prefer for schema documentation and examples?

## Appendix

### Sample Kafka Messages

#### Current Untyped Format
```json
{
  "envelope": { ... },
  "payload": {
    "event": {
      "kind": "series",
      "venue_id": "polymarket",
      "event_id": "10244",
      "title": "Concacaf",
      "active": true,
      "closed": false
    },
    "event_id": "evt_1762203892042327000_0",
    "event_type": "discovered",
    "timestamp": "2025-11-03T16:04:52.042327-05:00",
    "venue_id": "polymarket"
  }
}
```

#### Proposed Typed Format
```json
{
  "envelope": { ... },
  "payload": {
    "event": {
      "kind": "series",
      "venue_id": "polymarket",
      "event_id": "10244",
      "title": "Concacaf",
      "active": true,
      "closed": false,
      "discovered_at": "2025-11-03T21:04:51.761033Z",
      "last_seen": "2025-11-03T21:04:51.761033Z",
      "series_data": {
        "series_type": "single",
        "timestamps": {
          "created_at": "2025-09-03T03:07:56.295896Z",
          "updated_at": "2025-11-03T21:01:11.954948Z"
        }
      }
    },
    "event_id": "evt_1762203892042327000_0",
    "event_type": "discovered",
    "timestamp": "2025-11-03T21:04:52.042327Z",
    "venue_id": "polymarket"
  }
}
```

---

**Document Version**: 1.0
**Created**: 2025-11-03
**Author**: sunday-connectors team
**Target Audience**: sunday-schemas maintainers