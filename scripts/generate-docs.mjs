#!/usr/bin/env node

/**
 * Sunday Schemas - Automated Documentation Generator
 *
 * Generates usage documentation by introspecting actual generated code
 * to ensure docs are always accurate and up-to-date.
 */

import fs from 'fs/promises';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const rootDir = path.resolve(__dirname, '..');

// Configuration
const DOCS_DIR = path.join(rootDir, 'docs');
const CODEGEN_TS_DIR = path.join(rootDir, 'codegen', 'ts');
const CODEGEN_GO_DIR = path.join(rootDir, 'codegen', 'go');

/**
 * Introspect TypeScript codegen to extract types and constants
 */
async function introspectTypeScript() {
    const indexPath = path.join(CODEGEN_TS_DIR, 'index.ts');
    const content = await fs.readFile(indexPath, 'utf8');

    // Extract type exports
    const typeExports = [];
    const aliasExports = [];
    const constantExports = [];

    // Parse export statements
    const exportRegex = /export type \{ (.+?) \}/g;
    const aliasRegex = /export type \{ (.+?) as (.+?) \}/g;
    const unionRegex = /export type (\w+) =/g;
    const constRegex = /export const (\w+):/g;

    let match;

    // Extract type exports with aliases
    while ((match = aliasRegex.exec(content)) !== null) {
        const [, fullName, alias] = match;
        aliasExports.push({ fullName: fullName.trim(), alias: alias.trim() });
    }

    // Extract regular type exports
    while ((match = exportRegex.exec(content)) !== null) {
        const types = match[1].split(',').map(t => t.trim());
        typeExports.push(...types);
    }

    // Extract union types
    while ((match = unionRegex.exec(content)) !== null) {
        typeExports.push(match[1]);
    }

    // Extract constants
    while ((match = constRegex.exec(content)) !== null) {
        constantExports.push(match[1]);
    }

    // Extract schema constants from content
    const schemaConstants = [];
    const venueConstants = [];

    if (content.includes("'raw.v0'")) {
        schemaConstants.push("'raw.v0'");
    }
    if (content.includes("'md.orderbook.delta.v1'")) {
        schemaConstants.push("'md.orderbook.delta.v1'");
    }
    if (content.includes("'md.trade.v1'")) {
        schemaConstants.push("'md.trade.v1'");
    }
    // Add more as needed...

    if (content.includes("'polymarket'")) {
        venueConstants.push("'polymarket'");
    }
    if (content.includes("'kalshi'")) {
        venueConstants.push("'kalshi'");
    }

    return {
        types: typeExports,
        aliases: aliasExports,
        constants: constantExports,
        schemaConstants,
        venueConstants
    };
}

/**
 * Introspect Go codegen to extract types and functions
 */
