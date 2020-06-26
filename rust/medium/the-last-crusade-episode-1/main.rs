use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let w = parse_input!(inputs[0], i32); // number of columns.
    let h = parse_input!(inputs[1], i32); // number of rows.
    for i in 0..h as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let line = input_line.trim_matches('\n').to_string(); // represents a line in the grid and contains W integers. Each integer represents one room of a given type.
    }
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let ex = parse_input!(input_line, i32); // the coordinate along the X axis of the exit (not useful for this first mission, but must be read).

    // game loop
    loop {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let xi = parse_input!(inputs[0], i32);
        let yi = parse_input!(inputs[1], i32);
        let pos = inputs[2].trim().to_string();

        // Write an action using println!("message...");
        // To debug: eprintln!("Debug message...");


        // One line containing the X Y coordinates of the room in which you believe Indy will be on the next turn.
        println!("0 0");
    }
}
