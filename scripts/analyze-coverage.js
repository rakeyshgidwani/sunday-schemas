#!/usr/bin/env node
/**
 * Analyzes example coverage across all schemas
 * Usage: node scripts/analyze-coverage.js
 */

const fs = require('fs');
const path = require('path');

const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');
const EXAMPLES_DIR = path.join(__dirname, '../schemas/examples');

function getSchemas() {
  if (!fs.existsSync(SCHEMAS_DIR)) return [];

  return fs.readdirSync(SCHEMAS_DIR)
    .filter(f => f.endsWith('.json'))
    .map(f => {
      const content = JSON.parse(fs.readFileSync(path.join(SCHEMAS_DIR, f), 'utf8'));
      const schemaId = content.properties?.schema?.const || f.replace('.schema.json', '');
      return {
        file: f,
        schemaId: schemaId,
        path: path.join(SCHEMAS_DIR, f)
      };
    });
}

function getExamples() {
  if (!fs.existsSync(EXAMPLES_DIR)) return [];

  return fs.readdirSync(EXAMPLES_DIR)
    .filter(f => f.endsWith('.json'))
    .map(f => {
      const content = JSON.parse(fs.readFileSync(path.join(EXAMPLES_DIR, f), 'utf8'));
      return {
        file: f,
        schemaId: content.schema,
        path: path.join(EXAMPLES_DIR, f)
      };
    });
}

function analyzeCoverage() {
  const schemas = getSchemas();
  const examples = getExamples();

  console.log('📊 Example Coverage Analysis\n');

  // Group examples by schema
  const examplesBySchema = examples.reduce((acc, example) => {
    if (!acc[example.schemaId]) acc[example.schemaId] = [];
    acc[example.schemaId].push(example);
    return acc;
  }, {});

  let totalCovered = 0;
  let totalSchemas = schemas.length;

  console.log('📋 Schema Coverage Details:\n');

  for (const schema of schemas) {
    const schemaExamples = examplesBySchema[schema.schemaId] || [];
    const isCovered = schemaExamples.length > 0;

    if (isCovered) totalCovered++;

    const status = isCovered ? '✅' : '❌';
    const count = schemaExamples.length;

    console.log(`${status} ${schema.schemaId} (${count} example${count !== 1 ? 's' : ''})`);

    if (schemaExamples.length > 0) {
      schemaExamples.forEach(ex => {
        console.log(`   📄 ${ex.file}`);
      });
    } else {
      console.log(`   ⚠️  No examples found`);
    }
    console.log();
  }

  // Coverage summary
  const coveragePercent = totalSchemas > 0 ? Math.round((totalCovered / totalSchemas) * 100) : 0;
  const totalExamples = examples.length;
  const coverageRatio = totalSchemas > 0 ? (totalExamples / totalSchemas).toFixed(1) : '0.0';

  console.log('📈 Coverage Summary:');
  console.log(`   Total Schemas: ${totalSchemas}`);
  console.log(`   Covered Schemas: ${totalCovered}`);
  console.log(`   Coverage Percentage: ${coveragePercent}%`);
  console.log(`   Total Examples: ${totalExamples}`);
  console.log(`   Coverage Ratio: ${coverageRatio}x`);
  console.log();

  // Coverage quality assessment
  if (coveragePercent === 100) {
    console.log('🎯 Coverage Status: EXCELLENT - All schemas covered');
  } else if (coveragePercent >= 80) {
    console.log('✅ Coverage Status: GOOD - Most schemas covered');
  } else if (coveragePercent >= 60) {
    console.log('⚠️  Coverage Status: FAIR - Needs improvement');
  } else {
    console.log('❌ Coverage Status: POOR - Many schemas uncovered');
  }

  // Find orphaned examples (examples without matching schemas)
  const orphanedExamples = examples.filter(ex =>
    !schemas.find(s => s.schemaId === ex.schemaId)
  );

  if (orphanedExamples.length > 0) {
    console.log('\n🔍 Orphaned Examples (no matching schema):');
    orphanedExamples.forEach(ex => {
      console.log(`   ⚠️  ${ex.file} (schema: ${ex.schemaId})`);
    });
  }

  // Recommendations
  console.log('\n💡 Recommendations:');

  const uncoveredSchemas = schemas.filter(s => !examplesBySchema[s.schemaId]);
  if (uncoveredSchemas.length > 0) {
    console.log('   📝 Add examples for:');
    uncoveredSchemas.forEach(s => {
      console.log(`      - ${s.schemaId}`);
    });
  }

  const multiExampleSchemas = Object.entries(examplesBySchema)
    .filter(([_, examples]) => examples.length > 1)
    .length;

  console.log(`   🎯 ${multiExampleSchemas} schemas have multiple examples (good coverage)`);

  if (totalExamples < totalSchemas) {
    console.log('   📈 Consider adding more examples for better testing');
  }

  return {
    totalSchemas,
    totalCovered,
    coveragePercent,
    totalExamples,
    coverageRatio: parseFloat(coverageRatio)
  };
}

function main() {
  try {
    const results = analyzeCoverage();

    // Exit with error if coverage is below threshold
    if (results.coveragePercent < 100) {
      console.log('\n❌ Coverage below 100% - some schemas lack examples');
      process.exit(1);
    } else {
      console.log('\n✅ All schemas have example coverage');
    }
  } catch (error) {
    console.error('❌ Error analyzing coverage:', error.message);
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}