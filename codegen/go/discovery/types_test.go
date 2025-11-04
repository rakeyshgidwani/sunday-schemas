package discovery

import (
	"encoding/json"
	"testing"
	"time"
)

func TestEventDiscoveryPayloadV0_JSONMarshalUnmarshal(t *testing.T) {
	original := EventDiscoveryPayloadV0{
		Event: EventMetadataV0{
			Kind:         DiscoveryKindEvent,
			VenueID:      VenueIDPolymarket,
			EventID:      "test-event-123",
			Title:        "Test Event Title",
			Description:  stringPtr("Test event description"),
			Category:     stringPtr("Politics"),
			Active:       boolPtr(true),
			Closed:       boolPtr(false),
			StartDate:    timePtr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			EndDate:      timePtr(time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)),
			DiscoveredAt: time.Date(2025, 11, 4, 10, 0, 0, 0, time.UTC),
			LastSeen:     time.Date(2025, 11, 4, 12, 0, 0, 0, time.UTC),
			Tags:         []string{"politics", "test"},
		},
		EventID:   "evt_test_123",
		EventType: EventTypeDiscovered,
		Timestamp: time.Date(2025, 11, 4, 12, 0, 0, 0, time.UTC),
		VenueID:   VenueIDPolymarket,
		DiscoveryMeta: &DiscoveryMetaV0{
			BatchID:          "batch_001",
			BatchSequence:    1,
			BatchTotalCount:  10,
			DiscoveryRunID:   "run_001",
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal EventDiscoveryPayloadV0: %v", err)
	}

	// Unmarshal back to struct
	var unmarshaled EventDiscoveryPayloadV0
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal EventDiscoveryPayloadV0: %v", err)
	}

	// Compare critical fields
	if unmarshaled.Event.Kind != original.Event.Kind {
		t.Errorf("Kind mismatch: got %v, want %v", unmarshaled.Event.Kind, original.Event.Kind)
	}
	if unmarshaled.Event.VenueID != original.Event.VenueID {
		t.Errorf("VenueID mismatch: got %v, want %v", unmarshaled.Event.VenueID, original.Event.VenueID)
	}
	if unmarshaled.Event.EventID != original.Event.EventID {
		t.Errorf("EventID mismatch: got %v, want %v", unmarshaled.Event.EventID, original.Event.EventID)
	}
	if unmarshaled.EventType != original.EventType {
		t.Errorf("EventType mismatch: got %v, want %v", unmarshaled.EventType, original.EventType)
	}
}

func TestSeriesDiscoveryPayloadV0_JSONMarshalUnmarshal(t *testing.T) {
	original := SeriesDiscoveryPayloadV0{
		Event: SeriesMetadataV0{
			Kind:          DiscoveryKindSeries,
			VenueID:       VenueIDKalshi,
			EventID:       "test-series-456",
			Title:         "Test Series Title",
			Description:   stringPtr("Test series description"),
			Category:      stringPtr("Sports"),
			Active:        boolPtr(true),
			Closed:        boolPtr(false),
			Tags:          []string{"sports", "test", "series"},
			ChildEventIDs: []string{"event_1", "event_2", "event_3"},
			DiscoveredAt:  time.Date(2025, 11, 4, 9, 0, 0, 0, time.UTC),
			LastSeen:      time.Date(2025, 11, 4, 12, 0, 0, 0, time.UTC),
			SeriesData: &SeriesDataV0{
				Ticker:     stringPtr("TEST-SERIES"),
				Slug:       stringPtr("test-series-slug"),
				Subtitle:   stringPtr("Test Series Subtitle"),
				SeriesType: stringPtr("competitive"),
				Financial: &FinancialDataV0{
					Volume24hUSD:      floatPtr(1000.50),
					VolumeTotalUSD:    floatPtr(50000.75),
					Currency:          stringPtr("USD"),
				},
			},
		},
		EventID:   "evt_series_456",
		EventType: EventTypeUpdated,
		Timestamp: time.Date(2025, 11, 4, 12, 0, 0, 0, time.UTC),
		VenueID:   VenueIDKalshi,
	}

	// Marshal to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal SeriesDiscoveryPayloadV0: %v", err)
	}

	// Unmarshal back to struct
	var unmarshaled SeriesDiscoveryPayloadV0
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal SeriesDiscoveryPayloadV0: %v", err)
	}

	// Compare critical fields
	if unmarshaled.Event.Kind != original.Event.Kind {
		t.Errorf("Kind mismatch: got %v, want %v", unmarshaled.Event.Kind, original.Event.Kind)
	}
	if unmarshaled.Event.VenueID != original.Event.VenueID {
		t.Errorf("VenueID mismatch: got %v, want %v", unmarshaled.Event.VenueID, original.Event.VenueID)
	}
	if len(unmarshaled.Event.ChildEventIDs) != len(original.Event.ChildEventIDs) {
		t.Errorf("ChildEventIDs length mismatch: got %d, want %d",
			len(unmarshaled.Event.ChildEventIDs), len(original.Event.ChildEventIDs))
	}
}

