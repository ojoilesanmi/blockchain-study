extern crate anyinput;

use anyinput::anyinput;

#[anyinput]
pub fn encrypt(input: String, rot_key: i32) -> String {
    /* TODO write these:
    ->  Use same ASCII code for upper and lower case.
    */
    let mut result = String::new();

    for i in input.chars() {
        if i.is_lowercase() {
            let char = i
        }
    }
}

#[anyinput]
pub fn decrypt(input: String, rot_key: i32) -> String {
    /* TODO write these:
    ->  Use same ASCII code for upper and lower case.
    */
}