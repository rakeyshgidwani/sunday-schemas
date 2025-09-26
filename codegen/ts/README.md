# @sunday/schemas

Generated TypeScript types and API contracts for the Sunday platform event schemas.

## Installation

```bash
npm install @sunday/schemas
# or
yarn add @sunday/schemas
```

## Usage

### Event Types

```typescript
import type {
  SundayEvent,
  Trade,
  OrderbookDelta,
  EventBySchema
} from '@sunday/schemas';

// Type-safe event handling
function processEvent(event: SundayEvent) {
  switch (event.schema) {
    case 'md.trade.v1':
      // TypeScript knows this is a Trade
      console.log(`Trade: ${event.prob} @ ${event.size}`);
      break;
    case 'md.orderbook.delta.v1':
      // TypeScript knows this is OrderbookDelta
      console.log(`Orderbook: ${event.bids.length} bids, ${event.asks.length} asks`);
      break;
  }
}
```

### API Types

```typescript
import type { paths } from '@sunday/schemas';

type GetMarketsResponse = paths['/markets']['get']['responses']['200']['content']['application/json'];
```

### Constants and Utilities

```typescript
import { SCHEMA_CONSTANTS, VENUE_IDS, HealthStatus } from '@sunday/schemas';

// Validate schema
if (event.schema in SCHEMA_CONSTANTS) {
  // Valid schema
}

// Use venue constants
const supportedVenues: VenueId[] = VENUE_IDS; // ['polymarket', 'kalshi']
```

## Generated From

This package is automatically generated from the [Sunday Schemas](https://github.com/rakeyshgidwani/sunday-schemas) repository. Do not modify these files directly - instead update the source schemas and regenerate.

## Schema Versions

- Raw Events: `raw.v0`
- Market Data: `md.orderbook.delta.v1`, `md.trade.v1`
- Insights: `insights.arb.lite.v1`, `insights.movers.v1`, `insights.whales.lite.v1`, `insights.unusual.v1`
- Infrastructure: `infra.venue_health.v1`

## License

MIT