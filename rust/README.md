# CCXT Recipes — Rust

```bash
cargo run --bin price_snapshot            # BTC/USDT across binance, kraken, okx
cargo run --bin price_snapshot ETH/USDT
cargo run --bin ohlcv_to_csv BTC/USDT 1h 7
cargo run --bin orderbook_stream          # live top of book, Ctrl+C to stop
cargo run --bin funding_rates             # BTC/USDT:USDT perp funding
CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=... cargo run --bin sandbox_order
```
