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

## Available Types

- `schemas.RawEnvelopeV0` - Raw venue data envelope
- `schemas.NormalizedOrderBookDeltaV1` - Orderbook changes
- `schemas.NormalizedTradeV1` - Trade events
- `schemas.ArbitrageLiteV1` - Arbitrage opportunities
- `schemas.VenueHealthV1` - Health metrics
- `schemas.MoversV1` - Price movement events
- `schemas.UnusualActivityV1` - Unusual trading activity
- `schemas.WhaleFlowsLiteV1` - Large trade flows

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

Access embedded schema files (if implemented):
```go
import schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"

// Read embedded JSON schema
schemaContent := schemas.GetSchema("md.orderbook.delta.v1")

// List all available schemas
schemaList := schemas.ListSchemas()
```