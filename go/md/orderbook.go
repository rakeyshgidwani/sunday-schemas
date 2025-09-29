package md

// NormalizedOrderBookDeltaV1 represents normalized orderbook deltas with optional snapshots
type NormalizedOrderBookDeltaV1 struct {
	Schema       string      `json:"schema"`
	InstrumentID string      `json:"instrument_id"`
	VenueID      string      `json:"venue_id"`
	Seq          int64       `json:"seq"`
	TsMs         int64       `json:"ts_ms"`
	Bids         [][2]float64 `json:"bids"` // Array of [price, size] pairs
	Asks         [][2]float64 `json:"asks"` // Array of [price, size] pairs
	IsSnapshot   bool        `json:"is_snapshot"`
}

// NormalizedTradeV1 represents normalized trade events with implied probability pricing
type NormalizedTradeV1 struct {
	Schema       string   `json:"schema"`
	InstrumentID string   `json:"instrument_id"`
	VenueID      string   `json:"venue_id"`
	TsMs         int64    `json:"ts_ms"`
	Side         string   `json:"side"`
	Prob         float64  `json:"prob"`
	Size         float64  `json:"size"`
	NotionalUSD  *float64 `json:"notional_usd,omitempty"`
}

// Trade side enum values
const (
	SideBuy  = "buy"
	SideSell = "sell"
)