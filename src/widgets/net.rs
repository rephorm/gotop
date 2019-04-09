use tui::buffer::Buffer;
use tui::layout::Rect;
use tui::widgets::Widget;

use crate::widgets::block;

const TITLE: &str = " Network Usage ";

pub struct NetWidget {}

impl NetWidget {
    pub fn new() -> NetWidget {
        NetWidget {}
    }
}

impl Widget for NetWidget {
    fn draw(&mut self, area: Rect, buf: &mut Buffer) {
        block::new().title(TITLE).draw(area, buf);
    }
}
