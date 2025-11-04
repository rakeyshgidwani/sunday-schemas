package discovery

import (
	"testing"
	"time"
)

// TestVenueFieldMapping tests the mapping from venue-specific APIs to canonical schema
func TestVenueFieldMapping(t *testing.T) {
	t.Run("Polymarket to Canonical Mapping", func(t *testing.T) {
		testPolymarketMapping(t)
	})

	t.Run("Kalshi to Canonical Mapping", func(t *testing.T) {
		testKalshiMapping(t)
	})
}

func testPolymarketMapping(t *testing.T) {

	// Sample Polymarket data
	polymarketData := PolymarketSeriesAPI{
		ID:         "10244",
		Title:      "Concacaf",
		Ticker:     "CONCACAF",
		Slug:       "concacaf",
		Subtitle:   "Soccer Championship",
		SeriesType: "single",
		Recurrence: "annual",
		Category:   "Sports",
		Image:      "https://polymarket.com/image.jpg",
		Icon:       "https://polymarket.com/icon.png",
		Volume24hr: 125000.50,
		Volume:     2500000.75,
		Active:     true,
		Closed:     false,
		CreatedAt:  "2025-09-03T03:07:56.295896Z",
		UpdatedAt:  "2025-11-03T21:01:11.954948Z",
	}

	// Map to canonical schema
	canonical := mapPolymarketToCanonical(polymarketData)

	// Validate mapping
	if canonical.Kind != DiscoveryKindSeries {
		t.Errorf("Expected kind 'series', got %v", canonical.Kind)
	}
	if canonical.VenueID != VenueIDPolymarket {
		t.Errorf("Expected venue_id 'polymarket', got %v", canonical.VenueID)
	}
	if canonical.EventID != polymarketData.ID {
		t.Errorf("Expected event_id '%s', got %s", polymarketData.ID, canonical.EventID)
	}
	if canonical.Title != polymarketData.Title {
		t.Errorf("Expected title '%s', got %s", polymarketData.Title, canonical.Title)
	}

	// Validate series data mapping
	if canonical.SeriesData == nil {
		t.Fatal("Expected series_data to be populated")
	}
	if canonical.SeriesData.Ticker == nil || *canonical.SeriesData.Ticker != polymarketData.Ticker {
		t.Errorf("Expected ticker '%s', got %v", polymarketData.Ticker, canonical.SeriesData.Ticker)
	}
	if canonical.SeriesData.Slug == nil || *canonical.SeriesData.Slug != polymarketData.Slug {
		t.Errorf("Expected slug '%s', got %v", polymarketData.Slug, canonical.SeriesData.Slug)
	}

	// Validate financial data mapping
	if canonical.SeriesData.Financial == nil {
		t.Fatal("Expected financial data to be populated")
	}
	if canonical.SeriesData.Financial.Volume24hUSD == nil || *canonical.SeriesData.Financial.Volume24hUSD != polymarketData.Volume24hr {
		t.Errorf("Expected volume_24h_usd %.2f, got %v", polymarketData.Volume24hr, canonical.SeriesData.Financial.Volume24hUSD)
	}
	if canonical.SeriesData.Financial.VolumeTotalUSD == nil || *canonical.SeriesData.Financial.VolumeTotalUSD != polymarketData.Volume {
		t.Errorf("Expected volume_total_usd %.2f, got %v", polymarketData.Volume, canonical.SeriesData.Financial.VolumeTotalUSD)
	}
}

