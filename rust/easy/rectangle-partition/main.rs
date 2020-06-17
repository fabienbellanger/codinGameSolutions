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
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let w = parse_input!(inputs[0], i32);
    let h = parse_input!(inputs[1], i32);
    let count_x = parse_input!(inputs[2], usize);
    let count_y = parse_input!(inputs[3], usize);

    let mut inputs = String::new();
    let mut x_list: Vec<i32> = vec!(0);
    io::stdin().read_line(&mut inputs).unwrap();
    for i in inputs.split_whitespace() {
        let x = parse_input!(i, i32);
        x_list.push(x);
    }
    if !x_list.contains(&w) {
        x_list.push(w);
    }

    let mut inputs = String::new();
    let mut y_list: Vec<i32> = vec!(0);
    io::stdin().read_line(&mut inputs).unwrap();
    for i in inputs.split_whitespace() {
        let y = parse_input!(i, i32);
        y_list.push(y);
    }
    if !y_list.contains(&h) {
        y_list.push(h);
    }
    
    let mut i = 0;
    let mut squares = 0;
    while i < (count_y + 1) {
        let mut j = 0;
        while j < (count_x + 1) {
            let mut k = j + 1;
            while k < (count_x + 2) {
                let mut l = i + 1;
                while l < (count_y + 2) {
                    if (x_list[k] - x_list[j]) == (y_list[l] - y_list[i]) { // if = goulot d'étranglement !!
                        squares += 1;
                    }

                    l += 1;
                }

                k += 1;
            }

            j += 1;
        }

        i += 1;
    }

    // let mut i_index = 0;
    // let mut squares = 0;
    // for i in y_list.iter() {
    //     if i < &h {
    //         let mut j_index = 0;
    //         for j in x_list.iter() {
    //             if j < &w {
    //                 let x_partial: Vec<i32> = Vec::from(&x_list[j_index + 1..]);

    //                 for k in x_partial.iter() {
    //                     let y_partial: Vec<i32> = Vec::from(&y_list[i_index + 1..]);

    //                     for l in y_partial.iter() {
    //                         if (k - j) == (l - i) {
    //                             squares += 1;
    //                         }
    //                     }
    //                 }

    //                 j_index += 1;
    //             }
    //         }

    //         i_index += 1;
    //     }
    // }

    println!("{}", squares);

    // Write an answer using println!("message...");
    // To debug: eprintln!("Debug message...");
}
