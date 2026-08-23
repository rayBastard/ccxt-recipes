// Full order lifecycle on the exchange testnet: create -> inspect -> cancel.
// Needs Binance *testnet* keys (https://testnet.binance.vision):
//   export CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=...
import ccxt from 'ccxt';

const apiKey = process.env['CCXT_SANDBOX_APIKEY'];
const secret = process.env['CCXT_SANDBOX_SECRET'];
if (!apiKey || !secret) {
    console.error('Set CCXT_SANDBOX_APIKEY and CCXT_SANDBOX_SECRET (testnet keys only!)');
    process.exit(1);
}

const symbol = 'BTC/USDT';
const exchange = new ccxt.binance({ apiKey, secret });
exchange.setSandboxMode(true);
await exchange.loadMarkets();

const ticker = await exchange.fetchTicker(symbol);
// bid far below the market so the order rests instead of filling
const price = Number(exchange.priceToPrecision(symbol, ticker.last! * 0.8));
const amount = Number(exchange.amountToPrecision(symbol, 20 / price)); // ~20 USDT notional

const order = await exchange.createLimitBuyOrder(symbol, amount, price);
console.log(`created ${order.id}: ${amount} ${symbol} @ ${price}`);
try {
    const open = await exchange.fetchOpenOrders(symbol);
    console.log('open orders:', open.map((o) => o.id));
} finally {
    const canceled = await exchange.cancelOrder(order.id, symbol);
    console.log(`canceled ${canceled.id}, status: ${canceled.status}`);
}
