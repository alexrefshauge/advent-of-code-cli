package cmd

import (
	"fmt"
	"io"
)

const (
	solutionFile = `package day%[1]d\n
func main() {
	fmt.Printf("", %[1]d)
}
`
)

func renderSolutionFile(day int8, file io.Writer) {
	content := fmt.Sprintf(solutionFile, day)
	file.Write([]byte(content))
}
