# Sunday Schemas - Go Types

Generated Go types and constants for the Sunday platform event schemas.

## Installation

```bash
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@latest
```

## Usage

### Import

```go
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
```

### Event Processing

```go
package main

import (
    "encoding/json"
    "fmt"

    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func processEvent(rawEvent []byte) error {
    // Parse schema type
    var envelope struct {
        Schema string `json:"schema"`
    }

    if err := json.Unmarshal(rawEvent, &envelope); err != nil {
        return err
    }

    // Validate schema
    if err := schemas.ValidateSchema(envelope.Schema); err != nil {
        return fmt.Errorf("invalid schema: %w", err)
    }

    // Process by type
    switch schemas.EventSchema(envelope.Schema) {
    case schemas.SchemaMD_TRADE_V1:
        var trade schemas.Trade
        if err := json.Unmarshal(rawEvent, &trade); err != nil {
            return err
        }
        return processTrade(trade)

    case schemas.SchemaMD_ORDERBOOK_DELTA_V1:
        var orderbook schemas.OrderbookDelta
        if err := json.Unmarshal(rawEvent, &orderbook); err != nil {
            return err
        }
        return processOrderbook(orderbook)
    }

    return nil
}
```

### Constants

```go
// Schema validation
if err := schemas.ValidateSchema("md.trade.v1"); err != nil {
    // Invalid schema
}

// Venue validation
if err := schemas.ValidateVenue("polymarket"); err != nil {
    // Invalid venue
}

// Get all supported venues
venues := schemas.AllVenues() // []VenueID{"polymarket", "kalshi"}

// Health status
status := schemas.HealthConnected // "CONNECTED"
```

## Generated Types

### Event Types
- `RawEnvelope` - Raw venue data envelope (`raw.v0`)
- `Trade` - Normalized trade events (`md.trade.v1`)
- `OrderbookDelta` - Orderbook deltas (`md.orderbook.delta.v1`)
- `ArbitrageLite` - Arbitrage opportunities (`insights.arb.lite.v1`)
- `Movers` - Price movers (`insights.movers.v1`)
- `WhalesLite` - Large trades (`insights.whales.lite.v1`)
- `EventUnusual` - Unusual activity (`insights.unusual.v1`)
- `EventVenueHealth` - Venue health (`infra.venue_health.v1`)

### Constants
- Schema identifiers (e.g., `SchemaMD_TRADE_V1`)
- Venue IDs (`VenuePolymarket`, `VenueKalshi`)
- Health statuses (`HealthConnected`, `HealthDegraded`, `HealthStale`)

## Generated From

This module is automatically generated from [Sunday Schemas](https://github.com/rakeyshgidwani/sunday-schemas). Do not modify these files directly.

## License

MIT