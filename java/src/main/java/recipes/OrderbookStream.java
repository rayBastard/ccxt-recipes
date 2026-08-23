// Stream a live order book over WebSocket and print the top of book in place.
package recipes;

import io.github.ccxt.exchanges.pro.Binance;
import io.github.ccxt.types.OrderBook;

import java.util.List;

public class OrderbookStream {

    public static void main(String[] args) {
        String symbol = args.length > 0 ? args[0] : "BTC/USDT";
        Binance exchange = new Binance();
        while (true) {
            OrderBook orderbook = exchange.watchOrderBook(symbol, 10L);
            System.out.printf("\r%s  bid %s  ask %s   ", symbol, top(orderbook.bids), top(orderbook.asks));
        }
    }

    static String top(List<List<Double>> side) {
        if (side == null || side.isEmpty()) {
            return "n/a";
        }
        return side.get(0).get(0) + " (" + side.get(0).get(1) + ")";
    }
}
