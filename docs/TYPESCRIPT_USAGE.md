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
# Install dependency
npm install @rakeyshgidwani/sunday-schemas

# Install specific version
npm install @rakeyshgidwani/sunday-schemas@1.0.9

# Update to latest
npm update @rakeyshgidwani/sunday-schemas

# Check installed version
npm list @rakeyshgidwani/sunday-schemas
```

## Available Types and Constants

**Event Types (with aliases):**
```typescript
// Full names (backward compatibility)
RawEnvelopeV0                   // Raw venue data envelope
NormalizedOrderBookDeltaV1      // Orderbook changes
NormalizedTradeV1               // Trade events
ArbitrageLiteV1                 // Arbitrage opportunities
MoversV1                        // Price movement events
WhaleFlowsLiteV1                // Large trade flows
UnusualActivityV1               // Unusual trading activity
VenueHealthV1                   // Venue health metrics

// Short aliases (easier to use)
RawEnvelope                     // Alias for RawEnvelopeV0
OrderbookDelta                  // Alias for NormalizedOrderBookDeltaV1
Trade                           // Alias for NormalizedTradeV1
ArbitrageLite                   // Alias for ArbitrageLiteV1
Movers                          // Alias for MoversV1
WhalesLite                      // Alias for WhaleFlowsLiteV1
Unusual                         // Alias for UnusualActivityV1
VenueHealth                     // Alias for VenueHealthV1
```

**Union Types for Event Handling:**
```typescript
SundayEvent                     // Union of all event types
EventSchema                     // Union of all schema strings
EventBySchema<T>                // Extract event type by schema string
```

**Type Constants:**
```typescript
VenueId = 'polymarket' | 'kalshi'
TradeSide = 'buy' | 'sell'
HealthStatus = 'CONNECTED' | 'DEGRADED' | 'STALE'
```

**Runtime Constants:**
```typescript
SCHEMA_CONSTANTS                // Record<EventSchema, EventSchema>
VENUE_IDS                       // VenueId[] = ['polymarket', 'kalshi']
```

## Version Pinning

Pin to specific version in `package.json`:
```json
{
  "dependencies": {
    "@rakeyshgidwani/sunday-schemas": "1.0.8"
  }
}
```

Check available versions:
```bash
npm view @rakeyshgidwani/sunday-schemas versions --json
```

## Documentation Access for AI Agents

Access bundled documentation:
```bash
# View package info
npm info @rakeyshgidwani/sunday-schemas

# Show package contents
npm list @rakeyshgidwani/sunday-schemas --depth=0

# Locate package files
npm list @rakeyshgidwani/sunday-schemas --parseable
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

// Type-safe event handler
function handleEvent(rawData: unknown) {
  if (typeof rawData === 'object' && rawData !== null && 'schema' in rawData) {
    const event = rawData as { schema: string };

    if (isValidSchema(event.schema)) {
      // TypeScript now knows event.schema is EventSchema
      console.log('Valid schema:', event.schema);

      // Use mapped types for type-safe handling
      switch (event.schema) {
        case 'raw.v0':
          // TypeScript knows this is RawEnvelopeV0
          const rawEvent = event as RawEnvelopeV0;
          break;
        case 'md.trade.v1':
          // TypeScript knows this is NormalizedTradeV1
          const tradeEvent = event as NormalizedTradeV1;
          break;
        // ... handle other event types
      }
    }
  }
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
  switch (schema) {
    case 'raw.v0':
      // data is automatically typed as RawEnvelopeV0
      console.log('Processing raw data from venue:', data.venue_id);
      break;
    case 'md.trade.v1':
      // data is automatically typed as NormalizedTradeV1
      console.log('Processing trade:', data.side, data.prob);
      break;
  }
}

// Usage with full type checking
const tradeData: NormalizedTradeV1 = { /* ... */ };
processEvent('md.trade.v1', tradeData); // ✅ Type safe
// processEvent('md.trade.v1', rawData); // ❌ TypeScript error
```

**Union Type Handling:**
```typescript
import { SundayEvent } from '@rakeyshgidwani/sunday-schemas';

