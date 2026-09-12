# 🍳 CCXT Recipes

Practical, runnable recipes for [CCXT](https://github.com/ccxt/ccxt) — the same task solved in **TypeScript, Python, PHP, C#, Go, Java and Rust**.

CCXT is one TypeScript codebase transpiled to seven languages, so the unified API looks and behaves the same everywhere. These recipes show exactly how the idioms map from one language to another.

Maintained by [Roman Cuhari](https://github.com/rayBastard), CCXT core contributor. Questions about CCXT? Ping me on [Telegram](https://t.me/raybastard22).

## Recipes

| Recipe | What it demonstrates |
|---|---|
| **price-snapshot** | one symbol's ticker fetched from three exchanges, with spread comparison |
| **ohlcv-to-csv** | paginated OHLCV (candlestick) download → CSV file |
| **orderbook-stream** | live order book over WebSocket (`ccxt.pro`), top-of-book printed in place |
| **funding-rates** | perpetual-swap funding rates compared across derivatives exchanges |
| **sandbox-order** | full order lifecycle on the exchange **testnet**: create → inspect → cancel |

Every recipe exists in every language:

| Language | Directory | Run example |
|---|---|---|
| TypeScript | [`typescript/`](typescript/) | `npm install && npm run price-snapshot` |
| Python | [`python/`](python/) | `pip install -r requirements.txt && python price_snapshot.py` |
| PHP | [`php/`](php/) | `composer install && php price_snapshot.php` |
| C# | [`csharp/`](csharp/) | `dotnet run -- price-snapshot` |
| Go | [`go/`](go/) | `go mod tidy && go run ./cmd/price-snapshot` |
| Java | [`java/`](java/) | `gradle run -PmainClass=recipes.PriceSnapshot` |
| Rust | [`rust/`](rust/) | `cargo run --bin price_snapshot` |

All recipes accept an optional symbol argument, e.g. `npm run price-snapshot -- ETH/USDT`.

## Safety

- Four of the five recipes use **public endpoints only** — no API keys needed.
- `sandbox-order` talks to the exchange **testnet** (`setSandboxMode(true)`) and expects **testnet keys** in `CCXT_SANDBOX_APIKEY` / `CCXT_SANDBOX_SECRET`. Get free Binance testnet keys at <https://testnet.binance.vision>. Never put real trading keys into these variables.

## License

[MIT](LICENSE)
