package discovery

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestValidateEventDiscoveryPayload_Valid(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{
			name:     "Polymarket event payload",
			filename: "event-payload-polymarket-valid.json",
		},
		{
			name:     "Kalshi event payload",
			filename: "event-payload-kalshi-valid.json",
		},
		{
			name:     "Minimal event payload",
			filename: "minimal-event-payload.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := loadTestData(t, tt.filename)
			if err := ValidateEventDiscoveryPayload(payload); err != nil {
				t.Errorf("ValidateEventDiscoveryPayload() error = %v, want nil", err)
			}
		})
	}
}

func TestValidateEventDiscoveryPayload_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "Missing required fields",
			filename: "event-payload-invalid-missing-required.json",
			wantErr:  true,
		},
		{
			name:     "Wrong enum values",
			filename: "event-payload-invalid-wrong-enum.json",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := loadTestData(t, tt.filename)
			err := ValidateEventDiscoveryPayload(payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEventDiscoveryPayload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSeriesDiscoveryPayload_Valid(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{
			name:     "Polymarket series payload",
			filename: "series-payload-polymarket-valid.json",
		},
		{
			name:     "Kalshi series payload",
			filename: "series-payload-kalshi-valid.json",
		},
		{
			name:     "Minimal series payload",
			filename: "minimal-series-payload.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := loadTestData(t, tt.filename)
			if err := ValidateSeriesDiscoveryPayload(payload); err != nil {
				t.Errorf("ValidateSeriesDiscoveryPayload() error = %v, want nil", err)
			}
		})
	}
}

