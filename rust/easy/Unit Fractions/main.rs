use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, i64);

    // Write an action using println!("message...");
    // To debug: eprintln!("Debug message...");
    
    /*
    x, y > n
    n < x <= 2*n
    y = (n*x) / x - n
    
    x = 2*n
    tant que x > n
        x--
        y = (n*x) / (x-n)
        if (n*x) % (x-n) == 0 and y > n alors
            found solution: y = (n*x) / (x-n)
        sinon KO
    */
    // let mut x: i64;
    let mut y: i64;
    for x in (n+1)..(2*n+1) {
        if (n * x) % (x - n) == 0 {
            y = ((n * x) / (x - n)) as i64;
            println!("1/{} = 1/{} + 1/{}", n, y, x);
        }
    }
}
