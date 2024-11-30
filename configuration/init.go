package configuration

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	configPath := path.Join(home, ".aoc", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Printf("No config file found, creating a new one at %s\n", configPath)
		cobra.CheckErr(createConfig(configPath))
	} else {
		cobra.CheckErr(err)
	}

	viper.SetConfigFile(configPath)
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
