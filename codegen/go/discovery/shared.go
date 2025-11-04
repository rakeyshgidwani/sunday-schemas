// Code generated for discovery payload schemas. DO NOT EDIT.

package discovery

import "time"

// DiscoveryMetaV0 contains metadata about the discovery batch/run for monitoring and sequencing
type DiscoveryMetaV0 struct {
	BatchID          string `json:"batch_id"`
	BatchSequence    int    `json:"batch_sequence"`
	BatchTotalCount  int    `json:"batch_total_count"`
	DiscoveryRunID   string `json:"discovery_run_id"`
}

// RelationshipsV0 contains parent/child relationship mappings
type RelationshipsV0 struct {
	SeriesID      *string  `json:"series_id,omitempty"`
	EventIDs      []string `json:"event_ids,omitempty"`
	InstrumentIDs []string `json:"instrument_ids,omitempty"`
}

// FinancialDataV0 contains financial metrics and volume data
type FinancialDataV0 struct {
	Volume24hUSD         *float64 `json:"volume_24h_usd,omitempty"`
	VolumeTotalUSD       *float64 `json:"volume_total_usd,omitempty"`
	LiquidityTotalUSD    *float64 `json:"liquidity_total_usd,omitempty"`
	Volume24hContracts   *int     `json:"volume_24h_contracts,omitempty"`
	VolumeTotalContracts *int     `json:"volume_total_contracts,omitempty"`
	Score                *float64 `json:"score,omitempty"`
	Currency             *string  `json:"currency,omitempty"`
}

// StatusDataV0 contains status flags and metadata
type StatusDataV0 struct {
	Archived        *bool   `json:"archived,omitempty"`
	IsNew           *bool   `json:"is_new,omitempty"`
	Featured        *bool   `json:"featured,omitempty"`
	Restricted      *bool   `json:"restricted,omitempty"`
	IsTemplate      *bool   `json:"is_template,omitempty"`
	Competitive     *string `json:"competitive,omitempty"`
	CommentsEnabled *bool   `json:"comments_enabled,omitempty"`
}

// ContractDataV0 contains contract and settlement information
type ContractDataV0 struct {
	ContractURL             *string               `json:"contract_url,omitempty"`
	ContractTermsURL        *string               `json:"contract_terms_url,omitempty"`
	FeeType                 *string               `json:"fee_type,omitempty"`
	FeeMultiplier           *float64              `json:"fee_multiplier,omitempty"`
	AdditionalProhibitions  []string              `json:"additional_prohibitions,omitempty"`
	SettlementSources       []SettlementSourceV0  `json:"settlement_sources,omitempty"`
}

// SettlementSourceV0 contains settlement source information
type SettlementSourceV0 struct {
	Name string  `json:"name"`
	URL  *string `json:"url,omitempty"`
}

// TimestampDataV0 contains timestamp information
type TimestampDataV0 struct {
	PublishedAt *time.Time `json:"published_at,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// CreatorDataV0 contains creator information
type CreatorDataV0 struct {
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// SeriesDataV0 contains series-specific fields
type SeriesDataV0 struct {
	Ticker     *string           `json:"ticker,omitempty"`
	Slug       *string           `json:"slug,omitempty"`
	Subtitle   *string           `json:"subtitle,omitempty"`
	SeriesType *string           `json:"series_type,omitempty"`
	Recurrence *string           `json:"recurrence,omitempty"`
	ImageURL   *string           `json:"image_url,omitempty"`
	IconURL    *string           `json:"icon_url,omitempty"`
	Layout     *string           `json:"layout,omitempty"`
	Financial  *FinancialDataV0  `json:"financial,omitempty"`
	Status     *StatusDataV0     `json:"status,omitempty"`
	Contract   *ContractDataV0   `json:"contract,omitempty"`
	Timestamps *TimestampDataV0  `json:"timestamps,omitempty"`
	Creators   *CreatorDataV0    `json:"creators,omitempty"`
}