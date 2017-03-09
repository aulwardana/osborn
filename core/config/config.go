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
	MGO_HOSTS        string = "mgo.hosts"
	MGO_USER         string = "mgo.user"
	MGO_PASSWORD     string = "mgo.password"
	MGO_DBNAME       string = "mgo.dbname"
	RDS_HOST         string = "rds.host"
	RDS_PORT         string = "rds.port"
	RDS_PASSWORD     string = "rds.password"
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

type MgoConfig struct {
	Hosts    string
	User     string
	Password string
	DBName   string
}

type RdsConfig struct {
	Host     string
	Port     int
	Password string
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

	// db postgres (SQL DB)
	v.SetDefault(PG_HOST, "localhost")
	v.SetDefault(PG_PORT, 5432)
	v.SetDefault(PG_USER, "postgres")
	v.SetDefault(PG_PASSWORD, "postgres")
	v.SetDefault(PG_DBNAME, "postgres")
	v.SetDefault(PG_SSLMODE, "disable")

	// db mongo (NO-SQL DB)
	v.SetDefault(MGO_HOSTS, "localhost:27017")
	v.SetDefault(MGO_USER, "admin")
	v.SetDefault(MGO_PASSWORD, "admin")
	v.SetDefault(MGO_DBNAME, "admin")

	// db redis (KEY VALUE STORE DB)
	v.SetDefault(RDS_HOST, "localhost")
	v.SetDefault(RDS_PORT, 6379)
	v.SetDefault(RDS_PASSWORD, "admin")
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

func (c *Config) MongoDB() *MgoConfig {
	return &MgoConfig{
		Hosts:    c.GetString(MGO_HOSTS),
		User:     c.GetString(MGO_USER),
		Password: c.GetString(MGO_PASSWORD),
		DBName:   c.GetString(MGO_DBNAME),
	}
}

func (c *Config) Redis() *RdsConfig {
	return &RdsConfig{
		Host:     c.GetString(RDS_HOST),
		Port:     c.GetInt(RDS_PORT),
		Password: c.GetString(RDS_PASSWORD),
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
