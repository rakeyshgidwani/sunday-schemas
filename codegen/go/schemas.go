// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    venueHealthV1, err := UnmarshalVenueHealthV1(bytes)
//    bytes, err = venueHealthV1.Marshal()
//
//    arbitrageLiteV1, err := UnmarshalArbitrageLiteV1(bytes)
//    bytes, err = arbitrageLiteV1.Marshal()
//
//    moversV1, err := UnmarshalMoversV1(bytes)
//    bytes, err = moversV1.Marshal()
//
//    unusualActivityV1, err := UnmarshalUnusualActivityV1(bytes)
//    bytes, err = unusualActivityV1.Marshal()
//
//    whaleFlowsLiteV1, err := UnmarshalWhaleFlowsLiteV1(bytes)
//    bytes, err = whaleFlowsLiteV1.Marshal()
//
//    normalizedOrderBookDeltaV1, err := UnmarshalNormalizedOrderBookDeltaV1(bytes)
//    bytes, err = normalizedOrderBookDeltaV1.Marshal()
//
//    normalizedTradeV1, err := UnmarshalNormalizedTradeV1(bytes)
//    bytes, err = normalizedTradeV1.Marshal()
//
//    rawEnvelopeV0, err := UnmarshalRawEnvelopeV0(bytes)
//    bytes, err = rawEnvelopeV0.Marshal()

package sundayschemas

import "encoding/json"

