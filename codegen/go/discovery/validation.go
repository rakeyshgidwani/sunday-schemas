// Code generated for discovery payload schemas. DO NOT EDIT.

package discovery

import (
	"encoding/json"
	"fmt"
	"math"
)

// ValidationError represents a validation error with context
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

// ValidateEventDiscoveryPayload validates an event discovery payload
func ValidateEventDiscoveryPayload(payload []byte) error {
	var p EventDiscoveryPayloadV0
	if err := json.Unmarshal(payload, &p); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if err := validateEventMetadata(p.Event); err != nil {
		return fmt.Errorf("event metadata validation failed: %w", err)
	}

	if p.EventID == "" {
		return ValidationError{Field: "event_id", Message: "required field is empty"}
	}

	if !isValidEventType(p.EventType) {
		return ValidationError{Field: "event_type", Message: "invalid event type"}
	}

	if p.Timestamp.IsZero() {
		return ValidationError{Field: "timestamp", Message: "required field is zero"}
	}

	if !isValidVenueID(p.VenueID) {
		return ValidationError{Field: "venue_id", Message: "invalid venue ID"}
	}

	if p.DiscoveryMeta != nil {
		if err := validateDiscoveryMeta(*p.DiscoveryMeta); err != nil {
			return fmt.Errorf("discovery_meta validation failed: %w", err)
		}
	}

	return nil
}

// ValidateSeriesDiscoveryPayload validates a series discovery payload
func ValidateSeriesDiscoveryPayload(payload []byte) error {
	var p SeriesDiscoveryPayloadV0
	if err := json.Unmarshal(payload, &p); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if err := validateSeriesMetadata(p.Event); err != nil {
		return fmt.Errorf("series metadata validation failed: %w", err)
	}

	if p.EventID == "" {
		return ValidationError{Field: "event_id", Message: "required field is empty"}
	}

	if !isValidEventType(p.EventType) {
		return ValidationError{Field: "event_type", Message: "invalid event type"}
	}

	if p.Timestamp.IsZero() {
		return ValidationError{Field: "timestamp", Message: "required field is zero"}
	}

	if !isValidVenueID(p.VenueID) {
		return ValidationError{Field: "venue_id", Message: "invalid venue ID"}
	}

	if p.DiscoveryMeta != nil {
		if err := validateDiscoveryMeta(*p.DiscoveryMeta); err != nil {
			return fmt.Errorf("discovery_meta validation failed: %w", err)
		}
	}

	return nil
}

// ValidateEventMetadata validates event metadata
func ValidateEventMetadata(metadata EventMetadataV0) error {
	return validateEventMetadata(metadata)
}

// ValidateSeriesMetadata validates series metadata
func ValidateSeriesMetadata(metadata SeriesMetadataV0) error {
	return validateSeriesMetadata(metadata)
}

func validateEventMetadata(metadata EventMetadataV0) error {
	if metadata.Kind != DiscoveryKindEvent {
		return ValidationError{Field: "kind", Message: "must be 'event' for event metadata"}
	}

	if !isValidVenueID(metadata.VenueID) {
		return ValidationError{Field: "venue_id", Message: "invalid venue ID"}
	}

	if metadata.EventID == "" {
		return ValidationError{Field: "event_id", Message: "required field is empty"}
	}

	if metadata.Title == "" {
		return ValidationError{Field: "title", Message: "required field is empty"}
	}

	if metadata.Active == nil {
		return ValidationError{Field: "active", Message: "required field is missing"}
	}

	if metadata.Closed == nil {
		return ValidationError{Field: "closed", Message: "required field is missing"}
	}

	if metadata.DiscoveredAt.IsZero() {
		return ValidationError{Field: "discovered_at", Message: "required field is zero"}
	}

	if metadata.LastSeen.IsZero() {
		return ValidationError{Field: "last_seen", Message: "required field is zero"}
	}

	return nil
}

func validateSeriesMetadata(metadata SeriesMetadataV0) error {
	if metadata.Kind != DiscoveryKindSeries {
		return ValidationError{Field: "kind", Message: "must be 'series' for series metadata"}
	}

	if !isValidVenueID(metadata.VenueID) {
		return ValidationError{Field: "venue_id", Message: "invalid venue ID"}
	}

	if metadata.EventID == "" {
		return ValidationError{Field: "event_id", Message: "required field is empty"}
	}

	if metadata.Title == "" {
		return ValidationError{Field: "title", Message: "required field is empty"}
	}

	if metadata.Active == nil {
		return ValidationError{Field: "active", Message: "required field is missing"}
	}

	if metadata.Closed == nil {
		return ValidationError{Field: "closed", Message: "required field is missing"}
	}

	if metadata.DiscoveredAt.IsZero() {
		return ValidationError{Field: "discovered_at", Message: "required field is zero"}
	}

	if metadata.LastSeen.IsZero() {
		return ValidationError{Field: "last_seen", Message: "required field is zero"}
	}

	// Validate nested structures if present
	if metadata.SeriesData != nil {
		if err := validateSeriesData(*metadata.SeriesData); err != nil {
			return fmt.Errorf("series_data validation failed: %w", err)
		}
	}

	return nil
}

