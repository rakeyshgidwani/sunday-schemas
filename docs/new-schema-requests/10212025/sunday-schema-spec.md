# Sunday Schema Specification: Discovery Schemas

## Overview
This document specifies new schemas required for the Sunday Connectors discovery implementation. These schemas will handle Events, Series, and Categories discovery across multiple prediction market venues.

**📋 PURPOSE**: Define schema requirements for Events, Series, and Categories discovery to complement the existing market discovery schemas.

## Schema Requirements

### 1. `raw.events.v0`

**Purpose**: Capture event discovery data from prediction market venues
**Kafka Topics**:
- `dev.sunday.events.discovered` - New events found
- `dev.sunday.events.updated` - Event state changes
- `dev.sunday.events.expired` - Event closures/expiration

#### Schema Structure:
```json
{
  "type": "object",
  "required": ["envelope", "payload"],
  "properties": {
    "envelope": {
      "type": "object",
      "required": ["venue_id", "stream", "schema", "timestamp"],
      "properties": {
        "venue_id": {
          "type": "string",
          "enum": ["polymarket", "kalshi"]
        },
        "stream": {
          "type": "string",
          "const": "event_discovery"
        },
        "schema": {
          "type": "string",
          "const": "raw.events.v0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "metadata": {
          "type": "object",
          "properties": {
            "discovery_timestamp": {
              "type": "string",
              "format": "date-time"
            },
            "discovery_page": {
              "type": "integer"
            }
          }
        }
      }
    },
    "payload": {
      "type": "object",
      "description": "Raw venue-native event data - preserves original API response structure",
      "additionalProperties": true
    }
  }
}
```

#### Payload Examples:

**Polymarket Event:**
```json
{
  "envelope": {
    "venue_id": "polymarket",
    "stream": "event_discovery",
    "schema": "raw.events.v0",
    "timestamp": "2024-10-20T14:30:00Z",
    "metadata": {
      "discovery_timestamp": "2024-10-20T14:30:00Z",
      "discovery_page": 1
    }
  },
  "payload": {
    "id": "event_12345",
    "title": "2024 Presidential Election",
    "description": "Who will win the 2024 US Presidential Election?",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-11-05T23:59:59Z",
    "status": "active",
    "category": "Politics",
    "tags": ["politics", "election", "2024"],
    "volume": 1500000,
    "markets": ["market_1", "market_2"],
    "image_url": "https://...",
    "resolution_source": "AP News"
  }
}
```

**Kalshi Event:**
```json
{
  "envelope": {
    "venue_id": "kalshi",
    "stream": "event_discovery",
    "schema": "raw.events.v0",
    "timestamp": "2024-10-20T14:30:00Z",
    "metadata": {
      "discovery_timestamp": "2024-10-20T14:30:00Z",
      "discovery_page": 2
    }
  },
  "payload": {
    "event_ticker": "PRES24",
    "title": "2024 Presidential Election",
    "category": "Politics",
    "sub_title": "Republican vs Democratic nominee",
    "mutually_exclusive": true,
    "start_date": "2024-01-01T00:00:00Z",
    "settle_date": "2024-11-06T00:00:00Z",
    "status": "open",
    "markets_count": 5,
    "volume": 2300000
  }
}
```

---

### 2. `raw.series.v0`

**Purpose**: Capture series/collections discovery data
**Kafka Topics**:
- `dev.sunday.series.discovered` - New series found
- `dev.sunday.series.updated` - Series state changes
- `dev.sunday.series.expired` - Series closures/completion

#### Schema Structure:
```json
{
  "type": "object",
  "required": ["envelope", "payload"],
  "properties": {
    "envelope": {
      "type": "object",
      "required": ["venue_id", "stream", "schema", "timestamp"],
      "properties": {
        "venue_id": {
          "type": "string",
          "enum": ["polymarket", "kalshi"]
        },
        "stream": {
          "type": "string",
          "const": "series_discovery"
        },
        "schema": {
          "type": "string",
          "const": "raw.series.v0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "metadata": {
          "type": "object",
          "properties": {
            "discovery_timestamp": {
              "type": "string",
              "format": "date-time"
            },
            "discovery_page": {
              "type": "integer"
            }
          }
        }
      }
    },
    "payload": {
      "type": "object",
      "description": "Raw venue-native series data - preserves original API response structure",
      "additionalProperties": true
    }
  }
}
```

#### Payload Examples:

