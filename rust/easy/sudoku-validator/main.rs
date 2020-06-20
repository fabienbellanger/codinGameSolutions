use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    for i in 0..9 as usize {
        let mut inputs = String::new();
        io::stdin().read_line(&mut inputs).unwrap();
        for j in inputs.split_whitespace() {
            let n = parse_input!(j, i32);
        }
    }

    // Write an answer using println!("message...");
    // To debug: eprintln!("Debug message...");

    println!("true or false");
}
