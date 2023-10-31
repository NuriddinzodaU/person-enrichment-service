package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
	"strings"
)

var NewConfig = fx.Provide(newConfig)

type IConfig interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetString(key string) string
}

type config struct {
	cfg *viper.Viper
}

func newConfig() IConfig {

	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("json")

	cfg.AddConfigPath(getConfigPath())

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg.WatchConfig()

	return &config{cfg: cfg}
}

func (c *config) Get(key string) interface{} {
	return c.cfg.Get(key)
}

func (c *config) GetBool(key string) bool {
	return c.cfg.GetBool(key)
}

func (c *config) GetString(key string) string {
	return c.cfg.GetString(key)
}

func getConfigPath() (path string) {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	slice := strings.Split(wd, "person-enrichment-service")
	path = slice[0] + "person-enrichment-service/internal/config"
	return
}