func TestValidateSeriesDiscoveryPayload_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "Kind mismatch (event in series payload)",
			filename: "series-payload-invalid-kind-mismatch.json",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := loadTestData(t, tt.filename)
			err := ValidateSeriesDiscoveryPayload(payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSeriesDiscoveryPayload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEventMetadata(t *testing.T) {
	validMetadata := EventMetadataV0{
		Kind:         DiscoveryKindEvent,
		VenueID:      VenueIDPolymarket,
		EventID:      "test-event",
		Title:        "Test Event",
		Active:       boolPtr(true),
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
	}

	if err := ValidateEventMetadata(validMetadata); err != nil {
		t.Errorf("ValidateEventMetadata() error = %v, want nil", err)
	}

	// Test invalid kind
	invalidKind := validMetadata
	invalidKind.Kind = DiscoveryKindSeries
	if err := ValidateEventMetadata(invalidKind); err == nil {
		t.Error("ValidateEventMetadata() expected error for invalid kind, got nil")
	}

	// Test empty required fields
	emptyTitle := validMetadata
	emptyTitle.Title = ""
	if err := ValidateEventMetadata(emptyTitle); err == nil {
		t.Error("ValidateEventMetadata() expected error for empty title, got nil")
	}
}

func TestValidateSeriesMetadata(t *testing.T) {
	validMetadata := SeriesMetadataV0{
		Kind:         DiscoveryKindSeries,
		VenueID:      VenueIDKalshi,
		EventID:      "test-series",
		Title:        "Test Series",
		Active:       boolPtr(true),
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
	}

	if err := ValidateSeriesMetadata(validMetadata); err != nil {
		t.Errorf("ValidateSeriesMetadata() error = %v, want nil", err)
	}

	// Test invalid kind
	invalidKind := validMetadata
	invalidKind.Kind = DiscoveryKindEvent
	if err := ValidateSeriesMetadata(invalidKind); err == nil {
		t.Error("ValidateSeriesMetadata() expected error for invalid kind, got nil")
	}
}

func TestValidateDiscoveryMeta(t *testing.T) {
	validMeta := DiscoveryMetaV0{
		BatchID:          "batch_001",
		BatchSequence:    1,
		BatchTotalCount:  10,
		DiscoveryRunID:   "run_001",
	}

	if err := validateDiscoveryMeta(validMeta); err != nil {
		t.Errorf("validateDiscoveryMeta() error = %v, want nil", err)
	}

	// Test invalid batch sequence
	invalidSequence := validMeta
	invalidSequence.BatchSequence = 0
	if err := validateDiscoveryMeta(invalidSequence); err == nil {
		t.Error("validateDiscoveryMeta() expected error for invalid batch sequence, got nil")
	}

	// Test invalid batch total count
	invalidTotal := validMeta
	invalidTotal.BatchTotalCount = 0
	if err := validateDiscoveryMeta(invalidTotal); err == nil {
		t.Error("validateDiscoveryMeta() expected error for invalid batch total count, got nil")
	}
}

func TestValidateFinancialData(t *testing.T) {
	validData := FinancialDataV0{
		Volume24hUSD:         floatPtr(1000.50),
		VolumeTotalUSD:       floatPtr(10000.75),
		LiquidityTotalUSD:    floatPtr(5000.25),
		Volume24hContracts:   intPtr(100),
		VolumeTotalContracts: intPtr(1000),
		Score:                floatPtr(85.5),
		Currency:             stringPtr("USD"),
	}

	if err := ValidateFinancialData(validData); err != nil {
		t.Errorf("ValidateFinancialData() error = %v, want nil", err)
	}

	// Test negative values
	negativeVolume := validData
	negativeVolume.Volume24hUSD = floatPtr(-100.0)
	if err := ValidateFinancialData(negativeVolume); err == nil {
		t.Error("ValidateFinancialData() expected error for negative volume, got nil")
	}

	// Test invalid currency
	invalidCurrency := validData
	invalidCurrency.Currency = stringPtr("EUR")
	if err := ValidateFinancialData(invalidCurrency); err == nil {
		t.Error("ValidateFinancialData() expected error for invalid currency, got nil")
	}

	// Test multipleOf validation - values that are not multiples of 0.01
	invalidPrecision := validData
	invalidPrecision.Volume24hUSD = floatPtr(123.456) // 3 decimal places
	if err := ValidateFinancialData(invalidPrecision); err == nil {
		t.Error("ValidateFinancialData() expected error for invalid precision (123.456), got nil")
	}

	invalidPrecision2 := validData
	invalidPrecision2.VolumeTotalUSD = floatPtr(0.005) // Half cent
	if err := ValidateFinancialData(invalidPrecision2); err == nil {
		t.Error("ValidateFinancialData() expected error for invalid precision (0.005), got nil")
	}

	invalidPrecision3 := validData
	invalidPrecision3.LiquidityTotalUSD = floatPtr(1000.001) // Fraction of cent
	if err := ValidateFinancialData(invalidPrecision3); err == nil {
		t.Error("ValidateFinancialData() expected error for invalid precision (1000.001), got nil")
	}

	// Test valid precision values
	validPrecisionData := FinancialDataV0{
		Volume24hUSD:      floatPtr(0.01),  // Minimum valid
		VolumeTotalUSD:    floatPtr(0.00),  // Zero is valid
		LiquidityTotalUSD: floatPtr(99.99), // Two decimal places
	}
	if err := ValidateFinancialData(validPrecisionData); err != nil {
		t.Errorf("ValidateFinancialData() valid precision should pass: %v", err)
	}
}

func TestValidateSettlementSource(t *testing.T) {
	validSource := SettlementSourceV0{
		Name: "Test Source",
		URL:  stringPtr("https://example.com"),
	}

	if err := ValidateSettlementSource(validSource); err != nil {
		t.Errorf("ValidateSettlementSource() error = %v, want nil", err)
	}

	// Test empty name
	emptyName := validSource
	emptyName.Name = ""
	if err := ValidateSettlementSource(emptyName); err == nil {
		t.Error("ValidateSettlementSource() expected error for empty name, got nil")
	}
}

func TestIsValidEventType(t *testing.T) {
	validTypes := []EventType{
		EventTypeDiscovered,
		EventTypeUpdated,
		EventTypeExpired,
	}

	for _, eventType := range validTypes {
		if !isValidEventType(eventType) {
			t.Errorf("isValidEventType(%v) = false, want true", eventType)
		}
	}

	if isValidEventType("invalid") {
		t.Error("isValidEventType('invalid') = true, want false")
	}
}

func TestIsValidVenueID(t *testing.T) {
	validVenues := []VenueID{
		VenueIDPolymarket,
		VenueIDKalshi,
	}

	for _, venueID := range validVenues {
		if !isValidVenueID(venueID) {
			t.Errorf("isValidVenueID(%v) = false, want true", venueID)
		}
	}

	if isValidVenueID("invalid") {
		t.Error("isValidVenueID('invalid') = true, want false")
	}
}

// Helper functions
func loadTestData(t *testing.T, filename string) []byte {
	// Go up from codegen/go/discovery to find schemas/examples/discovery
	examplesPath := filepath.Join("..", "..", "..", "schemas", "examples", "discovery", filename)
	data, err := os.ReadFile(examplesPath)
	if err != nil {
		t.Fatalf("Failed to read test data file %s: %v", filename, err)
	}
	return data
}

func floatPtr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

// Benchmark tests
func BenchmarkValidateEventDiscoveryPayload(b *testing.B) {
	payload := createValidEventPayload()
	data, _ := json.Marshal(payload)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateEventDiscoveryPayload(data)
	}
}

