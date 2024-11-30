package cmd

const (
	solutionFile = `package day%[1]d

func main() {
	fmt.Printf("Running AoC day %%d solution\n", %[1]d)
}
`

	solutionTestFile = `package day%[1]d

func Test_day%[1](t *testing.T) {}
	
)`
)
