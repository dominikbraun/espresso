package core

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