// Handle any Sunday event with type narrowing
function handleAnyEvent(event: SundayEvent) {
  // Use discriminated union pattern
  switch (event.schema) {
    case 'raw.v0':
      // TypeScript narrows to RawEnvelopeV0
      console.log('Raw data stream:', event.stream);
      break;
    case 'md.trade.v1':
      // TypeScript narrows to NormalizedTradeV1
      console.log('Trade size:', event.size);
      break;
    case 'insights.arb.lite.v1':
      // TypeScript narrows to ArbitrageLiteV1
      console.log('Arbitrage opportunity profit:', event.profit_bps);
      break;
    // Handle other event types...
  }
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

## Complete Example

```typescript
import {
  // Event types (use aliases for cleaner code)
  RawEnvelope,
  Trade,
  OrderbookDelta,
  ArbitrageLite,

  // Full names for backward compatibility
  RawEnvelopeV0,
  NormalizedTradeV1,

  // Union types and utilities
  SundayEvent,
  EventSchema,
  EventBySchema,

  // Runtime constants
  SCHEMA_CONSTANTS,
  VENUE_IDS,

  // Type definitions
  VenueId,
  TradeSide,
  HealthStatus
} from '@rakeyshgidwani/sunday-schemas';

// Example 1: Creating type-safe events
function createRawEnvelope(venueId: VenueId, instrumentId: string, payload: Record<string, unknown>): RawEnvelope {
  return {
    schema: 'raw.v0',
    venue_id: venueId,
    stream: 'orderbook',
    instrument_native: instrumentId,
    partition_key: `${venueId}:${instrumentId}`,
    ts_event_ms: Date.now(),
    ts_ingest_ms: Date.now(),
    payload
  };
}

// Example 2: Runtime validation with type guards
function isValidEventData(data: unknown): data is SundayEvent {
  if (typeof data !== 'object' || data === null) return false;

  const event = data as Record<string, unknown>;

  // Check required fields
  if (typeof event.schema !== 'string') return false;

  // Validate schema
  return event.schema in SCHEMA_CONSTANTS;
}

// Example 3: Type-safe event processing pipeline
class EventProcessor {
  // Process different event types with full type safety
  processEvent(event: SundayEvent): void {
    switch (event.schema) {
      case 'raw.v0':
        this.handleRawData(event);
        break;
      case 'md.trade.v1':
        this.handleTradeEvent(event);
        break;
      case 'md.orderbook.delta.v1':
        this.handleOrderbookDelta(event);
        break;
      case 'insights.arb.lite.v1':
        this.handleArbitrageOpportunity(event);
        break;
      default:
        // TypeScript ensures exhaustive checking
        const _exhaustive: never = event;
        console.warn('Unhandled event type:', _exhaustive);
    }
  }

  private handleRawData(event: RawEnvelopeV0): void {
    console.log(`Raw data from ${event.venue_id} for ${event.instrument_native}`);

    // Validate venue
    if (!VENUE_IDS.includes(event.venue_id as VenueId)) {
      console.warn('Unknown venue:', event.venue_id);
      return;
    }

    // Process payload based on stream type
    switch (event.stream) {
      case 'orderbook':
        console.log('Processing orderbook data');
        break;
      case 'trades':
        console.log('Processing trade data');
        break;
      default:
        console.log('Processing other stream:', event.stream);
    }
  }

  private handleTradeEvent(event: NormalizedTradeV1): void {
    const tradeType = event.side === 'buy' ? 'purchase' : 'sale';
    const probability = (event.prob * 100).toFixed(1);

    console.log(`${tradeType} of ${event.size} shares at ${probability}% probability`);

    // Type-safe access to all trade properties
    if (event.instrument_id && event.venue_id) {
      console.log(`Trade on ${event.venue_id} for ${event.instrument_id}`);
    }
  }

  private handleOrderbookDelta(event: NormalizedOrderBookDeltaV1): void {
    console.log(`Orderbook update for ${event.instrument_id || 'unknown instrument'}`);

    // Process orderbook changes with type safety
    if (event.bids && event.bids.length > 0) {
      const bestBid = event.bids[0];
      console.log(`Best bid: ${bestBid[0]} @ size ${bestBid[1]}`);
    }

    if (event.asks && event.asks.length > 0) {
      const bestAsk = event.asks[0];
      console.log(`Best ask: ${bestAsk[0]} @ size ${bestAsk[1]}`);
    }
  }

  private handleArbitrageOpportunity(event: ArbitrageLiteV1): void {
    if (event.profit_bps !== undefined) {
      const profitPercent = (event.profit_bps / 100).toFixed(2);
      console.log(`Arbitrage opportunity: ${profitPercent}% profit potential`);
    }
  }
}

// Example 4: Generic utility functions with type safety
class SchemaValidator {
  // Validate and cast to specific event type
  static validateAndCast<T extends EventSchema>(
    data: unknown,
    expectedSchema: T
  ): EventBySchema<T> | null {
    if (!isValidEventData(data)) {
      return null;
    }

    if (data.schema !== expectedSchema) {
      console.error(`Expected schema ${expectedSchema}, got ${data.schema}`);
      return null;
    }

    return data as EventBySchema<T>;
  }

  // Get all valid schemas
  static getAllValidSchemas(): EventSchema[] {
    return Object.keys(SCHEMA_CONSTANTS) as EventSchema[];
  }

  // Check if venue is supported
  static isSupportedVenue(venue: string): venue is VenueId {
    return VENUE_IDS.includes(venue as VenueId);
  }
}

// Example usage
const processor = new EventProcessor();

// Simulate receiving event data
const rawEventData = {
  schema: 'md.trade.v1',
  instrument_id: 'polymarket-123',
  venue_id: 'polymarket',
  side: 'buy' as TradeSide,
  prob: 0.65,
  size: 100,
  ts_ms: Date.now()
};

// Validate and process
if (isValidEventData(rawEventData)) {
  processor.processEvent(rawEventData);
}

// Type-safe schema-specific validation
const tradeEvent = SchemaValidator.validateAndCast(rawEventData, 'md.trade.v1');
if (tradeEvent) {
  // TypeScript knows this is NormalizedTradeV1
  console.log(`Validated trade event with probability: ${tradeEvent.prob}`);
}

// Working with constants
console.log('Supported venues:', VENUE_IDS);
console.log('Available schemas:', SchemaValidator.getAllValidSchemas());
```