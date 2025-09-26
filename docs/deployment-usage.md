# Deployment & Usage Guide

This document explains how the Sunday Schema Registry gets deployed and used across other projects in the Sunday platform.

## 🏗️ Deployment Architecture

### 1. Schema Registry Repository (This Project)
```
sunday-schemas/
├── schemas/json/           # Source of truth for all schemas
├── codegen/ts/            # Generated TypeScript types
├── codegen/go/            # Generated Go types
├── openapi/               # API specifications
└── .github/workflows/     # CI/CD automation
```

**Deployment**: GitHub repository with automated publishing

### 2. Generated Packages

#### NPM Package: `@sunday-platform/schemas`
- **Published to**: npm registry
- **Trigger**: Git tags (e.g., `v1.2.3`)
- **Contains**: TypeScript types, API client types, validation helpers
- **Size**: ~50KB compressed

#### Go Module: `github.com/rakeyshgidwani/sunday-schemas/codegen/go`
- **Published to**: GitHub (Go modules)
- **Trigger**: Git tags (e.g., `v1.2.3`)
- **Contains**: Go structs, constants, validation functions
- **Import Path**: Standard Go module import

### 3. Schema Registry Service (Optional)
```
sunday-schema-registry-service/
├── api/                   # REST API for schema lookup
├── validation/            # Runtime schema validation
├── migration/             # Schema migration helpers
└── metrics/              # Usage analytics
```

## 📦 Package Consumption

### TypeScript/JavaScript Projects

#### Installation
```bash
# Install the generated types package
npm install @sunday-platform/schemas

# For development (if using schema validation)
npm install --save-dev ajv ajv-formats
```

#### Usage in Frontend (React/Next.js)
```typescript
// types/events.ts
import type {
  Trade,
  OrderbookDelta,
  SundayEvent,
  VenueId
} from '@sunday-platform/schemas';

// Real-time event handler
export function useEventStream() {
  const handleEvent = useCallback((event: SundayEvent) => {
    switch (event.schema) {
      case 'md.trade.v1':
        // TypeScript knows this is a Trade
        updateTrade(event.instrument_id, event.prob, event.size);
        break;
      case 'md.orderbook.delta.v1':
        // TypeScript knows this is OrderbookDelta
        updateOrderbook(event.instrument_id, event.bids, event.asks);
        break;
    }
  }, []);

  return { handleEvent };
}

// Component usage
export function TradingDashboard() {
  const { handleEvent } = useEventStream();

  useWebSocket('ws://api.sunday.com/events', {
    onMessage: (data) => {
      const event = JSON.parse(data) as SundayEvent;
      handleEvent(event);
    }
  });
}
```

#### Usage in Backend (Node.js/Express)
```typescript
// services/eventProcessor.ts
import {
  SundayEvent,
  SCHEMA_CONSTANTS,
  EventBySchema
} from '@sunday-platform/schemas';

export class EventProcessor {
  async processEvent(rawEvent: unknown): Promise<void> {
    // Runtime validation
    if (!this.isValidEvent(rawEvent)) {
      throw new Error('Invalid event structure');
    }

    const event = rawEvent as SundayEvent;

    switch (event.schema) {
      case 'md.trade.v1':
        await this.processTrade(event);
        break;
      case 'insights.arb.lite.v1':
        await this.processArbitrage(event);
        break;
    }
  }

  private isValidEvent(event: unknown): event is SundayEvent {
    return typeof event === 'object' &&
           event !== null &&
           'schema' in event &&
           typeof event.schema === 'string' &&
           event.schema in SCHEMA_CONSTANTS;
  }
}
```

#### API Client Usage
```typescript
// api/sundayClient.ts
import type { paths } from '@sunday-platform/schemas';

type ApiClient = {
  getMarkets(): Promise<paths['/markets']['get']['responses']['200']['content']['application/json']>;
  getArbitrageOpportunities(): Promise<paths['/arb']['get']['responses']['200']['content']['application/json']>;
};

export const sundayApi: ApiClient = {
  async getMarkets() {
    const response = await fetch('/api/markets');
    return response.json();
  },

  async getArbitrageOpportunities() {
    const response = await fetch('/api/arb');
    return response.json();
  }
};
```

### Go Projects

#### Installation
```bash
# Add to go.mod
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@latest

# For API client (if needed)
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go/api@latest
```

