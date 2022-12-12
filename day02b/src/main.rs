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
    // A - rock, B - peper, C - scissors
    // rock(1),  peper(2), scissors(3)
    // Y - draw, X - lose, Z - win
    let possibilities: HashMap<&str, u8> = HashMap::from([
        ("A Y", 3 + 1),
        ("A X", 0 + 3),
        ("A Z", 6 + 2),
        ("B Y", 3 + 2),
        ("B X", 0 + 1),
        ("B Z", 6 + 3),
        ("C Y", 3 + 3),
        ("C X", 0 + 2),
        ("C Z", 6 + 1),
    ]);

    let m1: HashMap<&str, &str> = HashMap::from([("A", "rock"), ("B", "peper"), ("C", "scissors")]);
    let m2: HashMap<&str, &str> = HashMap::from([("X", "rock"), ("Y", "peper"), ("Z", "scissors")]);

    let mut sum: u32 = 0;
    for line in lines {
        let value = *possibilities.get(line.as_str()).unwrap() as u32;
        println!("{line} - {value}");
        sum += value;
    }
    println!("{sum}")
}