**Polymarket Series:**
```json
{
  "envelope": {
    "venue_id": "polymarket",
    "stream": "series_discovery",
    "schema": "raw.series.v0",
    "timestamp": "2024-10-20T14:30:00Z"
  },
  "payload": {
    "id": "series_nfl_2024",
    "name": "NFL 2024 Season",
    "description": "All NFL-related markets for the 2024 season",
    "category": "Sports",
    "tags": ["nfl", "football", "2024"],
    "events": ["event_1", "event_2", "event_3"],
    "start_date": "2024-09-01T00:00:00Z",
    "end_date": "2025-02-15T00:00:00Z",
    "status": "active"
  }
}
```

**Kalshi Series:**
```json
{
  "envelope": {
    "venue_id": "kalshi",
    "stream": "series_discovery",
    "schema": "raw.series.v0",
    "timestamp": "2024-10-20T14:30:00Z"
  },
  "payload": {
    "series_ticker": "NFL24",
    "title": "NFL 2024",
    "frequency": "weekly",
    "category": "Sports",
    "settlement_source": "ESPN",
    "tags": ["nfl", "sports"],
    "event_count": 25,
    "status": "open"
  }
}
```

---

### 3. `raw.categories.v0`

**Purpose**: Capture category/tag discovery data for unified taxonomy
**Kafka Topics**:
- `dev.sunday.categories.discovered` - New categories found
- `dev.sunday.categories.updated` - Category changes
- `dev.sunday.categories.expired` - Category removal/deprecation

#### Schema Structure:
```json
{
  "type": "object",
  "required": ["envelope", "payload"],
  "properties": {
    "envelope": {
      "type": "object",
      "required": ["venue_id", "stream", "schema", "timestamp"],
      "properties": {
        "venue_id": {
          "type": "string",
          "enum": ["polymarket", "kalshi"]
        },
        "stream": {
          "type": "string",
          "const": "category_discovery"
        },
        "schema": {
          "type": "string",
          "const": "raw.categories.v0"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "metadata": {
          "type": "object",
          "properties": {
            "discovery_timestamp": {
              "type": "string",
              "format": "date-time"
            }
          }
        }
      }
    },
    "payload": {
      "type": "object",
      "description": "Raw venue-native category/tag data - preserves original API response structure",
      "additionalProperties": true
    }
  }
}
```

#### Payload Examples:

**Polymarket Categories:**
```json
{
  "envelope": {
    "venue_id": "polymarket",
    "stream": "category_discovery",
    "schema": "raw.categories.v0",
    "timestamp": "2024-10-20T14:30:00Z"
  },
  "payload": {
    "id": "tag_politics",
    "name": "Politics",
    "slug": "politics",
    "description": "Political events and elections",
    "color": "#ff6b35",
    "parent_category": null,
    "subcategories": ["elections", "policy"],
    "market_count": 150,
    "volume": 5000000
  }
}
```

**Kalshi Categories (derived from events):**
```json
{
  "envelope": {
    "venue_id": "kalshi",
    "stream": "category_discovery",
    "schema": "raw.categories.v0",
    "timestamp": "2024-10-20T14:30:00Z"
  },
  "payload": {
    "category": "Politics",
    "event_count": 75,
    "market_count": 200,
    "total_volume": 3500000,
    "subcategories": ["Elections", "Policy", "International"],
    "derived_from": "event_aggregation"
  }
}
```

## Schema Validation Requirements

### Common Envelope Validation:
- `venue_id` must be one of the supported venues
- `timestamp` must be valid ISO 8601 datetime
- `schema` must match the exact schema version

### Payload Validation:
- Preserve venue-native structure (no transformation)
- Allow additional properties for forward compatibility
- Validate required fields exist but don't enforce structure



## Topic Naming Convention

All discovery schemas follow the pattern:
- **Development**: `dev.sunday.{type}.{lifecycle}`
- **Staging**: `staging.sunday.{type}.{lifecycle}`
- **Production**: `prod.sunday.{type}.{lifecycle}`

Where:
- `{type}` is: `events`, `series`, `categories`
- `{lifecycle}` is: `discovered`, `updated`, `expired`

**Examples**:
- `dev.sunday.events.discovered` (new)
- `staging.sunday.series.updated` (new)
- `prod.sunday.categories.expired` (new)

## Backward Compatibility

These schemas are net-new and do not impact existing schemas:
- `raw.v0` remains unchanged for real-time market data
- No breaking changes to current schema contracts


## Implementation Notes

### Envelope Consistency:
All discovery schemas share the same envelope structure for consistent processing pipelines.

### Raw Data Philosophy:
Payloads preserve venue-native formats to avoid data loss and enable flexible downstream processing, consistent with existing `raw.v0` approach.

### Schema Evolution:
Future versions (v1, v2) can add envelope metadata or validation rules without breaking v0 consumers, following the established versioning pattern.