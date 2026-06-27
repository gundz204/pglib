package base

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	vip *viper.Viper
}

func InitConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		return nil
	}

	v := viper.New()

	v.SetConfigFile("config.yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil
	}

	return &Config{
		vip: v,
	}
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
