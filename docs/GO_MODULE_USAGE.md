# Go Module Usage

## Quick Start

Add to your `go.mod`:
```go
require github.com/rakeyshgidwani/sunday-schemas/codegen/go v1.0.9
```

Import and use:
```go
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"

// Use generated types - all types in one package
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
- `schemas.RawEnvelopeV0` - Raw venue data envelope
- `schemas.NormalizedOrderBookDeltaV1` - Orderbook changes
- `schemas.NormalizedTradeV1` - Trade events
- `schemas.ArbitrageLiteV1` - Arbitrage opportunities
- `schemas.VenueHealthV1` - Health metrics
- `schemas.MoversV1` - Price movement events
- `schemas.UnusualActivityV1` - Unusual trading activity
- `schemas.WhaleFlowsLiteV1` - Large trade flows

**Schema Constants:**
```go
schemas.SchemaRAW_V0                    // "raw.v0"
schemas.SchemaMD_TRADE_V1               // "md.trade.v1"
schemas.SchemaMD_ORDERBOOK_DELTA_V1     // "md.orderbook.delta.v1"
schemas.SchemaINSIGHTS_ARB_LITE_V1      // "insights.arb.lite.v1"
schemas.SchemaINSIGHTS_MOVERS_V1        // "insights.movers.v1"
schemas.SchemaINSIGHTS_UNUSUAL_V1       // "insights.unusual.v1"
schemas.SchemaINSIGHTS_WHALES_LITE_V1   // "insights.whales.lite.v1"
schemas.SchemaINFRA_VENUE_HEALTH_V1     // "infra.venue_health.v1"
```

**Venue Constants:**
```go
schemas.VenuePolymarket   // "polymarket"
schemas.VenueKalshi       // "kalshi"
```

**Other Constants:**
```go
// Trade sides
schemas.TradeSideBuy      // "buy"
schemas.TradeSideSell     // "sell"

// Health statuses
schemas.HealthConnected   // "CONNECTED"
schemas.HealthDegraded    // "DEGRADED"
schemas.HealthStale       // "STALE"
```

## Version Pinning

Pin to specific version:
```go
require github.com/rakeyshgidwani/sunday-schemas/codegen/go v1.0.9
```

Check available versions:
```bash
go list -m -versions github.com/rakeyshgidwani/sunday-schemas/codegen/go
```

## Documentation Access

Access bundled documentation:
```bash
# View module documentation
go doc github.com/rakeyshgidwani/sunday-schemas/codegen/go

# List all exported types
go doc -all github.com/rakeyshgidwani/sunday-schemas/codegen/go

# View specific type documentation
go doc github.com/rakeyshgidwani/sunday-schemas/codegen/go.RawEnvelopeV0

# Get module info and location
go list -m -f '{{.Dir}}' github.com/rakeyshgidwani/sunday-schemas/codegen/go
```

## Available Functions

**Schema and Venue Validation:**
```go
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"

// Validate schema identifier
if err := schemas.ValidateSchema("raw.v0"); err != nil {
    log.Fatal("Invalid schema:", err)
}

// Validate venue identifier
if err := schemas.ValidateVenue("polymarket"); err != nil {
    log.Fatal("Invalid venue:", err)
}

// Get all available schemas
allSchemas := schemas.AllSchemas()
// Returns: []EventSchema containing all schema constants

// Get all available venues
allVenues := schemas.AllVenues()
// Returns: []VenueID containing all venue constants
```

**JSON Marshal/Unmarshal for Each Schema Type:**
```go
// Each schema type has dedicated marshal/unmarshal functions:

// Raw Envelope
data := []byte(`{"schema":"raw.v0", "venue_id":"polymarket"...}`)
envelope, err := schemas.UnmarshalRawEnvelopeV0(data)
bytes, err := envelope.Marshal()

// Trade Events
tradeData := []byte(`{"schema":"md.trade.v1"...}`)
trade, err := schemas.UnmarshalNormalizedTradeV1(tradeData)
bytes, err := trade.Marshal()

// Orderbook Events
orderbookData := []byte(`{"schema":"md.orderbook.delta.v1"...}`)
orderbook, err := schemas.UnmarshalNormalizedOrderBookDeltaV1(orderbookData)
bytes, err := orderbook.Marshal()

// Insights Events
arbData := []byte(`{"schema":"insights.arb.lite.v1"...}`)
arb, err := schemas.UnmarshalArbitrageLiteV1(arbData)
bytes, err := arb.Marshal()
```

**Helper Functions (Compatibility Layer):**
```go
// Create raw envelope with helper function
envelope := schemas.NewRawEnvelope(
    "polymarket",
    "orderbook",
    "instrument-123",
    time.Now(),
    payloadData,
)

// Convert between envelope formats
v0Envelope := envelope.ToRawEnvelopeV0()
backToEnvelope := schemas.FromRawEnvelopeV0(v0Envelope)
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "time"

    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func main() {
    // Example 1: Create and validate a raw envelope
    envelope := schemas.NewRawEnvelope(
        "polymarket",
        "orderbook",
        "some-instrument-id",
        time.Now(),
        map[string]interface{}{
            "bids": [][]float64{{0.51, 1000.0}},
            "asks": [][]float64{{0.49, 500.0}},
        },
    )

    // Validate the venue
    if err := schemas.ValidateVenue(string(envelope.VenueID)); err != nil {
        log.Fatal("Invalid venue:", err)
    }

    // Convert to V0 format and marshal to JSON
    v0Envelope := envelope.ToRawEnvelopeV0()
    jsonData, err := v0Envelope.Marshal()
    if err != nil {
        log.Fatal("Failed to marshal:", err)
    }

    fmt.Printf("Raw envelope JSON: %s\n", jsonData)

    // Example 2: Unmarshal and validate a trade event
    tradeJSON := `{
        "schema": "md.trade.v1",
        "instrument_id": "poly-123",
        "venue_id": "polymarket",
        "ts_ms": 1640995200000,
        "side": "buy",
        "prob": 0.52,
        "size": 100.0
    }`

    trade, err := schemas.UnmarshalNormalizedTradeV1([]byte(tradeJSON))
    if err != nil {
        log.Fatal("Failed to unmarshal trade:", err)
    }

    // Validate using constants
    if err := schemas.ValidateSchema(trade.Schema); err != nil {
        log.Fatal("Invalid schema:", err)
    }

    if err := schemas.ValidateVenue(trade.VenueID); err != nil {
        log.Fatal("Invalid venue:", err)
    }

    fmt.Printf("Trade: %+v\n", trade)

    // Example 3: List all available schemas and venues
    fmt.Println("Available schemas:")
    for _, schema := range schemas.AllSchemas() {
        fmt.Printf("  - %s\n", schema)
    }

    fmt.Println("Available venues:")
    for _, venue := range schemas.AllVenues() {
        fmt.Printf("  - %s\n", venue)
    }
}
```