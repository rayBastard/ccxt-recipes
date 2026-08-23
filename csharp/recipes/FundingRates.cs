// Compare perpetual-swap funding rates across derivatives exchanges.
using ccxt;

namespace Recipes;

public static class FundingRates
{
    public static async Task Run(string[] args)
    {
        var symbol = args.Length > 0 ? args[0] : "BTC/USDT:USDT";
        await Print("binance", new Binance().FetchFundingRate(symbol));
        await Print("bybit", new Bybit().FetchFundingRate(symbol));
        await Print("okx", new Okx().FetchFundingRate(symbol));
    }

    private static async Task Print(string id, Task<FundingRate> request)
    {
        try
        {
            var rate = await request;
            var pct = (rate.fundingRate != null) ? $"{rate.fundingRate * 100:F4}%" : "n/a";
            Console.WriteLine($"{id,-8} funding {pct}  next {rate.fundingDatetime ?? "n/a"}");
        }
        catch (Exception e)
        {
            Console.WriteLine($"{id,-8} error: {e.Message}");
        }
    }
}
