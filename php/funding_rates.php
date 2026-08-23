<?php
// Compare perpetual-swap funding rates across derivatives exchanges.
// Usage: php funding_rates.php [SYMBOL]
require __DIR__ . '/vendor/autoload.php';

$symbol = $argv[1] ?? 'BTC/USDT:USDT';

foreach (['binance', 'bybit', 'okx'] as $id) {
    $class = "\\ccxt\\{$id}";
    $exchange = new $class();
    $rate = $exchange->fetch_funding_rate($symbol);
    $funding = $rate['fundingRate'];
    $pct = ($funding !== null) ? number_format($funding * 100, 4) . '%' : 'n/a';
    printf("%-8s funding %s  next %s\n", $id, $pct, $rate['fundingDatetime'] ?? 'n/a');
}
