use std::io;
use std::collections::HashMap;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

#[derive(Debug)]
struct Diff {
    before: i32,
    after: i32,
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let a1 = parse_input!(input_line, i32);
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, i32);

    let mut sequence = vec![a1];
    let mut values: HashMap<i32, Diff> = HashMap::new();
    values.insert(a1, Diff{before: 0, after: 0});
    for i in 1..n {
        let last = match sequence.last() {
            Some(l) => *l,
            None => 0i32,
        };
        
        let next = match values.get(&last) {
            Some(next) => next,
            None => &Diff{before: 0, after: 0},
        };
        
        let new_val = next.after - next.before;
        sequence.push(new_val);

        // let new_diff = if let Some(v) = values.get(&new_val) {
        //     Diff{before: v.after, after: i}
        // } else {
        //     Diff{before: i, after: i}
        // };
        // values.insert(new_val, new_diff);
        //values.entry(new_val).or_insert(Diff{before: i, after: i});
        //values.insert(new_val, Diff{before: values[&new_val].after, after: i});
        *values.entry(new_val).or_insert(Diff{before: i, after: i}) = Diff{before: values[&new_val].after, after: i};

    }

    let last = match sequence.last() {
        Some(l) => l,
        None => &0i32,
    };
    
    println!("{}", last);
}
