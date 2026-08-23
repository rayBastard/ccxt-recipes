// Compare perpetual-swap funding rates across derivatives exchanges.
// Usage: npm run funding-rates [-- SYMBOL]
import ccxt, { Exchange } from 'ccxt';

const symbol = process.argv[2] ?? 'BTC/USDT:USDT';
const ids = ['binance', 'bybit', 'okx'];

for (const id of ids) {
    const exchange: Exchange = new (ccxt as any)[id]();
    const rate = await exchange.fetchFundingRate(symbol);
    const pct = (rate.fundingRate !== undefined) ? (rate.fundingRate * 100).toFixed(4) + '%' : 'n/a';
    console.log(`${id.padEnd(8)} funding ${pct}  next ${rate.fundingDatetime ?? 'n/a'}`);
}
