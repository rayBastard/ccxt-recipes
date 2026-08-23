<?php
// Stream a live order book over WebSocket and print the top of book in place.
// Usage: php orderbook_stream.php [SYMBOL]   (Ctrl+C to stop)
require __DIR__ . '/vendor/autoload.php';

use React\Async;

$symbol = $argv[1] ?? 'BTC/USDT';

function stream($symbol) {
    return Async\async(function () use ($symbol) {
        $exchange = new \ccxt\pro\binance();
        while (true) {
            $orderbook = Async\await($exchange->watch_order_book($symbol, 10));
            $bid = $orderbook['bids'][0] ?? [null, null];
            $ask = $orderbook['asks'][0] ?? [null, null];
            printf("\r%s  bid %s (%s)  ask %s (%s)   ", $symbol, $bid[0], $bid[1], $ask[0], $ask[1]);
        }
    })();
}

Async\await(stream($symbol));
