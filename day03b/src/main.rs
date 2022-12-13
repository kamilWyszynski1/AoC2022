#![feature(array_windows)]

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
    let mut groups: Vec<(String, String, String)> = vec![];

    for i in (0..lines.len()).step_by(3) {
        groups.push((
            lines[i].to_string(),
            lines[i + 1].to_string(),
            lines[i + 2].to_string(),
        ));
    }

    for (first, second, third) in groups {
        // let (first, second, third) = (&line[0], &line[1], &line[2]);

        println!("group: {},{},{}", first, second, third);

        let fh: HashSet<char> = HashSet::from_iter(first.chars());
        let sh: HashSet<char> = HashSet::from_iter(second.chars());
        let th: HashSet<char> = HashSet::from_iter(third.chars());

        for v in fh {
            if sh.contains(&v) && th.contains(&v) {
                let value = if v.is_lowercase() {
                    v as u32 - 96
                } else {
                    v as u32 - 38
                };
                println!("value of {v} is {value}");
                sum += value;
            }
        }
    }
    println!("{sum}")
}
