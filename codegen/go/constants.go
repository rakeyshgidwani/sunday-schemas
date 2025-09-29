// Package sundayschemas provides constants and validation for Sunday platform schemas
//
// This file was automatically generated from JSON Schema definitions.
// DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema files,
// and run 'npm run generate-go' to regenerate this file.
package sundayschemas

import (
	"fmt"
)

// EventSchema represents all valid schema identifiers
type EventSchema string

const (
	SchemaINFRA_VENUE_HEALTH_V1 EventSchema = "infra.venue_health.v1"
	SchemaINSIGHTS_ARB_LITE_V1 EventSchema = "insights.arb.lite.v1"
	SchemaINSIGHTS_MOVERS_V1 EventSchema = "insights.movers.v1"
	SchemaINSIGHTS_UNUSUAL_V1 EventSchema = "insights.unusual.v1"
	SchemaINSIGHTS_WHALES_LITE_V1 EventSchema = "insights.whales.lite.v1"
	SchemaMD_ORDERBOOK_DELTA_V1 EventSchema = "md.orderbook.delta.v1"
	SchemaMD_TRADE_V1 EventSchema = "md.trade.v1"
	SchemaRAW_V0 EventSchema = "raw.v0"
)

// Additional venue constants (VenueID type is defined in schemas.go)
const (
	VenuePolymarket VenueID = "polymarket"
	VenueKalshi     VenueID = "kalshi"
)

// TradeSide represents trade directions
type TradeSide string

const (
	TradeSideBuy  TradeSide = "buy"
	TradeSideSell TradeSide = "sell"
)

// HealthStatus represents venue health states
type HealthStatus string

const (
	HealthConnected HealthStatus = "CONNECTED"
	HealthDegraded  HealthStatus = "DEGRADED"
	HealthStale     HealthStatus = "STALE"
)

// ValidateSchema checks if a schema string is valid
func ValidateSchema(schema string) error {
	switch EventSchema(schema) {
	case SchemaINFRA_VENUE_HEALTH_V1, SchemaINSIGHTS_ARB_LITE_V1, SchemaINSIGHTS_MOVERS_V1, SchemaINSIGHTS_UNUSUAL_V1, SchemaINSIGHTS_WHALES_LITE_V1, SchemaMD_ORDERBOOK_DELTA_V1, SchemaMD_TRADE_V1, SchemaRAW_V0:
		return nil
	default:
		return fmt.Errorf("invalid schema: %s", schema)
	}
}

// ValidateVenue checks if a venue ID is valid
func ValidateVenue(venue string) error {
	switch VenueID(venue) {
	case VenuePolymarket, VenueKalshi:
		return nil
	default:
		return fmt.Errorf("invalid venue: %s", venue)
	}
}

// AllSchemas returns all valid schema constants
func AllSchemas() []EventSchema {
	return []EventSchema{
		SchemaINFRA_VENUE_HEALTH_V1,
		SchemaINSIGHTS_ARB_LITE_V1,
		SchemaINSIGHTS_MOVERS_V1,
		SchemaINSIGHTS_UNUSUAL_V1,
		SchemaINSIGHTS_WHALES_LITE_V1,
		SchemaMD_ORDERBOOK_DELTA_V1,
		SchemaMD_TRADE_V1,
		SchemaRAW_V0,
	}
}

// AllVenues returns all valid venue IDs
func AllVenues() []VenueID {
	return []VenueID{VenuePolymarket, VenueKalshi}
}
