#!/usr/bin/env node
/**
 * Checks for backward compatibility breaking changes
 * Usage: node scripts/check-compatibility.js [--base-ref=main]
 */

const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');
const OPENAPI_FILE = path.join(__dirname, '../openapi/ui.v1.yaml');
const TOPICS_FILE = path.join(__dirname, '../schemas/topics.json');
const VENUES_FILE = path.join(__dirname, '../schemas/registries/venues.json');

function getGitRef(ref = 'main') {
  try {
    execSync(`git show-ref --verify --quiet refs/heads/${ref}`, { stdio: 'ignore' });
    return ref;
  } catch {
    // If main doesn't exist, this might be the first commit
    return null;
  }
}

function getFileAtRef(filePath, ref) {
  if (!ref) return null;

  try {
    const relativePath = path.relative(process.cwd(), filePath);
    const content = execSync(`git show ${ref}:${relativePath}`, { encoding: 'utf8' });
    return content;
  } catch {
    return null; // File didn't exist at that ref
  }
}

function parseJson(content) {
  if (!content) return null;
  try {
    return JSON.parse(content);
  } catch {
    return null;
  }
}

function checkSchemaCompatibility(currentSchema, previousSchema, filename) {
  const issues = [];

  if (!previousSchema) {
    console.log(`✅ ${filename}: New schema (no compatibility check needed)`);
    return [];
  }

  // Check if required fields were added (breaking change)
  const currentRequired = currentSchema.required || [];
  const previousRequired = previousSchema.required || [];
  const addedRequired = currentRequired.filter(field => !previousRequired.includes(field));

  if (addedRequired.length > 0) {
    issues.push(`❌ ${filename}: Added required fields: ${addedRequired.join(', ')}`);
  }

  // Check if required fields were removed (breaking change)
  const removedRequired = previousRequired.filter(field => !currentRequired.includes(field));
  if (removedRequired.length > 0) {
    issues.push(`❌ ${filename}: Removed required fields: ${removedRequired.join(', ')}`);
  }

  // Check if properties were removed (breaking change)
  const currentProps = Object.keys(currentSchema.properties || {});
  const previousProps = Object.keys(previousSchema.properties || {});
  const removedProps = previousProps.filter(prop => !currentProps.includes(prop));

  if (removedProps.length > 0) {
    issues.push(`❌ ${filename}: Removed properties: ${removedProps.join(', ')}`);
  }

  // Check enum narrowing (breaking change)
  for (const prop of currentProps) {
    const currentProp = currentSchema.properties[prop];
    const previousProp = previousSchema.properties?.[prop];

    if (currentProp?.enum && previousProp?.enum) {
      const removedValues = previousProp.enum.filter(val => !currentProp.enum.includes(val));
      if (removedValues.length > 0) {
        issues.push(`❌ ${filename}.${prop}: Removed enum values: ${removedValues.join(', ')}`);
      }
    }
  }

  if (issues.length === 0) {
    console.log(`✅ ${filename}: Backward compatible`);
  }

  return issues;
}

function checkTopicsCompatibility(currentTopics, previousTopics) {
  const issues = [];

  if (!previousTopics) {
    console.log(`✅ topics.json: New file (no compatibility check needed)`);
    return [];
  }

  // Check if schema->topic mappings changed (breaking change)
  for (const [schemaId, currentMapping] of Object.entries(currentTopics)) {
    const previousMapping = previousTopics[schemaId];

    if (previousMapping) {
      const currentTopic = currentMapping.topic || (Array.isArray(currentMapping.topics) ? currentMapping.topics[0] : null);
      const previousTopic = previousMapping.topic || (Array.isArray(previousMapping.topics) ? previousMapping.topics[0] : null);

      if (currentTopic !== previousTopic) {
        issues.push(`❌ topics.json: Schema ${schemaId} topic changed from ${previousTopic} to ${currentTopic}`);
      }
    }
  }

  // Check if schema mappings were removed (breaking change)
  const removedSchemas = Object.keys(previousTopics).filter(schemaId => !currentTopics[schemaId]);
  if (removedSchemas.length > 0) {
    issues.push(`❌ topics.json: Removed schema mappings: ${removedSchemas.join(', ')}`);
  }

  if (issues.length === 0) {
    console.log(`✅ topics.json: Backward compatible`);
  }

  return issues;
}

