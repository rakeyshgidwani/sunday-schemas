// Package sundayschemas provides compatibility types for legacy integrations
//
// This file provides compatibility aliases for projects that were using
// the old type definitions. New projects should use the types from schemas.go directly.
//
// DO NOT MODIFY THIS FILE BY HAND. This file maintains backward compatibility
// and will be updated when the underlying schemas change.
package sundayschemas

import "time"

// Legacy type aliases for backward compatibility
// These provide the interface that sunday-connectors expects

// Stream is a compatibility alias for RawEnvelopeV0Stream
// Maintains backward compatibility with existing sunday-connectors code
type Stream = RawEnvelopeV0Stream

// Stream constants for backward compatibility
const (
	StreamOrderbook         Stream = Orderbook
	StreamTrades           Stream = Trades
	StreamStatus           Stream = Status
	StreamEventDiscovery   Stream = "event_discovery"
	StreamSeriesDiscovery  Stream = "series_discovery"
	StreamCategoryDiscovery Stream = "category_discovery"
)

// RawEnvelope is a compatibility alias for RawEnvelopeV0
// This maintains compatibility with existing code that expects the old field names
type RawEnvelope struct {
	Schema           string      `json:"schema"`
	VenueID          string      `json:"venue_id"`
	Stream           string      `json:"stream"`
	InstrumentNative string      `json:"instrument_native"`
	PartitionKey     string      `json:"partition_key"`
	TsEventMs        int64       `json:"ts_event_ms"`   // Note: lowercase 's' for compatibility
	TsIngestMs       int64       `json:"ts_ingest_ms"`  // Note: lowercase 's' for compatibility
	IsHistorical     *bool       `json:"is_historical,omitempty"`
	BackfillTsMs     *int64      `json:"backfill_ts_ms,omitempty"`
	Payload          interface{} `json:"payload"`
}

// NewRawEnvelope creates a RawEnvelope using the new schema types
// This function bridges between the old interface and new implementation
func NewRawEnvelope(venueID, stream, instrument string, eventTS time.Time, payload interface{}) RawEnvelope {
	now := time.Now().UTC()
	return RawEnvelope{
		Schema:           string(SchemaRAW_V0),
		VenueID:          venueID,
		Stream:           stream,
		InstrumentNative: instrument,
		PartitionKey:     venueID + ":" + instrument,
		TsEventMs:        eventTS.UnixMilli(),
		TsIngestMs:       now.UnixMilli(),
		Payload:          payload,
	}
}

// ToRawEnvelopeV0 converts a legacy RawEnvelope to the canonical RawEnvelopeV0
func (r RawEnvelope) ToRawEnvelopeV0() RawEnvelopeV0 {
	return RawEnvelopeV0{
		Schema:           RawV0, // Use the canonical schema constant
		VenueID:          VenueID(r.VenueID),
		Stream:           RawEnvelopeV0Stream(r.Stream),
		InstrumentNative: r.InstrumentNative,
		PartitionKey:     r.PartitionKey,
		TsEventMS:        r.TsEventMs,   // Convert lowercase to uppercase
		TsIngestMS:       r.TsIngestMs,  // Convert lowercase to uppercase
		IsHistorical:     r.IsHistorical,
		BackfillTsMS:     r.BackfillTsMs,
		Payload:          r.Payload.(map[string]interface{}),
	}
}

// FromRawEnvelopeV0 converts a canonical RawEnvelopeV0 to legacy RawEnvelope
func FromRawEnvelopeV0(env RawEnvelopeV0) RawEnvelope {
	return RawEnvelope{
		Schema:           string(env.Schema),
		VenueID:          string(env.VenueID),
		Stream:           string(env.Stream),
		InstrumentNative: env.InstrumentNative,
		PartitionKey:     env.PartitionKey,
		TsEventMs:        env.TsEventMS,   // Convert uppercase to lowercase
		TsIngestMs:       env.TsIngestMS,  // Convert uppercase to lowercase
		IsHistorical:     env.IsHistorical,
		BackfillTsMs:     env.BackfillTsMS,
		Payload:          env.Payload,
	}
}

// OrderbookDelta provides compatibility type for market data
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

// Trade provides compatibility type for trade events
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

// ToNormalizedTradeV1 converts a legacy Trade to canonical NormalizedTradeV1
func (t Trade) ToNormalizedTradeV1() NormalizedTradeV1 {
	return NormalizedTradeV1{
		Schema:       MdTradeV1,
		InstrumentID: t.InstrumentID,
		VenueID:      VenueID(t.VenueID),
		TsMS:         t.TsMs,
		Side:         Direction(t.Side),
		Prob:         t.Prob,
		Size:         t.Size,
		NotionalUsd:  t.NotionalUsd,
	}
}

// FromNormalizedTradeV1 converts canonical NormalizedTradeV1 to legacy Trade
func FromNormalizedTradeV1(trade NormalizedTradeV1) Trade {
	return Trade{
		Schema:       string(trade.Schema),
		InstrumentID: trade.InstrumentID,
		VenueID:      string(trade.VenueID),
		TsMs:         trade.TsMS,
		Side:         string(trade.Side),
		Prob:         trade.Prob,
		Size:         trade.Size,
		NotionalUsd:  trade.NotionalUsd,
	}
}