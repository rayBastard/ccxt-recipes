// Download paginated OHLCV candles and write them to a CSV file.
// Usage: cargo run --bin ohlcv_to_csv [SYMBOL TIMEFRAME DAYS]
use chrono::DateTime;
use std::io::Write;

use ccxt::{Binance, Params};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<String> = std::env::args().collect();
    let symbol = args.get(1).map(String::as_str).unwrap_or("BTC/USDT");
    let timeframe = args.get(2).map(String::as_str).unwrap_or("1h");
    let days: i64 = args.get(3).and_then(|d| d.parse().ok()).unwrap_or(7);

    let mut exchange = Binance::new(None);
    let mut cursor = chrono::Utc::now().timestamp_millis() - days * 86400 * 1000;
    let mut candles: Vec<ccxt::types::OHLCV> = Vec::new();
    loop {
        let batch = exchange
            .fetch_ohlcv(symbol, Some(timeframe), Some(cursor), Some(1000), Params::none())
            .await?;
        if batch.is_empty() {
            break;
        }
        let last = batch[batch.len() - 1][0] as i64;
        let batch_len = batch.len();
        candles.extend(batch);
        if batch_len < 1000 || last <= cursor {
            break; // reached the newest candle
        }
        cursor = last + 1;
    }

    let filename = format!("{}-{}.csv", symbol.replace('/', "-").replace(':', "-"), timeframe);
    let mut file = std::fs::File::create(&filename)?;
    writeln!(file, "timestamp,datetime,open,high,low,close,volume")?;
    for c in &candles {
        let datetime = DateTime::from_timestamp_millis(c[0] as i64).map(|d| d.to_rfc3339()).unwrap_or_default();
        writeln!(file, "{},{},{},{},{},{},{}", c[0] as i64, datetime, c[1], c[2], c[3], c[4], c[5])?;
    }
    println!("{} candles -> {}", candles.len(), filename);
    Ok(())
}
