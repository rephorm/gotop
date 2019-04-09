use structopt::StructOpt;
use strum_macros::EnumString;

#[derive(EnumString)]
enum Colorscheme {
    default,
    default_dark,
    monokai,
    solarized,
    vice,
}

#[derive(StructOpt)]
pub struct Args {
    /// Set a colorscheme
    #[structopt(short = "c", default_value = "default")]
    colorscheme: Colorscheme,

    /// Only show CPU, Mem, and Process widgets
    #[structopt(short = "m")]
    minimal: bool,

    /// Number of times per second to update CPU and Mem widgets
    #[structopt(short = "r", default_value = "1")]
    rate: f64,

    /// Show each CPU in the CPU widget
    #[structopt(short = "p")]
    percpu: bool,

    /// Show average CPU in the CPU widget
    #[structopt(short = "a")]
    averagecpu: bool,

    /// Show temperatures in fahrenheit
    #[structopt(short = "f")]
    fahrenheit: bool,

    /// Show a statusbar with the time
    #[structopt(short = "s")]
    statusbar: bool,

    /// Show battery level widget ('minimal' turns off)
    #[structopt(short = "b")]
    battery: bool,
}
