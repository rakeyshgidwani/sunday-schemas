# sunday-schemas — End‑to‑End Release Spec (TypeScript & Go)

> **Objective:** Ship **versioned, reproducible** releases for the `sunday-schemas` TypeScript package (npm) and Go module (Go proxy), tied to the same source tag and validated against back‑compat rules.
>
> **Out of scope:** UI docs hosting. (Docs-as-assets are covered in the Release Assets spec; this document focuses on **publishing the TS and Go modules end‑to‑end**.)

---

## 1) Scope & Principles
- **Single source of truth:** JSON Schemas (events) + OpenAPI (UI). Code in TS/Go is generated from those.
- **One tag → many artifacts:** Creating **`vX.Y.Z`** publishes npm package **`@sunday/schemas@X.Y.Z`** and creates/validates the Go module tag **`go/vX.Y.Z`**.
- **Reproducible:** Linted, validated, diffed against previous tag; checksums and SBOM attached as release assets (per separate spec).
- **SemVer:** Back‑compat is enforced in CI. Breaking changes require MAJOR and coordinated rollout (not allowed in Phase‑1 by policy).

---

## 2) Versioning & Tagging
- **Repo tag (root):** `vX.Y.Z` (e.g., `v1.0.1`). Drives npm release and release assets.
- **Go submodule tag:** `go/vX.Y.Z` (same commit). Go tooling expects module‑path‑prefixed tags for submodules.
- **Prereleases:** Use `vX.Y.Z-rc.N` for candidates.
  - npm dist‑tag mapping: `-rc.N` → `next`, stable → `latest`.
  - Go: publish `go/vX.Y.Z-rc.N` (supported by proxy); consumers must opt‑in explicitly.

---

## 3) Repository Layout (relevant parts)
```
/sunday-schemas
  /schemas
    /json                 # JSON Schemas (events)
    /examples             # Validated examples
    /registries           # venues.json, instruments.json (optional)
    topics.json           # schema ↔ topic mapping
  /openapi
    ui.v1.yaml            # OpenAPI 3.1
  /packages/ts            # TypeScript package workspace
    package.json
    src/
      events/             # generated d.ts from JSON Schemas
      ui/                 # generated d.ts from OpenAPI
      index.d.ts          # barrel export (generated)
    openapi/
      ui.v1.yaml
      ui.v1.bundled.json
    docs/                 # (optional) bundled docs copied in package
  /go                     # Go module root: github.com/sunday-xyz/schemas/go
    go.mod
    rawv0/                # structs/types for raw.v0 envelope
    md/                   # md.* types
    insights/             # insights.* types
    validate/             # JSON Schema validation helpers (if included)
  /scripts
    check-version-matches-tag.mjs
    rollup-ts-index.mjs
  CHANGELOG.md
```

---

## 4) TypeScript Package (npm)
**Package name:** `@sunday/schemas`  
**Version:** must match the repo tag (e.g., `1.0.1`).

### 4.1 Codegen & Contents
- **Events (JSON Schema → d.ts):** `json-schema-to-typescript` generates `.d.ts` into `packages/ts/src/events/`.
- **OpenAPI → d.ts:** `openapi-typescript openapi/ui.v1.yaml -o packages/ts/src/ui/index.d.ts`.
- **Bundle OpenAPI into package:** copy `openapi/ui.v1.yaml` and `openapi/ui.v1.bundled.json` into `packages/ts/openapi/` and export as subpaths.
- **Barrel export:** `rollup-ts-index.mjs` creates `packages/ts/src/index.d.ts` that re‑exports `events/*` and `ui/*` types.

