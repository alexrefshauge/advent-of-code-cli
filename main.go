/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/alexrefshauge/advent-of-code-cli/cmd"
	"github.com/alexrefshauge/advent-of-code-cli/configuration"
	"github.com/spf13/viper"
)

func main() {
	configuration.Init()
	fmt.Println(viper.AllKeys())
	cmd.Execute()
}
