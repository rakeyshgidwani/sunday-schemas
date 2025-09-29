#!/usr/bin/env node

import { readdir, writeFile, access } from 'fs/promises';
import { join, basename, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const rootDir = join(__dirname, '..');
const tsPackageDir = join(rootDir, 'packages/ts');
const eventsDir = join(tsPackageDir, 'src/events');
const uiDir = join(tsPackageDir, 'src/ui');
const outputFile = join(tsPackageDir, 'src/index.d.ts');

async function fileExists(path) {
  try {
    await access(path);
    return true;
  } catch {
    return false;
  }
}

async function generateBarrelExport() {
  const exports = [];

  // Header comment
  exports.push('// Generated barrel file - do not edit manually');
  exports.push('// Run `npm run build` to regenerate');
  exports.push('');

  // Export events types
  try {
    if (await fileExists(eventsDir)) {
      const eventFiles = await readdir(eventsDir);
      const typeFiles = eventFiles.filter(f => f.endsWith('.d.ts') && f !== 'index.d.ts');

      if (typeFiles.length > 0) {
        exports.push('// Event schemas');
        for (const file of typeFiles) {
          const name = basename(file, '.d.ts');
          // Convert schema file names to export names (e.g., md.orderbook.delta.v1.schema -> MdOrderbookDeltaV1)
          const exportName = name
            .replace(/\.schema$/, '')
            .split('.')
            .map(part => part.charAt(0).toUpperCase() + part.slice(1))
            .join('');
          exports.push(`export * as ${exportName} from './events/${name}.js';`);
        }
        exports.push('');
      }
    }
  } catch (error) {
    console.warn(`Warning: Could not read events directory: ${error.message}`);
  }

  // Export UI types (single file)
  try {
    if (await fileExists(join(uiDir, 'index.d.ts'))) {
      exports.push('// UI/API schemas');
      exports.push("export * as UI from './ui/index.js';");
      exports.push('');
    }
  } catch (error) {
    console.warn(`Warning: Could not check UI types: ${error.message}`);
  }

  // Add namespace export for convenience
  exports.push('// Namespace exports for convenience');
  exports.push('export * as Events from \'./events/index.js\';');
  exports.push('export * as Api from \'./ui/index.js\';');

  const content = exports.join('\n');

  try {
    await writeFile(outputFile, content, 'utf8');
    console.log(`✅ Generated barrel export at ${outputFile}`);
    console.log(`   Exported ${exports.filter(line => line.startsWith('export')).length} modules`);
  } catch (error) {
    console.error(`❌ Error writing barrel export: ${error.message}`);
    process.exit(1);
  }
}

// Create events index file if it doesn't exist
async function ensureEventsIndex() {
  const eventsIndexPath = join(eventsDir, 'index.d.ts');

  if (!(await fileExists(eventsIndexPath))) {
    try {
      if (await fileExists(eventsDir)) {
        const eventFiles = await readdir(eventsDir);
        const typeFiles = eventFiles.filter(f => f.endsWith('.d.ts') && f !== 'index.d.ts');

        const eventsExports = [
          '// Generated events index - do not edit manually',
          ''
        ];

        for (const file of typeFiles) {
          const name = basename(file, '.d.ts');
          eventsExports.push(`export * from './${name}.js';`);
        }

        await writeFile(eventsIndexPath, eventsExports.join('\n'), 'utf8');
        console.log(`✅ Generated events index at ${eventsIndexPath}`);
      }
    } catch (error) {
      console.warn(`Warning: Could not create events index: ${error.message}`);
    }
  }
}

try {
  await ensureEventsIndex();
  await generateBarrelExport();
} catch (error) {
  console.error(`❌ Error generating barrel export: ${error.message}`);
  process.exit(1);
}