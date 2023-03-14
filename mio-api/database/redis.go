package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"mio-api/config"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.GetAddr(),
		Password: config.Conf.Redis.Password, // no password set
		DB:       config.Conf.Redis.DB,       // use default DB
	})

	fmt.Println("[Success] Redis数据库连接成功！！！")
	Client = client
}
