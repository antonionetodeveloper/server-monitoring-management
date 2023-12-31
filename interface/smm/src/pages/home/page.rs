slint::include_modules!();
use slint::Weak;

pub fn content() -> Result<(), slint::PlatformError> {
    let app: AppWindow = AppWindow::new().unwrap();
    let weak: Weak<AppWindow> = app.as_weak();

    

    app.run().unwrap();
}
