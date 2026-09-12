// Full order lifecycle on the exchange testnet: create -> inspect -> cancel.
// Needs Binance *testnet* keys (https://testnet.binance.vision):
//   export CCXT_SANDBOX_APIKEY=... CCXT_SANDBOX_SECRET=...
use ccxt::{Binance, Config, Params};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let (Ok(api_key), Ok(secret)) = (std::env::var("CCXT_SANDBOX_APIKEY"), std::env::var("CCXT_SANDBOX_SECRET")) else {
        eprintln!("Set CCXT_SANDBOX_APIKEY and CCXT_SANDBOX_SECRET (testnet keys only!)");
        std::process::exit(1);
    };

    let symbol = "BTC/USDT";
    let mut exchange = Binance::with_config(Config::new().api_key(&api_key).secret(&secret).sandbox(true));
    exchange.try_load_markets(false).await?;

    let ticker = exchange.fetch_ticker(symbol, Params::none()).await?;
    // bid far below the market so the order rests instead of filling
    // (rounded manually here; production code should use price_to_precision)
    let price = (ticker.last.expect("no last price") * 0.8 * 100.0).round() / 100.0;
    let amount = (20.0 / price * 1e5).round() / 1e5; // ~20 USDT notional

    let order = exchange.create_limit_buy_order(symbol, amount, price, Params::none()).await?;
    let order_id = order.id.clone().unwrap_or_default();
    println!("created {}: {} {} @ {}", order_id, amount, symbol, price);

    let open = exchange.fetch_open_orders(Some(symbol), None, None, Params::none()).await?;
    for o in &open {
        println!("open order: {}", o.id.clone().unwrap_or_default());
    }

    let canceled = exchange.cancel_order(&order_id, Some(symbol), Params::none()).await?;
    println!("canceled {}, status: {:?}", canceled.id.clone().unwrap_or_default(), canceled.status);
    Ok(())
}
