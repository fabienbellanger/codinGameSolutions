use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

const ORDERED_ROW: [usize; 9] = [1, 2, 3, 4, 5, 6, 7, 8, 9];

fn check_line(line: &Vec<usize>) -> bool {
    let v = ORDERED_ROW.to_vec();
    let mut k = line.clone();
    k.sort();
    k == v
}

fn main() {
    let mut all: Vec<usize> = Vec::with_capacity(81);
    let mut ok = true;

    for _ in 0..9 as usize {
        let mut inputs = String::new();
        io::stdin().read_line(&mut inputs).unwrap();

        for j in inputs.split_whitespace() {
            let l = parse_input!(j, usize);
            all.push(l);
        }
    }

    eprintln!("all={:?}", all);
    
    let mut ok = true;
    for i in 0..9 {
        // Rows
        // ----
        let mut row: Vec<usize> = Vec::with_capacity(9);
        for j in 0..9 {
            row.push(all[9 * i + j]);
        }
        // eprintln!("row={:?}", row);
        if !check_line(&row) {
            ok = false;
            break;
        }

        // Columns
        // -------
        let mut col: Vec<usize> = Vec::with_capacity(9);
        for j in 0..9 {
            col.push(all[i + 9 * j]);
        }
        // eprintln!("col={:?}", col);
        if !check_line(&col) {
            ok = false;
            break;
        }

        // Squares
        // -------
        let mut square: Vec<usize> = Vec::with_capacity(9);
        // for j in 0..9 {
        //     square.push(all[i + 9 * j]);
        // }
    };

    if ok {
        println!("true");
    } else {
        println!("false");
    }
}
