#!/usr/bin/env node
/**
 * Check for deprecated schemas and validate deprecation metadata
 * Usage: node scripts/check-deprecations.js [--format=json|table]
 */

const fs = require('fs');
const path = require('path');

const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');

function checkDeprecations(outputFormat = 'table') {
  if (!fs.existsSync(SCHEMAS_DIR)) {
    console.error('❌ Schemas directory not found:', SCHEMAS_DIR);
    process.exit(1);
  }

  const schemaFiles = fs.readdirSync(SCHEMAS_DIR)
    .filter(f => f.endsWith('.schema.json'))
    .sort();

  console.log('🔍 Checking schemas for deprecation status...\n');

  const results = {
    total: schemaFiles.length,
    deprecated: 0,
    active: 0,
    warnings: [],
    errors: [],
    deprecations: []
  };

  for (const schemaFile of schemaFiles) {
    const schemaPath = path.join(SCHEMAS_DIR, schemaFile);
    const baseName = schemaFile.replace('.schema.json', '');

    try {
      const content = JSON.parse(fs.readFileSync(schemaPath, 'utf8'));
      const deprecationInfo = content['x-deprecated'];

      if (deprecationInfo && deprecationInfo.deprecated === true) {
        results.deprecated++;
        const deprecation = analyzeDeprecation(baseName, deprecationInfo, content);
        results.deprecations.push(deprecation);

        // Validate deprecation metadata
        const validation = validateDeprecationMetadata(deprecation);
        results.warnings.push(...validation.warnings);
        results.errors.push(...validation.errors);

        if (outputFormat === 'table') {
          console.log(`⚠️  ${baseName} - DEPRECATED`);
          console.log(`   Since: ${deprecation.deprecatedInVersion || 'Unknown'}`);
          console.log(`   Removal: ${deprecation.removalPlannedInVersion || 'Not specified'}`);
          console.log(`   Reason: ${deprecation.reason || 'No reason specified'}`);
          if (deprecation.replacedBy) {
            console.log(`   Replaced by: ${deprecation.replacedBy}`);
          }
          if (deprecation.migrationGuide) {
            console.log(`   Migration guide: ${deprecation.migrationGuide}`);
          }
          console.log();
        }
      } else {
        results.active++;
        if (outputFormat === 'table') {
          console.log(`✅ ${baseName} - Active`);
        }
      }
    } catch (error) {
      results.errors.push(`Failed to parse ${schemaFile}: ${error.message}`);
      if (outputFormat === 'table') {
        console.error(`❌ ${baseName} - Parse error: ${error.message}`);
      }
    }
  }

  if (outputFormat === 'json') {
    console.log(JSON.stringify(results, null, 2));
    return results;
  }

  // Table format summary
  console.log('\n📊 Deprecation Summary:');
  console.log(`   Total schemas: ${results.total}`);
  console.log(`   Active: ${results.active}`);
  console.log(`   Deprecated: ${results.deprecated}`);
  console.log();

  if (results.warnings.length > 0) {
    console.log('⚠️  Deprecation Warnings:');
    results.warnings.forEach(warning => console.log(`   ${warning}`));
    console.log();
  }

  if (results.errors.length > 0) {
    console.log('❌ Deprecation Errors:');
    results.errors.forEach(error => console.log(`   ${error}`));
    console.log();
  }

  // Check for urgent deprecations
  const urgentDeprecations = results.deprecations.filter(d => isUrgentDeprecation(d));
  if (urgentDeprecations.length > 0) {
    console.log('🚨 Urgent Deprecations (removal imminent):');
    urgentDeprecations.forEach(d => {
      console.log(`   ${d.schema} - removal planned: ${d.plannedRemovalDate || d.removalPlannedInVersion}`);
    });
    console.log();
  }

  // Generate migration recommendations
  if (results.deprecated > 0) {
    console.log('💡 Recommendations:');
    results.deprecations.forEach(d => {
      if (!d.migrationGuide) {
        console.log(`   📝 Create migration guide for ${d.schema}`);
      }
      if (!d.replacedBy) {
        console.log(`   🔄 Specify replacement schema for ${d.schema}`);
      }
      if (!d.plannedRemovalDate && !d.removalPlannedInVersion) {
        console.log(`   📅 Set removal timeline for ${d.schema}`);
      }
    });
  }

  return results;
}

function analyzeDeprecation(schema, deprecationInfo, schemaContent) {
  return {
    schema,
    deprecated: true,
    deprecatedInVersion: deprecationInfo.deprecatedInVersion,
    removalPlannedInVersion: deprecationInfo.removalPlannedInVersion,
    reason: deprecationInfo.reason,
    replacedBy: deprecationInfo.replacedBy,
    migrationGuide: deprecationInfo.migrationGuide,
    deprecationDate: deprecationInfo.deprecationDate,
    plannedRemovalDate: deprecationInfo.plannedRemovalDate,
    contact: deprecationInfo.contact,
    urgency: deprecationInfo.urgency || 'medium',
    impactLevel: deprecationInfo.impactLevel || 'medium',
    description: schemaContent.description || '',
    hasWarningInDescription: schemaContent.description?.includes('DEPRECATED') || false
  };
}

