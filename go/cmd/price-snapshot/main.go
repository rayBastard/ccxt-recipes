// Compare one symbol's price across several exchanges.
// Usage: go run ./cmd/price-snapshot [SYMBOL]
package main

import (
	"fmt"
	"os"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func main() {
	symbol := "BTC/USDT"
	if len(os.Args) > 1 {
		symbol = os.Args[1]
	}
	fmt.Println(symbol)
	printTicker("binance", func() (ccxt.Ticker, error) { return ccxt.NewBinance(nil).FetchTicker(symbol) })
	printTicker("kraken", func() (ccxt.Ticker, error) { return ccxt.NewKraken(nil).FetchTicker(symbol) })
	printTicker("okx", func() (ccxt.Ticker, error) { return ccxt.NewOkx(nil).FetchTicker(symbol) })
}

func printTicker(id string, fetch func() (ccxt.Ticker, error)) {
	ticker, err := fetch()
	if err != nil {
		fmt.Printf("%-8s error: %v\n", id, err)
		return
	}
	spread := "n/a"
	if ticker.Bid != nil && ticker.Ask != nil {
		spread = fmt.Sprintf("%.4f%%", (*ticker.Ask-*ticker.Bid)/(*ticker.Ask)*100)
	}
	fmt.Printf("%-8s last %v  bid %v  ask %v  spread %s\n", id, deref(ticker.Last), deref(ticker.Bid), deref(ticker.Ask), spread)
}

func deref(value *float64) any {
	if value == nil {
		return "n/a"
	}
	return *value
}
