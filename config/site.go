// Package config provides user-defined type-safe config as structs
// as well as functions for populating them from configuration files.
package config

import (
	"github.com/spf13/viper"
)

// Site concludes all user-defined site settings which are typically
// defined in the site.yml file. It holds content-related configuration
// values for complementing and overriding generated default values.
type Site struct {
	Title       string
	Subtitle    string
	BaseURL     string
	Description string
	Author      string
	Nav         struct {
		Items []struct {
			Label  string
			Target string
		}
		Override bool
	}
	Footer struct {
		Text  string
		Items []struct {
			Label  string
			Target string
		}
	}
}

// FromFile parses any configuration file (YAML, TOML or JSON) with the
// specified name in the specified path and unmarshals its values into
// the destination. dest has to be a pointer value.
//
// FromFile does not return an error if the configuration file could not
// be found since Espresso doesn't require any configuration.
func FromFile(path, filename string, dest interface{}) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return err
	}

	return viper.Unmarshal(dest)
}