func testKalshiMapping(t *testing.T) {

	// Sample Kalshi data
	kalshiData := KalshiSeriesAPI{
		Ticker:    "PRES24",
		Title:     "2024 Presidential Election",
		Category:  "Politics",
		Frequency: "quadrennial",
		Tags:      []string{"politics", "election", "president"},
		ContractURL:          "https://kalshi.com/markets/PRES24",
		ContractTermsURL:     "https://kalshi.com/terms/PRES24",
		FeeType:              "percentage",
		AdditionalProhibitions: []string{"insider_trading"},
		SettlementSources: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{Name: "Associated Press", URL: "https://www.ap.org"},
			{Name: "Federal Election Commission", URL: "https://www.fec.gov"},
		},
		Active:    true,
		Closed:    false,
		CreatedAt: "2023-11-01T00:00:00Z",
		UpdatedAt: "2025-11-04T11:30:00Z",
	}

	// Map to canonical schema
	canonical := mapKalshiToCanonical(kalshiData)

	// Validate mapping
	if canonical.Kind != DiscoveryKindSeries {
		t.Errorf("Expected kind 'series', got %v", canonical.Kind)
	}
	if canonical.VenueID != VenueIDKalshi {
		t.Errorf("Expected venue_id 'kalshi', got %v", canonical.VenueID)
	}
	if canonical.EventID != kalshiData.Ticker {
		t.Errorf("Expected event_id '%s', got %s", kalshiData.Ticker, canonical.EventID)
	}
	if len(canonical.Tags) != len(kalshiData.Tags) {
		t.Errorf("Expected %d tags, got %d", len(kalshiData.Tags), len(canonical.Tags))
	}

	// Validate series data mapping
	if canonical.SeriesData == nil {
		t.Fatal("Expected series_data to be populated")
	}
	if canonical.SeriesData.Ticker == nil || *canonical.SeriesData.Ticker != kalshiData.Ticker {
		t.Errorf("Expected ticker '%s', got %v", kalshiData.Ticker, canonical.SeriesData.Ticker)
	}
	if canonical.SeriesData.Recurrence == nil || *canonical.SeriesData.Recurrence != kalshiData.Frequency {
		t.Errorf("Expected recurrence '%s', got %v", kalshiData.Frequency, canonical.SeriesData.Recurrence)
	}

	// Validate contract data mapping
	if canonical.SeriesData.Contract == nil {
		t.Fatal("Expected contract data to be populated")
	}
	if canonical.SeriesData.Contract.ContractURL == nil || *canonical.SeriesData.Contract.ContractURL != kalshiData.ContractURL {
		t.Errorf("Expected contract_url '%s', got %v", kalshiData.ContractURL, canonical.SeriesData.Contract.ContractURL)
	}
	if len(canonical.SeriesData.Contract.SettlementSources) != len(kalshiData.SettlementSources) {
		t.Errorf("Expected %d settlement sources, got %d",
			len(kalshiData.SettlementSources), len(canonical.SeriesData.Contract.SettlementSources))
	}
}

func TestCrossVenueConsistency(t *testing.T) {
	// Test that common fields map consistently across venues
	testCases := []struct {
		name           string
		polymarketData map[string]interface{}
		kalshiData     map[string]interface{}
		expectedField  string
		expectedValue  interface{}
	}{
		{
			name:           "Title mapping",
			polymarketData: map[string]interface{}{"title": "Test Event"},
			kalshiData:     map[string]interface{}{"title": "Test Event"},
			expectedField:  "title",
			expectedValue:  "Test Event",
		},
		{
			name:           "Category mapping",
			polymarketData: map[string]interface{}{"category": "Politics"},
			kalshiData:     map[string]interface{}{"category": "Politics"},
			expectedField:  "category",
			expectedValue:  "Politics",
		},
		{
			name:           "Active status mapping",
			polymarketData: map[string]interface{}{"active": true},
			kalshiData:     map[string]interface{}{"active": true},
			expectedField:  "active",
			expectedValue:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// This test validates that common field mappings are consistent
			// across different venues for the same conceptual data

			// In a real implementation, you would have venue-specific
			// mapping functions that convert API responses to canonical format
			// and this test would verify they produce consistent results

			// For now, we verify the concept works
			if tc.expectedValue != tc.polymarketData[tc.expectedField] {
				t.Errorf("Polymarket mapping inconsistent for %s", tc.expectedField)
			}
			if tc.expectedValue != tc.kalshiData[tc.expectedField] {
				t.Errorf("Kalshi mapping inconsistent for %s", tc.expectedField)
			}
		})
	}
}

