# Sunday Schemas Integration Guide for Development Teams

> **Comprehensive guide for integrating sunday-schemas into your project**

## 🚨 **IMPORTANT: Field Name Corrections**

**If you encounter type mismatches**, see the [**Corrected Integration Guide**](./CORRECTED_INTEGRATION_GUIDE.md) for verified field names and working examples.

**Key corrections for Go usage:**
- Use `schemas.RawEnvelope` (not `RawEnvelopeV0`)
- Use `TsEventMs` (not `TsEventMS`)
- Use `TsIngestMs` (not `TsIngestMS`)
- Use `Bids/Asks` directly (not `BidsDelta/AsksDelta`)

## 📋 Table of Contents

1. [Overview](#overview)
2. [Quick Start](#quick-start)
3. [Installation](#installation)
4. [Schema Architecture](#schema-architecture)
5. [TypeScript/JavaScript Usage](#typescriptjavascript-usage)
6. [Go Usage](#go-usage)
7. [Event Schema Reference](#event-schema-reference)
8. [Validation & Compliance](#validation--compliance)
9. [Best Practices](#best-practices)
10. [Troubleshooting](#troubleshooting)
11. [Examples](#examples)
12. [Support & Contributing](#support--contributing)

---

## Overview

**sunday-schemas** is the centralized schema registry for the Sunday platform, providing:

- **Event Schemas** (JSON Schema) for all Kafka topics and event streams
- **OpenAPI Specifications** for HTTP APIs
- **Generated Types** for TypeScript and Go with full type safety
- **Validation Functions** for runtime schema compliance
- **Venue & Instrument Registries** for canonical data definitions

### 🎯 Why Use Sunday Schemas?

✅ **Type Safety**: Compile-time guarantees across all services
✅ **Consistency**: Same types used platform-wide
✅ **Validation**: Built-in schema compliance checking
✅ **Future-Proof**: Automatic updates when schemas evolve
✅ **Documentation**: Self-documenting with comprehensive examples

---

## Quick Start

### TypeScript/JavaScript
```bash
npm install sunday-schemas
```

```typescript
import { RawEnvelopeV0, VENUE_IDS, validateVenue } from 'sunday-schemas';

// Type-safe event envelope
const event: RawEnvelopeV0 = {
  schema: 'raw.v0',
  venue_id: 'polymarket',
  stream: 'orderbook',
  instrument_native: 'some-instrument',
  partition_key: 'polymarket:some-instrument',
  ts_event_ms: Date.now(),
  ts_ingest_ms: Date.now(),
  payload: { /* venue-specific data */ }
};

// Validation
if (validateVenue('polymarket')) {
  console.log('Valid venue');
}
```

### Go
```bash
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1
```

```go
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"

// Type-safe event envelope
event := schemas.RawEnvelope{
    Schema:           string(schemas.SchemaRAW_V0),
    VenueID:          string(schemas.Polymarket),
    Stream:           "orderbook",
    InstrumentNative: "some-instrument",
    PartitionKey:     "polymarket:some-instrument",
    TsEventMs:        time.Now().UnixMilli(),  // Note: TsEventMs (not TsEventMS)
    TsIngestMs:       time.Now().UnixMilli(),  // Note: TsIngestMs (not TsIngestMS)
    Payload:          map[string]interface{}{/* venue data */},
}

// Validation
if err := schemas.ValidateVenue(string(schemas.Polymarket)); err != nil {
    log.Fatal("Invalid venue:", err)
}
```

---

## Installation

### NPM Package
```bash
# Install latest version
npm install sunday-schemas

# Install specific version
npm install sunday-schemas@1.0.0

# Development dependency (if only using for types)
npm install -D sunday-schemas
```

### Go Module

⚠️ **Important**: This is a private repository. See [Private Repository Setup](./PRIVATE_REPOSITORY_SETUP.md) if you get permission errors.

```bash
# Configure for private repository (one-time setup)
go env -w GOPRIVATE=github.com/rakeyshgidwani/sunday-schemas

# Install latest version
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go

# Install specific version
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1

# Add to go.mod
require github.com/rakeyshgidwani/sunday-schemas/codegen/go v1.0.1
```

**If you get 404/permission errors**: Follow the [Private Repository Setup Guide](./PRIVATE_REPOSITORY_SETUP.md) for detailed configuration steps.

### Version Compatibility

| Schema Version | NPM Package | Go Module | Release Date |
|----------------|-------------|-----------|--------------|
| v1.0.1 ✅       | `sunday-schemas@1.0.0` | `codegen/go@v1.0.1` | Latest |
| v1.0.0         | `sunday-schemas@1.0.0` | `codegen/go@v1.0.0` | Initial |

> **Recommendation**: Always use the latest version for bug fixes and new features.

---

## Schema Architecture

### Event Schema Hierarchy

```
Sunday Platform Events
├── raw.v0                      # Raw venue data envelope
├── md.*.v1                     # Market data (normalized)
│   ├── md.orderbook.delta.v1   # Orderbook updates
│   └── md.trade.v1             # Trade executions
├── insights.*.v1               # Analytics & insights
│   ├── insights.arb.lite.v1    # Arbitrage opportunities
│   ├── insights.movers.v1      # Price movers
│   ├── insights.whales.lite.v1 # Large trades
│   └── insights.unusual.v1     # Unusual activity
└── infra.*.v1                  # Infrastructure
    └── infra.venue_health.v1   # Venue connectivity
```

### Schema Naming Convention
- **Format**: `{category}.{type}.{version}`
- **Categories**: `raw`, `md`, `insights`, `infra`
- **Versioning**: Semantic versioning in schema ID (e.g., `v1`, `v2`)

### Supported Venues
- **Polymarket** (`polymarket`) - Prediction markets
- **Kalshi** (`kalshi`) - Event trading platform

---

## TypeScript/JavaScript Usage

### Basic Imports
```typescript
// Type definitions
import type {
  RawEnvelopeV0,
  NormalizedOrderBookDeltaV1,
  NormalizedTradeV1,
  HealthStatus
} from 'sunday-schemas';

// Constants and utilities
import {
  SCHEMA_CONSTANTS,
  VENUE_IDS,
  validateVenue,
  validateSchema
} from 'sunday-schemas';
```

### Working with Event Envelopes
```typescript
import type { RawEnvelopeV0 } from 'sunday-schemas';

// Creating events
const createRawEvent = (
  venueId: string,
  stream: string,
  instrument: string,
  payload: any
): RawEnvelopeV0 => ({
  schema: 'raw.v0',
  venue_id: venueId,
  stream,
  instrument_native: instrument,
  partition_key: `${venueId}:${instrument}`,
  ts_event_ms: Date.now(),
  ts_ingest_ms: Date.now(),
  payload
});

// Type-safe event processing
const processEvent = (event: RawEnvelopeV0) => {
  // TypeScript ensures all required fields are present
  console.log(`Processing ${event.schema} from ${event.venue_id}`);

  // Optional fields are properly typed
  if (event.is_historical) {
    console.log('Historical data processing');
  }
};
```

### Schema Validation
```typescript
import { validateVenue, validateSchema, VENUE_IDS } from 'sunday-schemas';

// Venue validation
const isValidVenue = (venue: string): boolean => {
  return VENUE_IDS.includes(venue);
};

// Runtime validation
const validateEvent = (event: any): event is RawEnvelopeV0 => {
  // Check required fields
  if (!event.schema || !event.venue_id || !event.stream) {
    return false;
  }

  // Validate venue
  if (!validateVenue(event.venue_id)) {
    return false;
  }

  // Validate schema
  if (!validateSchema(event.schema)) {
    return false;
  }

  return true;
};
```

### Advanced Type Usage
```typescript
// Union types for all event schemas
import type { SundayEvent } from 'sunday-schemas';

// Event processors with type discrimination
const processAnyEvent = (event: SundayEvent) => {
  switch (event.schema) {
    case 'raw.v0':
      // TypeScript knows this is RawEnvelopeV0
      processRawEvent(event);
      break;
    case 'md.trade.v1':
      // TypeScript knows this is NormalizedTradeV1
      processTradeEvent(event);
      break;
    default:
      console.log('Unknown event type');
  }
};

// Custom type guards
const isRawEvent = (event: SundayEvent): event is RawEnvelopeV0 => {
  return event.schema === 'raw.v0';
};
```

---

## Go Usage

### Basic Imports
```go
import (
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
    "encoding/json"
    "time"
)
```

### Working with Event Envelopes
```go
// Creating events (verified working pattern from sunday-connectors)
func CreateRawEvent(venueID, stream, instrument string, payload interface{}) schemas.RawEnvelope {
    now := time.Now().UTC()
    return schemas.RawEnvelope{
        Schema:           string(schemas.SchemaRAW_V0),
        VenueID:          venueID,
        Stream:           stream,
        InstrumentNative: instrument,
        PartitionKey:     venueID + ":" + instrument,
        TsEventMs:        now.UnixMilli(),  // ✅ Correct: TsEventMs
        TsIngestMs:       now.UnixMilli(),  // ✅ Correct: TsIngestMs
        Payload:          payload,
    }
}

// Type-safe event processing
func ProcessEvent(event schemas.RawEnvelope) error {
    // All fields are properly typed
    log.Printf("Processing %s from %s", event.Schema, event.VenueID)

    // Optional fields are handled safely
    if event.IsHistorical != nil && *event.IsHistorical {
        log.Println("Processing historical data")
    }

    return nil
}
```

### Schema Validation
```go
// Venue validation
func ValidateEventVenue(venueID string) error {
    return schemas.ValidateVenue(venueID)
}

// Schema validation
func ValidateEventSchema(schema string) error {
    return schemas.ValidateSchema(schema)
}

// Complete event validation
func ValidateEvent(event schemas.RawEnvelope) error {
    if err := schemas.ValidateVenue(event.VenueID); err != nil {
        return fmt.Errorf("invalid venue: %w", err)
    }

    if err := schemas.ValidateSchema(event.Schema); err != nil {
        return fmt.Errorf("invalid schema: %w", err)
    }

    return nil
}
```

### Working with Constants
```go
// Using venue constants
func ProcessByVenue(event schemas.RawEnvelope) {
    switch schemas.VenueID(event.VenueID) {
    case schemas.Polymarket:
        processPolymarketEvent(event)
    case schemas.Kalshi:
        processKalshiEvent(event)
    default:
        log.Printf("Unknown venue: %s", event.VenueID)
    }
}

// Getting all available schemas
func ListSupportedSchemas() []schemas.EventSchema {
    return schemas.AllSchemas()
}

// Getting all supported venues
func ListSupportedVenues() []schemas.VenueID {
    return schemas.AllVenues()
}
```

### JSON Serialization
```go
// Marshaling events
func SerializeEvent(event schemas.RawEnvelope) ([]byte, error) {
    return json.Marshal(event)
}

// Unmarshaling events
func DeserializeEvent(data []byte) (schemas.RawEnvelope, error) {
    var event schemas.RawEnvelope
    err := json.Unmarshal(data, &event)
    return event, err
}

// Working with generic payloads
func ProcessGenericPayload(event schemas.RawEnvelope) error {
    // Payload is interface{} - can be marshaled to specific types
    payloadBytes, err := json.Marshal(event.Payload)
    if err != nil {
        return err
    }

    // Unmarshal to venue-specific struct
    var polymarketData PolymarketOrderbook
    if err := json.Unmarshal(payloadBytes, &polymarketData); err != nil {
        return err
    }

    return nil
}
```

---

## Event Schema Reference

### Raw Event Envelope (`raw.v0`)

**Purpose**: Universal wrapper for all venue-specific event data

**Go Type Definition** (✅ **Verified Working**):
```go
type RawEnvelope struct {
    Schema           string      `json:"schema"`                    // "raw.v0"
    VenueID          string      `json:"venue_id"`                  // "polymarket" | "kalshi"
    Stream           string      `json:"stream"`                    // "orderbook" | "trades"
    InstrumentNative string      `json:"instrument_native"`         // Venue-specific instrument ID
    PartitionKey     string      `json:"partition_key"`             // Usually venue:instrument
    TsEventMs        int64       `json:"ts_event_ms"`              // Event timestamp (milliseconds)
    TsIngestMs       int64       `json:"ts_ingest_ms"`             // Ingestion timestamp (milliseconds)
    Payload          interface{} `json:"payload"`                   // Venue-specific event data

    // Optional fields
    IsHistorical     *bool       `json:"is_historical,omitempty"`   // Historical/backfill data
    BackfillTsMs     *int64      `json:"backfill_ts_ms,omitempty"`  // Backfill timestamp
}
```

**TypeScript Type Definition**:
```typescript
{
  schema: "raw.v0",              // Schema identifier
  venue_id: string,              // Venue identifier ("polymarket" | "kalshi")
  stream: string,                // Event stream type ("orderbook" | "trades")
  instrument_native: string,     // Venue-specific instrument ID
  partition_key: string,         // Kafka partition key (usually venue:instrument)
  ts_event_ms: number,          // Event timestamp (milliseconds)
  ts_ingest_ms: number,         // Ingestion timestamp (milliseconds)
  payload: object,               // Venue-specific event data
  is_historical?: boolean,       // True if historical/backfill data
  backfill_ts_ms?: number       // Backfill timestamp if applicable
}
```

### Market Data Events

#### Orderbook Delta (`md.orderbook.delta.v1`)
```go
// Go type definition (✅ Verified)
type OrderbookDelta struct {
    Schema       string      `json:"schema"`           // "md.orderbook.delta.v1"
    InstrumentID string      `json:"instrument_id"`    // Normalized instrument ID
    VenueID      string      `json:"venue_id"`         // "polymarket" | "kalshi"
    TsMs         int64       `json:"ts_ms"`           // Timestamp (milliseconds)
    Seq          int64       `json:"seq"`             // Sequence number for ordering
    Bids         [][]float64 `json:"bids"`            // [price, size] pairs
    Asks         [][]float64 `json:"asks"`            // [price, size] pairs
    IsSnapshot   bool        `json:"is_snapshot"`     // True if full snapshot vs delta
}
```

#### Trade (`md.trade.v1`)
```go
// Go type definition (✅ Verified)
type Trade struct {
    Schema       string   `json:"schema"`             // "md.trade.v1"
    InstrumentID string   `json:"instrument_id"`      // Normalized instrument ID
    VenueID      string   `json:"venue_id"`           // "polymarket" | "kalshi"
    TsMs         int64    `json:"ts_ms"`             // Timestamp (milliseconds)
    Side         string   `json:"side"`               // "buy" | "sell"
    Prob         float64  `json:"prob"`               // Implied probability [0.0, 1.0]
    Size         float64  `json:"size"`               // Trade size in base units
    NotionalUsd  *float64 `json:"notional_usd,omitempty"` // Optional USD value
}
```

### Insights Events

#### Arbitrage Opportunities (`insights.arb.lite.v1`)
```typescript
{
  schema: "insights.arb.lite.v1",
  instrument_group_id: string,   // Cross-venue instrument group
  ts_event_ms: number,
  opportunities: [{
    venue_a: string,
    venue_b: string,
    price_diff: number,          // Price difference (implied probability)
    profit_potential: number,    // Estimated profit percentage
    liquidity_a: number,        // Available liquidity venue A
    liquidity_b: number        // Available liquidity venue B
  }]
}
```

### Infrastructure Events

#### Venue Health (`infra.venue_health.v1`)
```typescript
{
  schema: "infra.venue_health.v1",
  venue_id: string,
  ts_event_ms: number,
  status: "CONNECTED" | "DEGRADED" | "STALE", // Connection status
  latency_ms?: number,          // Connection latency
  error_rate?: number,          // Error rate percentage
  last_message_ts?: number     // Last successful message timestamp
}
```

---

## Validation & Compliance

### Schema Validation

**TypeScript**:
```typescript
import { validateSchema, validateVenue } from 'sunday-schemas';

// Validate schema identifier
if (!validateSchema('raw.v0')) {
  throw new Error('Invalid schema');
}

// Validate venue identifier
if (!validateVenue('polymarket')) {
  throw new Error('Invalid venue');
}
```

**Go**:
```go
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"

// Validate schema
if err := schemas.ValidateSchema("raw.v0"); err != nil {
    return fmt.Errorf("invalid schema: %w", err)
}

// Validate venue
if err := schemas.ValidateVenue("polymarket"); err != nil {
    return fmt.Errorf("invalid venue: %w", err)
}
```

### Runtime Type Checking

**TypeScript**:
```typescript
import type { RawEnvelopeV0 } from 'sunday-schemas';

// Runtime validation function
function isValidRawEvent(obj: any): obj is RawEnvelopeV0 {
  return (
    typeof obj === 'object' &&
    obj.schema === 'raw.v0' &&
    typeof obj.venue_id === 'string' &&
    typeof obj.stream === 'string' &&
    typeof obj.instrument_native === 'string' &&
    typeof obj.partition_key === 'string' &&
    typeof obj.ts_event_ms === 'number' &&
    typeof obj.ts_ingest_ms === 'number' &&
    obj.payload !== undefined
  );
}

// Usage
const data = JSON.parse(eventJson);
if (isValidRawEvent(data)) {
  // TypeScript now knows data is RawEnvelopeV0
  processEvent(data);
}
```

**Go**:
```go
// Validation with error details
func ValidateRawEnvelope(event schemas.RawEnvelope) error {
    if event.Schema == "" {
        return errors.New("schema is required")
    }

    if event.VenueID == "" {
        return errors.New("venue_id is required")
    }

    if event.Stream == "" {
        return errors.New("stream is required")
    }

    if event.TsEventMs <= 0 {
        return errors.New("ts_event_ms must be positive")
    }

    // Validate using schema registry
    if err := schemas.ValidateVenue(event.VenueID); err != nil {
        return fmt.Errorf("invalid venue_id: %w", err)
    }

    if err := schemas.ValidateSchema(event.Schema); err != nil {
        return fmt.Errorf("invalid schema: %w", err)
    }

    return nil
}
```

---

## Best Practices

### 1. Version Management
```bash
# Pin to specific versions in production
npm install sunday-schemas@1.0.0
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1

# Update regularly for new features and fixes
npm update sunday-schemas
go get -u github.com/rakeyshgidwani/sunday-schemas/codegen/go
```

### 2. Import Organization

**TypeScript**:
```typescript
// Group imports by type vs runtime
import type {
  RawEnvelopeV0,
  NormalizedTradeV1
} from 'sunday-schemas';

import {
  VENUE_IDS,
  validateVenue
} from 'sunday-schemas';
```

**Go**:
```go
// Use alias for cleaner code
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"

// Not: github.com/rakeyshgidwani/sunday-schemas/codegen/go.SchemaRAW_V0
// But: schemas.SchemaRAW_V0
```

### 3. Error Handling

**TypeScript**:
```typescript
// Graceful degradation
const processEvent = (event: any) => {
  try {
    if (!validateVenue(event.venue_id)) {
      console.warn(`Unknown venue: ${event.venue_id}`);
      return; // Skip unknown venues gracefully
    }

    // Process event...
  } catch (error) {
    console.error('Event processing failed:', error);
    // Log but don't crash
  }
};
```

**Go**:
```go
// Structured error handling
func ProcessEvent(event schemas.RawEnvelope) error {
    if err := ValidateRawEnvelope(event); err != nil {
        return fmt.Errorf("event validation failed: %w", err)
    }

    // Process event...
    return nil
}

// Caller handles appropriately
if err := ProcessEvent(event); err != nil {
    log.Printf("Failed to process event: %v", err)
    // Continue processing other events
}
```

### 4. Performance Considerations

**TypeScript**:
```typescript
// Cache validation results for hot paths
const venueCache = new Set(VENUE_IDS);
const isValidVenue = (venue: string) => venueCache.has(venue);

// Batch process events
const processEventBatch = (events: RawEnvelopeV0[]) => {
  const validEvents = events.filter(event =>
    isValidVenue(event.venue_id)
  );

  // Process valid events...
};
```

**Go**:
```go
// Pre-compile validation maps for performance
var validVenues = make(map[string]bool)

func init() {
    for _, venue := range schemas.AllVenues() {
        validVenues[string(venue)] = true
    }
}

func IsValidVenue(venue string) bool {
    return validVenues[venue]
}
```

### 5. Testing

**TypeScript**:
```typescript
// Create test fixtures
const createTestEvent = (overrides?: Partial<RawEnvelopeV0>): RawEnvelopeV0 => ({
  schema: 'raw.v0',
  venue_id: 'polymarket',
  stream: 'orderbook',
  instrument_native: 'test-instrument',
  partition_key: 'polymarket:test-instrument',
  ts_event_ms: Date.now(),
  ts_ingest_ms: Date.now(),
  payload: {},
  ...overrides
});

// Test with different venues
describe('Event Processing', () => {
  test('handles all supported venues', () => {
    VENUE_IDS.forEach(venue => {
      const event = createTestEvent({ venue_id: venue });
      expect(() => processEvent(event)).not.toThrow();
    });
  });
});
```

**Go**:
```go
// Helper functions for tests
func NewTestRawEnvelope(venue schemas.VenueID) schemas.RawEnvelope {
    return schemas.RawEnvelope{
        Schema:           string(schemas.SchemaRAW_V0),
        VenueID:          string(venue),
        Stream:           "test",
        InstrumentNative: "test-instrument",
        PartitionKey:     string(venue) + ":test-instrument",
        TsEventMs:        time.Now().UnixMilli(),
        TsIngestMs:       time.Now().UnixMilli(),
        Payload:          map[string]interface{}{"test": true},
    }
}

// Table-driven tests
func TestEventProcessing(t *testing.T) {
    venues := schemas.AllVenues()

    for _, venue := range venues {
        t.Run(string(venue), func(t *testing.T) {
            event := NewTestRawEnvelope(venue)
            err := ProcessEvent(event)
            assert.NoError(t, err)
        })
    }
}
```

---

## Troubleshooting

### Common Issues

#### 1. **Import Errors**

**Problem**: Cannot import sunday-schemas
```bash
Error: Cannot resolve module 'sunday-schemas'
```

**Solution**:
```bash
# Ensure package is installed
npm install sunday-schemas

# Check node_modules
ls node_modules/sunday-schemas

# Clear cache if needed
npm cache clean --force
```

#### 2. **Type Errors**

**Problem**: TypeScript cannot find type definitions
```typescript
// Error: Could not find declaration file for 'sunday-schemas'
```

**Solution**:
```bash
# Reinstall with TypeScript definitions
npm install sunday-schemas

# Check package includes .d.ts files
ls node_modules/sunday-schemas/codegen/ts/*.d.ts
```

#### 3. **Go Module Issues**

**Problem**: Cannot fetch Go module (404 Not Found)
```bash
go: github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1: verifying module
fatal: could not read Username for 'https://github.com': terminal prompts disabled
```

**Solution**: This is a private repository access issue. Follow the **[Private Repository Setup Guide](./PRIVATE_REPOSITORY_SETUP.md)** for complete setup instructions.

**Quick fix**:
```bash
# Configure GOPRIVATE (tells Go to bypass public proxy)
go env -w GOPRIVATE=github.com/rakeyshgidwani/sunday-schemas

# Configure Git authentication (choose one):
# Option A: Personal Access Token
git config --global url."https://USERNAME:TOKEN@github.com/".insteadOf "https://github.com/"

# Option B: SSH
git config --global url."git@github.com:".insteadOf "https://github.com/"

# Then try again
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1
```

#### 4. **Version Mismatches**

**Problem**: Different schema versions in different services

**Solution**:
```bash
# Use exact versions in package.json/go.mod
"sunday-schemas": "1.0.0"  # Not "^1.0.0"

# Update all services to same version
npm update sunday-schemas
go get -u github.com/rakeyshgidwani/sunday-schemas/codegen/go
```

### Debugging Tips

#### 1. **Validate Installation**
```bash
# NPM
npm list sunday-schemas

# Go
go list -m github.com/rakeyshgidwani/sunday-schemas/codegen/go
```

#### 2. **Check Available Types**
```typescript
// List all available exports
import * as schemas from 'sunday-schemas';
console.log(Object.keys(schemas));
```

```go
// Print available constants
package main

import (
    "fmt"
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func main() {
    fmt.Println("Schemas:", schemas.AllSchemas())
    fmt.Println("Venues:", schemas.AllVenues())
}
```

#### 3. **Runtime Validation**
```typescript
// Debug event structure
const debugEvent = (event: any) => {
  console.log('Event keys:', Object.keys(event));
  console.log('Schema:', event.schema);
  console.log('Venue valid:', validateVenue(event.venue_id));
};
```

---

## Examples

### Example 1: Event Stream Processing

**TypeScript/Node.js**:
```typescript
import type { RawEnvelopeV0 } from 'sunday-schemas';
import { validateVenue, VENUE_IDS } from 'sunday-schemas';

class EventProcessor {
  private handlers = new Map<string, (event: RawEnvelopeV0) => void>();

  constructor() {
    this.handlers.set('orderbook', this.handleOrderbook.bind(this));
    this.handlers.set('trades', this.handleTrades.bind(this));
  }

  async processEvent(rawEvent: any): Promise<void> {
    // Validate event structure
    if (!this.isValidEvent(rawEvent)) {
      console.warn('Invalid event structure:', rawEvent);
      return;
    }

    const event = rawEvent as RawEnvelopeV0;

    // Get handler for stream type
    const handler = this.handlers.get(event.stream);
    if (!handler) {
      console.warn(`No handler for stream: ${event.stream}`);
      return;
    }

    try {
      handler(event);
    } catch (error) {
      console.error(`Error processing ${event.stream} event:`, error);
    }
  }

  private isValidEvent(event: any): event is RawEnvelopeV0 {
    return (
      event &&
      event.schema === 'raw.v0' &&
      validateVenue(event.venue_id) &&
      typeof event.stream === 'string' &&
      typeof event.ts_event_ms === 'number'
    );
  }

  private handleOrderbook(event: RawEnvelopeV0): void {
    console.log(`Orderbook update from ${event.venue_id}:${event.instrument_native}`);
    // Process orderbook payload...
  }

  private handleTrades(event: RawEnvelopeV0): void {
    console.log(`Trade from ${event.venue_id}:${event.instrument_native}`);
    // Process trade payload...
  }
}

// Usage
const processor = new EventProcessor();

// Process events from Kafka/WebSocket/etc
const eventStream = getEventStream();
eventStream.on('data', (event) => {
  processor.processEvent(event);
});
```

**Go**:
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

type EventProcessor struct {
    handlers map[string]func(schemas.RawEnvelope) error
}

func NewEventProcessor() *EventProcessor {
    ep := &EventProcessor{
        handlers: make(map[string]func(schemas.RawEnvelope) error),
    }

    ep.handlers["orderbook"] = ep.handleOrderbook
    ep.handlers["trades"] = ep.handleTrades

    return ep
}

func (ep *EventProcessor) ProcessEvent(rawEvent []byte) error {
    var event schemas.RawEnvelope
    if err := json.Unmarshal(rawEvent, &event); err != nil {
        return fmt.Errorf("failed to unmarshal event: %w", err)
    }

    // Validate event
    if err := ep.validateEvent(event); err != nil {
        return fmt.Errorf("invalid event: %w", err)
    }

    // Get handler
    handler, exists := ep.handlers[event.Stream]
    if !exists {
        log.Printf("No handler for stream: %s", event.Stream)
        return nil
    }

    return handler(event)
}

func (ep *EventProcessor) validateEvent(event schemas.RawEnvelope) error {
    if event.Schema != string(schemas.SchemaRAW_V0) {
        return fmt.Errorf("unsupported schema: %s", event.Schema)
    }

    if err := schemas.ValidateVenue(event.VenueID); err != nil {
        return fmt.Errorf("invalid venue: %w", err)
    }

    return nil
}

func (ep *EventProcessor) handleOrderbook(event schemas.RawEnvelope) error {
    log.Printf("Orderbook update from %s:%s", event.VenueID, event.InstrumentNative)
    // Process orderbook payload...
    return nil
}

func (ep *EventProcessor) handleTrades(event schemas.RawEnvelope) error {
    log.Printf("Trade from %s:%s", event.VenueID, event.InstrumentNative)
    // Process trade payload...
    return nil
}

func main() {
    processor := NewEventProcessor()

    // Example event
    exampleEvent := `{
        "schema": "raw.v0",
        "venue_id": "polymarket",
        "stream": "orderbook",
        "instrument_native": "test-instrument",
        "partition_key": "polymarket:test-instrument",
        "ts_event_ms": 1234567890000,
        "ts_ingest_ms": 1234567890000,
        "payload": {}
    }`

    if err := processor.ProcessEvent([]byte(exampleEvent)); err != nil {
        log.Fatal(err)
    }
}
```

### Example 2: Data Pipeline Integration

**TypeScript - Kafka Consumer**:
```typescript
import { KafkaConsumer } from 'your-kafka-library';
import type { RawEnvelopeV0, SundayEvent } from 'sunday-schemas';
import { validateSchema, validateVenue } from 'sunday-schemas';

class SundayDataPipeline {
  private consumer: KafkaConsumer;

  constructor() {
    this.consumer = new KafkaConsumer({
      topics: ['raw-events', 'market-data', 'insights'],
      groupId: 'sunday-data-processor'
    });
  }

  async start(): Promise<void> {
    this.consumer.on('message', async (message) => {
      try {
        const event = JSON.parse(message.value.toString());
        await this.processEvent(event);
      } catch (error) {
        console.error('Failed to process message:', error);
      }
    });

    await this.consumer.connect();
    console.log('Sunday Data Pipeline started');
  }

  private async processEvent(event: any): Promise<void> {
    // Route by schema
    switch (event.schema) {
      case 'raw.v0':
        await this.processRawEvent(event as RawEnvelopeV0);
        break;
      case 'md.trade.v1':
        await this.processTradeEvent(event);
        break;
      case 'insights.arb.lite.v1':
        await this.processArbEvent(event);
        break;
      default:
        console.log(`Unknown schema: ${event.schema}`);
    }
  }

  private async processRawEvent(event: RawEnvelopeV0): Promise<void> {
    // Validate venue
    if (!validateVenue(event.venue_id)) {
      console.warn(`Invalid venue: ${event.venue_id}`);
      return;
    }

    // Store raw event
    await this.storeRawEvent(event);

    // Trigger normalization pipeline
    await this.triggerNormalization(event);
  }

  private async storeRawEvent(event: RawEnvelopeV0): Promise<void> {
    // Store in data warehouse/lake
    console.log(`Storing raw event: ${event.venue_id}:${event.instrument_native}`);
  }

  private async triggerNormalization(event: RawEnvelopeV0): Promise<void> {
    // Trigger downstream normalization
    console.log(`Triggering normalization for: ${event.stream}`);
  }
}

// Usage
const pipeline = new SundayDataPipeline();
pipeline.start().catch(console.error);
```

### Example 3: API Integration

**Go - HTTP API**:
```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

type APIServer struct {
    eventStore EventStore
}

type EventStore interface {
    StoreEvent(event schemas.RawEnvelope) error
    GetEvents(venue string, limit int) ([]schemas.RawEnvelope, error)
}

func NewAPIServer(store EventStore) *APIServer {
    return &APIServer{eventStore: store}
}

func (s *APIServer) handlePostEvent(w http.ResponseWriter, r *http.Request) {
    var event schemas.RawEnvelope
    if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Validate using schemas
    if err := s.validateEvent(event); err != nil {
        http.Error(w, fmt.Sprintf("Invalid event: %v", err), http.StatusBadRequest)
        return
    }

    // Store event
    if err := s.eventStore.StoreEvent(event); err != nil {
        http.Error(w, "Failed to store event", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}

func (s *APIServer) handleGetEvents(w http.ResponseWriter, r *http.Request) {
    venue := r.URL.Query().Get("venue")

    // Validate venue if provided
    if venue != "" {
        if err := schemas.ValidateVenue(venue); err != nil {
            http.Error(w, "Invalid venue", http.StatusBadRequest)
            return
        }
    }

    events, err := s.eventStore.GetEvents(venue, 100)
    if err != nil {
        http.Error(w, "Failed to get events", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(events)
}

func (s *APIServer) validateEvent(event schemas.RawEnvelope) error {
    // Schema validation
    if err := schemas.ValidateSchema(event.Schema); err != nil {
        return err
    }

    // Venue validation
    if err := schemas.ValidateVenue(event.VenueID); err != nil {
        return err
    }

    // Business logic validation
    if event.TsEventMs <= 0 {
        return fmt.Errorf("ts_event_ms must be positive")
    }

    if time.Unix(0, event.TsEventMs*int64(time.Millisecond)).After(time.Now().Add(time.Hour)) {
        return fmt.Errorf("event timestamp too far in future")
    }

    return nil
}

func (s *APIServer) SetupRoutes() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/events", s.handlePostEvent).Methods("POST")
    r.HandleFunc("/events", s.handleGetEvents).Methods("GET")

    // Add venue listing endpoint
    r.HandleFunc("/venues", func(w http.ResponseWriter, r *http.Request) {
        venues := schemas.AllVenues()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(venues)
    }).Methods("GET")

    return r
}

// Usage
func main() {
    // Implementation of EventStore would go here
    var store EventStore = &InMemoryEventStore{}

    server := NewAPIServer(store)
    router := server.SetupRoutes()

    fmt.Println("Sunday Data API listening on :8080")
    http.ListenAndServe(":8080", router)
}
```

---

## Support & Contributing

### Getting Help

1. **Documentation**: Check this guide and the [main README](../README.md)
2. **Issues**: Report bugs or request features in [GitHub Issues](https://github.com/rakeyshgidwani/sunday-schemas/issues)
3. **Code Review**: Review changes in [Pull Requests](https://github.com/rakeyshgidwani/sunday-schemas/pulls)

### Staying Updated

**Subscribe to releases**:
- Watch the GitHub repository for release notifications
- Follow semantic versioning for safe updates

**Update workflow**:
```bash
# Check current version
npm list sunday-schemas
go list -m github.com/rakeyshgidwani/sunday-schemas/codegen/go

# Update to latest
npm update sunday-schemas
go get -u github.com/rakeyshgidwani/sunday-schemas/codegen/go

# Test after updating
npm test
go test ./...
```

### Contributing Guidelines

1. **Schema Changes**: Propose new schemas or modifications via RFC (GitHub Issues)
2. **Backward Compatibility**: All changes must be backward compatible in Phase 1
3. **Testing**: Include examples and test cases for new schemas
4. **Documentation**: Update this guide for API changes

### Release Process

Schemas follow semantic versioning:
- **Patch** (`1.0.1`): Bug fixes, documentation updates
- **Minor** (`1.1.0`): New schemas, optional fields, additive changes
- **Major** (`2.0.0`): Breaking changes (not allowed in Phase 1)

---

## Quick Reference

### Package Information
- **NPM**: `sunday-schemas@1.0.0`
- **Go**: `github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1`
- **Repository**: https://github.com/rakeyshgidwani/sunday-schemas
- **Issues**: https://github.com/rakeyshgidwani/sunday-schemas/issues

### Key Types
- `RawEnvelopeV0` - Universal event wrapper
- `NormalizedTradeV1` - Standardized trade events
- `NormalizedOrderBookDeltaV1` - Standardized orderbook updates
- `HealthStatus` - Venue connectivity status

### Key Constants
- `VENUE_IDS` - All supported venues
- `SCHEMA_CONSTANTS` - All schema identifiers
- `Polymarket`, `Kalshi` - Venue constants (Go)
- `SchemaRAW_V0` - Schema constants (Go)

### Validation Functions
- `validateVenue(venue)` - Validate venue identifier
- `validateSchema(schema)` - Validate schema identifier
- `ValidateVenue(venue)` - Go venue validation
- `ValidateSchema(schema)` - Go schema validation

---

**Happy coding with Sunday Schemas! 🚀**

*For questions or support, reach out to the Platform Engineering team.*