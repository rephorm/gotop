mod block;
mod cpu;
mod disk;
mod mem;
mod net;
mod proc;
mod temp;

pub use self::cpu::CpuWidget;
pub use self::disk::DiskWidget;
pub use self::mem::MemWidget;
pub use self::net::NetWidget;
pub use self::proc::ProcWidget;
pub use self::temp::TempWidget;
