// Compare one symbol's price across several exchanges.
// Usage: cargo run --bin price_snapshot [SYMBOL]
use ccxt::types::Ticker;
use ccxt::{Binance, Kraken, Okx, Params};

fn print_row(id: &str, ticker: &Ticker) {
    let spread = match (ticker.bid, ticker.ask) {
        (Some(bid), Some(ask)) => format!("{:.4}%", (ask - bid) / ask * 100.0),
        _ => "n/a".to_string(),
    };
    println!("{:<8} last {:?}  bid {:?}  ask {:?}  spread {}", id, ticker.last, ticker.bid, ticker.ask, spread);
}

#[tokio::main]
async fn main() -> Result<(), ccxt::ExchangeError> {
    let symbol = std::env::args().nth(1).unwrap_or_else(|| "BTC/USDT".to_string());
    println!("{symbol}");
    let mut binance = Binance::new(None);
    print_row("binance", &binance.fetch_ticker(&symbol, Params::none()).await?);
    let mut kraken = Kraken::new(None);
    print_row("kraken", &kraken.fetch_ticker(&symbol, Params::none()).await?);
    let mut okx = Okx::new(None);
    print_row("okx", &okx.fetch_ticker(&symbol, Params::none()).await?);
    Ok(())
}
