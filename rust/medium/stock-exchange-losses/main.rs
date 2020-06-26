use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, usize);
    let mut inputs = String::new();
    io::stdin().read_line(&mut inputs).unwrap();

    let mut values: Vec<i32> = Vec::new();
    let mut diff: Vec<i32> = Vec::new();
    let mut j: usize = 0;
    for i in inputs.split_whitespace() {
        let v = parse_input!(i, i32);
        values.push(v);

        if values.len() > 1 {
            diff.push(v - values[j - 1]);
        }

        j += 1;
    }
    eprintln!("values={:?}, diff={:?}", values, diff);

    for i in 1..n-1 {
        eprintln!("diff={:?}", diff[i]);
    }

    let loss = 0i32;
    
    println!("{}", loss);
}