### 4.2 package.json (key fields)
```json
{
  "name": "@sunday/schemas",
  "version": "1.0.1",
  "type": "module",
  "files": [
    "src/**/*.d.ts",
    "openapi/ui.v1.yaml",
    "openapi/ui.v1.bundled.json",
    "docs/**"
  ],
  "exports": {
    ".": "./src/index.d.ts",
    "./ui": "./src/ui/index.d.ts",
    "./events": "./src/events/index.d.ts",
    "./openapi.yaml": "./openapi/ui.v1.yaml",
    "./openapi.json": "./openapi/ui.v1.bundled.json"
  },
  "publishConfig": { "access": "public" },
  "scripts": {
    "gen:events": "json-schema-to-typescript \"schemas/json/**/*.json\" --cwd . --outDir packages/ts/src/events",
    "gen:openapi": "openapi-typescript openapi/ui.v1.yaml -o packages/ts/src/ui/index.d.ts",
    "build": "npm run gen:events && npm run gen:openapi && node scripts/rollup-ts-index.mjs",
    "verify:versions": "node scripts/check-version-matches-tag.mjs",
    "prepublishOnly": "npm run build && npm run verify:versions"
  }
}
```

### 4.3 Version/Tag Guard
`scripts/check-version-matches-tag.mjs` verifies `package.json.version === $GITHUB_REF_NAME` (minus leading `v`); fails CI if not.

### 4.4 Dist‑tags
- Stable tag `vX.Y.Z`: `npm publish` (sets `latest`).
- RC tag `vX.Y.Z-rc.N`: `npm publish --tag next`.

---

## 5) Go Module
**Module path:** `github.com/sunday-xyz/schemas/go`  
**go.mod:**
```go
module github.com/sunday-xyz/schemas/go

go 1.23

require (
    // pin any validation/helper deps if used
)
```

### 5.1 Code Organization
- **rawv0**, **md**, **insights** packages with Go structs mirroring the schemas.  
- Optional `validate` helpers wrapping a JSON Schema validator (or keep validation in producers/consumers and ship types only).

### 5.2 Publishing
- **Tag creation = publish.** The Go proxy indexes the module when first requested. Tag must be **`go/vX.Y.Z`** at the same commit as `vX.Y.Z`.
- CI runs `go vet`/`go test` and optionally **warms the proxy**.

---

## 6) CI/CD — Workflows
> All jobs run on GitHub Actions; tokens stored as repo secrets. Jobs are idempotent.

### 6.1 On repo tag `v*.*.*` (root)
Runs **validation + docs build (separate workflow)**, **npm publish**, and **create Go tag**.

```yaml
name: Release (root tag)
on:
  push:
    tags: ["v*.*.*"]
jobs:
  validate-and-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with: { node-version: "20" }
      - uses: actions/setup-python@v5
        with: { python-version: "3.11" }
      - run: npm ci
      # Lint & validate (abbrev — full gates defined in release-assets spec)
      - run: npx spectral lint openapi/ui.v1.yaml
      - run: npx jsonlint -q schemas/json/**/*.json schemas/examples/**/*.json
      - run: npx ajv validate -s schemas/json/**/*.json -d schemas/examples/**/*.json || true
      # Bundle OpenAPI for both package + assets
      - run: npx redocly bundle openapi/ui.v1.yaml -o openapi/ui.v1.bundled.json
      # Keep as artifacts for downstream jobs if needed
      - uses: actions/upload-artifact@v4
        with:
          name: openapi-bundle
          path: openapi/ui.v1.bundled.json

  publish-npm:
    needs: validate-and-build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with: { node-version: "20", registry-url: "https://registry.npmjs.org" }
      - run: npm ci
      - run: npm run build
      - name: Verify version matches tag
        run: node scripts/check-version-matches-tag.mjs
      - name: Copy bundled OpenAPI into package
        run: |
          mkdir -p packages/ts/openapi
          cp openapi/ui.v1.yaml packages/ts/openapi/
          cp openapi/ui.v1.bundled.json packages/ts/openapi/
      - name: Publish to npm (stable)
        if: ${{ !contains(github.ref_name, '-rc.') }}
        env: { NODE_AUTH_TOKEN: ${{ secrets.NPM_TOKEN }} }
        run: npm publish --workspace=@sunday/schemas
      - name: Publish to npm (prerelease → next)
        if: ${{ contains(github.ref_name, '-rc.') }}
        env: { NODE_AUTH_TOKEN: ${{ secrets.NPM_TOKEN }} }
        run: npm publish --workspace=@sunday/schemas --tag next

  create-go-tag:
    needs: validate-and-build
    runs-on: ubuntu-latest
    permissions: { contents: write }
    steps:
      - uses: actions/checkout@v4
      - name: Create go/ tag matching root tag
        run: |
          GIT_TAG=${GITHUB_REF_NAME}
          GO_TAG="go/${GIT_TAG}"
          git tag -a "$GO_TAG" -m "Go module release $GO_TAG" "$GIT_TAG" || true
          git push origin "$GO_TAG" || true
```

