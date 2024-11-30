package configuration

import (
	"fmt"
	"github.com/alexrefshauge/advent-of-code-cli/configurationKeys"
	"github.com/spf13/cobra"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

const (
	PromptYes = "yes"
	PromptNo  = "no"
)

func UseAsSolutionDirectoryPrompt() error {
	prompt := promptui.Select{
		Label: "Use current directory as solution directory?",
		Items: []string{PromptYes, PromptNo},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	if result == PromptNo {
		return nil
	}

	currentDir, err := os.Getwd()
	fmt.Println(currentDir)
	if err != nil {
		return err
	}
	viper.Set(configurationKeys.SolutionDirectory, currentDir)
	cobra.CheckErr(viper.WriteConfig())
	return nil
}
