# CCXT Recipes — Java

Requires Java 21+ and Gradle.

```bash
gradle run -PmainClass=recipes.PriceSnapshot                       # BTC/USDT across binance, kraken, okx
gradle run -PmainClass=recipes.PriceSnapshot --args="ETH/USDT"
gradle run -PmainClass=recipes.OhlcvToCsv --args="BTC/USDT 1h 7"
gradle run -PmainClass=recipes.OrderbookStream                     # live top of book, Ctrl+C to stop
gradle run -PmainClass=recipes.FundingRates                        # BTC/USDT:USDT perp funding
CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=... gradle run -PmainClass=recipes.SandboxOrder
```
