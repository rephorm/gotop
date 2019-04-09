use tui::buffer::Buffer;
use tui::layout::Rect;
use tui::widgets::Widget;

use crate::widgets::block;

const TITLE: &str = " Disk Usage ";

pub struct DiskWidget {}

impl DiskWidget {
    pub fn new() -> DiskWidget {
        DiskWidget {}
    }
}

impl Widget for DiskWidget {
    fn draw(&mut self, area: Rect, buf: &mut Buffer) {
        block::new().title(TITLE).draw(area, buf);
    }
}
