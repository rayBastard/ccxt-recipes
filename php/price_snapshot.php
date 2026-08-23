<?php
// Compare one symbol's price across several exchanges.
// Usage: php price_snapshot.php [SYMBOL]
require __DIR__ . '/vendor/autoload.php';

$symbol = $argv[1] ?? 'BTC/USDT';

echo $symbol, "\n";
foreach (['binance', 'kraken', 'okx'] as $id) {
    $class = "\\ccxt\\{$id}";
    $exchange = new $class();
    $ticker = $exchange->fetch_ticker($symbol);
    $bid = $ticker['bid'];
    $ask = $ticker['ask'];
    $spread = ($bid && $ask) ? number_format(($ask - $bid) / $ask * 100, 4) . '%' : 'n/a';
    printf("%-8s last %s  bid %s  ask %s  spread %s\n", $id, $ticker['last'], $bid, $ask, $spread);
}
