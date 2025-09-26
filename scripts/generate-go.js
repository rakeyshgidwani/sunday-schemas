#!/usr/bin/env node
/**
 * Generate Go types from JSON schemas using quicktype
 * Usage: node scripts/generate-go.js
 */

const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');
const OUTPUT_DIR = path.join(__dirname, '../codegen/go');

function generateGoTypes() {
  if (!fs.existsSync(SCHEMAS_DIR)) {
    console.error('❌ Schemas directory not found:', SCHEMAS_DIR);
    process.exit(1);
  }

  // Ensure output directory exists
  if (!fs.existsSync(OUTPUT_DIR)) {
    fs.mkdirSync(OUTPUT_DIR, { recursive: true });
  }

  console.log('🔄 Generating Go types from JSON schemas...\n');

  const schemaFiles = fs.readdirSync(SCHEMAS_DIR)
    .filter(f => f.endsWith('.schema.json'))
    .sort();

  if (schemaFiles.length === 0) {
    console.warn('⚠️  No schema files found in', SCHEMAS_DIR);
    return;
  }

  // Generate all schemas in a single file to avoid conflicts
  const schemaPaths = schemaFiles.map(f => path.join(SCHEMAS_DIR, f));
  const outputFile = path.join(OUTPUT_DIR, 'schemas.go');

  console.log('📝 Generating all schema types in single file...');

  try {
    // Generate all schemas together to avoid constant conflicts
    const schemaPathsStr = schemaPaths.map(p => `"${p}"`).join(' ');
    const command = `quicktype --src-lang schema --lang go --package sundayschemas --top-level SundaySchemas ${schemaPathsStr} --out "${outputFile}"`;
    execSync(command, { stdio: 'pipe' });

    console.log(`✅ Generated: ${path.relative(process.cwd(), outputFile)}`);
  } catch (error) {
    console.error(`❌ Failed to generate types:`, error.message);
    process.exit(1);
  }

  // Generate a combined constants file
  generateConstantsFile(schemaFiles);

  console.log('\n✨ Go type generation completed successfully');
}

function generateConstantsFile(schemaFiles) {
  console.log('\n📦 Generating constants file...');

  const constantsContent = `// Package sundayschemas provides constants and validation for Sunday platform schemas
//
// This file was automatically generated from JSON Schema definitions.
// DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema files,
// and run 'npm run generate-go' to regenerate this file.
package sundayschemas

import (
	"fmt"
)

// EventSchema represents all valid schema identifiers
type EventSchema string

const (
	${schemaFiles.map(f => {
    const baseName = f.replace('.schema.json', '');
    const constName = baseName.toUpperCase().replace(/\./g, '_');
    return `Schema${constName} EventSchema = "${baseName}"`;
  }).join('\n\t')}
)

// VenueID represents supported venues
type VenueID string

const (
	VenuePolymarket VenueID = "polymarket"
	VenueKalshi     VenueID = "kalshi"
)

// TradeSide represents trade directions
type TradeSide string

const (
	TradeSideBuy  TradeSide = "buy"
	TradeSideSell TradeSide = "sell"
)

// HealthStatus represents venue health states
type HealthStatus string

const (
	HealthConnected HealthStatus = "connected"
	HealthDegraded  HealthStatus = "degraded"
	HealthStale     HealthStatus = "stale"
)

// ValidateSchema checks if a schema string is valid
func ValidateSchema(schema string) error {
	switch EventSchema(schema) {
	case ${schemaFiles.map(f => {
    const baseName = f.replace('.schema.json', '');
    const constName = baseName.toUpperCase().replace(/\./g, '_');
    return `Schema${constName}`;
  }).join(', ')}:
		return nil
	default:
		return fmt.Errorf("invalid schema: %s", schema)
	}
}

// ValidateVenue checks if a venue ID is valid
func ValidateVenue(venue string) error {
	switch VenueID(venue) {
	case VenuePolymarket, VenueKalshi:
		return nil
	default:
		return fmt.Errorf("invalid venue: %s", venue)
	}
}

// AllSchemas returns all valid schema constants
func AllSchemas() []EventSchema {
	return []EventSchema{
		${schemaFiles.map(f => {
    const baseName = f.replace('.schema.json', '');
    const constName = baseName.toUpperCase().replace(/\./g, '_');
    return `Schema${constName}`;
  }).join(',\n\t\t')},
	}
}

// AllVenues returns all valid venue IDs
func AllVenues() []VenueID {
	return []VenueID{VenuePolymarket, VenueKalshi}
}
`;

  const outputPath = path.join(OUTPUT_DIR, 'constants.go');
  fs.writeFileSync(outputPath, constantsContent);
  console.log(`✅ Generated constants file: ${path.relative(process.cwd(), outputPath)}`);
}

function pascalCase(str) {
  return str.split(/[\._\-]/)
    .map(part => part.charAt(0).toUpperCase() + part.slice(1).toLowerCase())
    .join('');
}

if (require.main === module) {
  generateGoTypes();
}

module.exports = { generateGoTypes };