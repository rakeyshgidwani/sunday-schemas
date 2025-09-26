// Package sundayschemas provides generated Go types for Sunday platform event schemas
//
// This file was automatically generated from JSON Schema definitions.
// DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema files,
// and run 'npm run generate-go' to regenerate this file.
package sundayschemas

// Schema type definitions for Sunday event platform
// These types correspond to the JSON schemas defined in the schemas/json directory

// RawEnvelope represents the raw.v0.envelope schema
type RawEnvelope struct {
	Schema           string      `json:"schema"`
	VenueID          string      `json:"venue_id"`
	InstrumentNative string      `json:"instrument_native"`
	TsEventMs        int64       `json:"ts_event_ms"`
	TsReceivedMs     int64       `json:"ts_received_ms"`
	SeqNum           int64       `json:"seq_num,omitempty"`
	Payload          interface{} `json:"payload"`
}

// OrderbookDelta represents the md.orderbook.delta.v1 schema
type OrderbookDelta struct {
	Schema       string      `json:"schema"`
	InstrumentID string      `json:"instrument_id"`
	VenueID      string      `json:"venue_id"`
	TsMs         int64       `json:"ts_ms"`
	Seq          int64       `json:"seq"`
	Bids         [][]float64 `json:"bids"`
	Asks         [][]float64 `json:"asks"`
	IsSnapshot   bool        `json:"is_snapshot"`
}

// Trade represents the md.trade.v1 schema
type Trade struct {
	Schema       string   `json:"schema"`
	InstrumentID string   `json:"instrument_id"`
	VenueID      string   `json:"venue_id"`
	TsMs         int64    `json:"ts_ms"`
	Side         string   `json:"side"`
	Prob         float64  `json:"prob"`
	Size         float64  `json:"size"`
	NotionalUsd  *float64 `json:"notional_usd,omitempty"`
}

// ArbitrageLite represents the insights.arb.lite.v1 schema
type ArbitrageLite struct {
	Schema        string  `json:"schema"`
	InstrumentID  string  `json:"instrument_id"`
	LongVenue     string  `json:"long_venue"`
	ShortVenue    string  `json:"short_venue"`
	EdgeBps       float64 `json:"edge_bps"`
	DepthTier     string  `json:"depth_tier"`
	PersistenceMs int64   `json:"persistence_ms"`
	LastSeenMs    int64   `json:"last_seen_ms"`
	FeesIncluded  bool    `json:"fees_included"`
}

// Movers represents the insights.movers.v1 schema
type Movers struct {
	Schema         string  `json:"schema"`
	InstrumentID   string  `json:"instrument_id"`
	Window         string  `json:"window"`
	ProbNow        float64 `json:"prob_now"`
	ProbPrev       float64 `json:"prob_prev"`
	DeltaBps       int     `json:"delta_bps"`
	ImbalanceIndex int     `json:"imbalance_index"`
	TsMs           int64   `json:"ts_ms"`
}

// WhalesLite represents the insights.whales.lite.v1 schema
type WhalesLite struct {
	Schema      string `json:"schema"`
	InstrumentID string `json:"instrument_id"`
	VenueID     string `json:"venue_id"`
	Impact      string `json:"impact"`
	Direction   string `json:"direction"`
	PostMoveBps int    `json:"post_move_bps"`
	TsMs        int64  `json:"ts_ms"`
}

// EventUnusual represents the insights.unusual.v1 schema
// Renamed to avoid conflict with API generated types
type EventUnusual struct {
	Schema       string  `json:"schema"`
	InstrumentID string  `json:"instrument_id"`
	Metric       string  `json:"metric"`
	Window       string  `json:"window"`
	Zscore       float64 `json:"zscore"`
	TsMs         int64   `json:"ts_ms"`
}

// EventVenueHealth represents the infra.venue_health.v1 schema
// Renamed to avoid conflict with API generated types
type EventVenueHealth struct {
	Schema              string   `json:"schema"`
	VenueID             string   `json:"venue_id"`
	Status              string   `json:"status"`
	LastEventTsMs       int64    `json:"last_event_ts_ms"`
	MessagesPerSecond   *float64 `json:"messages_per_second,omitempty"`
	StalenessSeconds    *float64 `json:"staleness_seconds,omitempty"`
	ObservedAtMs        int64    `json:"observed_at_ms"`
}