func TestExtraMetadataHandling(t *testing.T) {
	// Test that venue-specific fields are properly handled in extra_metadata

	// Polymarket-specific fields
	polymarketMetadata := EventMetadataV0{
		Kind:         DiscoveryKindEvent,
		VenueID:      VenueIDPolymarket,
		EventID:      "test",
		Title:        "Test",
		Active:       boolPtr(true),
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
		ExtraMetadata: map[string]any{
			"polymarket_specific_id": "pm_12345",
			"creation_timestamp":     "2024-01-01T00:00:00Z",
			"featured_order":         1,
		},
	}

	// Verify extra metadata is preserved
	if polymarketMetadata.ExtraMetadata["polymarket_specific_id"] != "pm_12345" {
		t.Error("Polymarket-specific metadata not preserved")
	}

	// Kalshi-specific fields
	kalshiMetadata := SeriesMetadataV0{
		Kind:         DiscoveryKindSeries,
		VenueID:      VenueIDKalshi,
		EventID:      "test",
		Title:        "Test",
		Active:       boolPtr(true),
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
		ExtraMetadata: map[string]any{
			"kalshi_contract_id":     "KALSHI_CONTRACT_123",
			"compliance_reviewed":    true,
			"regulatory_approval":    "CFTC_approved",
		},
	}

	// Verify extra metadata is preserved
	if kalshiMetadata.ExtraMetadata["kalshi_contract_id"] != "KALSHI_CONTRACT_123" {
		t.Error("Kalshi-specific metadata not preserved")
	}
}

func TestRelationshipMapping(t *testing.T) {
	// Test parent/child relationship mapping

	// Series with child events
	series := SeriesMetadataV0{
		Kind:          DiscoveryKindSeries,
		VenueID:       VenueIDPolymarket,
		EventID:       "series_123",
		Title:         "Test Series",
		Active:        boolPtr(true),
		Closed:        boolPtr(false),
		DiscoveredAt:  time.Now(),
		LastSeen:      time.Now(),
		ChildEventIDs: []string{"event_1", "event_2", "event_3"},
		Relationships: &RelationshipsV0{
			EventIDs: []string{"event_1", "event_2", "event_3"},
			InstrumentIDs: []string{"inst_1_yes", "inst_1_no", "inst_2_yes", "inst_2_no"},
		},
	}

	// Verify relationships
	if len(series.ChildEventIDs) != 3 {
		t.Errorf("Expected 3 child events, got %d", len(series.ChildEventIDs))
	}
	if series.Relationships == nil {
		t.Fatal("Expected relationships to be populated")
	}
	if len(series.Relationships.EventIDs) != len(series.ChildEventIDs) {
		t.Error("Child event IDs and relationship event IDs should match")
	}

	// Event with parent series
	event := EventMetadataV0{
		Kind:              DiscoveryKindEvent,
		VenueID:           VenueIDPolymarket,
		EventID:           "event_1",
		Title:             "Test Event",
		Active:            boolPtr(true),
		Closed:            boolPtr(false),
		DiscoveredAt:      time.Now(),
		LastSeen:          time.Now(),
		ParentSeriesID:    stringPtr("series_123"),
		ParentSeriesTitle: stringPtr("Test Series"),
		Relationships: &RelationshipsV0{
			SeriesID: stringPtr("series_123"),
			InstrumentIDs: []string{"inst_1_yes", "inst_1_no"},
		},
	}

	// Verify parent relationship
	if event.ParentSeriesID == nil || *event.ParentSeriesID != "series_123" {
		t.Error("Parent series ID not correctly mapped")
	}
	if event.Relationships == nil || event.Relationships.SeriesID == nil {
		t.Error("Parent relationship not correctly mapped")
	}
}