func BenchmarkValidateSeriesDiscoveryPayload(b *testing.B) {
	payload := createValidSeriesPayload()
	data, _ := json.Marshal(payload)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateSeriesDiscoveryPayload(data)
	}
}

func createValidEventPayload() EventDiscoveryPayloadV0 {
	return EventDiscoveryPayloadV0{
		Event: EventMetadataV0{
			Kind:         DiscoveryKindEvent,
			VenueID:      VenueIDPolymarket,
			EventID:      "test-event",
			Title:        "Test Event",
			Active:       boolPtr(true),
			Closed:       boolPtr(false),
			DiscoveredAt: time.Now(),
			LastSeen:     time.Now(),
		},
		EventID:   "evt_test",
		EventType: EventTypeDiscovered,
		Timestamp: time.Now(),
		VenueID:   VenueIDPolymarket,
	}
}

func createValidSeriesPayload() SeriesDiscoveryPayloadV0 {
	return SeriesDiscoveryPayloadV0{
		Event: SeriesMetadataV0{
			Kind:         DiscoveryKindSeries,
			VenueID:      VenueIDKalshi,
			EventID:      "test-series",
			Title:        "Test Series",
			Active:       boolPtr(true),
			Closed:       boolPtr(false),
			DiscoveredAt: time.Now(),
			LastSeen:     time.Now(),
		},
		EventID:   "evt_series_test",
		EventType: EventTypeUpdated,
		Timestamp: time.Now(),
		VenueID:   VenueIDKalshi,
	}
}

func TestValidateSeriesDiscoveryPayload_NestedValidation(t *testing.T) {
	// Test that nested structures are properly validated

	basePayload := SeriesDiscoveryPayloadV0{
		Event: SeriesMetadataV0{
			Kind:         DiscoveryKindSeries,
			VenueID:      VenueIDKalshi,
			EventID:      "test-series",
			Title:        "Test Series",
			Active:       boolPtr(true),
			Closed:       boolPtr(false),
			DiscoveredAt: time.Now(),
			LastSeen:     time.Now(),
		},
		EventID:   "evt_series_test",
		EventType: EventTypeDiscovered,
		Timestamp: time.Now(),
		VenueID:   VenueIDKalshi,
	}

	// Test invalid settlement source (missing name)
	invalidSettlementPayload := basePayload
	invalidSettlementPayload.Event.SeriesData = &SeriesDataV0{
		Contract: &ContractDataV0{
			SettlementSources: []SettlementSourceV0{
				{Name: ""}, // Invalid: empty name
			},
		},
	}

	data, _ := json.Marshal(invalidSettlementPayload)
	if err := ValidateSeriesDiscoveryPayload(data); err == nil {
		t.Error("Expected validation error for empty settlement source name, got nil")
	}

	// Test invalid financial precision
	invalidFinancialPayload := basePayload
	invalidFinancialPayload.Event.SeriesData = &SeriesDataV0{
		Financial: &FinancialDataV0{
			Volume24hUSD: floatPtr(123.456), // Invalid: not multiple of 0.01
		},
	}

	data, _ = json.Marshal(invalidFinancialPayload)
	if err := ValidateSeriesDiscoveryPayload(data); err == nil {
		t.Error("Expected validation error for invalid financial precision, got nil")
	}

	// Test valid nested structures
	validPayload := basePayload
	validPayload.Event.SeriesData = &SeriesDataV0{
		Financial: &FinancialDataV0{
			Volume24hUSD: floatPtr(123.45), // Valid: multiple of 0.01
			Currency:     stringPtr("USD"),
		},
		Contract: &ContractDataV0{
			SettlementSources: []SettlementSourceV0{
				{Name: "Valid Source", URL: stringPtr("https://example.com")},
			},
		},
	}

	data, _ = json.Marshal(validPayload)
	if err := ValidateSeriesDiscoveryPayload(data); err != nil {
		t.Errorf("Expected valid nested structures to pass validation, got: %v", err)
	}
}

