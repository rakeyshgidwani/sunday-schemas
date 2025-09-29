#!/usr/bin/env node

import { readFileSync } from 'fs';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const rootDir = join(__dirname, '..');

try {
  // Read package.json from packages/ts
  const packagePath = join(rootDir, 'packages/ts/package.json');
  const pkg = JSON.parse(readFileSync(packagePath, 'utf8'));

  // Get tag from environment (GitHub Actions sets GITHUB_REF_NAME)
  const tag = process.env.GITHUB_REF_NAME || process.env.npm_config_tag || process.argv[2];

  if (!tag) {
    console.error('No tag provided. Set GITHUB_REF_NAME or pass as argument.');
    process.exit(1);
  }

  // Remove 'v' prefix if present (e.g., v1.0.1 -> 1.0.1)
  const version = tag.startsWith('v') ? tag.slice(1) : tag;

  if (pkg.version !== version) {
    console.error(`❌ Version mismatch:`);
    console.error(`   package.json version: ${pkg.version}`);
    console.error(`   git tag version:      ${version}`);
    console.error(`   Please update packages/ts/package.json version to match the tag.`);
    process.exit(1);
  }

  console.log(`✅ Version OK: ${pkg.version} matches tag ${tag}`);
} catch (error) {
  console.error(`❌ Error checking version: ${error.message}`);
  process.exit(1);
}