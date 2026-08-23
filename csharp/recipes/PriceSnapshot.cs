// Compare one symbol's price across several exchanges.
using ccxt;

namespace Recipes;

public static class PriceSnapshot
{
    public static async Task Run(string[] args)
    {
        var symbol = args.Length > 0 ? args[0] : "BTC/USDT";
        var exchanges = new List<Exchange> { new Binance(), new Kraken(), new Okx() };
        Console.WriteLine(symbol);
        foreach (var exchange in exchanges)
        {
            var ticker = await exchange.FetchTicker(symbol);
            var spread = (ticker.bid != null && ticker.ask != null)
                ? $"{(ticker.ask - ticker.bid) / ticker.ask * 100:F4}%"
                : "n/a";
            Console.WriteLine($"{exchange.id,-8} last {ticker.last}  bid {ticker.bid}  ask {ticker.ask}  spread {spread}");
        }
    }
}
