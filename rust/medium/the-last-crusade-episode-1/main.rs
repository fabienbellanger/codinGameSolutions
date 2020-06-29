use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

#[derive(Debug)]
struct Path {
    In: String,
    Coords: (isize, isize),
}

fn generate_zones() -> Vec<Vec<Path>> {
    let mut zones: Vec<Vec<Path>> = Vec::with_capacity(14);

    // Zone type 0
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::new(), Coords: (0, 0)});
    zones.push(zone);

    // Zone type 1
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("TOP"), Coords: (0, 1)});
    zone.push(Path{In: String::from("RIGHT"), Coords: (0, 1)});
    zone.push(Path{In: String::from("LEFT"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 2
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("RIGHT"), Coords: (-1, 0)});
    zone.push(Path{In: String::from("LEFT"), Coords: (1, 0)});
    zones.push(zone);

    // Zone type 3
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("TOP"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 4
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("TOP"), Coords: (-1, 0)});
    zone.push(Path{In: String::from("RIGHT"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 5
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("TOP"), Coords: (1, 0)});
    zone.push(Path{In: String::from("LEFT"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 6
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("RIGHT"), Coords: (-1, 0)});
    zone.push(Path{In: String::from("LEFT"), Coords: (1, 0)});
    zones.push(zone);

    // Zone type 7
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("RIGHT"), Coords: (0, 1)});
    zone.push(Path{In: String::from("TOP"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 8
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("RIGHT"), Coords: (0, 1)});
    zone.push(Path{In: String::from("LEFT"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 9
    // -----------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("LEFT"), Coords: (0, 1)});
    zone.push(Path{In: String::from("TOP"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 10
    // ------------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("TOP"), Coords: (-1, 0)});
    zones.push(zone);

    // Zone type 11
    // ------------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("TOP"), Coords: (1, 0)});
    zones.push(zone);

    // Zone type 12
    // ------------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("RIGHT"), Coords: (0, 1)});
    zones.push(zone);

    // Zone type 13
    // ------------
    let mut zone: Vec<Path> = Vec::new();
    zone.push(Path{In: String::from("LEFT"), Coords: (0, 1)});
    zones.push(zone);

    zones
}

fn calcul_next(x: isize, y: isize, pos: String, zone_index: i32, zones: &Vec<Vec<Path>>) -> (isize, isize) {
    let zone = &zones[zone_index as usize];

    let mut coords = (x, y);
    for path in zone {
        if path.In == pos {
            coords = (x + path.Coords.0, y + path.Coords.1);
            break;
        }
    }
    coords
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let w = parse_input!(inputs[0], usize); // number of columns.
    let h = parse_input!(inputs[1], usize); // number of rows.
    let mut lines: Vec<Vec<i32>> = Vec::with_capacity(h);

    for i in 0..h {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let line = input_line.trim_matches('\n').to_string(); // represents a line in the grid and contains W integers. Each integer represents one room of a given type.
        let line = line.split(" ").map(|x| parse_input!(x, i32)).collect::<Vec<_>>();
        lines.push(line);
    }
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let ex = parse_input!(input_line, i32); // the coordinate along the X axis of the exit (not useful for this first mission, but must be read).

    let zones = generate_zones();

    // game loop
    loop {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let xi = parse_input!(inputs[0], isize);
        let yi = parse_input!(inputs[1], isize);
        let pos = inputs[2].trim().to_string();
        let zone = lines[yi as usize][xi as usize];

        let coords = calcul_next(xi, yi, pos, zone, &zones);
        println!("{} {}", coords.0, coords.1);
    }
}
