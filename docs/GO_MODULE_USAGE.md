# Go Module Usage

There are **two Go modules** available with different structures:

## Option 1: Unified Package (Recommended)

**Best for:** Most use cases, matches documentation examples

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

## Option 2: Modular Packages

**Best for:** When you only need specific schema types

Add to your `go.mod`:
```go
require github.com/rakeyshgidwani/sunday-schemas/go v1.0.9
```

Import and use:
```go
import "github.com/rakeyshgidwani/sunday-schemas/go/rawv0"
import "github.com/rakeyshgidwani/sunday-schemas/go/md"

// Use types from separate packages
var event rawv0.RawEnvelopeV0
var trade md.NormalizedTradeV1
```

## Commands

**For Unified Package (Option 1):**
```bash
# Add dependency
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@latest

# Update to latest
go get -u github.com/rakeyshgidwani/sunday-schemas/codegen/go

# Specific version
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.9
```

**For Modular Packages (Option 2):**
```bash
# Add dependency
go get github.com/rakeyshgidwani/sunday-schemas/go@latest

# Update to latest
go get -u github.com/rakeyshgidwani/sunday-schemas/go

# Specific version
go get github.com/rakeyshgidwani/sunday-schemas/go@v1.0.9
```

**Common commands:**
```bash
# Download dependencies
go mod download

# Clean up
go mod tidy
```

## Available Types

**Unified Package (`codegen/go`):**
- `schemas.RawEnvelopeV0` - Raw venue data envelope
- `schemas.NormalizedOrderBookDeltaV1` - Orderbook changes
- `schemas.NormalizedTradeV1` - Trade events
- `schemas.ArbitrageLiteV1` - Arbitrage opportunities
- `schemas.VenueHealthV1` - Health metrics

**Modular Packages (`go/`):**
- `rawv0.RawEnvelopeV0` - Raw venue data envelope
- `md.OrderbookDeltaV1` - Orderbook changes
- `md.TradeV1` - Trade events
- `insights.ArbitrageLiteV1` - Arbitrage opportunities
- Package-specific constants and enums

## Version Pinning

**Unified Package:**
```go
require github.com/rakeyshgidwani/sunday-schemas/codegen/go v1.0.9
```

**Modular Packages:**
```go
require github.com/rakeyshgidwani/sunday-schemas/go v1.0.9
```

Check available versions:
```bash
# Unified package
go list -m -versions github.com/rakeyshgidwani/sunday-schemas/codegen/go

# Modular packages
go list -m -versions github.com/rakeyshgidwani/sunday-schemas/go
```

## Documentation Access

**Unified Package:**
```bash
# View module documentation
go doc github.com/rakeyshgidwani/sunday-schemas/codegen/go

# List all exported types
go doc -all github.com/rakeyshgidwani/sunday-schemas/codegen/go

# View specific type documentation
go doc github.com/rakeyshgidwani/sunday-schemas/codegen/go.RawEnvelopeV0
```

**Modular Packages:**
```bash
# View specific package documentation
go doc github.com/rakeyshgidwani/sunday-schemas/go/rawv0

# View specific type documentation
go doc github.com/rakeyshgidwani/sunday-schemas/go/rawv0.RawEnvelopeV0
```

**Both support embedded schema access** (if implemented):
```go
// Read embedded JSON schema
schemaContent := GetSchema("md.orderbook.delta.v1")

// List all available schemas
schemaList := ListSchemas()
```