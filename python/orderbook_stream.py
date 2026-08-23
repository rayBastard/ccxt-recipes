# Stream a live order book over WebSocket and print the top of book in place.
# Usage: python orderbook_stream.py [SYMBOL]   (Ctrl+C to stop)
import asyncio
import sys
import ccxt.pro as ccxtpro


async def main():
    symbol = sys.argv[1] if len(sys.argv) > 1 else 'BTC/USDT'
    exchange = ccxtpro.binance()
    try:
        while True:
            orderbook = await exchange.watch_order_book(symbol, 10)
            bid = orderbook['bids'][0] if orderbook['bids'] else (None, None)
            ask = orderbook['asks'][0] if orderbook['asks'] else (None, None)
            print(f"\r{symbol}  bid {bid[0]} ({bid[1]})  ask {ask[0]} ({ask[1]})   ", end='', flush=True)
    finally:
        await exchange.close()


try:
    asyncio.run(main())
except KeyboardInterrupt:
    pass
