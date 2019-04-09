use tui::buffer::Buffer;
use tui::layout::Rect;
use tui::widgets::Widget;

use crate::widgets::block;

const TITLE: &str = " Memory Usage ";

pub struct MemWidget {}

impl MemWidget {
    pub fn new() -> MemWidget {
        MemWidget {}
    }
}

impl Widget for MemWidget {
    fn draw(&mut self, area: Rect, buf: &mut Buffer) {
        block::new().title(TITLE).draw(area, buf);
    }
}
