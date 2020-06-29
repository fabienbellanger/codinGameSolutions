use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let r = parse_input!(input_line, i32);
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let l = parse_input!(input_line, i32);

    let mut list: Vec<i32> = vec![r];

    for i in 2..=l {
        let mut line: Vec<i32> = Vec::new();
        let mut current = list[0];
        let mut counter = 0i32;

        if list.len() == 1 {
            counter = 1;
        } else {
            for j in 0..list.len() {
                if list[j] == current {
                    counter += 1;
                } else {
                    line.push(counter);
                    line.push(current);
                    current = list[j];
                    counter = 1;
                }
            }
        }
        line.push(counter);
        line.push(current);

        list = line;
    }

    let result: String = list.iter().map(|&i| i.to_string() + " ").collect(); 
    print!("{}", result.trim());
}
