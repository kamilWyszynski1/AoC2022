use std::collections::HashMap;
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
    let mut tails = [(0, 0); 10];

    for line in lines {
        let (dir, amount) = line.split_once(' ').unwrap();
        let amount: usize = amount.parse().unwrap();

        println!("== {line} ==");

        let mut mvmt: (isize, isize);
        for a in 0..amount {
            match dir {
                "R" => {
                    {
                        tails[0].0 += 1;
                        mvmt = (1, 0);
                    };
                }
                "U" => {
                    {
                        tails[0].1 += 1;
                        mvmt = (0, 1);
                    };
                }
                "L" => {
                    tails[0].0 -= 1;
                    mvmt = (-1, 0);
                }
                "D" => {
                    tails[0].1 -= 1;
                    mvmt = (0, -1)
                }
                _ => {
                    panic!("invalid character")
                }
            }
            if a == 4 {
                println!()          
            }

            for i in 0..9 {
                if !are_close_enough(tails[i], tails[i + 1]) {
                    let previous = tails[i + 1];
                    if mvmt.0 == 0 || mvmt.1 == 0 {
                        // straight movement
                        tails[i + 1] = tails[i];
                        tails[i + 1].0 += -(mvmt.0);
                        tails[i + 1].1 += -(mvmt.1);
                    } else {
                        // diagonal movement
                        tails[i + 1] = find_closest(tails[i], tails[i + 1]);
                    }
                    mvmt = calc_mvmnt(previous, tails[i + 1]);
                }
            }

            visited.insert(tails[9], true);
            visualize(tails);
            println!("{a}")
        }
    }
    println!("{:?}", visited.len());
    // draw(visited)
}

fn visualize(tails: [(isize, isize); 10]) {
    let mut grid: [[String; 30]; 30] = Default::default();
    for (i, t) in tails.iter().enumerate() {
        let s = if i == 0 {
            String::from("H")
        } else {
            i.to_string()
        };
        grid[(14 - t.1) as usize][(14 + t.0) as usize] = s
    }
    grid[14][14] = "s".to_string();

    for g in grid {
        let mut builder = String::new();
        for v in g {
            if v.is_empty() {
                builder.push('.');
            } else {
                builder.push_str(v.as_str())
            }
        }
        println!("{builder}")
    }
}

fn calc_mvmnt(f: (isize, isize), s: (isize, isize)) -> (isize, isize) {
    (s.0 - f.0, s.1 - f.1)
}

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

fn find_closest(head: (isize, isize), tail: (isize, isize)) -> (isize, isize) {
    let mut m = HashMap::from([
        ((tail.0 - 1, tail.1 - 1), 0.),
        ((tail.0 - 1, tail.1 + 1), 0.),
        ((tail.0 + 1, tail.1 - 1), 0.),
        ((tail.0 + 1, tail.1 + 1), 0.),
    ]);
    m.iter_mut()
        .for_each(|((x, y), value)| *value = theorem((x - head.0) as f64, (y - head.1) as f64));

    m.iter()
        .min_by(|x, y| x.1.total_cmp(y.1))
        .map(|((x, y), value)| (*x, *y))
        .unwrap()
}
