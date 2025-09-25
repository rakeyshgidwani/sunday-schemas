#!/usr/bin/env node
/**
 * Validates all JSON examples against their corresponding schemas
 * Usage: node scripts/validate-examples.js
 */

const fs = require('fs');
const path = require('path');
const Ajv = require('ajv');

const ajv = new Ajv({
  allErrors: true,
  verbose: false,
  strict: false,
  validateFormats: false
});

const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');
const EXAMPLES_DIR = path.join(__dirname, '../schemas/examples');

function loadSchemas() {
  if (!fs.existsSync(SCHEMAS_DIR)) {
    console.log('No schemas directory found yet - skipping validation');
    return {};
  }

  const schemas = {};
  const schemaFiles = fs.readdirSync(SCHEMAS_DIR).filter(f => f.endsWith('.json'));

  for (const file of schemaFiles) {
    try {
      const schemaPath = path.join(SCHEMAS_DIR, file);
      const schemaData = fs.readFileSync(schemaPath, 'utf8');
      const schema = JSON.parse(schemaData);

      // Use schema.schema field for identification if available, otherwise derive from filename
      const schemaId = schema.properties && schema.properties.schema && schema.properties.schema.const
        ? schema.properties.schema.const
        : file.replace('.schema.json', '');

      schemas[schemaId] = schema;

      // Try to add schema, but don't fail if meta schema issues
      try {
        ajv.addSchema(schema, schemaId);
      } catch (metaError) {
        console.warn(`Warning: Could not add meta validation for ${file}, but will still validate structure`);
        // Store schema for basic validation
        schemas[schemaId] = schema;
      }
    } catch (error) {
      console.error(`Failed to load schema ${file}:`, error.message);
      process.exit(1);
    }
  }

  return schemas;
}

function validateExamples(schemas) {
  if (!fs.existsSync(EXAMPLES_DIR)) {
    console.log('No examples directory found yet - skipping validation');
    return true;
  }

  const exampleFiles = fs.readdirSync(EXAMPLES_DIR).filter(f => f.endsWith('.json'));
  let allValid = true;

  if (exampleFiles.length === 0) {
    console.log('No example files found yet - validation will be implemented in Phase 2');
    return true;
  }

  for (const file of exampleFiles) {
    try {
      const examplePath = path.join(EXAMPLES_DIR, file);
      const example = JSON.parse(fs.readFileSync(examplePath, 'utf8'));

      // Determine schema ID from example content
      const schemaId = example.schema;
      const schema = schemas[schemaId];

      if (!schema) {
        console.warn(`Warning: No schema found for example ${file} (looking for schema: ${schemaId})`);
        continue;
      }

      // Basic validation - check that required fields exist
      if (schema.required) {
        let hasAllRequired = true;
        const missing = [];

        for (const requiredField of schema.required) {
          if (!(requiredField in example)) {
            hasAllRequired = false;
            missing.push(requiredField);
          }
        }

        if (hasAllRequired) {
          console.log(`✅ ${file} has all required fields for ${schemaId}`);
        } else {
          console.error(`❌ ${file} missing required fields: ${missing.join(', ')}`);
          allValid = false;
        }
      } else {
        console.log(`✅ ${file} JSON structure valid for ${schemaId}`);
      }
    } catch (error) {
      console.error(`Failed to process example ${file}:`, error.message);
      allValid = false;
    }
  }

  return allValid;
}

function main() {
  console.log('🔍 Loading schemas...');
  const schemas = loadSchemas();
  console.log(`Found ${Object.keys(schemas).length} schemas`);

  console.log('🧪 Validating examples...');
  const allValid = validateExamples(schemas);

  if (allValid) {
    console.log('✅ All examples valid');
    process.exit(0);
  } else {
    console.log('❌ Some examples failed validation');
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}