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
    // X - rock(1), Y - peper(2), Z - scissors(3)
    let possibilities: HashMap<&str, u8> = HashMap::from([
        ("A Y", 6 + 2),
        ("A X", 3 + 1),
        ("A Z", 0 + 3),
        ("B Y", 3 + 2),
        ("B X", 0 + 1),
        ("B Z", 6 + 3),
        ("C Y", 0 + 2),
        ("C X", 6 + 1),
        ("C Z", 3 + 3),
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
