package config

import (
	"bytes"
	"io/ioutil"
	"os"

	"fmt"

	"github.com/spf13/viper"
)

const (
	WEB_ADDRESS      string = "web.address"
	WEB_PORT         string = "web.port"
	WEB_TEMPLATE_DIR string = "web.template-dir"
	PG_HOST          string = "pg.host"
	PG_PORT          string = "pg.port"
	PG_USER          string = "pg.user"
	PG_PASSWORD      string = "pg.password"
	PG_DBNAME        string = "pg.dbname"
	PG_SSLMODE       string = "pg.sslmode"
)

type WebConfig struct {
	Address     string
	Port        int
	TemplateDir string
}

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
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
	v.SetDefault(PG_HOST, "localhost")
	v.SetDefault(PG_PORT, 5432)
	v.SetDefault(PG_USER, "postgres")
	v.SetDefault(PG_PASSWORD, "postgres")
	v.SetDefault(PG_DBNAME, "postgres")
	v.SetDefault(PG_SSLMODE, "disable")
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
		Host:     c.GetString(PG_HOST),
		Port:     c.GetInt(PG_PORT),
		User:     c.GetString(PG_USER),
		Password: c.GetString(PG_PASSWORD),
		DBName:   c.GetString(PG_DBNAME),
		SSLMode:  c.GetString(PG_SSLMODE),
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
