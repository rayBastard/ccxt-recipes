// Compare perpetual-swap funding rates across derivatives exchanges.
// Usage: go run ./cmd/funding-rates [SYMBOL]
package main

import (
	"fmt"
	"os"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func main() {
	symbol := "BTC/USDT:USDT"
	if len(os.Args) > 1 {
		symbol = os.Args[1]
	}
	printRate("binance", func() (ccxt.FundingRate, error) { return ccxt.NewBinance(nil).FetchFundingRate(symbol) })
	printRate("bybit", func() (ccxt.FundingRate, error) { return ccxt.NewBybit(nil).FetchFundingRate(symbol) })
	printRate("okx", func() (ccxt.FundingRate, error) { return ccxt.NewOkx(nil).FetchFundingRate(symbol) })
}

func printRate(id string, fetch func() (ccxt.FundingRate, error)) {
	rate, err := fetch()
	if err != nil {
		fmt.Printf("%-8s error: %v\n", id, err)
		return
	}
	pct, next := "n/a", "n/a"
	if rate.FundingRate != nil {
		pct = fmt.Sprintf("%.4f%%", *rate.FundingRate*100)
	}
	if rate.FundingDatetime != nil {
		next = *rate.FundingDatetime
	}
	fmt.Printf("%-8s funding %s  next %s\n", id, pct, next)
}