func TestEnumValues_Consistency(t *testing.T) {
	// Test DiscoveryKind enum
	eventKind := DiscoveryKindEvent
	seriesKind := DiscoveryKindSeries

	if string(eventKind) != "event" {
		t.Errorf("DiscoveryKindEvent value: got %v, want 'event'", eventKind)
	}
	if string(seriesKind) != "series" {
		t.Errorf("DiscoveryKindSeries value: got %v, want 'series'", seriesKind)
	}

	// Test EventType enum
	discoveredType := EventTypeDiscovered
	updatedType := EventTypeUpdated
	expiredType := EventTypeExpired

	if string(discoveredType) != "discovered" {
		t.Errorf("EventTypeDiscovered value: got %v, want 'discovered'", discoveredType)
	}
	if string(updatedType) != "updated" {
		t.Errorf("EventTypeUpdated value: got %v, want 'updated'", updatedType)
	}
	if string(expiredType) != "expired" {
		t.Errorf("EventTypeExpired value: got %v, want 'expired'", expiredType)
	}

	// Test VenueID enum
	polymarket := VenueIDPolymarket
	kalshi := VenueIDKalshi

	if string(polymarket) != "polymarket" {
		t.Errorf("VenueIDPolymarket value: got %v, want 'polymarket'", polymarket)
	}
	if string(kalshi) != "kalshi" {
		t.Errorf("VenueIDKalshi value: got %v, want 'kalshi'", kalshi)
	}
}

func TestOptionalFields_NilHandling(t *testing.T) {
	// Test that optional fields can be nil without causing issues
	metadata := EventMetadataV0{
		Kind:         DiscoveryKindEvent,
		VenueID:      VenueIDPolymarket,
		EventID:      "test",
		Title:        "Test",
		Active:       boolPtr(true),
		Closed:       boolPtr(false),
		DiscoveredAt: time.Now(),
		LastSeen:     time.Now(),
		// All optional fields left as nil/empty
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("Failed to marshal metadata with nil optional fields: %v", err)
	}

	var unmarshaled EventMetadataV0
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal metadata with nil optional fields: %v", err)
	}

	// Verify nil fields are handled correctly
	if unmarshaled.Description != nil {
		t.Error("Expected Description to be nil")
	}
	if unmarshaled.Relationships != nil {
		t.Error("Expected Relationships to be nil")
	}
}

func TestJSONTags_Correctness(t *testing.T) {
	// Test that JSON tags match schema field names
	payload := EventDiscoveryPayloadV0{
		Event: EventMetadataV0{
			Kind:         DiscoveryKindEvent,
			VenueID:      VenueIDPolymarket,
			EventID:      "test",
			Title:        "Test",
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

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	// Unmarshal to map to check field names
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		t.Fatalf("Failed to unmarshal to map: %v", err)
	}

	// Check that expected fields exist with correct names
	expectedFields := []string{"event", "event_id", "event_type", "timestamp", "venue_id"}
	for _, field := range expectedFields {
		if _, exists := jsonMap[field]; !exists {
			t.Errorf("Expected field '%s' not found in JSON", field)
		}
	}

	// Check event metadata fields
	eventMap, ok := jsonMap["event"].(map[string]interface{})
	if !ok {
		t.Fatal("Event field is not a map")
	}

	expectedEventFields := []string{"kind", "venue_id", "event_id", "title", "active", "closed", "discovered_at", "last_seen"}
	for _, field := range expectedEventFields {
		if _, exists := eventMap[field]; !exists {
			t.Errorf("Expected event field '%s' not found in JSON", field)
		}
	}
}

func TestFinancialData_NumericConstraints(t *testing.T) {
	// Test that financial data preserves precision
	financial := FinancialDataV0{
		Volume24hUSD:      floatPtr(1234.56),
		VolumeTotalUSD:    floatPtr(9876543.21),
		LiquidityTotalUSD: floatPtr(555555.55),
		Currency:          stringPtr("USD"),
	}

	data, err := json.Marshal(financial)
	if err != nil {
		t.Fatalf("Failed to marshal financial data: %v", err)
	}

	var unmarshaled FinancialDataV0
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal financial data: %v", err)
	}

	// Check precision preservation (within floating point tolerance)
	tolerance := 0.01
	if abs(*unmarshaled.Volume24hUSD - *financial.Volume24hUSD) > tolerance {
		t.Errorf("Volume24hUSD precision lost: got %v, want %v",
			*unmarshaled.Volume24hUSD, *financial.Volume24hUSD)
	}
}

