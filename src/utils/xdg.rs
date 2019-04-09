use std::env;
use std::path::PathBuf;

pub fn get_config_dir(name: &str) -> PathBuf {
    env::var("XDG_CONFIG_HOME")
        .map(PathBuf::from)
        .unwrap_or(dirs::home_dir().unwrap().join(".config"))
        .join(name)
}

pub fn get_log_dir(name: &str) -> PathBuf {
    env::var("XDG_STATE_HOME")
        .map(PathBuf::from)
        .unwrap_or(dirs::home_dir().unwrap().join(".local").join("state"))
        .join(name)
}
