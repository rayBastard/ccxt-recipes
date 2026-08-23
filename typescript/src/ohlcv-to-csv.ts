// Download paginated OHLCV candles and write them to a CSV file.
// Usage: npm run ohlcv-to-csv [-- SYMBOL TIMEFRAME DAYS]
import ccxt, { OHLCV } from 'ccxt';
import { writeFileSync } from 'fs';

const [symbol = 'BTC/USDT', timeframe = '1h', daysArg = '7'] = process.argv.slice(2);
const days = Number(daysArg);

const exchange = new ccxt.binance();
const all: OHLCV[] = [];
let cursor = exchange.milliseconds() - days * 86400 * 1000;
while (true) {
    const batch = await exchange.fetchOHLCV(symbol, timeframe, cursor, 1000);
    if (!batch.length) {
        break;
    }
    all.push(...batch);
    const lastTimestamp = batch[batch.length - 1][0]!;
    if (batch.length < 1000 || lastTimestamp <= cursor) {
        break; // reached the newest candle
    }
    cursor = lastTimestamp + 1;
}

const file = `${symbol.replace('/', '-').replace(':', '-')}-${timeframe}.csv`;
const header = 'timestamp,datetime,open,high,low,close,volume\n';
const lines = all.map((c) => [c[0], exchange.iso8601(c[0]), c[1], c[2], c[3], c[4], c[5]].join(','));
writeFileSync(file, header + lines.join('\n') + '\n');
console.log(`${all.length} candles -> ${file}`);
