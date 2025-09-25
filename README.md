# sunday-schemas

> Schema registry for Sunday platform event definitions and API contracts

This repository serves as the single source of truth for:
- **Event schemas** (JSON Schema) for Kafka topics across Sunday services
- **OpenAPI specification** for the UI BFF HTTP API
- **Generated types** published as npm (`@sunday/schemas`) and Go modules

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

## Development

**Phase 1 - Foundation Setup ✅**
- Repository structure created
- Basic tooling configured
- Registry files established

**Next Steps:**
- Phase 2: Schema Definitions
- Phase 3: OpenAPI Specification
- Phase 4: Validation & Tooling

See [IMPLEMENTATION_PLAN.md](./IMPLEMENTATION_PLAN.md) for detailed implementation roadmap.

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