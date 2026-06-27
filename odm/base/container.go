package base

import (
	"time"

	"github.com/gundz204/pglib/infra/database"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	vip *viper.Viper
	db  *sqlx.DB
}

func (c *Config) InitConfig(path string) *Config {
	err := godotenv.Load()

	if err != nil {
		return nil
	}

	v, err := c.init(path)

	if err != nil {
		return nil
	}

	db, err := database.Connect(c.vip)
	if err != nil {
		return nil
	}

	return &Config{
		vip: v,
		db:  db,
	}
}

func (c *Config) init(path string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func (c *Config) GetConfig() *viper.Viper {
	return c.vip
}

func (c *Config) GetString(key string) string {
	return c.vip.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.vip.GetInt(key)
}

func (c *Config) GetBool(key string) bool {
	return c.vip.GetBool(key)
}

func (c *Config) GetFloat64(key string) float64 {
	return c.vip.GetFloat64(key)
}

func (c *Config) GetDuration(key string) time.Duration {
	return c.vip.GetDuration(key)
}

func (c *Config) GetStringSlice(key string) []string {
	return c.vip.GetStringSlice(key)
}

func (c *Config) Get(key string) any {
	return c.vip.Get(key)
}
