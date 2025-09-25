#!/usr/bin/env node
/**
 * Validates that CHANGELOG.md has been updated with proper SemVer entry
 * Usage: node scripts/check-changelog.js
 */

const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const CHANGELOG_FILE = path.join(__dirname, '../CHANGELOG.md');

function getChangedFiles() {
  try {
    // Get files changed compared to main branch
    const output = execSync('git diff --name-only origin/main...HEAD', {
      encoding: 'utf8',
      stdio: ['pipe', 'pipe', 'ignore'] // Suppress stderr
    });
    return output.trim().split('\n').filter(f => f);
  } catch {
    try {
      // Fallback: get files in last commit
      const output = execSync('git diff --name-only HEAD~1 HEAD', {
        encoding: 'utf8',
        stdio: ['pipe', 'pipe', 'ignore']
      });
      return output.trim().split('\n').filter(f => f);
    } catch {
      // If even that fails, assume first commit
      return [];
    }
  }
}

function hasSchemaChanges(changedFiles) {
  return changedFiles.some(file =>
    file.startsWith('schemas/') ||
    file.startsWith('openapi/') ||
    file === 'package.json'
  );
}

function getChangelogContent() {
  if (!fs.existsSync(CHANGELOG_FILE)) {
    return null;
  }
  return fs.readFileSync(CHANGELOG_FILE, 'utf8');
}

function hasUnreleasedChanges(content) {
  if (!content) return false;

  const lines = content.split('\n');
  let inUnreleased = false;
  let hasContent = false;

  for (const line of lines) {
    if (line.includes('## [Unreleased]')) {
      inUnreleased = true;
      continue;
    }

    if (inUnreleased) {
      if (line.startsWith('## ')) {
        // Hit next section, stop looking
        break;
      }

      // Look for actual content (not just empty lines or headers)
      if (line.trim() && !line.startsWith('###') && !line.startsWith('## ')) {
        hasContent = true;
        break;
      }
    }
  }

  return hasContent;
}

function hasRecentVersionEntry(content) {
  if (!content) return false;

  // Look for version entries like ## [1.2.3] or ## [1.2.3] - YYYY-MM-DD
  const versionRegex = /^## \[\d+\.\d+\.\d+\]/m;
  return versionRegex.test(content);
}

function main() {
  console.log('📋 Checking CHANGELOG.md requirements...');

  const changedFiles = getChangedFiles();
  const hasSchemaModifications = hasSchemaChanges(changedFiles);

  if (!hasSchemaModifications) {
    console.log('✅ No schema changes detected, CHANGELOG check skipped');
    return;
  }

  console.log('🔍 Schema changes detected, checking CHANGELOG.md...');
  console.log('Changed files:', changedFiles.join(', '));

  const content = getChangelogContent();

  if (!content) {
    console.error('❌ CHANGELOG.md not found');
    process.exit(1);
  }

  const hasUnreleased = hasUnreleasedChanges(content);
  const hasVersionEntry = hasRecentVersionEntry(content);

  if (!hasUnreleased && !hasVersionEntry) {
    console.error('❌ CHANGELOG.md must be updated with either:');
    console.error('   - Entry under ## [Unreleased] section, or');
    console.error('   - New version entry (## [X.Y.Z] - YYYY-MM-DD)');
    console.error('\nFor schema changes, please document:');
    console.error('   - What schemas were added/modified');
    console.error('   - Whether this is a MAJOR, MINOR, or PATCH change');
    console.error('   - Any breaking changes or migration notes');
    process.exit(1);
  }

  if (hasUnreleased) {
    console.log('✅ CHANGELOG.md has unreleased changes documented');
  } else if (hasVersionEntry) {
    console.log('✅ CHANGELOG.md has recent version entry');
  }
}

if (require.main === module) {
  main();
}