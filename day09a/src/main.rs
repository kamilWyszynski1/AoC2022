use std::collections::{HashMap, HashSet};
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    // File hosts must exist in current path before this produces output
    if let Ok(lines) = read_lines("./input.txt") {
        // Consumes the iterator, returns an (Optional) String
        let mut lines_str = vec![];

        for line in lines {
            lines_str.push(line.unwrap().to_string())
        }
        calculate(lines_str)
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn calculate(lines: Vec<String>) {
    let mut visited: HashMap<(isize, isize), bool> = HashMap::new();
    let mut head = (0, 0);
    let mut tail = (0, 0);

    for line in lines {
        let (dir, amount) = line.split_once(' ').unwrap();
        let amount: usize = amount.parse().unwrap();

        println!("== {line} ==");

        for _ in 0..amount {
            match dir {
                "R" => {
                    head.0 += 1;
                }
                "U" => {
                    head.1 += 1;
                }
                "L" => head.0 -= 1,
                "D" => {
                    head.1 -= 1;
                }
                _ => {
                    panic!("invalid character")
                }
            }
            // calculate tail movement
            let are_close = are_close_enough(head, tail);
            if !are_close {
                println!("movement, \thead: {:?}, tail: {:?}", head, tail);
                // tail movement
                match dir {
                    "R" => tail = (head.0 - 1, head.1),
                    "U" => tail = (head.0, head.1 - 1),
                    "L" => tail = (head.0 + 1, head.1),
                    "D" => tail = (head.0, head.1 + 1),
                    _ => {
                        panic!("invalid character")
                    }
                }
                println!("\t\thead: {:?}, tail: {:?}", head, tail);
            }
            visited.insert(tail, true);
        }
    }
    println!("{:?}", visited.len());
    // draw(visited)
}

// fn draw(visited: HashSet<(isize, isize)>) {
//     let mut grid = [[false; 5]; 5];

//     for v in visited {
//         grid[v.0 as usize][v.1 as usize] = true;
//     }

//     for g in grid {
//         let mut line = String::new();
//         for v in g {
//             if v {
//                 line.push('#');
//             } else {
//                 line.push('.')
//             }
//         }
//         println!("{line}")
//     }
// }

fn are_close_enough(head: (isize, isize), tail: (isize, isize)) -> bool {
    let x: f64 = (head.0 - tail.0) as f64;
    let y: f64 = (head.1 - tail.1) as f64;

    let pyt = theorem(x, y);
    pyt < 1.5
}

pub fn theorem<T: Into<f64>>(a: T, b: T) -> f64 {
    let c: f64 = a.into().powi(2) + b.into().powi(2);
    c.sqrt()
}
