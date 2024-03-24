package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	ctx := context.Background()

	sub := rdb.Subscribe(ctx, "channel1")

	for message := range sub.Channel() {
		fmt.Println(message.Channel)
		fmt.Println(message.Payload)
	}
}
