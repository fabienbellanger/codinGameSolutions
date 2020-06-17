use std::io;
use std::collections::HashMap;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

// Calculate segments from a point
fn calculate_segments(final_list: &mut HashMap<i32, u32>, list: &Vec<i32>, point: i32) {
    for x in list {
        let size: i32 = (x - point).abs();

        let key = final_list.entry(size).or_insert(0);
        *key += 1;
    }
}

// Calculate squares number
fn calculate_squares(main: &HashMap<i32, u32>, second: &HashMap<i32, u32>) -> u32 {
    let mut squares = 0u32;

    for (size, n) in main {
        if let Some(m) = second.get(size) {
            squares += n * m;
        }
    }

    squares
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let w = parse_input!(inputs[0], i32);
    let h = parse_input!(inputs[1], i32);

    let mut inputs = String::new();
    let mut x_list: Vec<i32> = vec![0, w];
    let mut final_x_list: HashMap<i32, u32> = HashMap::new();
    final_x_list.insert(w, 1);
    io::stdin().read_line(&mut inputs).unwrap();
    for i in inputs.split_whitespace() {
        let x = parse_input!(i, i32);
        calculate_segments(&mut final_x_list, &x_list, x);
        x_list.push(x);
    }

    let mut inputs = String::new();
    let mut y_list: Vec<i32> = vec![0, h];
    let mut final_y_list: HashMap<i32, u32> = HashMap::new();
    final_y_list.insert(h, 1);
    io::stdin().read_line(&mut inputs).unwrap();
    for i in inputs.split_whitespace() {
        let y = parse_input!(i, i32);
        calculate_segments(&mut final_y_list, &y_list, y);
        y_list.push(y);
    }

    // Iterate on the smallest list
    let main_list: HashMap<i32, u32>;
    let second_list: HashMap<i32, u32>;
    
    if final_x_list.len() < final_y_list.len() {
        main_list = final_x_list;
        second_list = final_y_list;
    } else {
        main_list = final_y_list;
        second_list = final_x_list;
    }
    
    let squares = calculate_squares(&main_list, &second_list);
    println!("{}", squares);
}
