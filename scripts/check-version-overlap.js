#!/usr/bin/env node
/**
 * Check version overlap requirements for deprecated schemas
 * Usage: node scripts/check-version-overlap.js
 */

const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const SCHEMAS_DIR = path.join(__dirname, '../schemas/json');

function parseVersion(versionString) {
  if (!versionString) return null;

  // Handle both v1.2.3 and 1.2.3 formats
  const cleanVersion = versionString.replace(/^v/, '');
  const parts = cleanVersion.split('.').map(n => parseInt(n, 10));

  if (parts.length !== 3 || parts.some(isNaN)) {
    return null;
  }

  return {
    major: parts[0],
    minor: parts[1],
    patch: parts[2],
    original: versionString,
    toString: () => `${parts[0]}.${parts[1]}.${parts[2]}`
  };
}

function compareVersions(v1, v2) {
  if (!v1 || !v2) return null;

  if (v1.major !== v2.major) return v1.major - v2.major;
  if (v1.minor !== v2.minor) return v1.minor - v2.minor;
  return v1.patch - v2.patch;
}

function getMinorVersionDifference(deprecatedVersion, removalVersion) {
  if (!deprecatedVersion || !removalVersion) return null;

  const dep = parseVersion(deprecatedVersion);
  const rem = parseVersion(removalVersion);

  if (!dep || !rem) return null;

  // For same major version, calculate minor version difference
  if (dep.major === rem.major) {
    return rem.minor - dep.minor;
  }

  // For different major versions, consider it as sufficient overlap
  if (rem.major > dep.major) {
    return 10; // Arbitrary large number indicating major version change
  }

  return null;
}

function getCurrentVersion() {
  try {
    const packageJson = JSON.parse(fs.readFileSync(path.join(__dirname, '../package.json'), 'utf8'));
    return packageJson.version;
  } catch (error) {
    console.warn('Warning: Could not read current version from package.json');
    return null;
  }
}

function getVersionFromGit() {
  try {
    const gitTag = execSync('git describe --tags --abbrev=0', { encoding: 'utf8' }).trim();
    return gitTag;
  } catch (error) {
    console.warn('Warning: Could not get version from git tags');
    return null;
  }
}

