/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("config", "c", specifiedConfigFile, "Set config file")
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	configPath := path.Join(home, ".aoc", "config")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("No config file found, creating a new one at %s\n", configPath)
			cobra.CheckErr(createConfig(configPath))
		} else {
			cobra.CheckErr(err)
		}
	}

	cobra.CheckErr(viper.ReadInConfig())
	fmt.Println(viper.ConfigFileUsed())
}

func createConfig(configPath string) error {
	err := os.MkdirAll(path.Dir(configPath), os.ModePerm)
	if err != nil {
		return err
	}
	viper.Set("year", time.Now().Year())
	fmt.Println(viper.AllKeys())
	return viper.WriteConfigAs(configPath)
}
