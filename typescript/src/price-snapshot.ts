// Compare one symbol's price across several exchanges, fetched in parallel.
// Usage: npm run price-snapshot [-- SYMBOL]
import ccxt, { Exchange } from 'ccxt';

const symbol = process.argv[2] ?? 'BTC/USDT';
const ids = ['binance', 'kraken', 'okx'];

const rows = await Promise.all(ids.map(async (id) => {
    const exchange: Exchange = new (ccxt as any)[id]();
    const ticker = await exchange.fetchTicker(symbol);
    return { id, last: ticker.last, bid: ticker.bid, ask: ticker.ask };
}));

console.log(symbol);
for (const { id, last, bid, ask } of rows) {
    const spread = (bid && ask) ? (((ask - bid) / ask) * 100).toFixed(4) + '%' : 'n/a';
    console.log(`${id.padEnd(8)} last ${last}  bid ${bid}  ask ${ask}  spread ${spread}`);
}
