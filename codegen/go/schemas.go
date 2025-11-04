// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    eventMetadataV0, err := UnmarshalEventMetadataV0(bytes)
//    bytes, err = eventMetadataV0.Marshal()
//
//    eventDiscoveryPayloadV0, err := UnmarshalEventDiscoveryPayloadV0(bytes)
//    bytes, err = eventDiscoveryPayloadV0.Marshal()
//
//    seriesMetadataV0, err := UnmarshalSeriesMetadataV0(bytes)
//    bytes, err = seriesMetadataV0.Marshal()
//
//    seriesDiscoveryPayloadV0, err := UnmarshalSeriesDiscoveryPayloadV0(bytes)
//    bytes, err = seriesDiscoveryPayloadV0.Marshal()
//
//    discoverySharedTypesV0, err := UnmarshalDiscoverySharedTypesV0(bytes)
//    bytes, err = discoverySharedTypesV0.Marshal()
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
//    rawCategoriesDiscoveryV0, err := UnmarshalRawCategoriesDiscoveryV0(bytes)
//    bytes, err = rawCategoriesDiscoveryV0.Marshal()
//
//    rawEventsDiscoveryV0, err := UnmarshalRawEventsDiscoveryV0(bytes)
//    bytes, err = rawEventsDiscoveryV0.Marshal()
//
//    rawSeriesDiscoveryV0, err := UnmarshalRawSeriesDiscoveryV0(bytes)
//    bytes, err = rawSeriesDiscoveryV0.Marshal()
//
//    rawEnvelopeV0, err := UnmarshalRawEnvelopeV0(bytes)
//    bytes, err = rawEnvelopeV0.Marshal()

package sundayschemas

import "time"

import "encoding/json"