async function introspectGo() {
    const constantsPath = path.join(CODEGEN_GO_DIR, 'constants.go');
    const schemasPath = path.join(CODEGEN_GO_DIR, 'schemas.go');

    const constantsContent = await fs.readFile(constantsPath, 'utf8');
    const schemasContent = await fs.readFile(schemasPath, 'utf8');

    // Extract exported functions
    const functions = [];
    const funcRegex = /^func (\w+)\(/gm;
    let match;

    while ((match = funcRegex.exec(constantsContent)) !== null) {
        functions.push(match[1]);
    }
    while ((match = funcRegex.exec(schemasContent)) !== null) {
        functions.push(match[1]);
    }

    // Extract constants
    const constants = [];
    const constRegex = /^\s*(\w+)\s+\w+\s*=\s*"(.+?)"/gm;

    while ((match = constRegex.exec(constantsContent)) !== null) {
        constants.push({ name: match[1], value: match[2] });
    }

    // Extract type definitions
    const types = [];
    const typeRegex = /^type (\w+) struct/gm;

    while ((match = typeRegex.exec(schemasContent)) !== null) {
        types.push(match[1]);
    }

    return { functions, constants, types };
}

/**
 * Generate TypeScript usage documentation
 */
async function generateTypeScriptDocs() {
    const tsData = await introspectTypeScript();

    const packageName = '@rakeyshgidwani/sunday-schemas';

    return `# TypeScript Usage

## Quick Start

Install the package:
\`\`\`bash
npm install ${packageName}
\`\`\`

Import and use:
\`\`\`typescript
import {
  RawEnvelopeV0,
  NormalizedTradeV1,
  SCHEMA_CONSTANTS,
  VENUE_IDS
} from '${packageName}';

// Use generated types with full type safety
const envelope: RawEnvelopeV0 = {
  schema: 'raw.v0',
  venue_id: 'polymarket',
  stream: 'orderbook',
  instrument_native: 'some-instrument',
  partition_key: 'polymarket:some-instrument',
  ts_event_ms: Date.now(),
  ts_ingest_ms: Date.now(),
  payload: {}
};
\`\`\`

## Commands

\`\`\`bash
# Add dependency
npm install ${packageName}@latest

# Update to latest
npm update ${packageName}

# Specific version
npm install ${packageName}@1.0.9

# Check installed version
npm list ${packageName}
\`\`\`

## Available Types and Constants

**Event Types (with aliases):**
\`\`\`typescript
${tsData.aliases.map(a => `${a.fullName.padEnd(35)} // Alias: ${a.alias}`).join('\n')}
\`\`\`

**Union Types:**
\`\`\`typescript
${tsData.types.filter(t => ['SundayEvent', 'EventSchema', 'EventBySchema'].includes(t)).join('\n')}
\`\`\`

**Runtime Constants:**
\`\`\`typescript
${tsData.constants.map(c => `${c.padEnd(30)} // Available at runtime`).join('\n')}
\`\`\`

## Type-Safe Event Handling

**Runtime Schema Validation:**
\`\`\`typescript
import { SCHEMA_CONSTANTS, VENUE_IDS, EventSchema, VenueId } from '${packageName}';

// Validate schema at runtime
function isValidSchema(schema: string): schema is EventSchema {
  return schema in SCHEMA_CONSTANTS;
}

// Validate venue at runtime
function isValidVenue(venue: string): venue is VenueId {
  return VENUE_IDS.includes(venue as VenueId);
}
\`\`\`

**Generic Type-Safe Event Processing:**
\`\`\`typescript
import { EventBySchema, SundayEvent } from '${packageName}';

// Generic event processor with full type safety
function processEvent<T extends EventSchema>(
  schema: T,
  data: EventBySchema<T>
): void {
  // TypeScript knows the exact type of data based on schema
  console.log(\`Processing event: \${schema}\`);
}
\`\`\`

## TypeScript Configuration

Add to \`tsconfig.json\`:
\`\`\`json
{
  "compilerOptions": {
    "moduleResolution": "node",
    "esModuleInterop": true,
    "strict": true,
    "exactOptionalPropertyTypes": true
  }
}
\`\`\`

---

*This documentation is automatically generated from the actual TypeScript codegen output.*
*Last updated: ${new Date().toISOString()}*
`;
}

/**
 * Generate Go usage documentation
 */
async function generateGoDocs() {
    const goData = await introspectGo();

    const modulePath = 'github.com/rakeyshgidwani/sunday-schemas/codegen/go';

    return `# Go Module Usage

## Quick Start

Add to your \`go.mod\`:
\`\`\`go
require ${modulePath} v1.0.9
\`\`\`

Import and use:
\`\`\`go
import schemas "${modulePath}"

// Use generated types
var event schemas.RawEnvelopeV0
var trade schemas.NormalizedTradeV1
\`\`\`

## Commands

\`\`\`bash
# Add dependency
go get ${modulePath}@latest

# Update to latest
go get -u ${modulePath}

# Specific version
go get ${modulePath}@v1.0.9

# Download dependencies
go mod download

# Clean up
go mod tidy
\`\`\`

## Available Types and Constants

**Event Types:**
\`\`\`go
${goData.types.map(t => `schemas.${t.padEnd(35)} // Generated event type`).join('\n')}
\`\`\`

**Schema Constants:**
\`\`\`go
${goData.constants.filter(c => c.name.startsWith('Schema')).map(c => `schemas.${c.name.padEnd(30)} // "${c.value}"`).join('\n')}
\`\`\`

**Venue Constants:**
\`\`\`go
${goData.constants.filter(c => c.name.startsWith('Venue')).map(c => `schemas.${c.name.padEnd(30)} // "${c.value}"`).join('\n')}
\`\`\`

## Available Functions

**Schema and Venue Validation:**
\`\`\`go
${goData.functions.filter(f => f.startsWith('Validate') || f.startsWith('All')).map(f => `schemas.${f}()`).join('\n')}
\`\`\`

**JSON Marshal/Unmarshal Functions:**
\`\`\`go
${goData.functions.filter(f => f.startsWith('Unmarshal')).map(f => `schemas.${f}(data)`).join('\n')}
\`\`\`

**Helper Functions:**
\`\`\`go
${goData.functions.filter(f => !f.startsWith('Validate') && !f.startsWith('All') && !f.startsWith('Unmarshal')).map(f => `schemas.${f}()`).join('\n')}
\`\`\`

## Complete Example

\`\`\`go
package main

import (
    "fmt"
    "log"
    schemas "${modulePath}"
)

func main() {
    // Validate schema
    if err := schemas.ValidateSchema("raw.v0"); err != nil {
        log.Fatal("Invalid schema:", err)
    }

    // Get all available schemas
    allSchemas := schemas.AllSchemas()
    fmt.Printf("Available schemas: %v\\n", allSchemas)

    // Unmarshal JSON data
    data := []byte(\`{"schema":"md.trade.v1","venue_id":"polymarket"...}\`)
    trade, err := schemas.UnmarshalNormalizedTradeV1(data)
    if err != nil {
        log.Fatal("Failed to unmarshal:", err)
    }

    fmt.Printf("Trade: %+v\\n", trade)
}
\`\`\`

---

*This documentation is automatically generated from the actual Go codegen output.*
*Last updated: ${new Date().toISOString()}*
`;
}

/**
 * Main execution
 */
async function main() {
    console.log('🤖 Generating documentation from codegen artifacts...');

    try {
        // Generate TypeScript documentation
        console.log('📝 Generating TypeScript documentation...');
        const tsDocs = await generateTypeScriptDocs();
        await fs.writeFile(path.join(DOCS_DIR, 'TYPESCRIPT_USAGE.md'), tsDocs);
        console.log('✅ Generated docs/TYPESCRIPT_USAGE.md');

        // Generate Go documentation
        console.log('📝 Generating Go documentation...');
        const goDocs = await generateGoDocs();
        await fs.writeFile(path.join(DOCS_DIR, 'GO_MODULE_USAGE.md'), goDocs);
        console.log('✅ Generated docs/GO_MODULE_USAGE.md');

        console.log('');
        console.log('🎉 Documentation generation complete!');
        console.log('');
        console.log('📋 Next steps:');
        console.log('  1. Review generated documentation');
        console.log('  2. Add this script to your CI/CD pipeline');
        console.log('  3. Run after each codegen to keep docs in sync');
        console.log('');

    } catch (error) {
        console.error('❌ Documentation generation failed:', error.message);
        process.exit(1);
    }
}

if (import.meta.url === `file://${process.argv[1]}`) {
    main();
}