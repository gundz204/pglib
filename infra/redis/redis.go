package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func Connect(vip *viper.Viper) (*redis.Client, error) {
	addr := fmt.Sprintf(
		"%s:%d",
		vip.GetString("redis.host"),
		vip.GetInt("redis.port"),
	)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: vip.GetString("redis.password"),
		DB:       vip.GetInt("redis.db"),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
