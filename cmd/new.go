/*
Copyright © 2024 Alexander Refshauge
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/alexrefshauge/advent-of-code-cli/configuration"
	"github.com/alexrefshauge/advent-of-code-cli/configurationKeys"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new solution, will default to today",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(newSolution())
		},
	}
)

func init() {
	rootCmd.AddCommand(newCmd)
}

func newSolution() error {
	useToday, err := rootCmd.Flags().GetBool("today")
	if err != nil {
		return err
	}
	if !viper.IsSet(configurationKeys.SolutionDirectory) {
		fmt.Println("Please configure a solution directory before creating a new solution")
		err := configuration.UseAsSolutionDirectoryPrompt()
		if err != nil {
			return err
		}
	}

	solutionsDirectory := viper.GetString(configurationKeys.SolutionDirectory)
	year, day, err := configuration.GetDate()
	if err != nil {
		return err
	}

	if useToday {
		today := time.Now()
		year = today.Year()
		day = today.Day()
		viper.Set("year", year)
		viper.Set("day", day)
	}

	currentDirectory := path.Join(solutionsDirectory, fmt.Sprintf("%d", year), fmt.Sprintf("day%d", day))
	err = os.MkdirAll(currentDirectory, os.ModePerm)
	if err != nil {
		return err
	}

	err = writeCode(currentDirectory, day)
	if err != nil {
		return err
	}

	initMod := exec.Command("go", "mod", "init", fmt.Sprintf("day%d", day))
	initMod.Dir = currentDirectory
	initMod.Stdout = os.Stdout
	initMod.Stderr = os.Stderr
	err = initMod.Run()
	if err != nil {
		return err
	}

	return nil
}

func writeCode(dir string, day int) error {
	err := os.WriteFile(path.Join(dir, "main.go"), []byte(fmt.Sprintf(solutionFile, day)), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(dir, "solution_test.go"), []byte(fmt.Sprintf(solutionTestFile, day)), os.ModePerm)
	if err != nil {
		return err
	}

	inputBody, err := getPuzzleInput(dir)
	defer inputBody.Close()
	all, err := io.ReadAll(inputBody)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(dir, "input.txt"), all, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(dir, "example1.txt"), make([]byte, 0), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(dir, "example2.txt"), make([]byte, 0), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func getPuzzleInput(dir string) (io.ReadCloser, error) {
	if !viper.IsSet(configurationKeys.SessionCookie) {
		fmt.Println("Unable to fetch puzzle input, session cookie not configured")
		return nil, errors.New("session cookie not configured")
	}

	session := viper.GetString(configurationKeys.SessionCookie)
	request, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", viper.GetInt(configurationKeys.Year), viper.GetInt(configurationKeys.Day)), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Cookie", fmt.Sprintf("session=%s;", session))
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
