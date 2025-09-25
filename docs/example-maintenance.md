# Example Maintenance Guide

This document describes how to maintain and validate the example data files that accompany each JSON schema in the Sunday platform.

## Current Example Coverage

### ✅ Complete Coverage (13 examples for 8 schemas)

| Schema | Examples | Coverage |
|--------|----------|----------|
| `raw.v0.envelope` | `raw.polymarket.orderbook.example.json`<br>`raw.kalshi.trade.example.json` | Both venues ✅ |
| `md.orderbook.delta.v1` | `md.orderbook.delta.example.json`<br>`md.orderbook.snapshot.example.json` | Delta + snapshot ✅ |
| `md.trade.v1` | `md.trade.buy.example.json`<br>`md.trade.sell.example.json` | Both sides ✅ |
| `insights.arb.lite.v1` | `insights.arb.lite.example.json` | Basic coverage ✅ |
| `insights.movers.v1` | `insights.movers.example.json` | Basic coverage ✅ |
| `insights.whales.lite.v1` | `insights.whales.lite.example.json` | Basic coverage ✅ |
| `insights.unusual.v1` | `insights.unusual.example.json` | Basic coverage ✅ |
| `infra.venue_health.v1` | `infra.venue_health.connected.example.json`<br>`infra.venue_health.degraded.example.json`<br>`infra.venue_health.stale.example.json` | All health states ✅ |

## Example Naming Conventions

### File Naming Pattern
```
{schema_base_name}.{scenario}.example.json
```

**Examples:**
- `md.trade.buy.example.json` - Buy-side trade
- `md.trade.sell.example.json` - Sell-side trade
- `infra.venue_health.connected.example.json` - Connected health state
- `raw.polymarket.orderbook.example.json` - Polymarket raw data

### Scenario Types

**By Venue:**
- `.polymarket.` - Polymarket-specific examples
- `.kalshi.` - Kalshi-specific examples

**By State/Type:**
- `.connected.`, `.degraded.`, `.stale.` - Health states
- `.delta.`, `.snapshot.` - Orderbook update types
- `.buy.`, `.sell.` - Trade directions

**By Use Case:**
- `.basic.` - Simple, typical examples
- `.edge.` - Edge cases, boundary conditions
- `.error.` - Error scenarios (future)

## Example Quality Standards

### Required Fields
Every example must include:
- ✅ **Valid schema field** - Must match a defined schema constant
- ✅ **All required fields** - Per schema definition
- ✅ **Realistic data** - Representative of actual venue data
- ✅ **Proper formatting** - Valid JSON with consistent indentation

### Data Realism Requirements

**Instrument IDs:**
- Use realistic canonical format: `pm_us_election_2028_winner`
- Follow conventions from `docs/mapping.md`
- Be consistent across related examples

**Timestamps:**
- Use realistic future timestamps in milliseconds
- Maintain logical ordering in related examples
- Use consistent base time for example families

**Probabilities:**
- Stay within [0.0, 1.0] bounds
- Use realistic values (avoid 0.0, 1.0 except for edge cases)
- Maintain market consistency (complementary probabilities)

**Venue Data:**
- Reflect actual venue data formats and ranges
- Use realistic market sizes and prices
- Include venue-specific identifiers

## Validation Process

### Automated Validation

The CI system automatically validates all examples via:

```bash
npm run validate-examples
```

This checks:
1. **JSON syntax** - Valid JSON structure
2. **Schema compliance** - All required fields present
3. **Field types** - Basic type checking
4. **Schema references** - Valid schema field values

### Manual Validation

Before adding new examples:

1. **Review against schema** - Ensure all required fields present
2. **Check data realism** - Use realistic values and formats
3. **Test validation** - Run `npm run validate-examples` locally
4. **Cross-reference** - Check consistency with related examples

## Adding New Examples

### When to Add Examples

**Required:**
- 📋 New schema created → Add at least one basic example
- 🔄 Schema changes → Update affected examples
- 🐛 Bug reports → Add examples demonstrating the issue

**Recommended:**
- 🎯 New use cases → Add scenario-specific examples
- 🏢 New venues → Add venue-specific examples
- 📊 Edge cases → Add boundary condition examples

### Step-by-Step Process

