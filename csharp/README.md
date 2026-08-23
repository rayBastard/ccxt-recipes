# CCXT Recipes — C#

```bash
dotnet run -- price-snapshot            # BTC/USDT across binance, kraken, okx
dotnet run -- price-snapshot ETH/USDT
dotnet run -- ohlcv-to-csv BTC/USDT 1h 7
dotnet run -- orderbook-stream          # live top of book, Ctrl+C to stop
dotnet run -- funding-rates             # BTC/USDT:USDT perp funding
CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=... dotnet run -- sandbox-order
```