func TestBooleanFieldPresenceValidation(t *testing.T) {
	// Test that missing boolean fields are properly detected

	// Event metadata with missing Active field
	eventMissingActive := EventMetadataV0{
		Kind:         DiscoveryKindEvent,
		VenueID:      VenueIDPolymarket,
		EventID:      "test-event",
		Title:        "Test Event",
		// Active:       nil,  // Missing required field
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
	}

	if err := ValidateEventMetadata(eventMissingActive); err == nil {
		t.Error("Expected validation error for missing Active field, got nil")
	}

	// Event metadata with missing Closed field
	eventMissingClosed := EventMetadataV0{
		Kind:         DiscoveryKindEvent,
		VenueID:      VenueIDPolymarket,
		EventID:      "test-event",
		Title:        "Test Event",
		Active:       boolPtr(true),
		// Closed:       nil,  // Missing required field
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
	}

	if err := ValidateEventMetadata(eventMissingClosed); err == nil {
		t.Error("Expected validation error for missing Closed field, got nil")
	}

	// Series metadata with missing Active field
	seriesMissingActive := SeriesMetadataV0{
		Kind:         DiscoveryKindSeries,
		VenueID:      VenueIDKalshi,
		EventID:      "test-series",
		Title:        "Test Series",
		// Active:       nil,  // Missing required field
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
	}

	if err := ValidateSeriesMetadata(seriesMissingActive); err == nil {
		t.Error("Expected validation error for missing Active field in series, got nil")
	}

	// Series metadata with missing Closed field
	seriesMissingClosed := SeriesMetadataV0{
		Kind:         DiscoveryKindSeries,
		VenueID:      VenueIDKalshi,
		EventID:      "test-series",
		Title:        "Test Series",
		Active:       boolPtr(true),
		// Closed:       nil,  // Missing required field
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
	}

	if err := ValidateSeriesMetadata(seriesMissingClosed); err == nil {
		t.Error("Expected validation error for missing Closed field in series, got nil")
	}

	// Valid metadata with both fields present
	validEvent := EventMetadataV0{
		Kind:         DiscoveryKindEvent,
		VenueID:      VenueIDPolymarket,
		EventID:      "test-event",
		Title:        "Test Event",
		Active:       boolPtr(true),
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
	}

	if err := ValidateEventMetadata(validEvent); err != nil {
		t.Errorf("Valid event metadata should pass validation, got: %v", err)
	}
}

func TestIsValidCentsPrecision(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected bool
	}{
		{"Valid whole number", 100.00, true},
		{"Valid two decimals", 123.45, true},
		{"Valid zero", 0.00, true},
		{"Valid one cent", 0.01, true},
		{"Valid max precision", 99999.99, true},
		{"Invalid three decimals", 123.456, false},
		{"Invalid half cent", 0.005, false},
		{"Invalid tiny fraction", 100.001, false},
		{"Invalid many decimals", 1.23456789, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidCentsPrecision(tt.value); got != tt.expected {
				t.Errorf("isValidCentsPrecision(%v) = %v, want %v", tt.value, got, tt.expected)
			}
		})
	}
}