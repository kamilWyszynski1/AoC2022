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
    let mut grid = vec![];

    for line in lines {
        let mut l = vec![];
        for ch in line.chars() {
            let value: usize = ch.to_string().parse().unwrap();
            l.push(value);
        }
        grid.push(l)
    }
    println!("{}", calculate_visible(grid))
}

fn calculate_visible(grid: Vec<Vec<usize>>) -> usize {
    let mut visible = 0;

    let mut max = 0;
    for (i, line) in grid.iter().enumerate() {
        for (j, value) in line.iter().enumerate() {
            // top
            let t_visible = check_view(grid.clone(), *value, i, j, |x, y| (Some(x + 1), Some(y)));

            // bottom
            let b_visible = if i == 0 {
                0
            } else {
                check_view(grid.clone(), *value, i, j, |x, y| {
                    (x.checked_sub(1), Some(y))
                })
            };

            // left
            let l_visible = check_view(grid.clone(), *value, i, j, |x, y| {
                (Some(x), y.checked_sub(1))
            });

            // right
            let r_visible = if j == 0 {
                0
            } else {
                check_view(grid.clone(), *value, i, j, |x, y| (Some(x), Some(y + 1)))
            };

            let tree_view = t_visible * b_visible * l_visible * r_visible;
            if tree_view > max {
                max = tree_view
            }
        }
    }
    max
}

fn check_view(
    grid: Vec<Vec<usize>>,
    value: usize,
    mut x: usize,
    mut y: usize,
    f: fn(usize, usize) -> (Option<usize>, Option<usize>),
) -> usize {
    let mut view = 0;
    let mut should_break = false;

    loop {
        match f(x, y) {
            (None, None) | (Some(_), None) | (None, Some(_)) => break,
            (Some(i), Some(j)) => {
                x = i;
                y = j;
            }
        }

        let n = match grid.get(x) {
            Some(l) => match l.get(y) {
                Some(value) => value,
                None => {
                    should_break = true;
                    &0
                }
            },
            None => {
                should_break = true;
                &0
            }
        };
        if should_break {
            break;
        }
        if *n >= value {
            view += 1;
            break;
        }
        view += 1
    }
    view
}
