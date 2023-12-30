pub mod pages {
    pub mod home {
        pub mod page;
    }
}

fn main() {
    println!("Starting SSM...");
    router();
}

pub fn router() {
    pages::home::page::content();
}
