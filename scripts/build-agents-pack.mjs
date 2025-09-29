#!/usr/bin/env node

import { readdir, readFile, writeFile, mkdir, access } from 'fs/promises';
import { join, basename, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const rootDir = join(__dirname, '..');
const distDir = join(rootDir, 'dist');
const agentsDir = join(distDir, 'agents');
const schemasDir = join(rootDir, 'schemas/json');
const examplesDir = join(distDir, 'schemas/examples');

async function fileExists(path) {
  try {
    await access(path);
    return true;
  } catch {
    return false;
  }
}

async function safeReadJson(path) {
  try {
    const content = await readFile(path, 'utf8');
    return JSON.parse(content);
  } catch (error) {
    console.warn(`Warning: Could not read ${path}: ${error.message}`);
    return null;
  }
}

async function buildAgentsPack() {
  // Ensure agents directory exists
  await mkdir(agentsDir, { recursive: true });

  // Get metadata
  const version = (process.env.GITHUB_REF_NAME || '').replace(/^v/, '') || '0.0.0-dev';
  const commit_sha = process.env.GITHUB_SHA || '';
  const built_at = new Date().toISOString();

  // Load topics mapping
  const topics = await safeReadJson(join(rootDir, 'schemas/topics.json')) || {};

  // Load venues registry
  const venues = await safeReadJson(join(rootDir, 'schemas/registries/venues.json')) || {};

  // Process schemas
  const schemas = [];

  try {
    if (await fileExists(schemasDir)) {
      const schemaFiles = await readdir(schemasDir);
      const jsonFiles = schemaFiles.filter(f => f.endsWith('.json'));

      for (const file of jsonFiles) {
        const filePath = join(schemasDir, file);
        const schema = await safeReadJson(filePath);

        if (schema) {
          // Extract schema ID from $id or filename
          const id = (schema.$id || '').split('/').pop()?.replace('.schema.json', '') ||
                     basename(file, '.json').replace('.schema', '');

          // Find corresponding examples
          const exampleList = [];
          if (await fileExists(examplesDir)) {
            try {
              const exampleFiles = await readdir(examplesDir);
              const matchingExamples = exampleFiles.filter(ef =>
                ef.includes(id) || id.includes(ef.replace('.example.json', ''))
              );
              exampleList.push(...matchingExamples.map(ef => `dist/schemas/examples/${ef}`));
            } catch (error) {
              console.warn(`Warning: Could not read examples directory: ${error.message}`);
            }
          }

          schemas.push({
            id,
            file: `dist/schemas/json/${file}`,
            topic: topics[id] || null,
            examples: exampleList,
            summary: schema.title || schema.description || ''
          });
        }
      }
    }
  } catch (error) {
    console.error(`Error processing schemas: ${error.message}`);
  }

  // Build index.json
  const index = {
    version,
    commit_sha,
    built_at,
    openapi: {
      json: 'dist/openapi/ui.v1.bundled.json',
      yaml: 'dist/openapi/ui.v1.yaml'
    },
    schemas,
    registries: {
      venues: 'dist/schemas/registries/venues.json'
    },
    diff: {
      openapi_markdown: 'dist/diff/openapi-diff.md',
      schemas_markdown: 'dist/diff/json-schema-diff.md'
    }
  };

  // Write index.json
  const indexPath = join(agentsDir, 'index.json');
  await writeFile(indexPath, JSON.stringify(index, null, 2), 'utf8');
  console.log(`✅ Generated agents index at ${indexPath}`);

  // Generate cheatsheet.md
  let cheatsheet = `# Sunday Schemas — Cheatsheet

Generated at: ${built_at}
Version: ${version}
Commit: ${commit_sha}

## Overview

This cheatsheet provides a quick reference for all available schemas in the Sunday platform.

`;

  if (schemas.length > 0) {
    cheatsheet += `## Schemas\n\n`;

    for (const schema of schemas) {
      cheatsheet += `### ${schema.id}\n\n`;
      cheatsheet += `**Topic:** ${schema.topic || 'n/a'}\n\n`;
      if (schema.summary) {
        cheatsheet += `${schema.summary}\n\n`;
      }
      cheatsheet += `**Schema:** ${schema.file}\n`;
      if (schema.examples.length > 0) {
        cheatsheet += `**Examples:** ${schema.examples.join(', ')}\n`;
      } else {
        cheatsheet += `**Examples:** —\n`;
      }
      cheatsheet += `\n---\n\n`;
    }
  }

  cheatsheet += `## Usage

1. **Download the pack:** Get \`agents/pack.zip\` from the release assets
2. **Load schemas:** Use the paths in \`agents/index.json\` to load specific schemas
3. **Validate data:** Use examples as reference for expected data formats
4. **API integration:** Use the OpenAPI spec for UI BFF endpoints

## Resources

- **OpenAPI JSON:** ${index.openapi.json}
- **OpenAPI YAML:** ${index.openapi.yaml}
- **Venues Registry:** ${index.registries.venues}

`;

  const cheatsheetPath = join(agentsDir, 'cheatsheet.md');
  await writeFile(cheatsheetPath, cheatsheet, 'utf8');
  console.log(`✅ Generated agents cheatsheet at ${cheatsheetPath}`);

  console.log(`📦 Agents pack ready with ${schemas.length} schemas`);
}

try {
  await buildAgentsPack();
} catch (error) {
  console.error(`❌ Error building agents pack: ${error.message}`);
  process.exit(1);
}