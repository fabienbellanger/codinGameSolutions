use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let width = parse_input!(input_line, usize); // the number of cells on the X axis
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let height = parse_input!(input_line, usize); // the number of cells on the Y axis

    let mut grid: Vec<Vec<bool>> = Vec::with_capacity(height);
    
    for _ in 0..height as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let line = input_line.trim_matches('\n').to_string(); // width characters, each either 0 or .

        let mut w: Vec<bool> = Vec::with_capacity(width);
        for c in line.chars() {
            w.push(c == '0');
        }
        grid.push(w);
    }

    for y in 0..height {
        for x in 0..width {
            if grid[y][x] {
                let mut s = String::from(format!("{} {}", x, y));

                // Right
                // -----
                let mut i = x + 1;
                while !(i >= width || grid[y][i]) {
                    i += 1;
                }
                if i < width {
                    s.push_str(&format!(" {} {}", i, y));
                } else {
                    s.push_str(" -1 -1");
                }

                // Bottom
                // ------
                let mut i = y + 1;
                while !(i >= height || grid[i][x]) {
                    i += 1;
                }
                if i < height {
                    s.push_str(&format!(" {} {}", x, i));
                } else {
                    s.push_str(" -1 -1");
                }

                println!("{}", s);
            }
        }
    }
}
