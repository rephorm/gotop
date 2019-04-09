use tui::buffer::Buffer;
use tui::layout::Rect;
use tui::widgets::Widget;

use crate::widgets::block;

const TITLE: &str = " Temperatures ";

pub struct TempWidget {}

impl TempWidget {
    pub fn new() -> TempWidget {
        TempWidget {}
    }
}

impl Widget for TempWidget {
    fn draw(&mut self, area: Rect, buf: &mut Buffer) {
        block::new().title(TITLE).draw(area, buf);
    }
}