func UnmarshalVenueHealthV1(data []byte) (VenueHealthV1, error) {
	var r VenueHealthV1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VenueHealthV1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalArbitrageLiteV1(data []byte) (ArbitrageLiteV1, error) {
	var r ArbitrageLiteV1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ArbitrageLiteV1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMoversV1(data []byte) (MoversV1, error) {
	var r MoversV1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MoversV1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUnusualActivityV1(data []byte) (UnusualActivityV1, error) {
	var r UnusualActivityV1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UnusualActivityV1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhaleFlowsLiteV1(data []byte) (WhaleFlowsLiteV1, error) {
	var r WhaleFlowsLiteV1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhaleFlowsLiteV1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNormalizedOrderBookDeltaV1(data []byte) (NormalizedOrderBookDeltaV1, error) {
	var r NormalizedOrderBookDeltaV1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NormalizedOrderBookDeltaV1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNormalizedTradeV1(data []byte) (NormalizedTradeV1, error) {
	var r NormalizedTradeV1
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NormalizedTradeV1) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRawEnvelopeV0(data []byte) (RawEnvelopeV0, error) {
	var r RawEnvelopeV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RawEnvelopeV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Venue connector health monitoring
type VenueHealthV1 struct {
	LastEventTsMS     int64               `json:"last_event_ts_ms"`
	MessagesPerSecond *float64            `json:"messages_per_second,omitempty"`
	ObservedAtMS      int64               `json:"observed_at_ms"`
	Schema            VenueHealthV1Schema `json:"schema"`
	StalenessSeconds  *float64            `json:"staleness_seconds,omitempty"`
	Status            StatusEnum          `json:"status"`
	VenueID           VenueID             `json:"venue_id"`
}

// Arbitrage opportunities between venues (lite version)
type ArbitrageLiteV1 struct {
	DepthTier     DepthTier             `json:"depth_tier"`
	EdgeBps       float64               `json:"edge_bps"`
	FeesIncluded  bool                  `json:"fees_included"`
	InstrumentID  string                `json:"instrument_id"`
	LastSeenMS    int64                 `json:"last_seen_ms"`
	LongVenue     VenueID               `json:"long_venue"`
	PersistenceMS int64                 `json:"persistence_ms"`
	Schema        ArbitrageLiteV1Schema `json:"schema"`
	ShortVenue    VenueID               `json:"short_venue"`
}

// Price movers over time windows
type MoversV1 struct {
	DeltaBps       int64          `json:"delta_bps"`
	ImbalanceIndex int64          `json:"imbalance_index"`
	InstrumentID   string         `json:"instrument_id"`
	ProbNow        float64        `json:"prob_now"`
	ProbPrev       float64        `json:"prob_prev"`
	Schema         MoversV1Schema `json:"schema"`
	TsMS           int64          `json:"ts_ms"`
	Window         Window         `json:"window"`
}

// Unusual volume or volatility activity detection
type UnusualActivityV1 struct {
	InstrumentID string                  `json:"instrument_id"`
	Metric       Metric                  `json:"metric"`
	Schema       UnusualActivityV1Schema `json:"schema"`
	TsMS         int64                   `json:"ts_ms"`
	Window       Window                  `json:"window"`
	Zscore       float64                 `json:"zscore"`
}

// Large trade flow detection (lite version)
type WhaleFlowsLiteV1 struct {
	Direction    Direction              `json:"direction"`
	Impact       Impact                 `json:"impact"`
	InstrumentID string                 `json:"instrument_id"`
	PostMoveBps  int64                  `json:"post_move_bps"`
	Schema       WhaleFlowsLiteV1Schema `json:"schema"`
	TsMS         int64                  `json:"ts_ms"`
	VenueID      VenueID                `json:"venue_id"`
}

// Normalized orderbook deltas with optional snapshots. Prices are implied probability in
// [0.0, 1.0].
type NormalizedOrderBookDeltaV1 struct {
	// Array of [price, size] pairs where price is implied probability [0.0, 1.0]                                 
	Asks                                                                         [][]float64                      `json:"asks"`
	// Array of [price, size] pairs where price is implied probability [0.0, 1.0]                                 
	Bids                                                                         [][]float64                      `json:"bids"`
	InstrumentID                                                                 string                           `json:"instrument_id"`
	IsSnapshot                                                                   bool                             `json:"is_snapshot"`
	Schema                                                                       NormalizedOrderBookDeltaV1Schema `json:"schema"`
	Seq                                                                          int64                            `json:"seq"`
	TsMS                                                                         int64                            `json:"ts_ms"`
	VenueID                                                                      VenueID                          `json:"venue_id"`
}

// Normalized trade events with implied probability pricing
type NormalizedTradeV1 struct {
	InstrumentID string                  `json:"instrument_id"`
	NotionalUsd  *float64                `json:"notional_usd,omitempty"`
	Prob         float64                 `json:"prob"`
	Schema       NormalizedTradeV1Schema `json:"schema"`
	Side         Direction               `json:"side"`
	Size         float64                 `json:"size"`
	TsMS         int64                   `json:"ts_ms"`
	VenueID      VenueID                 `json:"venue_id"`
}

// Raw venue data envelope from connectors
type RawEnvelopeV0 struct {
	BackfillTsMS     *int64                 `json:"backfill_ts_ms,omitempty"`
	InstrumentNative string                 `json:"instrument_native"`
	IsHistorical     *bool                  `json:"is_historical,omitempty"`
	PartitionKey     string                 `json:"partition_key"`
	Payload          map[string]interface{} `json:"payload"`
	Schema           RawEnvelopeV0Schema    `json:"schema"`
	Stream           Stream                 `json:"stream"`
	TsEventMS        int64                  `json:"ts_event_ms"`
	TsIngestMS       int64                  `json:"ts_ingest_ms"`
	VenueID          VenueID                `json:"venue_id"`
}

type VenueHealthV1Schema string

const (
	InfraVenueHealthV1 VenueHealthV1Schema = "infra.venue_health.v1"
)

type StatusEnum string

const (
	Connected StatusEnum = "CONNECTED"
	Degraded  StatusEnum = "DEGRADED"
	Stale     StatusEnum = "STALE"
)

type VenueID string

const (
	Kalshi     VenueID = "kalshi"
	Polymarket VenueID = "polymarket"
)

type DepthTier string

const (
	L DepthTier = "L"
	M DepthTier = "M"
	S DepthTier = "S"
)

type ArbitrageLiteV1Schema string

const (
	InsightsArbLiteV1 ArbitrageLiteV1Schema = "insights.arb.lite.v1"
)

type MoversV1Schema string

const (
	InsightsMoversV1 MoversV1Schema = "insights.movers.v1"
)

type Window string

const (
	The1H  Window = "1h"
	The24H Window = "24h"
)

type Metric string

const (
	Volatility Metric = "volatility"
	Volume     Metric = "volume"
)

type UnusualActivityV1Schema string

const (
	InsightsUnusualV1 UnusualActivityV1Schema = "insights.unusual.v1"
)

type Direction string

const (
	Buy  Direction = "buy"
	Sell Direction = "sell"
)

type Impact string

const (
	High Impact = "HIGH"
	Low  Impact = "LOW"
	Med  Impact = "MED"
)

type WhaleFlowsLiteV1Schema string

const (
	InsightsWhalesLiteV1 WhaleFlowsLiteV1Schema = "insights.whales.lite.v1"
)

type NormalizedOrderBookDeltaV1Schema string

const (
	MdOrderbookDeltaV1 NormalizedOrderBookDeltaV1Schema = "md.orderbook.delta.v1"
)

type NormalizedTradeV1Schema string

const (
	MdTradeV1 NormalizedTradeV1Schema = "md.trade.v1"
)

type RawEnvelopeV0Schema string

const (
	RawV0 RawEnvelopeV0Schema = "raw.v0"
)

type Stream string

const (
	Orderbook Stream = "orderbook"
	Status    Stream = "status"
	Trades    Stream = "trades"
)
