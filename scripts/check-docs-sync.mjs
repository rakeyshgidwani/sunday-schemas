#!/usr/bin/env node

/**
 * Sunday Schemas - Documentation Sync Checker
 *
 * Verifies that documentation is in sync with generated code.
 * Used in CI/CD and pre-commit hooks to catch documentation drift.
 */

import fs from 'fs/promises';
import path from 'path';
import { fileURLToPath } from 'url';
import { execSync } from 'child_process';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const rootDir = path.resolve(__dirname, '..');

async function checkDocumentationSync() {
    console.log('🔍 Checking documentation sync with codegen...');

    try {
        // Get current documentation content
        const currentTsDocs = await fs.readFile(path.join(rootDir, 'docs', 'TYPESCRIPT_USAGE.md'), 'utf8');
        const currentGoDocs = await fs.readFile(path.join(rootDir, 'docs', 'GO_MODULE_USAGE.md'), 'utf8');

        // Generate fresh documentation
        console.log('📝 Generating fresh documentation...');
        execSync('npm run generate:docs', { cwd: rootDir, stdio: 'pipe' });

        // Read newly generated docs
        const newTsDocs = await fs.readFile(path.join(rootDir, 'docs', 'TYPESCRIPT_USAGE.md'), 'utf8');
        const newGoDocs = await fs.readFile(path.join(rootDir, 'docs', 'GO_MODULE_USAGE.md'), 'utf8');

        // Check for significant differences (ignore timestamps)
        const tsOutOfSync = removeTimestamps(currentTsDocs) !== removeTimestamps(newTsDocs);
        const goOutOfSync = removeTimestamps(currentGoDocs) !== removeTimestamps(newGoDocs);

        if (tsOutOfSync || goOutOfSync) {
            console.error('❌ Documentation is out of sync with codegen!');
            console.error('');

            if (tsOutOfSync) {
                console.error('📄 TYPESCRIPT_USAGE.md is outdated');
            }
            if (goOutOfSync) {
                console.error('📄 GO_MODULE_USAGE.md is outdated');
            }

            console.error('');
            console.error('🔧 To fix this, run:');
            console.error('   npm run generate:docs');
            console.error('');
            console.error('💡 Or add this to your build pipeline:');
            console.error('   npm run generate && npm run generate:docs');

            process.exit(1);
        } else {
            console.log('✅ Documentation is in sync with codegen');
        }

    } catch (error) {
        console.error('❌ Failed to check documentation sync:', error.message);
        process.exit(1);
    }
}

function removeTimestamps(content) {
    // Remove the "Last updated" timestamp to ignore minor changes
    return content.replace(/\*Last updated: [^*]+\*/g, '*Last updated: [TIMESTAMP]*');
}

if (import.meta.url === `file://${process.argv[1]}`) {
    checkDocumentationSync();
}