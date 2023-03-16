package redis

import (
	"context"
	"log"
	"strings"

	redisLib "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/nguyenvanxuanvu/register_course_check/redis/redisconfig"
	"go.uber.org/fx"
)

func NewCache(lifecycle fx.Lifecycle) redisconfig.Cache {
	addresses := viper.GetString("redis.addresses")

	if len(addresses) == 0 {
		log.Fatal("Invalid redis address")
	}

	client := redisLib.NewUniversalClient(&redisLib.UniversalOptions{
		Addrs: strings.Split(addresses, ","),
	})

	

	log.Println("Trying to connect redis...")
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("Can not connect redis, address=", addresses, err)
	}

	log.Println("Connect redis successfully")

	lifecycle.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		log.Println("Closing redis connection")
		return client.Close()
	}})

	return &redisconfig.RedisCache{UniversalClient: client}
}
