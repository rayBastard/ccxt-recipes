# CCXT Recipes — PHP

```bash
composer install
php price_snapshot.php            # BTC/USDT across binance, kraken, okx
php price_snapshot.php ETH/USDT
php ohlcv_to_csv.php BTC/USDT 1h 7
php orderbook_stream.php          # live top of book (async / ReactPHP), Ctrl+C to stop
php funding_rates.php             # BTC/USDT:USDT perp funding
CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=... php sandbox_order.php
```

> Note: parsing Binance's large `exchangeInfo` may exceed PHP's default 128M `memory_limit`.
> If you hit "Allowed memory size exhausted", run with `php -d memory_limit=512M <recipe>.php`.
