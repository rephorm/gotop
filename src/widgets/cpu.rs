use tui::buffer::Buffer;
use tui::layout::Rect;
use tui::style::{Color, Style};
use tui::widgets::{Axis, Chart, Dataset, Marker, Widget};

use crate::widgets::block;

const TITLE: &str = " CPU Usage ";

pub struct CpuWidget {
    cpu_count: usize,

    per_cpu_data: Vec<Vec<(f64, f64)>>,
    average_cpu_data: Vec<(f64, f64)>,

    show_average_cpu_load: bool,
    show_per_cpu_load: bool,
}

impl CpuWidget {
    pub fn new(show_average_cpu_load: bool, show_per_cpu_load: bool) -> CpuWidget {
        dbg!(psutil::system::cpu_percent_percpu(1.0).unwrap());
        dbg!(psutil::disk::disk_usage("foo"));
        // let num_cpus = num_cpus::get();
        // let mut per_cpu_data = Vec::new();
        // let mut cpu_names = Vec::new();
        // for i in 0..num_cpus {
        //     per_cpu_data.push(Vec::new());
        //     cpu_names.push(format!("CPU{}", i))
        // }
        CpuWidget {
            cpu_count: 0,

            per_cpu_data: Vec::new(),
            average_cpu_data: Vec::new(),

            show_average_cpu_load,
            show_per_cpu_load,
        }
    }

    pub fn update(&mut self) {
        // let procs = sys.get_processor_list();
        // self.average_cpu_data.push((
        //     self.len_data as f64,
        //     procs[0].get_cpu_usage() as f64 * 100.0,
        // ));
        // for (i, proc) in procs.iter().skip(1).enumerate() {
        //     self.per_cpu_data[i].push((self.len_data as f64, proc.get_cpu_usage() as f64 * 100.0));
        // }
        // self.len_data += 1;
    }
}

impl Widget for CpuWidget {
    fn draw(&mut self, area: Rect, buf: &mut Buffer) {
        // let x_bounds = [self.len_data as f64 - 25.0, self.len_data as f64 - 1.0];
        // let mut datasets = vec![];
        // if self.average_cpu {
        //     datasets.push(
        //         Dataset::default()
        //             .name("AVRG")
        //             .marker(Marker::Braille)
        //             .style(Style::default().fg(Color::Yellow))
        //             .data(&self.average_cpu_data[..]),
        //     )
        // }
        // if self.per_cpu {
        //     for (i, cpu) in self.per_cpu_data.iter().enumerate() {
        //         datasets.push(
        //             Dataset::default()
        //                 .name(&self.cpu_names[i])
        //                 .marker(Marker::Braille)
        //                 .style(Style::default().fg(Color::Yellow))
        //                 .data(cpu),
        //         )
        //     }
        // }

        // let mut chart: Chart<String, String> = Chart::default();
        // chart
        //     .block(block::new().title(TITLE))
        //     .x_axis(Axis::default().bounds(x_bounds))
        //     .y_axis(Axis::default().bounds([0.0, 100.0]))
        //     .datasets(&datasets)
        //     .draw(area, buf);
    }
}
