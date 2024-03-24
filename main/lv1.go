package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

func main() {

	ctx := context.Background()

	err := rdb.ZAdd(ctx, "key", redis.Z{Score: 3, Member: "first"},
		redis.Z{Score: 2, Member: "second"},
		redis.Z{Score: 1, Member: "third"}).Err()
	if err != nil {
		panic(err)
	}

	op := redis.ZRangeBy{
		Min:    "0",  // 最小分数
		Max:    "10", // 最大分数
		Offset: 0,    // 开始偏移量
		Count:  5,    // 一次返回多少数据
	}

	vals, err := rdb.ZRevRangeByScoreWithScores(ctx, "key", &op).Result()
	if err != nil {
		panic(err)
	}

	for _, val := range vals {
		fmt.Println("Name:", val.Member, "Score:", val.Score)
	}
}
