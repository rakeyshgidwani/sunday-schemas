# Go Module Usage

## Quick Start

Add to your `go.mod`:
```go
require github.com/rakeyshgidwani/sunday-schemas/codegen/go v1.0.9
```

Import and use:
```go
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"

// Use generated types
var event schemas.RawEnvelopeV0
var trade schemas.NormalizedTradeV1
```

## Commands

```bash
# Add dependency
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@latest

# Update to latest
go get -u github.com/rakeyshgidwani/sunday-schemas/codegen/go

# Specific version
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.9

# Download dependencies
go mod download

# Clean up
go mod tidy
```

## Available Types and Constants

**Event Types:**
```go
schemas.EventMetadataV0                     // Generated event type
schemas.Relationships                       // Generated event type
schemas.EventDiscoveryPayloadV0             // Generated event type
schemas.Discovery                           // Generated event type
schemas.EventClass                          // Generated event type
schemas.SeriesMetadataV0                    // Generated event type
schemas.SeriesData                          // Generated event type
schemas.Contract                            // Generated event type
schemas.DiscoverySharedV0Schema             // Generated event type
schemas.Creators                            // Generated event type
schemas.Financial                           // Generated event type
schemas.StatusClass                         // Generated event type
schemas.Timestamps                          // Generated event type
schemas.SeriesDiscoveryPayloadV0            // Generated event type
schemas.SeriesDiscoveryPayloadV0Event       // Generated event type
schemas.VenueHealthV1                       // Generated event type
schemas.ArbitrageLiteV1                     // Generated event type
schemas.MoversV1                            // Generated event type
schemas.UnusualActivityV1                   // Generated event type
schemas.WhaleFlowsLiteV1                    // Generated event type
schemas.NormalizedOrderBookDeltaV1          // Generated event type
schemas.NormalizedTradeV1                   // Generated event type
schemas.RawCategoriesDiscoveryV0            // Generated event type
schemas.RawCategoriesDiscoveryV0Envelope    // Generated event type
schemas.PurpleMetadata                      // Generated event type
schemas.RawEventsDiscoveryV0                // Generated event type
schemas.RawEventsDiscoveryV0Envelope        // Generated event type
schemas.FluffyMetadata                      // Generated event type
schemas.PayloadClass                        // Generated event type
schemas.RawSeriesDiscoveryV0                // Generated event type
schemas.RawSeriesDiscoveryV0Envelope        // Generated event type
schemas.TentacledMetadata                   // Generated event type
schemas.RawSeriesDiscoveryV0Payload         // Generated event type
schemas.RawEnvelopeV0                       // Generated event type
```

**Schema Constants:**
```go
schemas.SchemaINFRA_VENUE_HEALTH_V1    // "infra.venue_health.v1"
schemas.SchemaINSIGHTS_ARB_LITE_V1     // "insights.arb.lite.v1"
schemas.SchemaINSIGHTS_MOVERS_V1       // "insights.movers.v1"
schemas.SchemaINSIGHTS_UNUSUAL_V1      // "insights.unusual.v1"
schemas.SchemaINSIGHTS_WHALES_LITE_V1  // "insights.whales.lite.v1"
schemas.SchemaMD_ORDERBOOK_DELTA_V1    // "md.orderbook.delta.v1"
schemas.SchemaMD_TRADE_V1              // "md.trade.v1"
schemas.SchemaRAW_V0                   // "raw.v0"
```

**Venue Constants:**
```go
schemas.VenuePolymarket                // "polymarket"
schemas.VenueKalshi                    // "kalshi"
```

## Available Functions

**Schema and Venue Validation:**
```go
schemas.ValidateSchema()
schemas.ValidateVenue()
schemas.AllSchemas()
schemas.AllVenues()
```

**JSON Marshal/Unmarshal Functions:**
```go
schemas.UnmarshalEventMetadataV0(data)
schemas.UnmarshalEventDiscoveryPayloadV0(data)
schemas.UnmarshalSeriesMetadataV0(data)
schemas.UnmarshalSeriesDiscoveryPayloadV0(data)
schemas.UnmarshalDiscoverySharedTypesV0(data)
schemas.UnmarshalVenueHealthV1(data)
schemas.UnmarshalArbitrageLiteV1(data)
schemas.UnmarshalMoversV1(data)
schemas.UnmarshalUnusualActivityV1(data)
schemas.UnmarshalWhaleFlowsLiteV1(data)
schemas.UnmarshalNormalizedOrderBookDeltaV1(data)
schemas.UnmarshalNormalizedTradeV1(data)
schemas.UnmarshalRawCategoriesDiscoveryV0(data)
schemas.UnmarshalRawEventsDiscoveryV0(data)
schemas.UnmarshalRawSeriesDiscoveryV0(data)
schemas.UnmarshalRawEnvelopeV0(data)
```

**Helper Functions:**
```go

```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func main() {
    // Validate schema
    if err := schemas.ValidateSchema("raw.v0"); err != nil {
        log.Fatal("Invalid schema:", err)
    }

    // Get all available schemas
    allSchemas := schemas.AllSchemas()
    fmt.Printf("Available schemas: %v\n", allSchemas)

    // Unmarshal JSON data
    data := []byte(`{"schema":"md.trade.v1","venue_id":"polymarket"...}`)
    trade, err := schemas.UnmarshalNormalizedTradeV1(data)
    if err != nil {
        log.Fatal("Failed to unmarshal:", err)
    }

    fmt.Printf("Trade: %+v\n", trade)
}
```

---

*This documentation is automatically generated from the actual Go codegen output.*
*Last updated: 2025-11-04T22:32:45.498Z*
