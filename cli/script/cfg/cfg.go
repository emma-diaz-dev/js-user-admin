package cfg

import "github.com/joeshaw/envdecode"

var (
	config Config
)

type (
	Config struct {
		OwnURL string `env:"OWN_URL,default=http://127.0.0.1:5050"`
	}
)

func initCfg() {
	if err := envdecode.Decode(&config); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	initCfg()
	return config
}
