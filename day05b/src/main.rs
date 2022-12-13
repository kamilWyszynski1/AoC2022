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
    let mut stacks = HashMap::from([
        (1, vec!["V", "B", "J", "D"]),
        (2, vec!["P", "V", "B", "W", "R", "D", "F"]),
        (3, vec!["R", "G", "F", "L", "D", "C", "W", "Q"]),
        (4, vec!["W", "J", "P", "M", "L", "N", "D", "B"]),
        (5, vec!["H", "N", "B", "P", "C", "S", "Q"]),
        (6, vec!["R", "D", "B", "S", "N", "G"]),
        (7, vec!["Z", "B", "P", "M", "Q", "F", "S", "H"]),
        (8, vec!["W", "L", "F"]),
        (9, vec!["S", "V", "F", "M", "R"]),
    ]);

    for line in lines {
        let mut split = line.split(' ');

        let (amount, from, to) = (
            split.nth(1).unwrap(),
            split.nth(1).unwrap(),
            split.nth(1).unwrap(),
        );

        println!("amount: {}, from: {}, to: {}", amount, from, to);
        let amount: usize = amount
            .parse()
            .unwrap_or_else(|_| panic!("could not parse amount: {}, in line {line}", amount));
        let from: usize = from.parse().unwrap();
        let to: usize = to.parse().unwrap();

        let mut poped = pop_n(stacks.get_mut(&from).unwrap(), amount);
        stacks.get_mut(&to).unwrap().append(&mut poped);
    }
    let mut msg = String::new();
    for i in 1..10 {
        msg.push_str(stacks.get(&i).unwrap().last().unwrap())
    }
    println!("{msg}")
}

fn pop_n<'a>(v: &mut Vec<&'a str>, n: usize) -> Vec<&'a str> {
    let mut poped = vec![];
    for _ in 0..n {
        poped.push(v.pop().unwrap())
    }
    poped.reverse();
    poped
}
