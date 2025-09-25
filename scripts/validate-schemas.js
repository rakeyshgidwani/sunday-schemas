#!/usr/bin/env node
/**
 * Validates JSON Schema files for structural correctness
 * Usage: node scripts/validate-schemas.js
 */

const fs = require('fs');
const path = require('path');
const Ajv = require('ajv');
const addFormats = require('ajv-formats');

// Create AJV instance with 2020-12 support
const ajv = new Ajv({
  strict: false,
  allErrors: true,
  validateFormats: false,
  addUsedSchema: false
});

addFormats(ajv);

const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');

async function validateSchemas() {
  if (!fs.existsSync(SCHEMAS_DIR)) {
    console.log('No schemas directory found');
    return true;
  }

  const schemaFiles = fs.readdirSync(SCHEMAS_DIR).filter(f => f.endsWith('.json'));
  let allValid = true;

  console.log(`🔍 Validating ${schemaFiles.length} schema files...`);

  for (const file of schemaFiles) {
    try {
      const schemaPath = path.join(SCHEMAS_DIR, file);
      const schemaContent = fs.readFileSync(schemaPath, 'utf8');
      const schema = JSON.parse(schemaContent);

      // Basic structural validation
      const requiredFields = ['$id', '$schema', 'title', 'type', 'properties'];
      const missingFields = requiredFields.filter(field => !(field in schema));

      if (missingFields.length > 0) {
        console.error(`❌ ${file}: Missing required fields: ${missingFields.join(', ')}`);
        allValid = false;
        continue;
      }

      // Check that $schema is 2020-12
      if (!schema.$schema.includes('2020-12')) {
        console.warn(`⚠️  ${file}: Not using JSON Schema draft 2020-12`);
      }

      // Check that $id follows our URL pattern
      if (!schema.$id.startsWith('https://schemas.sunday.dev/')) {
        console.error(`❌ ${file}: $id should start with https://schemas.sunday.dev/`);
        allValid = false;
        continue;
      }

      // Validate that venue enums are consistent
      if (schema.properties) {
        const venueFields = ['venue_id', 'long_venue', 'short_venue'].filter(
          field => schema.properties[field] && schema.properties[field].enum
        );

        for (const field of venueFields) {
          const venueEnum = schema.properties[field].enum;
          const expectedVenues = ['polymarket', 'kalshi'];

          if (!venueEnum.every(v => expectedVenues.includes(v))) {
            console.error(`❌ ${file}: ${field} enum contains invalid venues`);
            allValid = false;
          }
        }
      }

      // Check required fields are present
      if (schema.required && Array.isArray(schema.required)) {
        const missingInProperties = schema.required.filter(
          field => !(schema.properties && schema.properties[field])
        );

        if (missingInProperties.length > 0) {
          console.error(`❌ ${file}: Required fields not in properties: ${missingInProperties.join(', ')}`);
          allValid = false;
          continue;
        }
      }

      console.log(`✅ ${file}: Schema structure valid`);

    } catch (error) {
      console.error(`❌ ${file}: ${error.message}`);
      allValid = false;
    }
  }

  return allValid;
}

async function main() {
  const valid = await validateSchemas();

  if (valid) {
    console.log('✅ All schemas structurally valid');
    process.exit(0);
  } else {
    console.log('❌ Some schemas have validation errors');
    process.exit(1);
  }
}

if (require.main === module) {
  main().catch(console.error);
}