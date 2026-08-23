# Compare perpetual-swap funding rates across derivatives exchanges.
# Usage: python funding_rates.py [SYMBOL]
import sys
import ccxt

symbol = sys.argv[1] if len(sys.argv) > 1 else 'BTC/USDT:USDT'

for exchange_id in ['binance', 'bybit', 'okx']:
    exchange = getattr(ccxt, exchange_id)()
    rate = exchange.fetch_funding_rate(symbol)
    funding = rate['fundingRate']
    pct = f"{funding * 100:.4f}%" if funding is not None else 'n/a'
    print(f"{exchange_id:<8} funding {pct}  next {rate['fundingDatetime'] or 'n/a'}")