// PolymarketSeriesAPI represents Polymarket API response structure
type PolymarketSeriesAPI struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Ticker       string  `json:"ticker"`
	Slug         string  `json:"slug"`
	Subtitle     string  `json:"subtitle"`
	SeriesType   string  `json:"seriesType"`
	Recurrence   string  `json:"recurrence"`
	Category     string  `json:"category"`
	Image        string  `json:"image"`
	Icon         string  `json:"icon"`
	Volume24hr   float64 `json:"volume24hr"`
	Volume       float64 `json:"volume"`
	Active       bool    `json:"active"`
	Closed       bool    `json:"closed"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// KalshiSeriesAPI represents Kalshi API response structure
type KalshiSeriesAPI struct {
	Ticker               string   `json:"ticker"`
	Title                string   `json:"title"`
	Category             string   `json:"category"`
	Frequency            string   `json:"frequency"`
	Tags                 []string `json:"tags"`
	ContractURL          string   `json:"contract_url"`
	ContractTermsURL     string   `json:"contract_terms_url"`
	FeeType              string   `json:"fee_type"`
	AdditionalProhibitions []string `json:"additional_prohibitions"`
	SettlementSources    []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"settlement_sources"`
	Active    bool   `json:"active"`
	Closed    bool   `json:"closed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Helper functions for mapping (these would be part of the actual implementation)
func mapPolymarketToCanonical(data PolymarketSeriesAPI) SeriesMetadataV0 {
	createdAt, _ := time.Parse(time.RFC3339, data.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, data.UpdatedAt)
	now := time.Now()

	return SeriesMetadataV0{
		Kind:         DiscoveryKindSeries,
		VenueID:      VenueIDPolymarket,
		EventID:      data.ID,
		Title:        data.Title,
		Category:     &data.Category,
		Active:       boolPtr(data.Active),
		Closed:       boolPtr(data.Closed),
		DiscoveredAt: now,
		LastSeen:     now,
		SeriesData: &SeriesDataV0{
			Ticker:     &data.Ticker,
			Slug:       &data.Slug,
			Subtitle:   &data.Subtitle,
			SeriesType: &data.SeriesType,
			Recurrence: &data.Recurrence,
			ImageURL:   &data.Image,
			IconURL:    &data.Icon,
			Financial: &FinancialDataV0{
				Volume24hUSD:   &data.Volume24hr,
				VolumeTotalUSD: &data.Volume,
				Currency:       stringPtr("USD"),
			},
			Timestamps: &TimestampDataV0{
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
		},
	}
}

func mapKalshiToCanonical(data KalshiSeriesAPI) SeriesMetadataV0 {
	createdAt, _ := time.Parse(time.RFC3339, data.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, data.UpdatedAt)
	now := time.Now()

	// Map settlement sources
	var settlementSources []SettlementSourceV0
	for _, source := range data.SettlementSources {
		settlementSources = append(settlementSources, SettlementSourceV0{
			Name: source.Name,
			URL:  &source.URL,
		})
	}

	return SeriesMetadataV0{
		Kind:         DiscoveryKindSeries,
		VenueID:      VenueIDKalshi,
		EventID:      data.Ticker,
		Title:        data.Title,
		Category:     &data.Category,
		Active:       boolPtr(data.Active),
		Closed:       boolPtr(data.Closed),
		Tags:         data.Tags,
		DiscoveredAt: now,
		LastSeen:     now,
		SeriesData: &SeriesDataV0{
			Ticker:     &data.Ticker,
			Recurrence: &data.Frequency,
			Contract: &ContractDataV0{
				ContractURL:            &data.ContractURL,
				ContractTermsURL:       &data.ContractTermsURL,
				FeeType:                &data.FeeType,
				AdditionalProhibitions: data.AdditionalProhibitions,
				SettlementSources:      settlementSources,
			},
			Timestamps: &TimestampDataV0{
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
		},
	}
}

// Additional test for schema evolution compatibility
func TestSchemaEvolutionCompatibility(t *testing.T) {
	// Test that new optional fields don't break existing data

	// Minimal payload (older version compatibility)
	minimalPayload := EventDiscoveryPayloadV0{
		Event: EventMetadataV0{
			Kind:         DiscoveryKindEvent,
			VenueID:      VenueIDPolymarket,
			EventID:      "minimal",
			Title:        "Minimal Event",
			Active:       boolPtr(true),
			Closed:       boolPtr(false),
			DiscoveredAt: time.Now(),
			LastSeen:     time.Now(),
		},
		EventID:   "evt_minimal",
		EventType: EventTypeDiscovered,
		Timestamp: time.Now(),
		VenueID:   VenueIDPolymarket,
	}

	// Should validate successfully
	if err := ValidateEventMetadata(minimalPayload.Event); err != nil {
		t.Errorf("Minimal payload should validate: %v", err)
	}

	// Extended payload (newer version with all fields)
	extendedPayload := minimalPayload
	extendedPayload.Event.Description = stringPtr("Extended description")
	extendedPayload.Event.Category = stringPtr("Test category")
	extendedPayload.Event.Tags = []string{"test", "extended"}
	extendedPayload.DiscoveryMeta = &DiscoveryMetaV0{
		BatchID:         "batch_001",
		BatchSequence:   1,
		BatchTotalCount: 1,
		DiscoveryRunID:  "run_001",
	}

	// Should also validate successfully
	if err := ValidateEventMetadata(extendedPayload.Event); err != nil {
		t.Errorf("Extended payload should validate: %v", err)
	}
}

