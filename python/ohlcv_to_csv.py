# Download paginated OHLCV candles and write them to a CSV file.
# Usage: python ohlcv_to_csv.py [SYMBOL TIMEFRAME DAYS]
import csv
import sys
import ccxt

symbol = sys.argv[1] if len(sys.argv) > 1 else 'BTC/USDT'
timeframe = sys.argv[2] if len(sys.argv) > 2 else '1h'
days = int(sys.argv[3]) if len(sys.argv) > 3 else 7

exchange = ccxt.binance()
cursor = exchange.milliseconds() - days * 86400 * 1000
candles = []
while True:
    batch = exchange.fetch_ohlcv(symbol, timeframe, cursor, 1000)
    if not batch:
        break
    candles += batch
    last_timestamp = batch[-1][0]
    if len(batch) < 1000 or last_timestamp <= cursor:
        break  # reached the newest candle
    cursor = last_timestamp + 1

filename = f"{symbol.replace('/', '-').replace(':', '-')}-{timeframe}.csv"
with open(filename, 'w', newline='') as f:
    writer = csv.writer(f)
    writer.writerow(['timestamp', 'datetime', 'open', 'high', 'low', 'close', 'volume'])
    for c in candles:
        writer.writerow([c[0], exchange.iso8601(c[0])] + c[1:])

print(f"{len(candles)} candles -> {filename}")
