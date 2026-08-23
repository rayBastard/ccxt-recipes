// Stream a live order book over WebSocket and print the top of book in place.
namespace Recipes;

public static class OrderbookStream
{
    public static async Task Run(string[] args)
    {
        var symbol = args.Length > 0 ? args[0] : "BTC/USDT";
        var exchange = new ccxt.pro.binance(new Dictionary<string, object>());
        while (true)
        {
            var orderbook = await exchange.WatchOrderBook(symbol, 10);
            var bid = orderbook.bids.Count > 0 ? orderbook.bids[0] as IList<object> : null;
            var ask = orderbook.asks.Count > 0 ? orderbook.asks[0] as IList<object> : null;
            Console.Write($"\r{symbol}  bid {bid?[0]} ({bid?[1]})  ask {ask?[0]} ({ask?[1]})   ");
        }
    }
}
