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
		panic(fmt.Errorf("unable to run part %%d", part))	
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

import "testing"

func Test_day%[1]d(t *testing.T) {
	
}`
)
