use std::io;
use std::collections::HashMap;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let n = parse_input!(inputs[0], i32); // the total number of nodes in the level, including the gateways
    let l = parse_input!(inputs[1], i32); // the number of links
    let e = parse_input!(inputs[2], i32); // the number of exit gateways

    let mut links: Vec<(i32, i32, bool)> = Vec::with_capacity(l as usize);
    for i in 0..l as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let n1 = parse_input!(inputs[0], i32); // N1 and N2 defines a link between these nodes
        let n2 = parse_input!(inputs[1], i32);
        
        links.push((n1, n2, true));
    }

    let mut gateways: HashMap<i32, bool> = HashMap::with_capacity(e as usize);
    for i in 0..e as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let ei = parse_input!(input_line, i32); // the index of a gateway node
        
        gateways.insert(ei, true);
    }

    // game loop
    loop {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let si = parse_input!(input_line, i32); // The index of the node on which the Skynet agent is positioned this turn

        let mut near_gateway: Option<(&i32, &bool)> = None;
        for gateway in &gateways {
            if *gateway.1 {
                let mut found_link = false;
                for link in &mut links {
                    if link.2 && ((link.0 == si && link.1 == *gateway.0) || (link.1 == si && link.0 == *gateway.0)) {
                        near_gateway = Some(gateway);
                        found_link = true;
                        break;
                    }
                }

                if found_link {
                    break;
                }
            }
        }

        if near_gateway == None {
            near_gateway = gateways.iter().find(|g| *g.1 == true);
        }

        if let Some(gateway) = near_gateway {
            eprintln!("valid gateway={:?}", gateway);

            let mut ok = false;
            for link in &mut links {
                if link.2 && ((link.0 == si && link.1 == *gateway.0) || (link.1 == si && link.0 == *gateway.0)) {
                    println!("{} {}", link.0, link.1);
                    link.2 = false;
                    ok = true;
                    break;
                }
            }

            if !ok {
                for link in &mut links {
                    if link.2 && (link.0 == si || link.1 == si) {
                        eprintln!("link 2={:?}", link);
                        println!("{} {}", link.0, link.1);
                        link.2 = false;
                        break;
                    }
                }
            }
        }
    }
}
