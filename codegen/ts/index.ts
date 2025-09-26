/**
 * Sunday Schemas - Generated TypeScript Types
 *
 * This file provides a unified export of all generated types from the
 * Sunday platform event schemas and API contracts.
 *
 * Generated from JSON Schema and OpenAPI specifications.
 * DO NOT MODIFY - regenerate using `npm run generate`
 */

// Event Schema Types
export type { RawV0Envelope as RawEnvelope } from './raw.v0.envelope.schema';
export type { NormalizedOrderbookDeltaV1 as OrderbookDelta } from './md.orderbook.delta.v1.schema';
export type { NormalizedTradeV1 as Trade } from './md.trade.v1.schema';
export type { ArbitrageLiteV1 as ArbitrageLite } from './insights.arb.lite.v1.schema';
export type { MoversV1 as Movers } from './insights.movers.v1.schema';
export type { WhalesLiteV1 as WhalesLite } from './insights.whales.lite.v1.schema';
export type { UnusualV1 as Unusual } from './insights.unusual.v1.schema';
export type { VenueHealthV1 as VenueHealth } from './infra.venue_health.v1.schema';

// Re-export all types with their original names for backward compatibility
export type { RawV0Envelope } from './raw.v0.envelope.schema';
export type { NormalizedOrderbookDeltaV1 } from './md.orderbook.delta.v1.schema';
export type { NormalizedTradeV1 } from './md.trade.v1.schema';
export type { ArbitrageLiteV1 } from './insights.arb.lite.v1.schema';
export type { MoversV1 } from './insights.movers.v1.schema';
export type { WhalesLiteV1 } from './insights.whales.lite.v1.schema';
export type { UnusualV1 } from './insights.unusual.v1.schema';
export type { VenueHealthV1 } from './infra.venue_health.v1.schema';

// API Types (OpenAPI generated)
export * from './api';

// Type unions for easier use
export type EventSchema =
  | 'raw.v0'
  | 'md.orderbook.delta.v1'
  | 'md.trade.v1'
  | 'insights.arb.lite.v1'
  | 'insights.movers.v1'
  | 'insights.whales.lite.v1'
  | 'insights.unusual.v1'
  | 'infra.venue_health.v1';

export type VenueId = 'polymarket' | 'kalshi';

export type TradeSide = 'buy' | 'sell';

export type HealthStatus = 'connected' | 'degraded' | 'stale';

// Event type union for type-safe event handling
export type SundayEvent =
  | RawV0Envelope
  | NormalizedOrderbookDeltaV1
  | NormalizedTradeV1
  | ArbitrageLiteV1
  | MoversV1
  | WhalesLiteV1
  | UnusualV1
  | VenueHealthV1;

// Utility type for extracting events by schema
export type EventBySchema<T extends EventSchema> =
  T extends 'raw.v0' ? RawV0Envelope :
  T extends 'md.orderbook.delta.v1' ? NormalizedOrderbookDeltaV1 :
  T extends 'md.trade.v1' ? NormalizedTradeV1 :
  T extends 'insights.arb.lite.v1' ? ArbitrageLiteV1 :
  T extends 'insights.movers.v1' ? MoversV1 :
  T extends 'insights.whales.lite.v1' ? WhalesLiteV1 :
  T extends 'insights.unusual.v1' ? UnusualV1 :
  T extends 'infra.venue_health.v1' ? VenueHealthV1 :
  never;

// Validation helpers
export const SCHEMA_CONSTANTS: Record<EventSchema, EventSchema> = {
  'raw.v0': 'raw.v0',
  'md.orderbook.delta.v1': 'md.orderbook.delta.v1',
  'md.trade.v1': 'md.trade.v1',
  'insights.arb.lite.v1': 'insights.arb.lite.v1',
  'insights.movers.v1': 'insights.movers.v1',
  'insights.whales.lite.v1': 'insights.whales.lite.v1',
  'insights.unusual.v1': 'insights.unusual.v1',
  'infra.venue_health.v1': 'infra.venue_health.v1',
} as const;

export const VENUE_IDS: VenueId[] = ['polymarket', 'kalshi'];