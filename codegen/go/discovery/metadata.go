// Code generated for discovery payload schemas. DO NOT EDIT.

package discovery

import "time"

// EventMetadataV0 contains structured metadata for individual prediction market events
type EventMetadataV0 struct {
	Kind               DiscoveryKind    `json:"kind"`
	VenueID            VenueID          `json:"venue_id"`
	EventID            string           `json:"event_id"`
	Title              string           `json:"title"`
	Description        *string          `json:"description,omitempty"`
	Category           *string          `json:"category,omitempty"`
	Active             *bool            `json:"active"`
	Closed             *bool            `json:"closed"`
	StartDate          *time.Time       `json:"start_date,omitempty"`
	EndDate            *time.Time       `json:"end_date,omitempty"`
	ParentSeriesID     *string          `json:"parent_series_id,omitempty"`
	ParentSeriesTitle  *string          `json:"parent_series_title,omitempty"`
	Tags               []string         `json:"tags,omitempty"`
	Relationships      *RelationshipsV0 `json:"relationships,omitempty"`
	DiscoveredAt       time.Time        `json:"discovered_at"`
	LastSeen           time.Time        `json:"last_seen"`
	ExtraMetadata      map[string]any   `json:"extra_metadata,omitempty"`
}

// SeriesMetadataV0 contains structured metadata for series/collections of prediction market events
type SeriesMetadataV0 struct {
	Kind            DiscoveryKind    `json:"kind"`
	VenueID         VenueID          `json:"venue_id"`
	EventID         string           `json:"event_id"`
	Title           string           `json:"title"`
	Description     *string          `json:"description,omitempty"`
	Category        *string          `json:"category,omitempty"`
	Active          *bool            `json:"active"`
	Closed          *bool            `json:"closed"`
	Tags            []string         `json:"tags,omitempty"`
	ChildEventIDs   []string         `json:"child_event_ids,omitempty"`
	Relationships   *RelationshipsV0 `json:"relationships,omitempty"`
	DiscoveredAt    time.Time        `json:"discovered_at"`
	LastSeen        time.Time        `json:"last_seen"`
	SeriesData      *SeriesDataV0    `json:"series_data,omitempty"`
	ExtraMetadata   map[string]any   `json:"extra_metadata,omitempty"`
}