#### Usage in Data Pipeline Service
```go
// internal/processor/events.go
package processor

import (
    "encoding/json"
    "fmt"

    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

type EventProcessor struct {
    tradeHandler      func(schemas.Trade) error
    orderbookHandler  func(schemas.OrderbookDelta) error
}

func (p *EventProcessor) ProcessEvent(rawEvent []byte) error {
    // First, determine event type by parsing schema field
    var eventEnvelope struct {
        Schema string `json:"schema"`
    }

    if err := json.Unmarshal(rawEvent, &eventEnvelope); err != nil {
        return fmt.Errorf("failed to parse event schema: %w", err)
    }

    // Validate schema
    if err := schemas.ValidateSchema(eventEnvelope.Schema); err != nil {
        return fmt.Errorf("invalid schema: %w", err)
    }

    // Process based on schema type
    switch schemas.EventSchema(eventEnvelope.Schema) {
    case schemas.SchemaMD_TRADE_V1:
        var trade schemas.Trade
        if err := json.Unmarshal(rawEvent, &trade); err != nil {
            return fmt.Errorf("failed to parse trade event: %w", err)
        }
        return p.tradeHandler(trade)

    case schemas.SchemaMD_ORDERBOOK_DELTA_V1:
        var orderbook schemas.OrderbookDelta
        if err := json.Unmarshal(rawEvent, &orderbook); err != nil {
            return fmt.Errorf("failed to parse orderbook event: %w", err)
        }
        return p.orderbookHandler(orderbook)

    default:
        return fmt.Errorf("unsupported event type: %s", eventEnvelope.Schema)
    }
}
```

#### Usage in API Service
```go
// cmd/api/main.go
package main

import (
    "context"
    "net/http"

    api "github.com/rakeyshgidwani/sunday-schemas/codegen/go/api"
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func main() {
    // Create API client for upstream Sunday services
    client, err := api.NewClient("https://internal-api.sunday.com")
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/markets", func(w http.ResponseWriter, r *http.Request) {
        // Use generated client
        resp, err := client.GetMarkets(context.Background(), &api.GetMarketsParams{
            Limit: &[]int{50}[0],
            Venues: &[]string{"polymarket,kalshi"}[0],
        })
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Response is properly typed
        if resp.JSON200 != nil {
            writeJSON(w, resp.JSON200.Markets)
        }
    })
}
```

## 🚀 Deployment Scenarios

### Scenario 1: Microservices Architecture

```yaml
# docker-compose.yml for Sunday platform
version: '3.8'
services:
  normalizer:
    build: ./sunday-data/normalizer
    environment:
      - SCHEMA_REGISTRY_URL=https://schemas.sunday.com
    volumes:
      - ./schemas:/app/schemas  # Mount schema files directly

  insights:
    build: ./sunday-data/insights
    depends_on: [normalizer]

  ui-bff:
    build: ./sunday-api/ui-bff
    environment:
      - SCHEMA_VERSION=v1.2.3  # Pin schema version
    ports:
      - "3000:3000"

  frontend:
    build: ./sunday-frontend/web
    # Uses @sunday-platform/schemas package
    depends_on: [ui-bff]
```

### Scenario 2: Kubernetes Deployment

```yaml
# k8s/schema-registry.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: schema-registry
spec:
  replicas: 3
  selector:
    matchLabels:
      app: schema-registry
  template:
    spec:
      containers:
      - name: schema-registry
        image: sunday/schema-registry:v1.2.3
        ports:
        - containerPort: 8080
        env:
        - name: SCHEMA_VERSION
          value: "v1.2.3"
        volumeMounts:
        - name: schemas
          mountPath: /app/schemas
      volumes:
      - name: schemas
        configMap:
          name: sunday-schemas
---
# Generated from schema files
apiVersion: v1
kind: ConfigMap
metadata:
  name: sunday-schemas
data:
  raw.v0.schema.json: |
    {{ .Files.Get "schemas/json/raw.v0.envelope.schema.json" }}
  # ... other schemas
```

### Scenario 3: Serverless (AWS Lambda/Vercel)

```typescript
// api/events/validate.ts (Vercel API route)
import { NextApiRequest, NextApiResponse } from 'next';
import { SundayEvent, SCHEMA_CONSTANTS } from '@sunday-platform/schemas';
import Ajv from 'ajv';

// Load schemas at runtime (cached)
const ajv = new Ajv();
const validators = new Map();

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  if (req.method !== 'POST') {
    return res.status(405).json({ error: 'Method not allowed' });
  }

  const event = req.body as SundayEvent;

  // Validate schema exists
  if (!event.schema || !(event.schema in SCHEMA_CONSTANTS)) {
    return res.status(400).json({
      error: 'Invalid schema',
      schema: event.schema
    });
  }

  // Get validator (cached)
  let validator = validators.get(event.schema);
  if (!validator) {
    // Load schema definition from deployed package
    const schema = await import(`@sunday-platform/schemas/schemas/${event.schema}.json`);
    validator = ajv.compile(schema);
    validators.set(event.schema, validator);
  }

  // Validate event
  const valid = validator(event);
  if (!valid) {
    return res.status(400).json({
      error: 'Validation failed',
      errors: validator.errors
    });
  }

  return res.status(200).json({ valid: true });
}
```

## 🔄 CI/CD Integration

### Consumer Project Setup

#### Package.json Dependencies
```json
{
  "dependencies": {
    "@sunday-platform/schemas": "^1.2.0"
  },
  "scripts": {
    "check-schema-compatibility": "node scripts/check-sunday-schemas.js",
    "update-schemas": "npm update @sunday-platform/schemas"
  }
}
```

