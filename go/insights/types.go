package insights

// ArbitrageLiteV1 represents arbitrage opportunities between venues (lite version)
type ArbitrageLiteV1 struct {
	Schema         string `json:"schema"`
	InstrumentID   string `json:"instrument_id"`
	LongVenue      string `json:"long_venue"`
	ShortVenue     string `json:"short_venue"`
	EdgeBps        int    `json:"edge_bps"`
	DepthTier      string `json:"depth_tier"`
	PersistenceMs  int64  `json:"persistence_ms"`
	LastSeenMs     int64  `json:"last_seen_ms"`
	FeesIncluded   bool   `json:"fees_included"`
}

// MoversV1 represents price movers over time windows
type MoversV1 struct {
	Schema         string  `json:"schema"`
	InstrumentID   string  `json:"instrument_id"`
	Window         string  `json:"window"`
	ProbNow        float64 `json:"prob_now"`
	ProbPrev       float64 `json:"prob_prev"`
	DeltaBps       int     `json:"delta_bps"`
	ImbalanceIndex float64 `json:"imbalance_index"`
	TsMs           int64   `json:"ts_ms"`
}

// UnusualActivityV1 represents unusual volume or volatility activity detection
type UnusualActivityV1 struct {
	Schema       string  `json:"schema"`
	InstrumentID string  `json:"instrument_id"`
	Metric       string  `json:"metric"`
	Window       string  `json:"window"`
	Zscore       float64 `json:"zscore"`
	TsMs         int64   `json:"ts_ms"`
}

// WhaleFlowsLiteV1 represents large trade flow detection (lite version)
type WhaleFlowsLiteV1 struct {
	Schema       string `json:"schema"`
	InstrumentID string `json:"instrument_id"`
	VenueID      string `json:"venue_id"`
	Impact       string `json:"impact"`
	Direction    string `json:"direction"`
	PostMoveBps  int    `json:"post_move_bps"`
	TsMs         int64  `json:"ts_ms"`
}

// Depth tier enum values
const (
	DepthTierS = "S"
	DepthTierM = "M"
	DepthTierL = "L"
)

// Time window enum values
const (
	Window1h  = "1h"
	Window24h = "24h"
)

// Metric enum values
const (
	MetricVolume     = "volume"
	MetricVolatility = "volatility"
)

// Impact enum values
const (
	ImpactLow  = "LOW"
	ImpactMed  = "MED"
	ImpactHigh = "HIGH"
)