package configuration

import (
	"errors"
	"fmt"
	"github.com/alexrefshauge/advent-of-code-cli/configurationKeys"
	"time"

	"github.com/spf13/viper"
)

// GetDate returns the configured year and day as integers
func GetDate() (int, int, error) {
	if !viper.IsSet(configurationKeys.Year) || !viper.IsSet(configurationKeys.Day) {
		fmt.Println("Current day has not been configured, using today")
		if err := SetDateToday(); err != nil {
			return 0, 0, nil
		}
	}
	return viper.GetInt(configurationKeys.Year),
		viper.GetInt(configurationKeys.Day), nil
}

func SetDateToday() error {
	today := time.Now()
	if today.Month() != time.December {
		return errors.New("Ooops! Can't use today, since it is not December")
	}
	viper.Set(configurationKeys.Year, today.Year())
	viper.Set(configurationKeys.Day, today.Day())
	viper.WriteConfig()
	return nil
}