function checkVenuesCompatibility(currentVenues, previousVenues) {
  const issues = [];

  if (!previousVenues) {
    console.log(`✅ venues.json: New file (no compatibility check needed)`);
    return [];
  }

  // Check if venues were removed (breaking change in Phase 1)
  const removedVenues = previousVenues.filter(venue => !currentVenues.includes(venue));
  if (removedVenues.length > 0) {
    issues.push(`❌ venues.json: Removed venues: ${removedVenues.join(', ')} (breaking in Phase 1)`);
  }

  // Adding venues is OK (minor version bump)
  const addedVenues = currentVenues.filter(venue => !previousVenues.includes(venue));
  if (addedVenues.length > 0) {
    console.log(`✅ venues.json: Added venues: ${addedVenues.join(', ')} (minor version bump)`);
  }

  if (issues.length === 0 && addedVenues.length === 0) {
    console.log(`✅ venues.json: No changes`);
  }

  return issues;
}

async function main() {
  const args = process.argv.slice(2);
  const baseRef = args.find(arg => arg.startsWith('--base-ref='))?.split('=')[1] || 'main';

  console.log('🔍 Checking backward compatibility...');

  const gitRef = getGitRef(baseRef);
  if (!gitRef) {
    console.log(`⚠️  No base reference found (${baseRef}), assuming first commit - skipping compatibility check`);
    return;
  }

  let allIssues = [];

  // Check JSON schemas
  if (fs.existsSync(SCHEMAS_DIR)) {
    const schemaFiles = fs.readdirSync(SCHEMAS_DIR).filter(f => f.endsWith('.json'));

    for (const file of schemaFiles) {
      const filePath = path.join(SCHEMAS_DIR, file);
      const currentContent = fs.readFileSync(filePath, 'utf8');
      const previousContent = getFileAtRef(filePath, gitRef);

      const currentSchema = parseJson(currentContent);
      const previousSchema = parseJson(previousContent);

      if (currentSchema) {
        const issues = checkSchemaCompatibility(currentSchema, previousSchema, file);
        allIssues.push(...issues);
      }
    }
  }

  // Check topics.json
  if (fs.existsSync(TOPICS_FILE)) {
    const currentContent = fs.readFileSync(TOPICS_FILE, 'utf8');
    const previousContent = getFileAtRef(TOPICS_FILE, gitRef);

    const currentTopics = parseJson(currentContent);
    const previousTopics = parseJson(previousContent);

    if (currentTopics) {
      const issues = checkTopicsCompatibility(currentTopics, previousTopics);
      allIssues.push(...issues);
    }
  }

  // Check venues.json
  if (fs.existsSync(VENUES_FILE)) {
    const currentContent = fs.readFileSync(VENUES_FILE, 'utf8');
    const previousContent = getFileAtRef(VENUES_FILE, gitRef);

    const currentVenues = parseJson(currentContent);
    const previousVenues = parseJson(previousContent);

    if (currentVenues) {
      const issues = checkVenuesCompatibility(currentVenues, previousVenues);
      allIssues.push(...issues);
    }
  }

  // TODO: Add OpenAPI compatibility checking with openapi-diff
  // For now, just log that OpenAPI checking is pending
  console.log('⚠️  OpenAPI compatibility checking with openapi-diff pending implementation');

  if (allIssues.length > 0) {
    console.log('\n💥 Compatibility Issues Found:');
    allIssues.forEach(issue => console.log(issue));
    console.log('\n❌ Compatibility check failed');
    process.exit(1);
  } else {
    console.log('\n✅ All compatibility checks passed');
  }
}

if (require.main === module) {
  main().catch(console.error);
}