// Code generated for discovery payload schemas. DO NOT EDIT.

package discovery

import "time"

// EventDiscoveryPayloadV0 contains structured payload for event discovery messages
type EventDiscoveryPayloadV0 struct {
	Event         EventMetadataV0  `json:"event"`
	EventID       string           `json:"event_id"`
	EventType     EventType        `json:"event_type"`
	Timestamp     time.Time        `json:"timestamp"`
	VenueID       VenueID          `json:"venue_id"`
	DiscoveryMeta *DiscoveryMetaV0 `json:"discovery_meta,omitempty"`
}

// SeriesDiscoveryPayloadV0 contains structured payload for series discovery messages
type SeriesDiscoveryPayloadV0 struct {
	Event         SeriesMetadataV0 `json:"event"`
	EventID       string           `json:"event_id"`
	EventType     EventType        `json:"event_type"`
	Timestamp     time.Time        `json:"timestamp"`
	VenueID       VenueID          `json:"venue_id"`
	DiscoveryMeta *DiscoveryMetaV0 `json:"discovery_meta,omitempty"`
}