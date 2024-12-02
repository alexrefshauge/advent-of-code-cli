# advent-of-code-cli

# Configuration
A configuration file `config.yaml` is plaed in your home directory in a `.aoc` directory.

To enable the tool to fetch puzzle inputs automatically, please configure the `sessionCookie` setting. This is the cookie stored in the browser, if you are logged in to the Advent of Code website.

Example configuration file
```yaml
sessionCookie: <sessionCookie>
solutiondirectory: /home/john/advent-of-code
year: 2024
day: 1
```

# Usage
To create a new solution for the current day:
`advent-of-code-cli new --today`
Use the `--today` flag, to ensure that the current day is used, if you have not configured day and year.

## Using the tests
To run the generated tests, you must first populate the test input files called `example1.txt` and `example2.txt`. These should contain the example input, often included in the puzzle description, followed by the solution prefixed by a `:`.
Example:
```
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
:2
```
