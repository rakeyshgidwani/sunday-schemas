# TypeScript Usage

## Quick Start

Install the package:
```bash
npm install @rakeyshgidwani/sunday-schemas
```

Import and use:
```typescript
import {
  RawEnvelopeV0,
  NormalizedTradeV1,
  SCHEMA_CONSTANTS,
  VENUE_IDS
} from '@rakeyshgidwani/sunday-schemas';

// Use generated types with full type safety
const envelope: RawEnvelopeV0 = {
  schema: 'raw.v0',
  venue_id: 'polymarket',
  stream: 'orderbook',
  instrument_native: 'some-instrument',
  partition_key: 'polymarket:some-instrument',
  ts_event_ms: Date.now(),
  ts_ingest_ms: Date.now(),
  payload: {}
};
```

## Commands

```bash
# Add dependency
npm install @rakeyshgidwani/sunday-schemas@latest

# Update to latest
npm update @rakeyshgidwani/sunday-schemas

# Specific version
npm install @rakeyshgidwani/sunday-schemas@1.0.9

# Check installed version
npm list @rakeyshgidwani/sunday-schemas
```

## Available Types and Constants

**Event Types (with aliases):**
```typescript
RawEnvelopeV0                       // Alias: RawEnvelope
NormalizedOrderBookDeltaV1          // Alias: OrderbookDelta
NormalizedTradeV1                   // Alias: Trade
ArbitrageLiteV1                     // Alias: ArbitrageLite
MoversV1                            // Alias: Movers
WhaleFlowsLiteV1                    // Alias: WhalesLite
UnusualActivityV1                   // Alias: Unusual
VenueHealthV1                       // Alias: VenueHealth
```

**Union Types:**
```typescript
EventSchema
SundayEvent
```

**Runtime Constants:**
```typescript
SCHEMA_CONSTANTS               // Available at runtime
VENUE_IDS                      // Available at runtime
```

## Type-Safe Event Handling

**Runtime Schema Validation:**
```typescript
import { SCHEMA_CONSTANTS, VENUE_IDS, EventSchema, VenueId } from '@rakeyshgidwani/sunday-schemas';

// Validate schema at runtime
function isValidSchema(schema: string): schema is EventSchema {
  return schema in SCHEMA_CONSTANTS;
}

// Validate venue at runtime
function isValidVenue(venue: string): venue is VenueId {
  return VENUE_IDS.includes(venue as VenueId);
}
```

**Generic Type-Safe Event Processing:**
```typescript
import { EventBySchema, SundayEvent } from '@rakeyshgidwani/sunday-schemas';

// Generic event processor with full type safety
function processEvent<T extends EventSchema>(
  schema: T,
  data: EventBySchema<T>
): void {
  // TypeScript knows the exact type of data based on schema
  console.log(`Processing event: ${schema}`);
}
```

## TypeScript Configuration

Add to `tsconfig.json`:
```json
{
  "compilerOptions": {
    "moduleResolution": "node",
    "esModuleInterop": true,
    "strict": true,
    "exactOptionalPropertyTypes": true
  }
}
```

---

*This documentation is automatically generated from the actual TypeScript codegen output.*
*Last updated: 2025-11-04T22:18:01.180Z*
