mod args;
mod utils;
mod widgets;

use std::io;
use std::path::PathBuf;

use args::Args;
use crossbeam_channel::unbounded;
use lazy_static::lazy_static;
use structopt::StructOpt;
use termion::event::Key;
use termion::input::{MouseTerminal, TermRead};
use termion::raw::IntoRawMode;
use tui::backend::Backend;
use tui::backend::TermionBackend;
use tui::style::{Color, Modifier, Style};
use tui::widgets::{Axis, Block, Borders, Chart, Dataset, Marker, Widget};
use tui::Terminal;

lazy_static! {
    static ref PROGRAM_NAME: &'static str = "gotop";
    static ref CONFIG_DIR: PathBuf = utils::xdg::get_config_dir(&PROGRAM_NAME);
    static ref LOG_DIR: PathBuf = utils::xdg::get_log_dir(&PROGRAM_NAME);
    static ref LOGFILE_PATH: PathBuf = LOG_DIR.join("errors.log");
}

struct App {
    cpu_data: Vec<Vec<f64>>,
}

fn draw<B: Backend>(terminal: &mut Terminal<B>, app: &App) {
    terminal
        .draw(|mut frame| {
            let size = frame.size();
            Chart::default()
                .block(Block::default().title("Chart"))
                .x_axis(
                    Axis::default()
                        .title("X Axis")
                        .title_style(Style::default().fg(Color::Red))
                        .style(Style::default().fg(Color::White))
                        .bounds([0.0, 10.0])
                        .labels(&["0.0", "5.0", "10.0"]),
                )
                .y_axis(
                    Axis::default()
                        .title("Y Axis")
                        .title_style(Style::default().fg(Color::Red))
                        .style(Style::default().fg(Color::White))
                        .bounds([0.0, 10.0])
                        .labels(&["0.0", "5.0", "10.0"]),
                )
                .datasets(&[
                    Dataset::default()
                        .name("data1")
                        .marker(Marker::Dot)
                        .style(Style::default().fg(Color::Cyan))
                        .data(&[(0.0, 5.0), (1.0, 6.0), (1.5, 6.434)]),
                    Dataset::default()
                        .name("data2")
                        .marker(Marker::Braille)
                        .style(Style::default().fg(Color::Magenta))
                        .data(&[(4.0, 5.0), (5.0, 8.0), (7.66, 13.5)]),
                ])
                .render(&mut frame, size);
        })
        .unwrap();
}

fn event_loop() {
    let stdin = io::stdin();
    // loop {
    for result in stdin.keys() {
        match result {
            Ok(key) => match key {
                Key::Char('q') => break,
                _ => {}
            },
            Err(_) => {}
        }
    }
    // ui::draw(&mut terminal, &app)
    // match events.next().unwrap() {
    //     Event::Input(key) => match key {
    //         Key::Char("q") => {
    //             break;
    //         }
    //     }
    // }
    // }
}

fn main() {
    let args = Args::from_args();

    widgets::CpuWidget::new(true, true);

    // let stdin = io::stdin();
    // for evt in stdin.keys()
    // let (tx, rx) = unbounded();

    let stdout = io::stdout().into_raw_mode().unwrap();
    let stdout = MouseTerminal::from(stdout);
    let backend = TermionBackend::new(stdout);
    let mut terminal = Terminal::new(backend).unwrap();
    terminal.hide_cursor().unwrap();

    event_loop()
}
