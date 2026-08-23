// Stream a live order book over WebSocket and print the top of book in place.
// Usage: npm run orderbook-stream [-- SYMBOL]   (Ctrl+C to stop)
import ccxt from 'ccxt';

const symbol = process.argv[2] ?? 'BTC/USDT';
const exchange = new ccxt.pro.binance();

process.on('SIGINT', () => {
    exchange.close().then(() => process.exit(0));
});

while (true) {
    const orderbook = await exchange.watchOrderBook(symbol, 10);
    const [bidPrice, bidSize] = orderbook.bids[0] ?? [];
    const [askPrice, askSize] = orderbook.asks[0] ?? [];
    process.stdout.write(`\r${symbol}  bid ${bidPrice} (${bidSize})  ask ${askPrice} (${askSize})   `);
}
