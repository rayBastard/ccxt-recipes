// Compare perpetual-swap funding rates across derivatives exchanges.
package recipes;

import io.github.ccxt.exchanges.Binance;
import io.github.ccxt.exchanges.Bybit;
import io.github.ccxt.exchanges.Okx;
import io.github.ccxt.types.FundingRate;

public class FundingRates {

    public static void main(String[] args) {
        String symbol = args.length > 0 ? args[0] : "BTC/USDT:USDT";
        print("binance", new Binance().fetchFundingRate(symbol));
        print("bybit", new Bybit().fetchFundingRate(symbol));
        print("okx", new Okx().fetchFundingRate(symbol));
    }

    static void print(String id, FundingRate rate) {
        String pct = (rate.fundingRate != null) ? String.format("%.4f%%", rate.fundingRate * 100) : "n/a";
        String next = (rate.fundingDatetime != null) ? rate.fundingDatetime : "n/a";
        System.out.printf("%-8s funding %s  next %s%n", id, pct, next);
    }
}