### 6.2 On Go tag `go/v*.*.*`
Builds/tests the Go module and optionally warms the Go proxy.

```yaml
name: Publish Go Module (go/ tag)
on:
  push:
    tags: ["go/v*.*.*"]
jobs:
  publish-go:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: "1.23" }
      - run: (cd go && go vet ./... && go test ./...)
      - name: Warm go proxy (optional)
        run: |
          VERSION=${GITHUB_REF_NAME#go/}
          MOD=github.com/sunday-xyz/schemas/go@${VERSION}
          curl -sSfL https://proxy.golang.org/${MOD}/info > /dev/null || true
```

> **Note:** If you maintain a Go **multi‑module** workspace, keep the `go` module self‑contained. No other submodules are published in Phase‑1.

---

## 7) Quality Gates (shared with Release Assets spec)
- **OpenAPI lint:** Spectral ruleset passes.
- **Schema lint:** jsonlint.
- **Examples validate:** `ajv` against corresponding schemas.
- **Back‑compat:** `json-schema-diff` (no narrowed enums/required removals) and `oasdiff` (no breaking API changes). Gate is informational for `-rc.*`, enforced for stable.
- **topics.json / venues.json coverage:** All `$id`s mapped; venue enum changes are only additive.

---

## 8) Security & Provenance
- **npm:** Use org‑scoped `NPM_TOKEN` with **publish only** rights; enable **2FA** for publishes on the org. Optionally sign tarballs with provenance (npm provenance / GitHub OIDC).
- **Go:** No credentials needed; tags are immutable; prefer signed tags.
- **SBOM & checksums:** Produced in release‑assets pipeline; include in GitHub Release.

---

## 9) Rollback & Remediation
- **npm:** If a bad version ships, run `npm deprecate @sunday/schemas@X.Y.Z "Deprecated due to <reason>; use X.Y.Z+1"` and publish a patched **X.Y.Z+1** with the fix.
- **Go:** Create a new tag **`go/vX.Y.Z+1`**. Avoid deleting tags; consumers may have cached sums.
- **Docs/Assets:** Re‑run the release workflow with a patched tag; assets are attached to the new release.

---

## 10) Acceptance Criteria
1. Creating tag **`vX.Y.Z`** publishes `@sunday/schemas@X.Y.Z` to npm (stable) or to `next` (if `-rc.*`), and creates/pushes the matching **`go/vX.Y.Z`** tag.
2. Creating tag **`go/vX.Y.Z`** runs Go tests and the module resolves via `go get github.com/sunday-xyz/schemas/go@vX.Y.Z`.
3. The npm package includes: generated TS types (`events`, `ui`) and bundled OpenAPI (`openapi.yaml/json`) available via `exports` paths.
4. All validation gates pass for stable releases; failures block publishing.
5. Release notes and CHANGELOG entry are updated for the version.

---

