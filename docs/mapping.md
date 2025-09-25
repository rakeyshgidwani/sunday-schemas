# Raw → Normalized Data Mapping

This document describes the transformation rules used by the `sunday-data/normalizer` service to convert raw venue data into normalized market data events.

## Overview

The Sunday platform ingests raw market data from multiple prediction markets (venues) and normalizes it into a consistent format for downstream analysis. This process involves:

1. **Instrument ID Canonicalization** - Map venue-specific identifiers to canonical `instrument_id` values
2. **Price Normalization** - Convert venue-specific pricing to implied probabilities [0.0, 1.0]
3. **Data Structure Standardization** - Transform venue-specific schemas to Sunday event schemas
4. **Metadata Enhancement** - Add timing, sequencing, and quality metadata

## Venue-Specific Mappings

### Polymarket

**Instrument ID Mapping:**
- **Source**: Polymarket market token addresses (e.g., `0x12345abcdef...`)
- **Target**: Canonical format `pm_{category}_{descriptor}_{outcome}`
- **Examples**:
  - `0x1234...` → `pm_us_election_2028_winner`
  - `0x5678...` → `pm_crypto_btc_100k_2025`
  - `0x9abc...` → `pm_fed_rate_cut_dec_2025`

**Price Conversion:**
- **Source**: Polymarket prices are already implied probabilities [0.0, 1.0]
- **Target**: No conversion needed, direct mapping
- **Validation**: Ensure prices are within [0.0, 1.0] bounds

**YES/NO Convention:**
- **YES outcomes**: Probability represents likelihood of event occurring
- **NO outcomes**: Probability = 1.0 - YES probability
- **Single-sided markets**: Only YES side is normalized (NO implied)

### Kalshi

**Instrument ID Mapping:**
- **Source**: Kalshi market tickers (e.g., `PRES-28`, `FEDR-25DEC`)
- **Target**: Canonical format `kalshi_{category}_{descriptor}`
- **Examples**:
  - `PRES-28` → `kalshi_pres_28_winner`
  - `FEDR-25DEC` → `kalshi_fed_rate_cut_dec_2025`
  - `BTC-100K` → `kalshi_crypto_btc_100k_2025`

**Price Conversion:**
- **Source**: Kalshi prices in cents (0-100, where 100 = $1.00)
- **Target**: Convert to implied probabilities: `prob = price_cents / 100.0`
- **Examples**:
  - 63 cents → 0.63 probability
  - 37 cents → 0.37 probability

**YES/NO Convention:**
- **YES markets**: Direct probability mapping
- **NO markets**: Invert probability: `prob = 1.0 - (price_cents / 100.0)`

## Cross-Venue Instrument Mapping

For markets that exist on multiple venues, the normalizer creates a single canonical `instrument_id`:

```json
{
  "pm_us_election_2028_winner": {
    "polymarket": "0x12345abcdef...",
    "kalshi": "PRES-28",
    "title": "US Presidential Election 2028 Winner",
    "category": "politics",
    "resolution_source": "Associated Press"
  }
}
```

This mapping is maintained in `/schemas/registries/instruments.json` and populated by the normalizer service.

## Data Transformation Rules

### Raw → Normalized Orderbook

**Input**: `raw.v0` envelope with venue-specific orderbook payload
**Output**: `md.orderbook.delta.v1` event

**Transformation Steps:**
1. **Extract timing**: Use `ts_event_ms` from raw envelope
2. **Map instrument**: Convert `instrument_native` to canonical `instrument_id`
3. **Normalize prices**: Apply venue-specific price conversion to [0.0, 1.0]
4. **Structure depth**: Convert to `[[price, size], ...]` array format
5. **Add sequencing**: Generate monotonic `seq` per `(instrument_id, venue_id)`
6. **Detect gaps**: Set `is_snapshot=true` when sequence gaps detected

**Example Transformation:**
```javascript
// Raw Polymarket data
{
  "schema": "raw.v0",
  "venue_id": "polymarket",
  "instrument_native": "0x12345...",
  "payload": {
    "bids": [{"price": "0.63", "size": "1000"}],
    "asks": [{"price": "0.64", "size": "800"}]
  }
}

// Normalized output
{
  "schema": "md.orderbook.delta.v1",
  "instrument_id": "pm_us_election_2028_winner",
  "venue_id": "polymarket",
  "seq": 12345,
  "bids": [[0.63, 1000.0]],
  "asks": [[0.64, 800.0]],
  "is_snapshot": false
}
```

### Raw → Normalized Trades

**Input**: `raw.v0` envelope with venue-specific trade payload
**Output**: `md.trade.v1` event

**Transformation Steps:**
1. **Extract timing**: Use `ts_event_ms` from raw envelope
2. **Map instrument**: Convert `instrument_native` to canonical `instrument_id`
3. **Normalize price**: Apply venue conversion to probability [0.0, 1.0]
4. **Standardize side**: Convert to "buy"/"sell" from taker perspective
5. **Calculate notional**: Compute USD value where possible

**Venue-Specific Trade Mapping:**

