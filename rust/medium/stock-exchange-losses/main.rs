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

    let mut losses: Vec<i32> = Vec::new();
    losses.push(diff[0]);
    let mut loss = 0i32;
    for i in 1..n-1 {
        let v: i32 = if losses[i - 1] >= 0 {
            diff[i]
        } else {
            diff[i] + losses[i - 1]
        };

        losses.push(v);
        if v < loss {
            loss = v;
        }
    }
    
    println!("{}", loss);
}