## 11) Optional Enhancements
- **Conventional Commits + auto changelog:** Use `changesets` or `semantic-release` (with `conventionalcommits`) to generate CHANGELOG and bump versions.
- **Provenance/sigstore:** Sign release zip and/or npm provenance; attach signatures to GitHub Release.
- **Monorepo compatibility:** If moving to a larger monorepo later, keep the `go/` submodule tagging rule and workspace isolation; npm workspace stays unchanged.

---

## 12) Developer Commands (local parity)
```bash
# Generate TS types
npm run gen:events && npm run gen:openapi && node scripts/rollup-ts-index.mjs

# Publish dry‑run (locally)
npm pack packages/ts

# Go module tests
( cd go && go vet ./... && go test ./... )
```

---

## 13) Helper Scripts (snippets)
**scripts/check-version-matches-tag.mjs**
```js
const fs = require('fs');
const pkg = JSON.parse(fs.readFileSync('packages/ts/package.json', 'utf8'));
const tag = process.env.GITHUB_REF_NAME; // e.g., v1.0.1 or v1.0.1-rc.1
const version = tag.startsWith('v') ? tag.slice(1) : tag;
if (pkg.version !== version) {
  console.error(`Version mismatch: package.json=${pkg.version} tag=${version}`);
  process.exit(1);
}
console.log(`Version OK: ${pkg.version}`);
```

**scripts/rollup-ts-index.mjs** (naive barrel example)
```js
const fs = require('fs');
const path = require('path');
const out = ['// generated barrel file'];
function reexportDir(dir, alias) {
  const files = fs.readdirSync(dir).filter(f => f.endsWith('.d.ts'));
  files.forEach(f => {
    const name = path.basename(f, '.d.ts');
    out.push(`export * as ${alias}_${name} from './${alias}/${name}.d';`);
  });
}
reexportDir('packages/ts/src/events', 'events');
// UI is a single index.d.ts already
out.push("export * as ui from './ui/index.d';");
fs.writeFileSync('packages/ts/src/index.d.ts', out.join('\n'));
```

> Replace with your preferred codegen/rollup approach; the key is that the package **only ships types + OpenAPI artifacts**, nothing executable.



---

## Addendum: Agents Pack (LLM‑friendly) — v1
**Goal:** Make every release consumable by AI agents and automation without HTML rendering. Ship a small machine‑readable index and a self‑contained zip.

### A1) New Release Assets
- `agents/index.json` — versioned index pointing to OpenAPI JSON/YAML, JSON Schemas, examples, topics/venues, and diffs.
- `agents/pack.zip` — bundle for offline/air‑gapped use: `index.json`, bundled `openapi.json`, `schemas/` (JSON), `examples.ndjson`, and `cheatsheet.md`.

**`agents/index.json` (v1 draft)**
```json
{
  "version": "<semver>",
  "commit_sha": "<40-hex>",
  "built_at": "<ISO8601>",
  "openapi": {
    "json": "dist/openapi/ui.v1.bundled.json",
    "yaml": "dist/openapi/ui.v1.yaml"
  },
  "schemas": [
    {
      "id": "md.orderbook.delta.v1",
      "file": "dist/schemas/json/md.orderbook.delta.v1.schema.json",
      "topic": "md.normalized.orderbook",
      "examples": ["dist/schemas/examples/md.orderbook.delta.example.json"],
      "summary": "Order book deltas with [prob,size] pairs; seq monotonic; set is_snapshot=true after gaps."
    }
  ],
  "registries": {"venues": "dist/schemas/registries/venues.json"},
  "diff": {
    "openapi_markdown": "dist/diff/openapi-diff.md",
    "schemas_markdown": "dist/diff/json-schema-diff.md"
  }
}
```

### A2) CI Additions
Add steps to **build the agents pack** after docs are generated and sources copied into `dist/`:
1. **Index & cheatsheet**
   - Run `node scripts/build-agents-pack.mjs` to create `dist/agents/index.json` and `dist/agents/cheatsheet.md` by scanning `schemas/json`, `schemas/examples`, `schemas/topics.json`, and `dist/openapi/ui.v1.bundled.json`.
