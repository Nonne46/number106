package config

type ConfigModel struct {
	Num NumConfigModel `json:"num"`
}

type NumConfigModel struct {
	URL   string `json:"url"`
	Token string `json:"token"`
	Cache string `json:"cache"`
}

var DefaultConfig = ConfigModel{
	Num: NumConfigModel{
		URL:   "http://127.0.0.1:2386/",
		Cache: "./client_cache",
		Token: "",
	},
}
