// Compare perpetual-swap funding rates across derivatives exchanges.
// Usage: cargo run --bin funding_rates [SYMBOL]
use ccxt::types::FundingRate;
use ccxt::{Binance, Bybit, Okx, Params};

fn print_row(id: &str, rate: &FundingRate) {
    let pct = match rate.funding_rate {
        Some(funding) => format!("{:.4}%", funding * 100.0),
        None => "n/a".to_string(),
    };
    let next = rate.funding_datetime.as_deref().unwrap_or("n/a");
    println!("{:<8} funding {}  next {}", id, pct, next);
}

#[tokio::main]
async fn main() -> Result<(), ccxt::ExchangeError> {
    let symbol = std::env::args().nth(1).unwrap_or_else(|| "BTC/USDT:USDT".to_string());
    let mut binance = Binance::new(None);
    print_row("binance", &binance.fetch_funding_rate(&symbol, Params::none()).await?);
    let mut bybit = Bybit::new(None);
    print_row("bybit", &bybit.fetch_funding_rate(&symbol, Params::none()).await?);
    let mut okx = Okx::new(None);
    print_row("okx", &okx.fetch_funding_rate(&symbol, Params::none()).await?);
    Ok(())
}
