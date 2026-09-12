// Stream a live order book over WebSocket and print the top of book in place.
// Usage: cargo run --bin orderbook_stream [SYMBOL]   (Ctrl+C to stop)
use std::io::Write;

use ccxt::Params;
use ccxt_pro::Binance;

fn top(side: &[[f64; 2]]) -> String {
    match side.first() {
        Some(level) => format!("{} ({})", level[0], level[1]),
        None => "n/a".to_string(),
    }
}

#[tokio::main]
async fn main() -> Result<(), ccxt::ExchangeError> {
    let symbol = std::env::args().nth(1).unwrap_or_else(|| "BTC/USDT".to_string());
    let mut exchange = Binance::new(None);
    loop {
        let book = exchange.watch_order_book(&symbol, Some(10), Params::none()).await?;
        print!("\r{}  bid {}  ask {}   ", symbol, top(&book.bids), top(&book.asks));
        std::io::stdout().flush().ok();
    }
}
