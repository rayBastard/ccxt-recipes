<?php
// Download paginated OHLCV candles and write them to a CSV file.
// Usage: php ohlcv_to_csv.php [SYMBOL TIMEFRAME DAYS]
require __DIR__ . '/vendor/autoload.php';

$symbol = $argv[1] ?? 'BTC/USDT';
$timeframe = $argv[2] ?? '1h';
$days = (int) ($argv[3] ?? 7);

$exchange = new \ccxt\binance();
$cursor = $exchange->milliseconds() - $days * 86400 * 1000;
$candles = [];
while (true) {
    $batch = $exchange->fetch_ohlcv($symbol, $timeframe, $cursor, 1000);
    if (count($batch) === 0) {
        break;
    }
    $candles = array_merge($candles, $batch);
    $lastTimestamp = $batch[count($batch) - 1][0];
    if (count($batch) < 1000 || $lastTimestamp <= $cursor) {
        break; // reached the newest candle
    }
    $cursor = $lastTimestamp + 1;
}

$filename = str_replace(['/', ':'], '-', $symbol) . "-{$timeframe}.csv";
$file = fopen($filename, 'w');
fputcsv($file, ['timestamp', 'datetime', 'open', 'high', 'low', 'close', 'volume']);
foreach ($candles as $c) {
    fputcsv($file, [$c[0], $exchange->iso8601($c[0]), $c[1], $c[2], $c[3], $c[4], $c[5]]);
}
fclose($file);
echo count($candles), " candles -> {$filename}\n";
