// Download paginated OHLCV candles and write them to a CSV file.
package recipes;

import io.github.ccxt.exchanges.Binance;
import io.github.ccxt.types.OHLCV;

import java.io.PrintWriter;
import java.time.Instant;
import java.util.ArrayList;
import java.util.List;

public class OhlcvToCsv {

    public static void main(String[] args) throws Exception {
        String symbol = args.length > 0 ? args[0] : "BTC/USDT";
        String timeframe = args.length > 1 ? args[1] : "1h";
        int days = args.length > 2 ? Integer.parseInt(args[2]) : 7;

        Binance exchange = new Binance();
        long cursor = System.currentTimeMillis() - (long) days * 86400 * 1000;
        List<OHLCV> candles = new ArrayList<>();
        while (true) {
            List<OHLCV> batch = exchange.fetchOHLCV(symbol, timeframe, cursor, 1000L);
            if (batch.isEmpty()) {
                break;
            }
            candles.addAll(batch);
            long lastTimestamp = batch.get(batch.size() - 1).timestamp;
            if (batch.size() < 1000 || lastTimestamp <= cursor) {
                break; // reached the newest candle
            }
            cursor = lastTimestamp + 1;
        }

        String filename = symbol.replace("/", "-").replace(":", "-") + "-" + timeframe + ".csv";
        try (PrintWriter writer = new PrintWriter(filename)) {
            writer.println("timestamp,datetime,open,high,low,close,volume");
            for (OHLCV c : candles) {
                writer.printf("%d,%s,%s,%s,%s,%s,%s%n",
                    c.timestamp, Instant.ofEpochMilli(c.timestamp), c.open, c.high, c.low, c.close, c.volume);
            }
        }
        System.out.println(candles.size() + " candles -> " + filename);
    }
}
