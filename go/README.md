# CCXT Recipes — Go

```bash
go mod tidy
go run ./cmd/price-snapshot            # BTC/USDT across binance, kraken, okx
go run ./cmd/price-snapshot ETH/USDT
go run ./cmd/ohlcv-to-csv BTC/USDT 1h 7
go run ./cmd/orderbook-stream          # live top of book, Ctrl+C to stop
go run ./cmd/funding-rates             # BTC/USDT:USDT perp funding
CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=... go run ./cmd/sandbox-order
```
