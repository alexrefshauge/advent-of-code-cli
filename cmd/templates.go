package cmd

const (
	solutionFile = `package main

import "fmt"

func main() {
	fmt.Printf("Running AoC day %%d solution\n", %[1]d)
}
`

	solutionTestFile = `package main

import "testing"

func Test_day%[1]d(t *testing.T) {
	
}`
)
