// Code generated for discovery payload schemas. DO NOT EDIT.

package discovery

// DiscoveryKind discriminates between event and series metadata
type DiscoveryKind string

const (
	DiscoveryKindEvent  DiscoveryKind = "event"
	DiscoveryKindSeries DiscoveryKind = "series"
)

// EventType represents the type of discovery event
type EventType string

const (
	EventTypeDiscovered EventType = "discovered"
	EventTypeUpdated    EventType = "updated"
	EventTypeExpired    EventType = "expired"
)

// VenueID represents supported prediction market venues
type VenueID string

const (
	VenueIDPolymarket VenueID = "polymarket"
	VenueIDKalshi     VenueID = "kalshi"
)

// Currency represents supported currency units
type Currency string

const (
	CurrencyUSD Currency = "USD"
)