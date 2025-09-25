# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is `sunday-schemas`, a schema registry for event definitions and API specifications for the Sunday platform. This repository serves as the single source of truth for:

- **Event schemas** (JSON Schema) for Kafka topics across Sunday services
- **OpenAPI specification** for the UI BFF HTTP API (`/ui/*` endpoints)
- **Generated types** published as npm (`@sunday/schemas`) and Go modules (`github.com/sunday-xyz/schemas/go`)

**Important**: This repository contains NO business logic - only schema definitions, validation, and code generation.

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
  /docs/
    mapping.md              # Raw→normalized transformation rules
  /scripts/                 # Build and validation scripts
  /specification/           # Technical specifications and requirements
```

## Development Commands

**Note**: This repository currently has no package.json, Makefile, or other build configuration. Based on the technical specification, the following commands should be implemented:

- JSON Schema validation: `ajv` against examples in `/schemas/examples/`
- OpenAPI linting: `spectral` for API consistency
- Structure validation: `jsonlint` and `yamllint`
- Compatibility checking: `json-schema-diff` and `oas-diff` for backward compatibility
- Code generation: `json-schema-to-typescript` and `openapi-typescript` for TS; `oapi-codegen` and `quicktype` for Go

## Schema Architecture

### Event Schemas (JSON Schema draft/2020-12)
- **`raw.v0`**: Raw venue data envelope from connectors
- **`md.*`**: Normalized market data (orderbook deltas, trades)
- **`insights.*`**: Analytics (arbitrage, movers, whales, unusual activity)
- **`infra.*`**: Infrastructure health metrics

### Key Principles
- **Versioning**: Each schema has a version in its ID (e.g., `md.orderbook.delta.v1`)
- **Backward compatibility**: Only additive changes allowed in Phase 1
- **Enums from registries**: venue_id values sourced from `/schemas/registries/venues.json`
- **Topic mapping**: Schema IDs mapped to Kafka topics in `/schemas/topics.json`

### Price Conventions
- All prices represented as **implied probability** in range `[0.0, 1.0]`
- Orderbook depth arrays are `[price, size]` pairs
- Sequencing must be monotonic per `(instrument_id, venue_id)`

## Validation Requirements

When modifying schemas:
1. All examples in `/schemas/examples/` must validate against their schemas
2. Backward compatibility must be maintained (use diff tools)
3. venue_id enums must reference `/schemas/registries/venues.json`
4. Schema IDs must map correctly in `/schemas/topics.json`
5. OpenAPI endpoints should reuse event schema components where possible

## Publishing Process

Based on the specification:
1. Changes require Schema Working Group review
2. PRs must include CHANGELOG.md entries with SemVer bump
3. CI gates prevent breaking changes
4. Git tags trigger automated publishing to npm and Go modules
5. Generated artifacts include types, checksums, and SBOM

## Phase 1 Constraints

- **Breaking changes blocked**: No field removal, type changes, or enum narrowing
- **Supported venues**: Currently `["polymarket", "kalshi"]`
- **Auth not modeled**: UI endpoints assume gateway-level auth
- **Single venue messages**: Raw payloads must be single objects, not arrays