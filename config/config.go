// Package config provides the global Espresso configuration.
package config

import (
	"github.com/spf13/viper"
)

// Config represents the Espresso configuration provided by the
// user in espresso.yml.
type Config struct {
	Site struct {
		Meta struct {
			Title    string
			Subtitle string
			Author   string
			Base     string
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

// ReadConfig reads a YAML, TOML or JSON configuration file from
// the filesystem and converts it into a Config instance.
//
// Returns an error if the file can't be found or parsed.
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
