/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	specifiedConfigFile string

	rootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "Handle Advent of Code solutions",
		Long:  `The AoC cli can be used to create and run Advent of Code solutions`,
	}
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", specifiedConfigFile, "Set config file")
	rootCmd.PersistentFlags().Bool("today", false, "use Viper for configuration")
}
