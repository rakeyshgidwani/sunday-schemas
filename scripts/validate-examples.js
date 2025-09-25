#!/usr/bin/env node
/**
 * Validates all JSON examples against their corresponding schemas
 * Usage: node scripts/validate-examples.js
 */

const fs = require('fs');
const path = require('path');
const Ajv = require('ajv');

const ajv = new Ajv({ allErrors: true, verbose: true });

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
      const schema = JSON.parse(fs.readFileSync(schemaPath, 'utf8'));
      const schemaId = schema.$id || file.replace('.schema.json', '');
      schemas[schemaId] = schema;
      ajv.addSchema(schema, schemaId);
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

      // Determine schema ID from example content or filename
      const schemaId = example.schema || file.split('.')[0];
      const schema = schemas[schemaId];

      if (!schema) {
        console.warn(`Warning: No schema found for example ${file} (looking for schema: ${schemaId})`);
        continue;
      }

      const validate = ajv.getSchema(schemaId);
      if (!validate) {
        console.error(`Failed to get validator for schema: ${schemaId}`);
        allValid = false;
        continue;
      }

      const valid = validate(example);
      if (valid) {
        console.log(`✅ ${file} validates against ${schemaId}`);
      } else {
        console.error(`❌ ${file} validation failed:`, validate.errors);
        allValid = false;
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