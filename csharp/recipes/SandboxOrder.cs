// Full order lifecycle on the exchange testnet: create -> inspect -> cancel.
// Needs Binance *testnet* keys (https://testnet.binance.vision):
//   export CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=...
using ccxt;

namespace Recipes;

public static class SandboxOrder
{
    public static async Task Run(string[] args)
    {
        var apiKey = Environment.GetEnvironmentVariable("CCXT_SANDBOX_APIKEY");
        var secret = Environment.GetEnvironmentVariable("CCXT_SANDBOX_SECRET");
        if (string.IsNullOrEmpty(apiKey) || string.IsNullOrEmpty(secret))
        {
            Console.Error.WriteLine("Set CCXT_SANDBOX_APIKEY and CCXT_SANDBOX_SECRET (testnet keys only!)");
            return;
        }

        var symbol = "BTC/USDT";
        var exchange = new Binance(new Dictionary<string, object> { { "apiKey", apiKey }, { "secret", secret } });
        exchange.setSandboxMode(true);
        await exchange.LoadMarkets();

        var ticker = await exchange.FetchTicker(symbol);
        // bid far below the market so the order rests instead of filling
        var price = Convert.ToDouble(exchange.priceToPrecision(symbol, (double)ticker.last! * 0.8));
        var amount = Convert.ToDouble(exchange.amountToPrecision(symbol, 20 / price)); // ~20 USDT notional

        var order = await exchange.CreateOrder(symbol, "limit", "buy", amount, price);
        Console.WriteLine($"created {order.id}: {amount} {symbol} @ {price}");
        try
        {
            var open = await exchange.FetchOpenOrders(symbol);
            Console.WriteLine("open orders: " + string.Join(", ", open.Select(o => o.id)));
        }
        finally
        {
            var canceled = await exchange.CancelOrder(order.id, symbol);
            Console.WriteLine($"canceled {canceled.id}, status: {canceled.status}");
        }
    }
}
