using Recipes;

var recipe = args.Length > 0 ? args[0] : "price-snapshot";
var rest = args.Skip(1).ToArray();

switch (recipe)
{
    case "price-snapshot": await PriceSnapshot.Run(rest); break;
    case "ohlcv-to-csv": await OhlcvToCsv.Run(rest); break;
    case "orderbook-stream": await OrderbookStream.Run(rest); break;
    case "funding-rates": await FundingRates.Run(rest); break;
    case "sandbox-order": await SandboxOrder.Run(rest); break;
    default: Console.WriteLine($"unknown recipe: {recipe}"); break;
}
