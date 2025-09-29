# TypeScript Usage

## Quick Start

Install the package:
```bash
npm install @rakeyshgidwani/sunday-schemas
```

Import and use:
```typescript
import { MdOrderbookDeltaV1, RawV0 } from '@rakeyshgidwani/sunday-schemas';

// Use generated types
const event: MdOrderbookDeltaV1 = {
  // Type-safe schema structure
};
```

## Commands

```bash
# Install dependency
npm install @rakeyshgidwani/sunday-schemas

# Install specific version
npm install @rakeyshgidwani/sunday-schemas@1.0.9

# Update to latest
npm update @rakeyshgidwani/sunday-schemas

# Check installed version
npm list @rakeyshgidwani/sunday-schemas
```

## Available Types

- `RawV0` - Raw venue data
- `MdOrderbookDeltaV1` - Orderbook changes
- `MdTradeV1` - Trade events
- `InsightsArbitrageV1` - Arbitrage opportunities
- `InfraHealthV1` - Health metrics

## Version Pinning

Pin to specific version in `package.json`:
```json
{
  "dependencies": {
    "@rakeyshgidwani/sunday-schemas": "1.0.8"
  }
}
```

Check available versions:
```bash
npm view @rakeyshgidwani/sunday-schemas versions --json
```

## Documentation Access for AI Agents

Access bundled documentation:
```bash
# View package info
npm info @rakeyshgidwani/sunday-schemas

# Show package contents
npm list @rakeyshgidwani/sunday-schemas --depth=0

# Locate package files
npm list @rakeyshgidwani/sunday-schemas --parseable
```

Access embedded schema files:
```typescript
import { getSchema, listSchemas } from '@rakeyshgidwani/sunday-schemas';

// Read embedded JSON schema
const schemaContent = getSchema('md.orderbook.delta.v1');

// List all available schemas
const schemaList = listSchemas();
```

## TypeScript Configuration

Add to `tsconfig.json`:
```json
{
  "compilerOptions": {
    "moduleResolution": "node",
    "esModuleInterop": true
  }
}
```