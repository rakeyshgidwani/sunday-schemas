# sunday-schemas

> Schema registry for Sunday platform event definitions and API contracts

This repository serves as the single source of truth for:
- **Event schemas** (JSON Schema) for Kafka topics across Sunday services
- **OpenAPI specification** for the UI BFF HTTP API
- **Generated types** published as npm (`@rakeyshgidwani/sunday-schemas`) and Go modules

## Repository Structure

```
/sunday-schemas
  /schemas
    /json/                    # JSON Schema files for all events
    /examples/                # Example data for validation testing
    /registries/              # Canonical venue and instrument lists
    topics.json              # Schema ID → Kafka topic mapping
  /openapi/
    ui.v1.yaml              # OpenAPI spec for UI BFF endpoints
  /codegen/
    /ts/                    # Generated TypeScript types
    /go/                    # Generated Go types
  /docs/                    # Documentation
  /scripts/                 # Build and validation scripts
```

## Usage

### Installation

**TypeScript/JavaScript:**
```bash
npm install @rakeyshgidwani/sunday-schemas
```

**Go:**
```bash
go get github.com/rakeyshgidwani/sunday-schemas/go@latest
```

### Development

**Available Commands:**
```bash
npm run validate                    # Validate schemas and examples
npm run build                       # Generate TypeScript and Go types
npm run deploy                      # Deploy new version (interactive)
npm run validate-deployment         # Validate deployed packages
```

**Deployment:**
```bash
# Deploy specific version
./scripts/deploy.sh --version 1.0.2

# Preview deployment
npm run deploy:dry-run --version 1.0.2
```

See [DEPLOYMENT.md](./DEPLOYMENT.md) for detailed deployment guide.

## Documentation

📚 **For Development Teams**: [Complete Integration Guide](./docs/INTEGRATION_GUIDE_FOR_TEAMS.md)
📊 **For Stakeholders**: [Executive Summary](./docs/EXECUTIVE_SUMMARY.md)
🚀 **For Deployment**: [Deployment Guide](./DEPLOYMENT.md)
🐹 **Go Usage**: [Go Module Integration](./docs/GO_MODULE_USAGE.md)
📦 **TypeScript Usage**: [TypeScript Integration](./docs/TYPESCRIPT_USAGE.md)

## Architecture

This repository contains **no business logic** - only schema definitions, validation, and code generation.

### Key Principles
- **Backward compatibility**: Only additive changes allowed in Phase 1
- **Versioning**: Each schema has version in ID (e.g., `md.orderbook.delta.v1`)
- **Price conventions**: All prices as implied probability [0.0, 1.0]
- **Registry-driven**: Venue enums sourced from `venues.json`

### Supported Venues
- Polymarket
- Kalshi

## License

ISC