func validateDiscoveryMeta(meta DiscoveryMetaV0) error {
	if meta.BatchID == "" {
		return ValidationError{Field: "batch_id", Message: "required field is empty"}
	}

	if meta.BatchSequence < 1 {
		return ValidationError{Field: "batch_sequence", Message: "must be >= 1"}
	}

	if meta.BatchTotalCount < 1 {
		return ValidationError{Field: "batch_total_count", Message: "must be >= 1"}
	}

	if meta.DiscoveryRunID == "" {
		return ValidationError{Field: "discovery_run_id", Message: "required field is empty"}
	}

	return nil
}

func isValidEventType(eventType EventType) bool {
	switch eventType {
	case EventTypeDiscovered, EventTypeUpdated, EventTypeExpired:
		return true
	default:
		return false
	}
}

func isValidVenueID(venueID VenueID) bool {
	switch venueID {
	case VenueIDPolymarket, VenueIDKalshi:
		return true
	default:
		return false
	}
}

// ValidateFinancialData validates financial data constraints
func ValidateFinancialData(data FinancialDataV0) error {
	if data.Volume24hUSD != nil {
		if *data.Volume24hUSD < 0 {
			return ValidationError{Field: "volume_24h_usd", Message: "must be >= 0"}
		}
		if !isValidCentsPrecision(*data.Volume24hUSD) {
			return ValidationError{Field: "volume_24h_usd", Message: "must be a multiple of 0.01"}
		}
	}

	if data.VolumeTotalUSD != nil {
		if *data.VolumeTotalUSD < 0 {
			return ValidationError{Field: "volume_total_usd", Message: "must be >= 0"}
		}
		if !isValidCentsPrecision(*data.VolumeTotalUSD) {
			return ValidationError{Field: "volume_total_usd", Message: "must be a multiple of 0.01"}
		}
	}

	if data.LiquidityTotalUSD != nil {
		if *data.LiquidityTotalUSD < 0 {
			return ValidationError{Field: "liquidity_total_usd", Message: "must be >= 0"}
		}
		if !isValidCentsPrecision(*data.LiquidityTotalUSD) {
			return ValidationError{Field: "liquidity_total_usd", Message: "must be a multiple of 0.01"}
		}
	}

	if data.Volume24hContracts != nil && *data.Volume24hContracts < 0 {
		return ValidationError{Field: "volume_24h_contracts", Message: "must be >= 0"}
	}

	if data.VolumeTotalContracts != nil && *data.VolumeTotalContracts < 0 {
		return ValidationError{Field: "volume_total_contracts", Message: "must be >= 0"}
	}

	if data.Score != nil && *data.Score < 0 {
		return ValidationError{Field: "score", Message: "must be >= 0"}
	}

	if data.Currency != nil && *data.Currency != string(CurrencyUSD) {
		return ValidationError{Field: "currency", Message: "must be 'USD'"}
	}

	return nil
}

// ValidateSettlementSource validates settlement source data
func ValidateSettlementSource(source SettlementSourceV0) error {
	if source.Name == "" {
		return ValidationError{Field: "name", Message: "required field is empty"}
	}

	return nil
}

// validateSeriesData validates series-specific data structures
func validateSeriesData(data SeriesDataV0) error {
	// Validate financial data if present
	if data.Financial != nil {
		if err := ValidateFinancialData(*data.Financial); err != nil {
			return fmt.Errorf("financial data validation failed: %w", err)
		}
	}

	// Validate contract data if present
	if data.Contract != nil {
		if err := validateContractData(*data.Contract); err != nil {
			return fmt.Errorf("contract data validation failed: %w", err)
		}
	}

	return nil
}

// validateContractData validates contract data structures
func validateContractData(data ContractDataV0) error {
	// Validate settlement sources
	for i, source := range data.SettlementSources {
		if err := ValidateSettlementSource(source); err != nil {
			return fmt.Errorf("settlement_sources[%d] validation failed: %w", i, err)
		}
	}

	return nil
}

// isValidCentsPrecision checks if a value is a multiple of 0.01 (cents precision)
func isValidCentsPrecision(value float64) bool {
	// Round to cents and compare with original
	rounded := math.Round(value*100) / 100
	return math.Abs(value-rounded) < 1e-10 // Use small epsilon for floating point comparison
}