| Field | Polymarket | Kalshi |
|-------|------------|---------|
| Price | Direct probability | `price_cents / 100.0` |
| Size | Share count | Contract count |
| Side | "buy"/"sell" | Map "yes"→"buy", "no"→"sell" |
| Notional | `price × size` | `(price_cents/100) × size × $1` |

## Probability Conventions

### Standard YES/NO Markets

**YES Outcome Probability:**
- Range: [0.0, 1.0] where 1.0 = 100% certainty
- Interpretation: Likelihood that the stated event will occur
- Examples:
  - 0.63 = 63% chance event happens
  - 0.37 = 37% chance event happens

**NO Outcome (when separate):**
- Probability = 1.0 - YES probability
- Usually not separately tracked (implied)

### Multi-Outcome Markets

For markets with >2 outcomes, probabilities should sum to ≤ 1.0:

```json
{
  "democrat_winner": 0.45,
  "republican_winner": 0.42,
  "third_party_winner": 0.08,
  "other": 0.05
  // Sum = 1.00
}
```

### Probability Validation Rules

1. **Range Check**: All probabilities must be in [0.0, 1.0]
2. **Precision**: Round to 6 decimal places maximum
3. **Boundary Handling**:
   - 0.0 = impossible (but tradeable)
   - 1.0 = certain (trading should cease)
4. **Sum Validation**: Multi-outcome probabilities ≤ 1.0

## Instrument ID Conventions

### Format Standards

**General Pattern**: `{venue_prefix}_{category}_{descriptor}_{outcome?}`

**Venue Prefixes:**
- `pm_` = Polymarket markets
- `kalshi_` = Kalshi markets
- `manifold_` = Manifold markets (future)

**Categories:**
- `us_election` = US political elections
- `crypto` = Cryptocurrency price predictions
- `fed` = Federal Reserve policy decisions
- `weather` = Weather and climate events
- `sports` = Sports outcomes
- `econ` = Economic indicators

**Descriptors:**
- Use lowercase with underscores
- Include year for time-sensitive events
- Be specific but concise

**Examples:**
```
pm_us_election_2028_winner
pm_crypto_btc_100k_2025
kalshi_fed_rate_cut_dec_2025
pm_weather_hurricane_season_2024
kalshi_sports_superbowl_winner_2025
```

### Cross-Venue Canonicalization

When the same market exists on multiple venues:

1. **Choose primary venue** based on liquidity/volume
2. **Use primary venue prefix** for canonical ID
3. **Map all venues** to same canonical ID in instruments registry
4. **Maintain consistency** across all event schemas

## Sequence Number Management

### Per-Instrument Sequencing

Each `(instrument_id, venue_id)` pair maintains independent sequence numbers:

```javascript
// Example sequence tracking
{
  "pm_us_election_2028_winner": {
    "polymarket": 12345,
    "kalshi": 8901
  },
  "pm_crypto_btc_100k_2025": {
    "polymarket": 5432
  }
}
```

### Gap Detection & Recovery

**Gap Detection:**
- Compare incoming `seq` with `last_seq + 1`
- If gap detected, request snapshot from venue
- Mark next message with `is_snapshot=true`

**Recovery Process:**
1. Detect sequence gap
2. Request full orderbook snapshot
3. Emit snapshot event with `is_snapshot=true`
4. Resume normal delta processing
5. Validate sequence continuity

## Error Handling

### Invalid Data

**Price Out of Bounds:**
- Log error, drop event
- Alert monitoring system
- Request fresh snapshot

**Unknown Instrument:**
- Log warning
- Attempt fuzzy matching
- Add to manual review queue

**Malformed Payload:**
- Log error with full context
- Drop event, continue processing
- Track error rates by venue

### Venue Connectivity

**Connection Loss:**
- Mark venue as STALE in health monitoring
- Continue processing other venues
- Attempt reconnection with backoff

**Data Staleness:**
- Compare `ts_event_ms` with current time
- Mark events >5min old as stale
- Emit venue health updates

## Validation & Quality Assurance

### Real-time Validation

**Price Sanity Checks:**
- Probability bounds [0.0, 1.0]
- No dramatic price jumps (>50% in 1 minute)
- Cross-venue arbitrage bounds (<500 bps difference)

**Sequence Validation:**
- Monotonic sequence numbers
- No duplicate sequences
- Timely gap detection

**Volume Validation:**
- Reasonable trade sizes
- Volume spikes detection
- Wash trading patterns

### Data Quality Metrics

Track and alert on:
- Message processing latency
- Invalid event percentage
- Sequence gap frequency
- Cross-venue price divergence
- Venue uptime and staleness

## Future Enhancements

### Phase 2+ Roadmap

1. **Additional Venues**: Manifold, Augur, others
2. **Complex Instruments**: Multi-outcome, conditional markets
3. **Advanced Pricing**: Options pricing models, volatility surfaces
4. **ML Enhancement**: Automated instrument matching, anomaly detection
5. **Real-time Arbitrage**: Cross-venue opportunity detection and routing