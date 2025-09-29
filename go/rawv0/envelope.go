package rawv0

// RawEnvelopeV0 represents the raw venue data envelope from connectors
type RawEnvelopeV0 struct {
	Schema          string                 `json:"schema"`
	VenueID         string                 `json:"venue_id"`
	Stream          string                 `json:"stream"`
	InstrumentNative string                `json:"instrument_native"`
	PartitionKey    string                 `json:"partition_key"`
	TsEventMs       int64                  `json:"ts_event_ms"`
	TsIngestMs      int64                  `json:"ts_ingest_ms"`
	IsHistorical    *bool                  `json:"is_historical,omitempty"`
	BackfillTsMs    *int64                 `json:"backfill_ts_ms,omitempty"`
	Payload         map[string]interface{} `json:"payload"`
}

// VenueID enum values
const (
	VenuePolymarket = "polymarket"
	VenueKalshi     = "kalshi"
)

// Stream enum values
const (
	StreamOrderbook = "orderbook"
	StreamTrades    = "trades"
	StreamStatus    = "status"
)