func UnmarshalEventMetadataV0(data []byte) (EventMetadataV0, error) {
	var r EventMetadataV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventMetadataV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEventDiscoveryPayloadV0(data []byte) (EventDiscoveryPayloadV0, error) {
	var r EventDiscoveryPayloadV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventDiscoveryPayloadV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSeriesMetadataV0(data []byte) (SeriesMetadataV0, error) {
	var r SeriesMetadataV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SeriesMetadataV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSeriesDiscoveryPayloadV0(data []byte) (SeriesDiscoveryPayloadV0, error) {
	var r SeriesDiscoveryPayloadV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SeriesDiscoveryPayloadV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DiscoverySharedTypesV0 map[string]interface{}

func UnmarshalDiscoverySharedTypesV0(data []byte) (DiscoverySharedTypesV0, error) {
	var r DiscoverySharedTypesV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscoverySharedTypesV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

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

func UnmarshalRawCategoriesDiscoveryV0(data []byte) (RawCategoriesDiscoveryV0, error) {
	var r RawCategoriesDiscoveryV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RawCategoriesDiscoveryV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRawEventsDiscoveryV0(data []byte) (RawEventsDiscoveryV0, error) {
	var r RawEventsDiscoveryV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RawEventsDiscoveryV0) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRawSeriesDiscoveryV0(data []byte) (RawSeriesDiscoveryV0, error) {
	var r RawSeriesDiscoveryV0
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RawSeriesDiscoveryV0) Marshal() ([]byte, error) {
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

// Structured metadata for individual prediction market events
type EventMetadataV0 struct {
	// Whether the event is currently active                                       
	Active                                                  bool                   `json:"active"`
	// Event category                                                              
	Category                                                *string                `json:"category,omitempty"`
	// Whether the event is closed for trading                                     
	Closed                                                  bool                   `json:"closed"`
	// Event description                                                           
	Description                                             *string                `json:"description,omitempty"`
	// When this event was first discovered                                        
	DiscoveredAt                                            time.Time              `json:"discovered_at"`
	// Event end date/time                                                         
	EndDate                                                 *time.Time             `json:"end_date,omitempty"`
	// Venue-specific event identifier                                             
	EventID                                                 string                 `json:"event_id"`
	// Venue-specific fields that don't fit canonical schema                       
	ExtraMetadata                                           map[string]interface{} `json:"extra_metadata,omitempty"`
	// Discriminator field for event metadata                                      
	Kind                                                    EventMetadataV0Kind    `json:"kind"`
	// When this event was last seen in discovery                                  
	LastSeen                                                time.Time              `json:"last_seen"`
	// Series ID this event belongs to                                             
	ParentSeriesID                                          *string                `json:"parent_series_id,omitempty"`
	// Series title for convenience                                                
	ParentSeriesTitle                                       *string                `json:"parent_series_title,omitempty"`
	Relationships                                           *Relationships         `json:"relationships,omitempty"`
	// Event start date/time                                                       
	StartDate                                               *time.Time             `json:"start_date,omitempty"`
	// Structured tags for categorization                                          
	Tags                                                    []string               `json:"tags,omitempty"`
	// Event title                                                                 
	Title                                                   string                 `json:"title"`
	// Venue identifier from venues.json registry                                  
	VenueID                                                 VenueID                `json:"venue_id"`
}

// Parent/child relationship mappings
type Relationships struct {
	// Child event identifiers                      
	EventIDS                               []string `json:"event_ids,omitempty"`
	// Related venue instrument identifiers         
	InstrumentIDS                          []string `json:"instrument_ids,omitempty"`
	// Parent series identifier                     
	SeriesID                               *string  `json:"series_id,omitempty"`
}

// Structured payload for event discovery messages
type EventDiscoveryPayloadV0 struct {
	DiscoveryMeta                                        *Discovery `json:"discovery_meta,omitempty"`
	Event                                                EventClass `json:"event"`
	// Unique event identifier for this discovery message           
	EventID                                              string     `json:"event_id"`
	// Type of discovery event                                      
	EventType                                            EventType  `json:"event_type"`
	// When this discovery event occurred                           
	Timestamp                                            time.Time  `json:"timestamp"`
	// Venue identifier from venues.json registry                   
	VenueID                                              VenueID    `json:"venue_id"`
}

// Metadata about the discovery batch/run for monitoring and sequencing
type Discovery struct {
	// Unique identifier for this discovery batch           
	BatchID                                          string `json:"batch_id"`
	// Position of this item within the batch               
	BatchSequence                                    int64  `json:"batch_sequence"`
	// Total number of items in this batch                  
	BatchTotalCount                                  int64  `json:"batch_total_count"`
	// Unique identifier for the entire discovery run       
	DiscoveryRunID                                   string `json:"discovery_run_id"`
}

// Structured metadata for individual prediction market events
type EventClass struct {
	// Whether the event is currently active                                       
	Active                                                  bool                   `json:"active"`
	// Event category                                                              
	Category                                                *string                `json:"category,omitempty"`
	// Whether the event is closed for trading                                     
	Closed                                                  bool                   `json:"closed"`
	// Event description                                                           
	Description                                             *string                `json:"description,omitempty"`
	// When this event was first discovered                                        
	DiscoveredAt                                            time.Time              `json:"discovered_at"`
	// Event end date/time                                                         
	EndDate                                                 *time.Time             `json:"end_date,omitempty"`
	// Venue-specific event identifier                                             
	EventID                                                 string                 `json:"event_id"`
	// Venue-specific fields that don't fit canonical schema                       
	ExtraMetadata                                           map[string]interface{} `json:"extra_metadata,omitempty"`
	// Discriminator field for event metadata                                      
	Kind                                                    EventMetadataV0Kind    `json:"kind"`
	// When this event was last seen in discovery                                  
	LastSeen                                                time.Time              `json:"last_seen"`
	// Series ID this event belongs to                                             
	ParentSeriesID                                          *string                `json:"parent_series_id,omitempty"`
	// Series title for convenience                                                
	ParentSeriesTitle                                       *string                `json:"parent_series_title,omitempty"`
	Relationships                                           *Relationships         `json:"relationships,omitempty"`
	// Event start date/time                                                       
	StartDate                                               *time.Time             `json:"start_date,omitempty"`
	// Structured tags for categorization                                          
	Tags                                                    []string               `json:"tags,omitempty"`
	// Event title                                                                 
	Title                                                   string                 `json:"title"`
	// Venue identifier from venues.json registry                                  
	VenueID                                                 VenueID                `json:"venue_id"`
}

// Structured metadata for series/collections of prediction market events
type SeriesMetadataV0 struct {
	// Whether the series is currently active                                            
	Active                                                        bool                   `json:"active"`
	// Series category                                                                   
	Category                                                      *string                `json:"category,omitempty"`
	// Event IDs that belong to this series                                              
	ChildEventIDS                                                 []string               `json:"child_event_ids,omitempty"`
	// Whether the series is closed                                                      
	Closed                                                        bool                   `json:"closed"`
	// Series description                                                                
	Description                                                   *string                `json:"description,omitempty"`
	// When this series was first discovered                                             
	DiscoveredAt                                                  time.Time              `json:"discovered_at"`
	// Series identifier (note: field name kept for compatibility)                       
	EventID                                                       string                 `json:"event_id"`
	// Venue-specific fields that don't fit canonical schema                             
	ExtraMetadata                                                 map[string]interface{} `json:"extra_metadata,omitempty"`
	// Discriminator field for series metadata                                           
	Kind                                                          SeriesMetadataV0Kind   `json:"kind"`
	// When this series was last seen in discovery                                       
	LastSeen                                                      time.Time              `json:"last_seen"`
	Relationships                                                 *Relationships         `json:"relationships,omitempty"`
	SeriesData                                                    *SeriesData            `json:"series_data,omitempty"`
	// Structured tags for categorization                                                
	Tags                                                          []string               `json:"tags,omitempty"`
	// Series title                                                                      
	Title                                                         string                 `json:"title"`
	// Venue identifier from venues.json registry                                        
	VenueID                                                       VenueID                `json:"venue_id"`
}

// Series-specific fields
type SeriesData struct {
	Contract                         *Contract    `json:"contract,omitempty"`
	Creators                         *Creators    `json:"creators,omitempty"`
	Financial                        *Financial   `json:"financial,omitempty"`
	IconURL                          *string      `json:"icon_url,omitempty"`
	ImageURL                         *string      `json:"image_url,omitempty"`
	// UI layout hint                             
	Layout                           *string      `json:"layout,omitempty"`
	// Series recurrence pattern                  
	Recurrence                       *string      `json:"recurrence,omitempty"`
	// Type/classification of series              
	SeriesType                       *string      `json:"series_type,omitempty"`
	// URL-friendly series identifier             
	Slug                             *string      `json:"slug,omitempty"`
	Status                           *StatusClass `json:"status,omitempty"`
	// Short series subtitle                      
	Subtitle                         *string      `json:"subtitle,omitempty"`
	// Series ticker/symbol                       
	Ticker                           *string      `json:"ticker,omitempty"`
	Timestamps                       *Timestamps  `json:"timestamps,omitempty"`
}

type Contract struct {
	AdditionalProhibitions   []string                  `json:"additional_prohibitions,omitempty"`
	ContractTermsURL         *string                   `json:"contract_terms_url,omitempty"`
	ContractURL              *string                   `json:"contract_url,omitempty"`
	FeeMultiplier            *float64                  `json:"fee_multiplier,omitempty"`
	// Fee calculation method                          
	FeeType                  *string                   `json:"fee_type,omitempty"`
	SettlementSources        []DiscoverySharedV0Schema `json:"settlement_sources,omitempty"`
}

type DiscoverySharedV0Schema struct {
	Name string  `json:"name"`
	URL  *string `json:"url,omitempty"`
}

type Creators struct {
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

type Financial struct {
	// Currency unit for monetary values                           
	Currency                                             *Currency `json:"currency,omitempty"`
	// Total USD liquidity, decimal precision to 2 places          
	LiquidityTotalUsd                                    *float64  `json:"liquidity_total_usd,omitempty"`
	// Ranking/scoring metric (unitless)                           
	Score                                                *float64  `json:"score,omitempty"`
	// 24-hour contract volume count                               
	Volume24HContracts                                   *int64    `json:"volume_24h_contracts,omitempty"`
	// 24-hour USD volume, decimal precision to 2 places           
	Volume24HUsd                                         *float64  `json:"volume_24h_usd,omitempty"`
	// Total contract volume count                                 
	VolumeTotalContracts                                 *int64    `json:"volume_total_contracts,omitempty"`
	// Total USD volume, decimal precision to 2 places             
	VolumeTotalUsd                                       *float64  `json:"volume_total_usd,omitempty"`
}

type StatusClass struct {
	Archived                        *bool   `json:"archived,omitempty"`
	CommentsEnabled                 *bool   `json:"comments_enabled,omitempty"`
	// Competitive mode/flag                
	Competitive                     *string `json:"competitive,omitempty"`
	Featured                        *bool   `json:"featured,omitempty"`
	// Newly featured series                
	IsNew                           *bool   `json:"is_new,omitempty"`
	// Template for event generation        
	IsTemplate                      *bool   `json:"is_template,omitempty"`
	// Access restrictions apply            
	Restricted                      *bool   `json:"restricted,omitempty"`
}

type Timestamps struct {
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// Structured payload for series discovery messages
type SeriesDiscoveryPayloadV0 struct {
	DiscoveryMeta                                        *Discovery                    `json:"discovery_meta,omitempty"`
	Event                                                SeriesDiscoveryPayloadV0Event `json:"event"`
	// Unique event identifier for this discovery message                              
	EventID                                              string                        `json:"event_id"`
	// Type of discovery event                                                         
	EventType                                            EventType                     `json:"event_type"`
	// When this discovery event occurred                                              
	Timestamp                                            time.Time                     `json:"timestamp"`
	// Venue identifier from venues.json registry                                      
	VenueID                                              VenueID                       `json:"venue_id"`
}

// Structured metadata for series/collections of prediction market events
type SeriesDiscoveryPayloadV0Event struct {
	// Whether the series is currently active                                            
	Active                                                        bool                   `json:"active"`
	// Series category                                                                   
	Category                                                      *string                `json:"category,omitempty"`
	// Event IDs that belong to this series                                              
	ChildEventIDS                                                 []string               `json:"child_event_ids,omitempty"`
	// Whether the series is closed                                                      
	Closed                                                        bool                   `json:"closed"`
	// Series description                                                                
	Description                                                   *string                `json:"description,omitempty"`
	// When this series was first discovered                                             
	DiscoveredAt                                                  time.Time              `json:"discovered_at"`
	// Series identifier (note: field name kept for compatibility)                       
	EventID                                                       string                 `json:"event_id"`
	// Venue-specific fields that don't fit canonical schema                             
	ExtraMetadata                                                 map[string]interface{} `json:"extra_metadata,omitempty"`
	// Discriminator field for series metadata                                           
	Kind                                                          SeriesMetadataV0Kind   `json:"kind"`
	// When this series was last seen in discovery                                       
	LastSeen                                                      time.Time              `json:"last_seen"`
	Relationships                                                 *Relationships         `json:"relationships,omitempty"`
	SeriesData                                                    *SeriesData            `json:"series_data,omitempty"`
	// Structured tags for categorization                                                
	Tags                                                          []string               `json:"tags,omitempty"`
	// Series title                                                                      
	Title                                                         string                 `json:"title"`
	// Venue identifier from venues.json registry                                        
	VenueID                                                       VenueID                `json:"venue_id"`
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

// Category/tag discovery data for unified taxonomy from prediction market venues
type RawCategoriesDiscoveryV0 struct {
	Envelope RawCategoriesDiscoveryV0Envelope `json:"envelope"`
	Payload  map[string]interface{}           `json:"payload"`
}

type RawCategoriesDiscoveryV0Envelope struct {
	Metadata  *PurpleMetadata `json:"metadata,omitempty"`
	Schema    PurpleSchema    `json:"schema"`
	Stream    PurpleStream    `json:"stream"`
	Timestamp time.Time       `json:"timestamp"`
	VenueID   VenueID         `json:"venue_id"`
}

type PurpleMetadata struct {
	DiscoveryTimestamp *time.Time `json:"discovery_timestamp,omitempty"`
}

// Event discovery data from prediction market venues
type RawEventsDiscoveryV0 struct {
	Envelope RawEventsDiscoveryV0Envelope `json:"envelope"`
	Payload  PayloadClass                 `json:"payload"`
}

type RawEventsDiscoveryV0Envelope struct {
	Metadata  *FluffyMetadata `json:"metadata,omitempty"`
	Schema    FluffySchema    `json:"schema"`
	Stream    FluffyStream    `json:"stream"`
	Timestamp time.Time       `json:"timestamp"`
	VenueID   VenueID         `json:"venue_id"`
}

type FluffyMetadata struct {
	DiscoveryPage      *int64     `json:"discovery_page,omitempty"`
	DiscoveryTimestamp *time.Time `json:"discovery_timestamp,omitempty"`
}

// Structured payload for event discovery messages
type PayloadClass struct {
	DiscoveryMeta                                        *Discovery  `json:"discovery_meta,omitempty"`
	Event                                                *EventClass `json:"event,omitempty"`
	// Unique event identifier for this discovery message            
	EventID                                              *string     `json:"event_id,omitempty"`
	// Type of discovery event                                       
	EventType                                            *EventType  `json:"event_type,omitempty"`
	// When this discovery event occurred                            
	Timestamp                                            *time.Time  `json:"timestamp,omitempty"`
	// Venue identifier from venues.json registry                    
	VenueID                                              *VenueID    `json:"venue_id,omitempty"`
}

// Series/collections discovery data from prediction market venues
type RawSeriesDiscoveryV0 struct {
	Envelope RawSeriesDiscoveryV0Envelope `json:"envelope"`
	Payload  RawSeriesDiscoveryV0Payload  `json:"payload"`
}

type RawSeriesDiscoveryV0Envelope struct {
	Metadata  *TentacledMetadata `json:"metadata,omitempty"`
	Schema    TentacledSchema    `json:"schema"`
	Stream    TentacledStream    `json:"stream"`
	Timestamp time.Time          `json:"timestamp"`
	VenueID   VenueID            `json:"venue_id"`
}

type TentacledMetadata struct {
	DiscoveryPage      *int64     `json:"discovery_page,omitempty"`
	DiscoveryTimestamp *time.Time `json:"discovery_timestamp,omitempty"`
}

// Structured payload for series discovery messages
type RawSeriesDiscoveryV0Payload struct {
	DiscoveryMeta                                        *Discovery                     `json:"discovery_meta,omitempty"`
	Event                                                *SeriesDiscoveryPayloadV0Event `json:"event,omitempty"`
	// Unique event identifier for this discovery message                               
	EventID                                              *string                        `json:"event_id,omitempty"`
	// Type of discovery event                                                          
	EventType                                            *EventType                     `json:"event_type,omitempty"`
	// When this discovery event occurred                                               
	Timestamp                                            *time.Time                     `json:"timestamp,omitempty"`
	// Venue identifier from venues.json registry                                       
	VenueID                                              *VenueID                       `json:"venue_id,omitempty"`
}

// Raw venue data envelope from connectors
type RawEnvelopeV0 struct {
	BackfillTsMS     *int64                 `json:"backfill_ts_ms,omitempty"`
	InstrumentNative string                 `json:"instrument_native"`
	IsHistorical     *bool                  `json:"is_historical,omitempty"`
	PartitionKey     string                 `json:"partition_key"`
	Payload          map[string]interface{} `json:"payload"`
	Schema           RawEnvelopeV0Schema    `json:"schema"`
	Stream           RawEnvelopeV0Stream    `json:"stream"`
	TsEventMS        int64                  `json:"ts_event_ms"`
	TsIngestMS       int64                  `json:"ts_ingest_ms"`
	VenueID          VenueID                `json:"venue_id"`
}

type EventMetadataV0Kind string

const (
	Event EventMetadataV0Kind = "event"
)

// Venue identifier from venues.json registry
type VenueID string

const (
	Kalshi     VenueID = "kalshi"
	Polymarket VenueID = "polymarket"
)

// Type of discovery event
type EventType string

const (
	Discovered EventType = "discovered"
	Expired    EventType = "expired"
	Updated    EventType = "updated"
)

type SeriesMetadataV0Kind string

const (
	Series SeriesMetadataV0Kind = "series"
)

// Currency unit for monetary values
type Currency string

const (
	Usd Currency = "USD"
)

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

type PurpleSchema string

const (
	RawCategoriesV0 PurpleSchema = "raw.categories.v0"
)

type PurpleStream string

const (
	CategoryDiscovery PurpleStream = "category_discovery"
)

type FluffySchema string

const (
	RawEventsV0 FluffySchema = "raw.events.v0"
)

type FluffyStream string

const (
	EventDiscovery FluffyStream = "event_discovery"
)

type TentacledSchema string

const (
	RawSeriesV0 TentacledSchema = "raw.series.v0"
)

type TentacledStream string

const (
	SeriesDiscovery TentacledStream = "series_discovery"
)

type RawEnvelopeV0Schema string

const (
	RawV0 RawEnvelopeV0Schema = "raw.v0"
)

type RawEnvelopeV0Stream string

const (
	Orderbook RawEnvelopeV0Stream = "orderbook"
	Status    RawEnvelopeV0Stream = "status"
	Trades    RawEnvelopeV0Stream = "trades"
)
