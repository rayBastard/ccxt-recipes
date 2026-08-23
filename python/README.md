# CCXT Recipes — Python

```bash
pip install -r requirements.txt
python price_snapshot.py            # BTC/USDT across binance, kraken, okx
python price_snapshot.py ETH/USDT
python ohlcv_to_csv.py BTC/USDT 1h 7
python orderbook_stream.py          # live top of book, Ctrl+C to stop
python funding_rates.py             # BTC/USDT:USDT perp funding
CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=... python sandbox_order.py
```
