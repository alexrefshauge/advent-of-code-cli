package cmd

const (
	solutionFile = `package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %%d solution (part %%d)\n", %[1]d, part)

	var answer string
	if part == 1 {
		answer = part1(input)
	} else if part == 2 {
		answer = part2(input)
	} else {
		panic(fmt.Errorf("part must be 1 or 2 %%d", part))	
	}

	fmt.Println("Output:", answer)
	err := os.WriteFile(fmt.Sprintf("output%%d.txt", part), []byte(answer), 0644)
	if err != nil {
		panic(err)
	}
}

func part1(input string) string {
	return input
}

func part2(input string) string {
	return input
}
`

	solutionTestFile = `package main

import (
	_ "embed"
	"testing"
)

//go:embed example1.txt
var example1 string

//go:embed example2.txt
var example2 string

func Test_day1_part1(t *testing.T) {
	innerTest(t, example1, 1, part1)
}

func Test_day1_part2(t *testing.T) {
	innerTest(t, example2, 2, part2)
}

func innerTest(t *testing.T, input string, part int, solutionFunc func(string) string) {
	if input == "" {
		t.Fatalf("please provide example input (part %d)", part)
	}
	in, out := parseTestInput(input)
	answer := solutionFunc(in)
	if answer != out {
		t.Fatalf("wanted %s, got %s", out, answer)
	}
}

func parseTestInput(input string) (string, string) {
	var token rune
	i := len(input)-1
	for token != ':' {
		token = rune(input[i])
		i--
	}

	return input[:i], input[i+2:]
}`
)
