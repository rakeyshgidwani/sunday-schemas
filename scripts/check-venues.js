#!/usr/bin/env node
/**
 * Validates that all schemas use venue enums that match venues.json registry
 * Usage: node scripts/check-venues.js
 */

const fs = require('fs');
const path = require('path');

const VENUES_REGISTRY = path.join(__dirname, '../schemas/registries/venues.json');
const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');

function loadVenues() {
  if (!fs.existsSync(VENUES_REGISTRY)) {
    console.error('❌ venues.json registry not found');
    return null;
  }

  try {
    return JSON.parse(fs.readFileSync(VENUES_REGISTRY, 'utf8'));
  } catch (error) {
    console.error('❌ Failed to parse venues.json:', error.message);
    return null;
  }
}

function checkSchemaVenues(schemas, validVenues) {
  let allValid = true;

  for (const [schemaId, schema] of Object.entries(schemas)) {
    // Check venue_id enums in schema properties
    if (schema.properties && schema.properties.venue_id) {
      const venueEnum = schema.properties.venue_id.enum;
      if (venueEnum) {
        const invalidVenues = venueEnum.filter(v => !validVenues.includes(v));
        if (invalidVenues.length > 0) {
          console.error(`❌ Schema ${schemaId} has invalid venues: ${invalidVenues.join(', ')}`);
          allValid = false;
        } else {
          console.log(`✅ Schema ${schemaId} venues valid: ${venueEnum.join(', ')}`);
        }
      }
    }

    // Check other venue-related fields (long_venue, short_venue, etc.)
    ['long_venue', 'short_venue'].forEach(field => {
      if (schema.properties && schema.properties[field] && schema.properties[field].enum) {
        const venueEnum = schema.properties[field].enum;
        const invalidVenues = venueEnum.filter(v => !validVenues.includes(v));
        if (invalidVenues.length > 0) {
          console.error(`❌ Schema ${schemaId} field ${field} has invalid venues: ${invalidVenues.join(', ')}`);
          allValid = false;
        }
      }
    });
  }

  return allValid;
}

function loadSchemas() {
  if (!fs.existsSync(SCHEMAS_DIR)) {
    console.log('No schemas directory found yet - skipping venue validation');
    return {};
  }

  const schemas = {};
  const schemaFiles = fs.readdirSync(SCHEMAS_DIR).filter(f => f.endsWith('.json'));

  for (const file of schemaFiles) {
    try {
      const schemaPath = path.join(SCHEMAS_DIR, file);
      const schema = JSON.parse(fs.readFileSync(schemaPath, 'utf8'));
      const schemaId = schema.$id || file.replace('.schema.json', '');
      schemas[schemaId] = schema;
    } catch (error) {
      console.error(`Failed to load schema ${file}:`, error.message);
      process.exit(1);
    }
  }

  return schemas;
}

function main() {
  console.log('🏢 Loading venue registry...');
  const validVenues = loadVenues();
  if (!validVenues) {
    process.exit(1);
  }
  console.log(`Valid venues: ${validVenues.join(', ')}`);

  console.log('📋 Loading schemas...');
  const schemas = loadSchemas();
  console.log(`Found ${Object.keys(schemas).length} schemas`);

  if (Object.keys(schemas).length === 0) {
    console.log('✅ No schemas to validate yet - venue validation will be implemented in Phase 2');
    process.exit(0);
  }

  console.log('🔍 Checking venue consistency...');
  const allValid = checkSchemaVenues(schemas, validVenues);

  if (allValid) {
    console.log('✅ All schema venues are valid');
    process.exit(0);
  } else {
    console.log('❌ Some schemas have invalid venue references');
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}