func TestTimeFieldHandling(t *testing.T) {
	// Test RFC3339 time format handling
	testTime := time.Date(2025, 11, 4, 12, 30, 45, 0, time.UTC)

	metadata := EventMetadataV0{
		Kind:         DiscoveryKindEvent,
		VenueID:      VenueIDPolymarket,
		EventID:      "test",
		Title:        "Test",
		Active:       boolPtr(true),
		Closed:       boolPtr(false),
		StartDate:    &testTime,
		EndDate:      &testTime,
		DiscoveredAt: testTime,
		LastSeen:     testTime,
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("Failed to marshal metadata with times: %v", err)
	}

	var unmarshaled EventMetadataV0
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal metadata with times: %v", err)
	}

	// Check time preservation
	if !unmarshaled.DiscoveredAt.Equal(testTime) {
		t.Errorf("DiscoveredAt time mismatch: got %v, want %v",
			unmarshaled.DiscoveredAt, testTime)
	}
	if !unmarshaled.StartDate.Equal(testTime) {
		t.Errorf("StartDate time mismatch: got %v, want %v",
			*unmarshaled.StartDate, testTime)
	}
}

// Helper functions
func timePtr(t time.Time) *time.Time {
	return &t
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}


// Property-based testing helpers
func TestRoundTripPropertyEvent(t *testing.T) {
	// Property: Any valid EventDiscoveryPayloadV0 should survive JSON round-trip
	for i := 0; i < 100; i++ {
		original := generateRandomEventPayload(i)

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("Failed to marshal random payload %d: %v", i, err)
		}

		var unmarshaled EventDiscoveryPayloadV0
		if err := json.Unmarshal(data, &unmarshaled); err != nil {
			t.Fatalf("Failed to unmarshal random payload %d: %v", i, err)
		}

		// Validate unmarshaled data
		marshaledAgain, err := json.Marshal(unmarshaled)
		if err != nil {
			t.Fatalf("Failed to marshal unmarshaled payload %d: %v", i, err)
		}

		if err := ValidateEventDiscoveryPayload(marshaledAgain); err != nil {
			t.Errorf("Round-trip payload %d failed validation: %v", i, err)
		}
	}
}

func generateRandomEventPayload(seed int) EventDiscoveryPayloadV0 {
	// Simple deterministic "random" generation for testing
	venues := []VenueID{VenueIDPolymarket, VenueIDKalshi}
	eventTypes := []EventType{EventTypeDiscovered, EventTypeUpdated, EventTypeExpired}

	return EventDiscoveryPayloadV0{
		Event: EventMetadataV0{
			Kind:         DiscoveryKindEvent,
			VenueID:      venues[seed%len(venues)],
			EventID:      "event_" + string(rune(seed)),
			Title:        "Test Event " + string(rune(seed)),
			Active:       boolPtr(seed%2 == 0),
			Closed:       boolPtr(seed%3 == 0),
			DiscoveredAt: time.Now().Add(time.Duration(seed) * time.Minute),
			LastSeen:     time.Now(),
		},
		EventID:   "evt_" + string(rune(seed)),
		EventType: eventTypes[seed%len(eventTypes)],
		Timestamp: time.Now(),
		VenueID:   venues[seed%len(venues)],
	}
}