2. **Flatten examples**
   - Create `dist/agents/examples.ndjson` (one JSON per line) from all example files.
3. **Zip bundle**
   - `(cd dist && zip -r agents/pack.zip agents/index.json agents/cheatsheet.md openapi/ui.v1.bundled.json schemas/json agents/examples.ndjson)`
4. **Upload assets**
   - Include `dist/agents/**` in the release upload list.

**Workflow patch (excerpt):**
```yaml
      - name: Build agents index & cheatsheet
        run: node scripts/build-agents-pack.mjs
      - name: Flatten examples to NDJSON
        run: |
          jq -c . dist/schemas/examples/*.json > dist/agents/examples.ndjson || true
      - name: Pack agents bundle
        run: |
          (cd dist && zip -r agents/pack.zip agents/index.json agents/cheatsheet.md openapi/ui.v1.bundled.json schemas/json agents/examples.ndjson 2>/dev/null || true)
      - uses: softprops/action-gh-release@v2
        with:
          files: |
            # …existing assets…
            dist/agents/**
```

### A3) Repo Scripts
Add a Node script to `scripts/build-agents-pack.mjs`:
```js
import { promises as fs } from 'node:fs';
import path from 'node:path';
const DIST = 'dist';
const SCHEMAS = 'schemas/json';
const EXAMPLES = 'dist/schemas/examples';
const TOPICS = JSON.parse(await fs.readFile('schemas/topics.json', 'utf8'));
const VENUES = await fs.readFile('schemas/registries/venues.json', 'utf8').then(JSON.parse).catch(()=>({}));
const version = (process.env.GITHUB_REF_NAME || '').replace(/^v/, '') || '0.0.0-dev';
const commit_sha = process.env.GITHUB_SHA || '';
const built_at = new Date().toISOString();
const schemaFiles = (await fs.readdir(SCHEMAS)).filter(f=>f.endsWith('.json'));
const schemas = [];
for (const f of schemaFiles) {
  const p = path.join(SCHEMAS, f);
  const j = JSON.parse(await fs.readFile(p, 'utf8'));
  const id = (j.$id || '').split('/').pop()?.replace('.schema.json','') || f.replace('.json','');
  const exampleList = await fs.readdir(EXAMPLES).then(list => list.filter(x => x.includes(id))).catch(()=>[]);
  schemas.push({ id, file: `dist/schemas/json/${f}`, topic: TOPICS[id] || null, examples: exampleList.map(e=>`dist/schemas/examples/${e}`), summary: j.title || '' });
}
await fs.mkdir(path.join(DIST,'agents'), { recursive: true });
await fs.writeFile(path.join(DIST,'agents/index.json'), JSON.stringify({ version, commit_sha, built_at, openapi:{ json:'dist/openapi/ui.v1.bundled.json', yaml:'dist/openapi/ui.v1.yaml' }, schemas, registries:{ venues:'dist/schemas/registries/venues.json' }, diff:{ openapi_markdown:'dist/diff/openapi-diff.md', schemas_markdown:'dist/diff/json-schema-diff.md' } }, null, 2));
let cheat = '# Sunday Schemas — Cheatsheet

';
schemas.forEach(s => { cheat += `## ${s.id}
Topic: ${s.topic || 'n/a'}

${s.summary}

Required: see ${s.file}
Examples: ${s.examples.join(', ') || '—'}

`; });
await fs.writeFile(path.join(DIST,'agents/cheatsheet.md'), cheat);
```

### A4) Acceptance Criteria Additions
- The release includes **`dist/agents/index.json`** and **`dist/agents/pack.zip`**.
- `index.json` references valid paths for OpenAPI JSON/YAML, all JSON Schemas, examples, topics/venues, and diff files.
- An internal tool/agent can fetch `index.json` and, without human intervention, load the OpenAPI and schemas it points to.

