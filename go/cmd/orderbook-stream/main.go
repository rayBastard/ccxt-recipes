// Stream a live order book over WebSocket and print the top of book in place.
// Usage: go run ./cmd/orderbook-stream [SYMBOL]   (Ctrl+C to stop)
package main

import (
	"fmt"
	"os"

	ccxt "github.com/ccxt/ccxt/go/v4"
	ccxtpro "github.com/ccxt/ccxt/go/v4/pro"
)

func main() {
	symbol := "BTC/USDT"
	if len(os.Args) > 1 {
		symbol = os.Args[1]
	}
	exchange := ccxtpro.NewBinance(nil)
	for {
		orderbook, err := exchange.WatchOrderBook(symbol, ccxt.WithWatchOrderBookLimit(10))
		if err != nil {
			panic(err)
		}
		bid, ask := top(orderbook.Bids), top(orderbook.Asks)
		fmt.Printf("\r%s  bid %s  ask %s   ", symbol, bid, ask)
	}
}

func top(side [][]float64) string {
	if len(side) == 0 {
		return "n/a"
	}
	return fmt.Sprintf("%v (%v)", side[0][0], side[0][1])
}
