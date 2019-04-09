use tui::buffer::Buffer;
use tui::layout::Rect;
use tui::widgets::Widget;

use crate::widgets::block;

const TITLE: &str = " Processes ";

pub struct ProcWidget {}

impl ProcWidget {
    pub fn new() -> ProcWidget {
        ProcWidget {}
    }
}

impl Widget for ProcWidget {
    fn draw(&mut self, area: Rect, buf: &mut Buffer) {
        block::new().title(TITLE).draw(area, buf);
    }
}
