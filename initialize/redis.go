package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"mtsw/global"
)

func Redis() {
	config := global.GVA_CONFIG.Redis
	fmt.Println("redis")
	client := redis.NewClient(&redis.Options{
		Addr: config.Addr,
		Password: config.Password,
		DB: config.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败")
	} else{
		fmt.Println("redis connect result is:", pong)
		global.GVA_REDIS = client
	}
}
