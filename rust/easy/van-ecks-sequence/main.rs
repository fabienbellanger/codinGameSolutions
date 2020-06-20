use std::io;
use std::collections::HashMap;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let a1 = parse_input!(input_line, i32);
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, i32);

    let mut sequence = vec![a1, 0];
    let mut values: HashMap<i32, i32> = HashMap::new();
    values.insert(a1, 0);

    for i in 1..n-1 {
        let v = sequence[i as usize];
        let mut new_val = 0i32;
        if let Some(t) = values.get(&v) {
            new_val = i - t;
        }
        values.insert(v, i);
        sequence.push(new_val);
    }

    let last = match sequence.last() {
        Some(l) => l,
        None => &0i32,
    };
    
    println!("{}", last);
}
