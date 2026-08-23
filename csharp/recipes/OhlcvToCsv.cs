// Download paginated OHLCV candles and write them to a CSV file.
using System.Text;
using ccxt;

namespace Recipes;

public static class OhlcvToCsv
{
    public static async Task Run(string[] args)
    {
        var symbol = args.Length > 0 ? args[0] : "BTC/USDT";
        var timeframe = args.Length > 1 ? args[1] : "1h";
        var days = args.Length > 2 ? int.Parse(args[2]) : 7;

        var exchange = new Binance();
        var cursor = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds() - (long)days * 86400 * 1000;
        var candles = new List<OHLCV>();
        while (true)
        {
            var batch = await exchange.FetchOHLCV(symbol, timeframe, cursor, 1000);
            if (batch.Count == 0)
            {
                break;
            }
            candles.AddRange(batch);
            var lastTimestamp = (long)batch[^1].timestamp!;
            if (batch.Count < 1000 || lastTimestamp <= cursor)
            {
                break; // reached the newest candle
            }
            cursor = lastTimestamp + 1;
        }

        var filename = $"{symbol.Replace("/", "-").Replace(":", "-")}-{timeframe}.csv";
        var sb = new StringBuilder("timestamp,datetime,open,high,low,close,volume\n");
        foreach (var c in candles)
        {
            var datetime = DateTimeOffset.FromUnixTimeMilliseconds((long)c.timestamp!).UtcDateTime.ToString("o");
            sb.AppendLine($"{c.timestamp},{datetime},{c.open},{c.high},{c.low},{c.close},{c.volume}");
        }
        File.WriteAllText(filename, sb.ToString());
        Console.WriteLine($"{candles.Count} candles -> {filename}");
    }
}