#### Schema Compatibility Check
```javascript
// scripts/check-sunday-schemas.js
const { execSync } = require('child_process');
const currentVersion = require('@sunday-platform/schemas/package.json').version;
const lastKnownGood = process.env.SCHEMA_VERSION || '1.1.0';

console.log(`Checking schema compatibility: ${lastKnownGood} -> ${currentVersion}`);

// Check if this is a breaking change
const [currentMajor, currentMinor] = currentVersion.split('.').map(Number);
const [lastMajor, lastMinor] = lastKnownGood.split('.').map(Number);

if (currentMajor > lastMajor) {
  console.error('❌ Major version change detected - manual review required');
  process.exit(1);
}

if (currentMajor === lastMajor && currentMinor > lastMinor + 1) {
  console.warn('⚠️ Minor version jump detected - review recommended');
}

console.log('✅ Schema compatibility check passed');
```

#### GitHub Actions Workflow
```yaml
# .github/workflows/test.yml
name: Test with Schema Validation

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '18'

      - name: Install dependencies
        run: npm ci

      - name: Check schema compatibility
        run: npm run check-schema-compatibility

      - name: Run tests with schema validation
        run: npm test
        env:
          SCHEMA_VALIDATION: strict
```

### Go Module Integration

#### Go.mod Management
```go
// go.mod
module github.com/sunday/service-name

go 1.22

require (
    github.com/rakeyshgidwani/sunday-schemas/codegen/go v1.2.3
    // Pin to specific version for stability
)
```

#### Build Validation
```makefile
# Makefile
.PHONY: schema-check
schema-check:
	@echo "Validating schema compatibility..."
	@go list -m github.com/rakeyshgidwani/sunday-schemas/codegen/go
	@go mod verify

.PHONY: build
build: schema-check
	go build ./...
```

## 📊 Monitoring & Observability

### Schema Usage Metrics
```typescript
// middleware/schemaMetrics.ts
import { SundayEvent } from '@sunday-platform/schemas';

export function trackSchemaUsage(event: SundayEvent) {
  // Track which schemas are being used
  metrics.increment('schema.usage', {
    schema: event.schema,
    service: process.env.SERVICE_NAME,
    version: process.env.SCHEMA_VERSION
  });

  // Track deprecated schema usage
  if (isDeprecatedSchema(event.schema)) {
    metrics.increment('schema.deprecated.usage', {
      schema: event.schema,
      service: process.env.SERVICE_NAME
    });

    console.warn(`Using deprecated schema: ${event.schema}`);
  }
}
```

### Health Checks
```go
// internal/health/schemas.go
func (h *HealthChecker) CheckSchemaCompatibility() error {
    // Verify we can parse all expected schema types
    testEvents := map[string][]byte{
        "trade": []byte(`{"schema":"md.trade.v1","instrument_id":"test",...}`),
        "orderbook": []byte(`{"schema":"md.orderbook.delta.v1",...}`),
    }

    for name, event := range testEvents {
        if err := h.eventProcessor.ProcessEvent(event); err != nil {
            return fmt.Errorf("schema health check failed for %s: %w", name, err)
        }
    }

    return nil
}
```

## 🚦 Version Management Strategy

### Semantic Versioning
- **Patch (v1.0.1)**: Bug fixes, documentation updates
- **Minor (v1.1.0)**: New schemas, backward-compatible changes
- **Major (v2.0.0)**: Breaking changes, schema removals

### Deployment Strategy
```bash
# Stage 1: Deploy schema registry
git tag v1.3.0
git push origin v1.3.0  # Triggers package publishing

# Stage 2: Update consuming services (in dependency order)
# 1. Data ingestion services (normalizer)
# 2. Processing services (insights)
# 3. API services (ui-bff)
# 4. Frontend applications

# Stage 3: Monitor and validate
kubectl logs -f deployment/normalizer | grep "schema.version"
```

This deployment architecture ensures the Sunday Schema Registry serves as a reliable foundation for the entire platform, with type safety, version management, and automated validation across all consuming services.

<function_calls>
<invoke name="TodoWrite">
<parameter name="todos">[{"content": "Explain deployment architecture and infrastructure setup", "status": "completed", "activeForm": "Explaining deployment architecture and infrastructure setup"}, {"content": "Document package consumption patterns for TypeScript/JavaScript projects", "status": "completed", "activeForm": "Documenting package consumption patterns for TypeScript/JavaScript projects"}, {"content": "Document Go module consumption and integration", "status": "completed", "activeForm": "Documenting Go module consumption and integration"}, {"content": "Explain CI/CD integration for consuming projects", "status": "completed", "activeForm": "Explaining CI/CD integration for consuming projects"}, {"content": "Document schema registry service deployment", "status": "completed", "activeForm": "Documenting schema registry service deployment"}]