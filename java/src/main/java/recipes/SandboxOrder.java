// Full order lifecycle on the exchange testnet: create -> inspect -> cancel.
// Needs Binance *testnet* keys (https://testnet.binance.vision):
//   export CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=...
package recipes;

import io.github.ccxt.exchanges.Binance;
import io.github.ccxt.types.Order;
import io.github.ccxt.types.Ticker;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class SandboxOrder {

    public static void main(String[] args) {
        String apiKey = System.getenv("CCXT_SANDBOX_APIKEY");
        String secret = System.getenv("CCXT_SANDBOX_SECRET");
        if (apiKey == null || secret == null) {
            System.err.println("Set CCXT_SANDBOX_APIKEY and CCXT_SANDBOX_SECRET (testnet keys only!)");
            System.exit(1);
        }

        String symbol = "BTC/USDT";
        Map<String, Object> config = new HashMap<>();
        config.put("apiKey", apiKey);
        config.put("secret", secret);
        Binance exchange = new Binance(config);
        exchange.setSandboxMode(true);
        exchange.loadMarkets(false);

        Ticker ticker = exchange.fetchTicker(symbol);
        // bid far below the market so the order rests instead of filling
        // (rounded manually here; production code should use exchange.priceToPrecision)
        double price = Math.round(ticker.last * 0.8 * 100) / 100.0;
        double amount = Math.round(20 / price * 1e5) / 1e5; // ~20 USDT notional

        Order order = exchange.createOrder(symbol, "limit", "buy", amount, price);
        System.out.printf("created %s: %s %s @ %s%n", order.id, amount, symbol, price);
        try {
            List<Order> open = exchange.fetchOpenOrders(symbol);
            open.forEach(o -> System.out.println("open order: " + o.id));
        } finally {
            Order canceled = exchange.cancelOrder(order.id, symbol);
            System.out.printf("canceled %s, status: %s%n", canceled.id, canceled.status);
        }
    }
}