1. **Identify the schema** that needs examples
2. **Choose scenario name** following naming conventions
3. **Create example file** in `schemas/examples/`
4. **Use realistic data** following quality standards
5. **Validate locally** with `npm run validate-examples`
6. **Test CI pipeline** - commit and check CI passes
7. **Update documentation** if needed

### Example Template

```json
{
  "schema": "{exact_schema_constant}",
  "instrument_id": "realistic_canonical_id",
  "venue_id": "polymarket_or_kalshi",
  "ts_ms": 1758763048123,
  // ... other required fields with realistic values
}
```

## Maintaining Examples

### Regular Maintenance Tasks

**Monthly Review:**
- [ ] Verify all examples still validate
- [ ] Check for schema changes requiring example updates
- [ ] Review example realism against current market data
- [ ] Update timestamps to stay reasonably current

**When Schemas Change:**
- [ ] Update affected examples immediately
- [ ] Test validation passes
- [ ] Consider adding examples for new fields
- [ ] Update documentation if conventions change

**Quality Audits:**
- [ ] Check instrument_id consistency across examples
- [ ] Verify venue enum usage matches registry
- [ ] Review probability ranges for realism
- [ ] Ensure timestamp progression is logical

### Common Maintenance Issues

**Schema Evolution:**
- ✅ **New optional fields** → May add to examples for completeness
- ⚠️ **New required fields** → Must update all affected examples
- ❌ **Removed fields** → Must remove from examples (breaking change)

**Data Staleness:**
- Update timestamps periodically (every 6 months)
- Refresh market references to current events
- Update pricing to reflect realistic ranges

**Validation Failures:**
- Usually indicate schema changes or bugs
- Fix examples to match schema requirements
- Never modify schemas to match examples

## Example Coverage Goals

### Current Status: ✅ Excellent (162% coverage ratio)
- 8 schemas with 13 examples
- All schemas have at least 1 example
- Critical schemas have multiple scenario coverage

### Coverage Targets

**Minimum (Phase 1):** ✅ Achieved
- Every schema has at least 1 basic example
- Examples validate successfully
- Examples use realistic data

**Good (Current):** ✅ Achieved
- Key schemas have multiple scenarios
- Both venues represented where applicable
- All enum states covered (health status, trade sides)

**Excellent (Future):**
- Edge case examples for boundary conditions
- Error scenario examples
- Performance test data sets
- Cross-venue arbitrage examples

## Troubleshooting

### Validation Failures

**"Schema not found" error:**
- Check schema field matches exact schema constant
- Verify schema file exists in `schemas/json/`
- Ensure schema constant is defined correctly

**"Missing required field" error:**
- Compare example against schema `required` array
- Add missing fields with appropriate values
- Check for typos in field names

**"Invalid venue" error:**
- Verify venue_id matches `venues.json` registry
- Check for typos in venue names
- Ensure new venues are added to registry first

### Example Development Tips

**Data Generation:**
- Use consistent base timestamps across examples
- Generate realistic probabilities using market research
- Reference actual venue data structures for authenticity

**Schema Changes:**
- Always update examples when schemas change
- Use git to track which examples need updates
- Test validation before committing changes

**Performance:**
- Keep examples focused and minimal
- Avoid large arrays unless testing array handling
- Use representative but not exhaustive data

## Integration with CI/CD

### Automated Checks

The CI pipeline runs comprehensive example validation:

```yaml
# In .github/workflows/validate.yml
- name: Validate examples against schemas
  run: npm run validate-examples

- name: Check venue registry consistency
  run: node scripts/check-venues.js
```

### Failure Handling

When CI fails on example validation:
1. **Check error messages** for specific validation failures
2. **Fix examples locally** and test with `npm run validate-examples`
3. **Commit fixes** and verify CI passes
4. **Never skip validation** - examples must always be valid

## Future Enhancements

### Phase 2+ Roadmap

**Enhanced Validation:**
- JSON Schema draft 2020-12 full validation
- Cross-example consistency checking
- Data quality metrics (realistic ranges)
- Automated example generation from schemas

**Expanded Coverage:**
- Multi-venue arbitrage examples
- Complex market scenarios
- Historical data examples
- Performance test datasets

**Documentation Integration:**
- Auto-generated example documentation
- Interactive example browser
- Schema-example cross-references
- Example-driven API documentation

**Quality Automation:**
- Automated example updates for schema changes
- Example freshness monitoring
- Realistic data validation rules
- Example coverage reporting