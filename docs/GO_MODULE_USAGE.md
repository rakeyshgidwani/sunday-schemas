# Go Module Usage

## Quick Start

Add to your `go.mod`:
```go
require github.com/rakeyshgidwani/sunday-schemas/go v1.0.9
```

Import and use:
```go
import "github.com/rakeyshgidwani/sunday-schemas/go"

// Use generated types
var event schemas.MdOrderbookDeltaV1
```

## Commands

```bash
# Add dependency
go get github.com/rakeyshgidwani/sunday-schemas/go@latest

# Update to latest
go get -u github.com/rakeyshgidwani/sunday-schemas/go

# Download dependencies
go mod download

# Clean up
go mod tidy
```

## Available Types

- `schemas.RawV0` - Raw venue data
- `schemas.MdOrderbookDeltaV1` - Orderbook changes
- `schemas.MdTradeV1` - Trade events
- `schemas.InsightsArbitrageV1` - Arbitrage opportunities
- `schemas.InfraHealthV1` - Health metrics

## Version Pinning

Pin to specific version:
```go
require github.com/rakeyshgidwani/sunday-schemas/go v1.0.8
```

Check available versions:
```bash
go list -m -versions github.com/rakeyshgidwani/sunday-schemas/go
```

## Documentation Access for AI Agents

Access bundled documentation:
```bash
# View module documentation
go doc github.com/rakeyshgidwani/sunday-schemas/go

# List all exported types
go doc -all github.com/rakeyshgidwani/sunday-schemas/go

# View specific type documentation
go doc github.com/rakeyshgidwani/sunday-schemas/go.MdOrderbookDeltaV1

# Get module info and location
go list -m -f '{{.Dir}}' github.com/rakeyshgidwani/sunday-schemas/go
```

Access embedded schema files:
```go
import "github.com/rakeyshgidwani/sunday-schemas/go"

// Read embedded JSON schema
schemaContent := schemas.GetSchema("md.orderbook.delta.v1")

// List all available schemas
schemaList := schemas.ListSchemas()
```