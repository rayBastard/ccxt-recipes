// Compare one symbol's price across several exchanges.
package recipes;

import io.github.ccxt.exchanges.Binance;
import io.github.ccxt.exchanges.Kraken;
import io.github.ccxt.exchanges.Okx;
import io.github.ccxt.types.Ticker;

public class PriceSnapshot {

    public static void main(String[] args) {
        String symbol = args.length > 0 ? args[0] : "BTC/USDT";
        System.out.println(symbol);
        print("binance", new Binance().fetchTicker(symbol));
        print("kraken", new Kraken().fetchTicker(symbol));
        print("okx", new Okx().fetchTicker(symbol));
    }

    static void print(String id, Ticker ticker) {
        String spread = (ticker.bid != null && ticker.ask != null)
            ? String.format("%.4f%%", (ticker.ask - ticker.bid) / ticker.ask * 100)
            : "n/a";
        System.out.printf("%-8s last %s  bid %s  ask %s  spread %s%n", id, ticker.last, ticker.bid, ticker.ask, spread);
    }
}
