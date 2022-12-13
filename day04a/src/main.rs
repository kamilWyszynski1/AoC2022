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
    let mut sum = 0;
    for line in lines {
        let (f, s) = line.split_once(',').unwrap();

        let first = range_to_tuple(f);
        let second = range_to_tuple(s);

        if contains_whole_range(first, second) || contains_whole_range(second, first) {
            sum += 1
        }
    }
    println!("{sum}")
}

fn range_to_tuple(r: &str) -> (usize, usize) {
    let (l, m) = r.split_once('-').unwrap();
    (l.parse().unwrap(), m.parse().unwrap())
}

fn contains_whole_range(f: (usize, usize), s: (usize, usize)) -> bool {
    f.0 <= s.0 && f.1 >= s.1
}
