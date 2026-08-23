// Full order lifecycle on the exchange testnet: create -> inspect -> cancel.
// Needs Binance *testnet* keys (https://testnet.binance.vision):
//   export CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=...
package main

import (
	"fmt"
	"math"
	"os"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func main() {
	apiKey := os.Getenv("CCXT_SANDBOX_APIKEY")
	secret := os.Getenv("CCXT_SANDBOX_SECRET")
	if apiKey == "" || secret == "" {
		fmt.Fprintln(os.Stderr, "Set CCXT_SANDBOX_APIKEY and CCXT_SANDBOX_SECRET (testnet keys only!)")
		os.Exit(1)
	}

	symbol := "BTC/USDT"
	exchange := ccxt.NewBinance(map[string]interface{}{"apiKey": apiKey, "secret": secret})
	exchange.SetSandboxMode(true)
	if _, err := exchange.LoadMarkets(); err != nil {
		panic(err)
	}

	ticker, err := exchange.FetchTicker(symbol)
	if err != nil {
		panic(err)
	}
	// bid far below the market so the order rests instead of filling
	// (rounded manually here; production code should use exchange.PriceToPrecision)
	price := math.Round(*ticker.Last*0.8*100) / 100
	amount := math.Round(20/price*1e5) / 1e5 // ~20 USDT notional

	order, err := exchange.CreateOrder(symbol, "limit", "buy", amount, ccxt.WithCreateOrderPrice(price))
	if err != nil {
		panic(err)
	}
	fmt.Printf("created %s: %v %s @ %v\n", *order.Id, amount, symbol, price)

	open, err := exchange.FetchOpenOrders(ccxt.WithFetchOpenOrdersSymbol(symbol))
	if err != nil {
		fmt.Println("fetchOpenOrders error:", err)
	} else {
		for _, o := range open {
			fmt.Println("open order:", *o.Id)
		}
	}

	canceled, err := exchange.CancelOrder(*order.Id, ccxt.WithCancelOrderSymbol(symbol))
	if err != nil {
		panic(err)
	}
	fmt.Printf("canceled %s, status: %v\n", *canceled.Id, deref(canceled.Status))
}

func deref(value *string) any {
	if value == nil {
		return "n/a"
	}
	return *value
}
