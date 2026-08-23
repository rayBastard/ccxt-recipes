# Compare one symbol's price across several exchanges.
# Usage: python price_snapshot.py [SYMBOL]
import sys
import ccxt

symbol = sys.argv[1] if len(sys.argv) > 1 else 'BTC/USDT'

print(symbol)
for exchange_id in ['binance', 'kraken', 'okx']:
    exchange = getattr(ccxt, exchange_id)()
    ticker = exchange.fetch_ticker(symbol)
    bid, ask = ticker['bid'], ticker['ask']
    spread = f"{(ask - bid) / ask * 100:.4f}%" if bid and ask else 'n/a'
    print(f"{exchange_id:<8} last {ticker['last']}  bid {bid}  ask {ask}  spread {spread}")
