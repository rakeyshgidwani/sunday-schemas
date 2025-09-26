/**
 * Sunday Schemas - Generated TypeScript Types
 *
 * This file provides a unified export of all generated types from the
 * Sunday platform event schemas and API contracts.
 *
 * Generated from JSON Schema and OpenAPI specifications.
 * DO NOT MODIFY - regenerate using `npm run generate`
 */
export type { RawEnvelopeV0 as RawEnvelope } from './raw.v0.envelope.schema';
export type { NormalizedOrderBookDeltaV1 as OrderbookDelta } from './md.orderbook.delta.v1.schema';
export type { NormalizedTradeV1 as Trade } from './md.trade.v1.schema';
export type { ArbitrageLiteV1 as ArbitrageLite } from './insights.arb.lite.v1.schema';
export type { MoversV1 as Movers } from './insights.movers.v1.schema';
export type { WhaleFlowsLiteV1 as WhalesLite } from './insights.whales.lite.v1.schema';
export type { UnusualActivityV1 as Unusual } from './insights.unusual.v1.schema';
export type { VenueHealthV1 as VenueHealth } from './infra.venue_health.v1.schema';
export type { RawEnvelopeV0 } from './raw.v0.envelope.schema';
export type { NormalizedOrderBookDeltaV1 } from './md.orderbook.delta.v1.schema';
export type { NormalizedTradeV1 } from './md.trade.v1.schema';
export type { ArbitrageLiteV1 } from './insights.arb.lite.v1.schema';
export type { MoversV1 } from './insights.movers.v1.schema';
export type { WhaleFlowsLiteV1 } from './insights.whales.lite.v1.schema';
export type { UnusualActivityV1 } from './insights.unusual.v1.schema';
export type { VenueHealthV1 } from './infra.venue_health.v1.schema';
export * from './api';
export type EventSchema = 'raw.v0' | 'md.orderbook.delta.v1' | 'md.trade.v1' | 'insights.arb.lite.v1' | 'insights.movers.v1' | 'insights.whales.lite.v1' | 'insights.unusual.v1' | 'infra.venue_health.v1';
export type VenueId = 'polymarket' | 'kalshi';
export type TradeSide = 'buy' | 'sell';
export type HealthStatus = 'CONNECTED' | 'DEGRADED' | 'STALE';
import type { RawEnvelopeV0 } from './raw.v0.envelope.schema';
import type { NormalizedOrderBookDeltaV1 } from './md.orderbook.delta.v1.schema';
import type { NormalizedTradeV1 } from './md.trade.v1.schema';
import type { ArbitrageLiteV1 } from './insights.arb.lite.v1.schema';
import type { MoversV1 } from './insights.movers.v1.schema';
import type { WhaleFlowsLiteV1 } from './insights.whales.lite.v1.schema';
import type { UnusualActivityV1 } from './insights.unusual.v1.schema';
import type { VenueHealthV1 } from './infra.venue_health.v1.schema';
export type SundayEvent = RawEnvelopeV0 | NormalizedOrderBookDeltaV1 | NormalizedTradeV1 | ArbitrageLiteV1 | MoversV1 | WhaleFlowsLiteV1 | UnusualActivityV1 | VenueHealthV1;
export type EventBySchema<T extends EventSchema> = T extends 'raw.v0' ? RawEnvelopeV0 : T extends 'md.orderbook.delta.v1' ? NormalizedOrderBookDeltaV1 : T extends 'md.trade.v1' ? NormalizedTradeV1 : T extends 'insights.arb.lite.v1' ? ArbitrageLiteV1 : T extends 'insights.movers.v1' ? MoversV1 : T extends 'insights.whales.lite.v1' ? WhaleFlowsLiteV1 : T extends 'insights.unusual.v1' ? UnusualActivityV1 : T extends 'infra.venue_health.v1' ? VenueHealthV1 : never;
export declare const SCHEMA_CONSTANTS: Record<EventSchema, EventSchema>;
export declare const VENUE_IDS: VenueId[];
