"use strict";
/**
 * Sunday Schemas - Generated TypeScript Types
 *
 * This file provides a unified export of all generated types from the
 * Sunday platform event schemas and API contracts.
 *
 * Generated from JSON Schema and OpenAPI specifications.
 * DO NOT MODIFY - regenerate using `npm run generate`
 */
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.VENUE_IDS = exports.SCHEMA_CONSTANTS = void 0;
// API Types (OpenAPI generated)
__exportStar(require("./api"), exports);
// Validation helpers
exports.SCHEMA_CONSTANTS = {
    'raw.v0': 'raw.v0',
    'md.orderbook.delta.v1': 'md.orderbook.delta.v1',
    'md.trade.v1': 'md.trade.v1',
    'insights.arb.lite.v1': 'insights.arb.lite.v1',
    'insights.movers.v1': 'insights.movers.v1',
    'insights.whales.lite.v1': 'insights.whales.lite.v1',
    'insights.unusual.v1': 'insights.unusual.v1',
    'infra.venue_health.v1': 'infra.venue_health.v1',
};
exports.VENUE_IDS = ['polymarket', 'kalshi'];
