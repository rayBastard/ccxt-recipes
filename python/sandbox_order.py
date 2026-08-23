# Full order lifecycle on the exchange testnet: create -> inspect -> cancel.
# Needs Binance *testnet* keys (https://testnet.binance.vision):
#   export CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=...
import os
import sys
import ccxt

api_key = os.environ.get('CCXT_SANDBOX_APIKEY')
secret = os.environ.get('CCXT_SANDBOX_SECRET')
if not api_key or not secret:
    sys.exit('Set CCXT_SANDBOX_APIKEY and CCXT_SANDBOX_SECRET (testnet keys only!)')

symbol = 'BTC/USDT'
exchange = ccxt.binance({'apiKey': api_key, 'secret': secret})
exchange.set_sandbox_mode(True)
exchange.load_markets()

ticker = exchange.fetch_ticker(symbol)
# bid far below the market so the order rests instead of filling
price = float(exchange.price_to_precision(symbol, ticker['last'] * 0.8))
amount = float(exchange.amount_to_precision(symbol, 20 / price))  # ~20 USDT notional

order = exchange.create_limit_buy_order(symbol, amount, price)
print(f"created {order['id']}: {amount} {symbol} @ {price}")
try:
    open_orders = exchange.fetch_open_orders(symbol)
    print('open orders:', [o['id'] for o in open_orders])
finally:
    canceled = exchange.cancel_order(order['id'], symbol)
    print(f"canceled {canceled['id']}, status: {canceled['status']}")
