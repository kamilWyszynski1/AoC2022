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
    let mut elves: Vec<i32> = Vec::new();
    let mut current = 0;
    for line in lines {
        if line.is_empty() {
            elves.push(current);
            current = 0;
            continue;
        }
        current += line.parse::<i32>().unwrap();
    }
    elves.push(current);
    elves.sort();
    let sum: i32 = elves.iter().rev().take(3).sum();
    println!("{}", sum)
}
