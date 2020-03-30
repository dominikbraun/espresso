package settings

import "github.com/spf13/viper"

type Site struct {
	Name string
	Nav  struct {
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

func FromFile(path, filename string, dest interface{}) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)

	return viper.Unmarshal(dest)
}
