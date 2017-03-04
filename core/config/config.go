package config

import (
	"bytes"
	"io/ioutil"
	"os"

	"fmt"

	"github.com/spf13/viper"
)

const (
	WEB_ADDRESS            string = "web.address"
	WEB_PORT               string = "web.port"
	WEB_TEMPLATE_DIR       string = "web.template-dir"
	PG_URI                 string = "pg.uri"
	PG_MAX_CONNECTION      string = "pg.max-connection"
	PG_MAX_IDLE_CONNECTION string = "pg.max-idle-connection"
)

type WebConfig struct {
	Address     string
	Port        int
	TemplateDir string
}

type PgConfig struct {
	Uri               string
	MaxConnection     int
	MaxIdleConnection int
}

type Config struct {
	*viper.Viper
}

func setDefaults(v *viper.Viper) {
	// set sensible default configuration
	// web configuration
	v.SetDefault(WEB_ADDRESS, "localhost")
	v.SetDefault(WEB_PORT, 8000)
	v.SetDefault(WEB_TEMPLATE_DIR, "./assets/")

	// db stores
	v.SetDefault(PG_URI, "")
	v.SetDefault(PG_MAX_CONNECTION, 10)
	v.SetDefault(PG_MAX_IDLE_CONNECTION, 8)
}

func loadConfPath(v *viper.Viper, path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return v.ReadConfig(bytes.NewBuffer(f))
}

func (c *Config) Web() *WebConfig {
	return &WebConfig{
		Address:     c.GetString(WEB_ADDRESS),
		Port:        c.GetInt(WEB_PORT),
		TemplateDir: c.GetString(WEB_TEMPLATE_DIR),
	}
}

func (c *Config) Postgres() *PgConfig {
	return &PgConfig{
		Uri:               c.GetString(PG_URI),
		MaxConnection:     c.GetInt(PG_MAX_CONNECTION),
		MaxIdleConnection: c.GetInt(PG_MAX_IDLE_CONNECTION),
	}
}

func New(path ...string) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	setDefaults(v)

	if len(path) > 0 {
		if err := loadConfPath(v, path[0]); err != nil {
			fmt.Printf("Failed load from file. Osborn using default configuration")
		}
	}

	return &Config{v}, nil
}
