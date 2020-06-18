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

    let mut sequence = vec![a1];
    for i in a1+1..n {
        let last = last_item(&sequence);
        let mut val: i32 = if let Some(index) = &sequence[..i as usize-2].iter().rposition(|&r| r == *last) {
            let j = *index as i32;
            if j > i - 1 {
                j
            } else {
                -1
            }
        } else {
            -1
        };

        if val == -1 {
            eprintln!("last={} - i={} : index={:?} => {}", last, i, val, 0);
            sequence.push(0);
        } else {
            eprintln!("last={} - i={} : index={:?} => {}", last, i, val, i - val);
            sequence.push(i - val);
        }
    }

    let last = last_item(&sequence);
    
    eprintln!("sequence={:?}\n", sequence);
    
    println!("{}", last);
}

fn last_item(list: &Vec<i32>) -> &i32 {
    match list.last() {
        Some(l) => l,
        None => &0i32,
    }
}
