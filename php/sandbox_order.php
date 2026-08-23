<?php
// Full order lifecycle on the exchange testnet: create -> inspect -> cancel.
// Needs Binance *testnet* keys (https://testnet.binance.vision):
//   export CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=...
require __DIR__ . '/vendor/autoload.php';

$apiKey = getenv('CCXT_SANDBOX_APIKEY');
$secret = getenv('CCXT_SANDBOX_SECRET');
if (!$apiKey || !$secret) {
    exit("Set CCXT_SANDBOX_APIKEY and CCXT_SANDBOX_SECRET (testnet keys only!)\n");
}

$symbol = 'BTC/USDT';
$exchange = new \ccxt\binance(['apiKey' => $apiKey, 'secret' => $secret]);
$exchange->set_sandbox_mode(true);
$exchange->load_markets();

$ticker = $exchange->fetch_ticker($symbol);
// bid far below the market so the order rests instead of filling
$price = (float) $exchange->price_to_precision($symbol, $ticker['last'] * 0.8);
$amount = (float) $exchange->amount_to_precision($symbol, 20 / $price); // ~20 USDT notional

$order = $exchange->create_limit_buy_order($symbol, $amount, $price);
echo "created {$order['id']}: {$amount} {$symbol} @ {$price}\n";
try {
    $open = $exchange->fetch_open_orders($symbol);
    echo 'open orders: ', implode(', ', array_map(fn ($o) => $o['id'], $open)), "\n";
} finally {
    $canceled = $exchange->cancel_order($order['id'], $symbol);
    echo "canceled {$canceled['id']}, status: {$canceled['status']}\n";
}