function validateDeprecationMetadata(deprecation) {
  const warnings = [];
  const errors = [];

  // Required fields
  if (!deprecation.reason) {
    errors.push(`${deprecation.schema}: Missing required field 'reason' in x-deprecated`);
  }

  if (!deprecation.deprecatedInVersion) {
    warnings.push(`${deprecation.schema}: Missing 'deprecatedInVersion' - when was this deprecated?`);
  }

  // Recommended fields
  if (!deprecation.replacedBy) {
    warnings.push(`${deprecation.schema}: No 'replacedBy' specified - what should users migrate to?`);
  }

  if (!deprecation.migrationGuide) {
    warnings.push(`${deprecation.schema}: No 'migrationGuide' - consider adding migration documentation`);
  }

  if (!deprecation.plannedRemovalDate && !deprecation.removalPlannedInVersion) {
    warnings.push(`${deprecation.schema}: No removal timeline specified - when will this be removed?`);
  }

  // Description validation
  if (!deprecation.hasWarningInDescription) {
    warnings.push(`${deprecation.schema}: Schema description should include deprecation warning`);
  }

  // Date validation
  if (deprecation.deprecationDate) {
    try {
      const depDate = new Date(deprecation.deprecationDate);
      if (isNaN(depDate.getTime())) {
        errors.push(`${deprecation.schema}: Invalid deprecationDate format`);
      }
    } catch (e) {
      errors.push(`${deprecation.schema}: Invalid deprecationDate: ${e.message}`);
    }
  }

  if (deprecation.plannedRemovalDate) {
    try {
      const removalDate = new Date(deprecation.plannedRemovalDate);
      if (isNaN(removalDate.getTime())) {
        errors.push(`${deprecation.schema}: Invalid plannedRemovalDate format`);
      } else {
        // Check if removal date is in the past
        if (removalDate < new Date()) {
          errors.push(`${deprecation.schema}: Planned removal date is in the past - remove schema or update date`);
        }
      }
    } catch (e) {
      errors.push(`${deprecation.schema}: Invalid plannedRemovalDate: ${e.message}`);
    }
  }

  return { warnings, errors };
}

function isUrgentDeprecation(deprecation) {
  if (deprecation.urgency === 'high') return true;

  // Check if removal date is within 30 days
  if (deprecation.plannedRemovalDate) {
    const removalDate = new Date(deprecation.plannedRemovalDate);
    const now = new Date();
    const thirtyDays = 30 * 24 * 60 * 60 * 1000; // 30 days in milliseconds
    return (removalDate - now) < thirtyDays;
  }

  return false;
}

function generateDeprecationReport() {
  console.log('📊 Generating Deprecation Report...\n');

  const results = checkDeprecations('json');

  const report = {
    reportDate: new Date().toISOString(),
    summary: {
      totalSchemas: results.total,
      activeSchemas: results.active,
      deprecatedSchemas: results.deprecated,
      urgentDeprecations: results.deprecations.filter(isUrgentDeprecation).length
    },
    deprecations: results.deprecations.map(d => ({
      schema: d.schema,
      status: 'deprecated',
      deprecatedInVersion: d.deprecatedInVersion,
      plannedRemoval: d.removalPlannedInVersion || d.plannedRemovalDate,
      urgency: d.urgency,
      replacedBy: d.replacedBy,
      migrationGuide: d.migrationGuide,
      needsAttention: !d.migrationGuide || !d.replacedBy || isUrgentDeprecation(d)
    })),
    recommendations: generateRecommendations(results.deprecations),
    validationIssues: {
      warnings: results.warnings.length,
      errors: results.errors.length
    }
  };

  const reportPath = path.join(__dirname, '../deprecation-report.json');
  fs.writeFileSync(reportPath, JSON.stringify(report, null, 2));

  console.log(`Report saved to: ${reportPath}`);
  return report;
}

function generateRecommendations(deprecations) {
  const recommendations = [];

  deprecations.forEach(d => {
    if (!d.migrationGuide) {
      recommendations.push({
        type: 'documentation',
        priority: 'medium',
        schema: d.schema,
        action: 'Create migration guide',
        description: `No migration guide available for ${d.schema}`
      });
    }

    if (!d.replacedBy) {
      recommendations.push({
        type: 'clarification',
        priority: 'high',
        schema: d.schema,
        action: 'Specify replacement',
        description: `No replacement schema specified for ${d.schema}`
      });
    }

    if (isUrgentDeprecation(d)) {
      recommendations.push({
        type: 'urgent',
        priority: 'critical',
        schema: d.schema,
        action: 'Immediate attention required',
        description: `${d.schema} removal is imminent - ensure migration is complete`
      });
    }
  });

  return recommendations;
}

function main() {
  const args = process.argv.slice(2);
  const formatArg = args.find(arg => arg.startsWith('--format='));
  const format = formatArg ? formatArg.split('=')[1] : 'table';

  const reportArg = args.includes('--report');

  try {
    if (reportArg) {
      generateDeprecationReport();
    } else {
      const results = checkDeprecations(format);

      // Exit with error if there are validation errors
      if (results.errors.length > 0) {
        process.exit(1);
      }
    }
  } catch (error) {
    console.error('❌ Error checking deprecations:', error.message);
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}

module.exports = { checkDeprecations, generateDeprecationReport };