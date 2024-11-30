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
	for _, key := range viper.AllKeys() {
		fmt.Printf("%s: %v\n", key, viper.Get(key))
	}
	cmd.Execute()
}
