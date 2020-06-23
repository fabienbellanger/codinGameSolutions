use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

const SIZE: usize = 3;
const LENGTH: usize = SIZE * SIZE;
const ORDERED_ROW: [usize; LENGTH] = [1, 2, 3, 4, 5, 6, 7, 8, 9];

// Checks if an array is valid
fn check_array(line: &Vec<usize>) -> bool {
    let v = ORDERED_ROW.to_vec();
    let mut k = line.clone();
    k.sort();
    k == v
}

fn main() {
    let mut all: Vec<usize> = Vec::with_capacity(81);

    for _ in 0..LENGTH as usize {
        let mut inputs = String::new();
        io::stdin().read_line(&mut inputs).unwrap();

        for j in inputs.split_whitespace() {
            all.push(parse_input!(j, usize));
        }
    }
    
    let mut ok = true;
    for i in 0..LENGTH {
        // Rows
        // ----
        let mut row: Vec<usize> = Vec::with_capacity(9);
        for j in 0..LENGTH {
            row.push(all[LENGTH * i + j]);
        }
        if !check_array(&row) {
            ok = false;
            break;
        }

        // Columns
        // -------
        let mut col: Vec<usize> = Vec::with_capacity(9);
        for j in 0..LENGTH {
            col.push(all[i + LENGTH * j]);
        }
        if !check_array(&col) {
            ok = false;
            break;
        }

        // Sub grids
        // ---------
        let mut sub_grid: Vec<usize> = Vec::with_capacity(9);
        let point = SIZE*(LENGTH*(i/SIZE) + i%SIZE);
        for j in 0..SIZE {
            for k in 0..SIZE {
                let index = point + LENGTH*j + k;
                sub_grid.push(all[index]);
            }
        }
        if !check_array(&sub_grid) {
            ok = false;
            break;
        }
    };

    println!("{}", ok);
}
