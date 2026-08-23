// Download paginated OHLCV candles and write them to a CSV file.
// Usage: go run ./cmd/ohlcv-to-csv [SYMBOL TIMEFRAME DAYS]
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func main() {
	symbol, timeframe, days := "BTC/USDT", "1h", 7
	if len(os.Args) > 1 {
		symbol = os.Args[1]
	}
	if len(os.Args) > 2 {
		timeframe = os.Args[2]
	}
	if len(os.Args) > 3 {
		days, _ = strconv.Atoi(os.Args[3])
	}

	exchange := ccxt.NewBinance(nil)
	cursor := time.Now().UnixMilli() - int64(days)*86400*1000
	var candles []ccxt.OHLCV
	for {
		batch, err := exchange.FetchOHLCV(symbol,
			ccxt.WithFetchOHLCVTimeframe(timeframe),
			ccxt.WithFetchOHLCVSince(cursor),
			ccxt.WithFetchOHLCVLimit(1000))
		if err != nil {
			panic(err)
		}
		if len(batch) == 0 {
			break
		}
		candles = append(candles, batch...)
		last := batch[len(batch)-1].Timestamp
		if len(batch) < 1000 || last <= cursor {
			break // reached the newest candle
		}
		cursor = last + 1
	}

	filename := strings.NewReplacer("/", "-", ":", "-").Replace(symbol) + "-" + timeframe + ".csv"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"timestamp", "datetime", "open", "high", "low", "close", "volume"})
	for _, c := range candles {
		writer.Write([]string{
			strconv.FormatInt(c.Timestamp, 10),
			time.UnixMilli(c.Timestamp).UTC().Format(time.RFC3339),
			fmt.Sprint(c.Open), fmt.Sprint(c.High), fmt.Sprint(c.Low), fmt.Sprint(c.Close), fmt.Sprint(c.Volume),
		})
	}
	fmt.Printf("%d candles -> %s\n", len(candles), filename)
}
