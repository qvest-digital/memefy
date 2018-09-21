package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// Set during build: use go build -X memefy/pkg/config.CommitShortHash ...
var (
	CommitShortHash string
	BuildNumber     string
	BuildTime       string
)

var config *Config
var once sync.Once

type Config struct {
	Server   Server
	Log      Log
	Security Security
}

type Server struct {
	Port int
}

type Log struct {
	Level string // one of debug, info, warn, error, fatal
}

type Security struct {
	EnableBasicAuth bool
	BasicAuthUser   string
	BasicAuthPass   string
}

// Get reads the config.yml, overrides values with environment variables
// and returns a singleton Config struct
func Get() *Config {
	once.Do(func() {
		viper.AddConfigPath("/var/service")
		viper.AddConfigPath("$HOME/.memefy")
		viper.AddConfigPath("$GOPATH/src/memefy/server")

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

		config = &Config{}
		err = viper.Unmarshal(config)

		if config.Log.Level == "" {
			config.Log.Level = "info"
		}

		if err != nil {
			panic(err)
		}
	})
	return config
}
