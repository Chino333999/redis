package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	fn := func(tx *redis.Tx) error {

		v, err := tx.Get(ctx, "key").Int()
		if err != nil && err != redis.Nil {
			return err
		}

		v++

		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, "key", v, 0)
			return nil
		})

		return err
	}

	for i := 0; i < 3; i++ {
		err := rdb.Watch(ctx, fn, "key")
		if err == nil {
			break
		}
		if err == redis.TxFailedErr {
			continue
		}
	}
}
