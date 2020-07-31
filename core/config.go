package core

import (
	"github.com/spf13/viper"
)

type Config struct {
	Site struct {
		Meta struct {
			Title    string
			Subtitle string
			Author   string
			BaseURL  string
		}
		Nav struct {
			Items []struct {
				Label  string
				Target string
			}
			Override bool
		}
		Footer struct {
			Items []struct {
				Label  string
				Target string
			}
			Override bool
		}
	}
}

func ReadConfig(path, filename string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