function checkVersionOverlap() {
  console.log('🔍 Checking version overlap requirements...\n');

  const currentVersion = getCurrentVersion() || getVersionFromGit() || 'v1.0.0';
  console.log(`📦 Current version: ${currentVersion}\n`);

  if (!fs.existsSync(SCHEMAS_DIR)) {
    console.error('❌ Schemas directory not found:', SCHEMAS_DIR);
    process.exit(1);
  }

  const schemaFiles = fs.readdirSync(SCHEMAS_DIR)
    .filter(f => f.endsWith('.schema.json'))
    .sort();

  const results = {
    total: 0,
    validOverlaps: 0,
    insufficientOverlaps: 0,
    missingTimelines: 0,
    violations: []
  };

  for (const schemaFile of schemaFiles) {
    const schemaPath = path.join(SCHEMAS_DIR, schemaFile);
    const baseName = schemaFile.replace('.schema.json', '');

    try {
      const content = JSON.parse(fs.readFileSync(schemaPath, 'utf8'));
      const deprecationInfo = content['x-deprecated'];

      if (deprecationInfo && deprecationInfo.deprecated === true) {
        results.total++;

        const deprecatedInVersion = deprecationInfo.deprecatedInVersion;
        const removalPlannedInVersion = deprecationInfo.removalPlannedInVersion;

        console.log(`📋 Checking ${baseName}:`);
        console.log(`   Deprecated in: ${deprecatedInVersion || 'Not specified'}`);
        console.log(`   Removal planned: ${removalPlannedInVersion || 'Not specified'}`);

        if (!deprecatedInVersion || !removalPlannedInVersion) {
          results.missingTimelines++;
          results.violations.push({
            schema: baseName,
            type: 'missing_timeline',
            message: `Missing ${!deprecatedInVersion ? 'deprecatedInVersion' : 'removalPlannedInVersion'}`,
            severity: 'warning'
          });
          console.log(`   ⚠️  Missing version timeline information`);
          console.log();
          continue;
        }

        const minorVersionDiff = getMinorVersionDifference(deprecatedInVersion, removalPlannedInVersion);

        if (minorVersionDiff === null) {
          results.violations.push({
            schema: baseName,
            type: 'invalid_version',
            message: `Invalid version format in deprecation metadata`,
            severity: 'error'
          });
          console.log(`   ❌ Invalid version format`);
        } else if (minorVersionDiff < 1) {
          results.insufficientOverlaps++;
          results.violations.push({
            schema: baseName,
            type: 'insufficient_overlap',
            message: `Only ${minorVersionDiff} minor version(s) overlap (minimum: 1)`,
            severity: 'error',
            deprecatedVersion: deprecatedInVersion,
            removalVersion: removalPlannedInVersion,
            actualOverlap: minorVersionDiff
          });
          console.log(`   ❌ Insufficient version overlap: ${minorVersionDiff} minor versions (minimum: 1)`);
        } else {
          results.validOverlaps++;
          console.log(`   ✅ Valid overlap: ${minorVersionDiff} minor versions`);
        }

        // Check if removal date has passed
        if (deprecationInfo.plannedRemovalDate) {
          const removalDate = new Date(deprecationInfo.plannedRemovalDate);
          const now = new Date();

          if (removalDate < now) {
            results.violations.push({
              schema: baseName,
              type: 'overdue_removal',
              message: `Removal date passed (${deprecationInfo.plannedRemovalDate}) - schema should be removed`,
              severity: 'error'
            });
            console.log(`   🚨 Overdue for removal: ${deprecationInfo.plannedRemovalDate}`);
          }
        }

        console.log();
      }
    } catch (error) {
      console.error(`❌ Failed to parse ${schemaFile}: ${error.message}`);
      results.violations.push({
        schema: baseName,
        type: 'parse_error',
        message: error.message,
        severity: 'error'
      });
    }
  }

  // Summary
  console.log('📊 Version Overlap Summary:');
  console.log(`   Total deprecated schemas: ${results.total}`);
  console.log(`   Valid overlaps: ${results.validOverlaps}`);
  console.log(`   Insufficient overlaps: ${results.insufficientOverlaps}`);
  console.log(`   Missing timelines: ${results.missingTimelines}`);
  console.log();

  // Violations
  const errors = results.violations.filter(v => v.severity === 'error');
  const warnings = results.violations.filter(v => v.severity === 'warning');

  if (errors.length > 0) {
    console.log('❌ Version Overlap Errors:');
    errors.forEach(error => {
      console.log(`   ${error.schema}: ${error.message}`);

      if (error.type === 'insufficient_overlap') {
        const suggestedRemoval = suggestRemovalVersion(error.deprecatedVersion, 1);
        console.log(`   → Suggested removal version: ${suggestedRemoval}`);
      }
    });
    console.log();
  }

  if (warnings.length > 0) {
    console.log('⚠️  Version Overlap Warnings:');
    warnings.forEach(warning => {
      console.log(`   ${warning.schema}: ${warning.message}`);
    });
    console.log();
  }

  // Recommendations
  if (results.violations.length > 0) {
    console.log('💡 Recommendations:');

    if (results.insufficientOverlaps > 0) {
      console.log('   📅 Extend removal timelines to ensure minimum 1 minor version overlap');
    }

    if (results.missingTimelines > 0) {
      console.log('   📝 Add version timeline information to deprecation metadata');
    }

    const overdueRemovals = results.violations.filter(v => v.type === 'overdue_removal');
    if (overdueRemovals.length > 0) {
      console.log('   🗑️  Remove schemas that are past their removal date');
    }
  }

  // Exit with error if there are violations
  if (errors.length > 0) {
    console.log('\n🚫 Version overlap validation failed');
    process.exit(1);
  } else if (results.total === 0) {
    console.log('✅ No deprecated schemas found');
  } else {
    console.log('✅ All deprecated schemas meet version overlap requirements');
  }

  return results;
}

function suggestRemovalVersion(deprecatedVersion, minMinorOverlap = 1) {
  const parsed = parseVersion(deprecatedVersion);
  if (!parsed) return 'Invalid version';

  return `v${parsed.major}.${parsed.minor + minMinorOverlap + 1}.0`;
}

function validateVersionOverlapPolicy(deprecatedVersion, removalVersion) {
  const minorDiff = getMinorVersionDifference(deprecatedVersion, removalVersion);

  return {
    valid: minorDiff !== null && minorDiff >= 1,
    actualOverlap: minorDiff,
    requiredOverlap: 1,
    suggestedRemoval: minorDiff !== null && minorDiff < 1 ? suggestRemovalVersion(deprecatedVersion, 1) : null
  };
}

function main() {
  try {
    checkVersionOverlap();
  } catch (error) {
    console.error('❌ Error checking version overlap:', error.message);
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}

module.exports = {
  checkVersionOverlap,
  validateVersionOverlapPolicy,
  parseVersion,
  compareVersions,
  getMinorVersionDifference
};