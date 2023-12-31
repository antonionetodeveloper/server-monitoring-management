slint::include_modules!();

pub mod pages {
    pub mod home {
        pub mod page;
    }
}


fn main() {
    println!("Starting SMM...");
    router();
}

fn router() {
    let home: Result<(), slint::PlatformError> = pages::home::page::content();
    return home.unwrap();
}