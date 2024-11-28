/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "Handle Advent of Code solutions",
		Long:  `The AoC cli can be used to create and run Advent of Code solutions`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("config", "c", configFile, "Set config file")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	configPath := path.Join(home, ".aoc", "config")

	viper.AddConfigPath(path.Dir(configPath))
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err = os.MkdirAll(path.Dir(configPath), os.ModePerm)
	cobra.CheckErr(err)
	viper.SafeWriteConfigAs(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
