/*
Copyright © 2024 Alexander Refshauge
*/
package cmd

import (
	"fmt"
	"github.com/alexrefshauge/advent-of-code-cli/configuration"
	"github.com/alexrefshauge/advent-of-code-cli/configurationKeys"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new solution, will default to today",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(newSolution())
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func newSolution() error {
	fmt.Println("new solution")
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

	currentDirectory := path.Join(solutionsDirectory, fmt.Sprintf("%d", year), fmt.Sprintf("%d", day))
	err = os.MkdirAll(currentDirectory, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
