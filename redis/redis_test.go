package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // redis地址
		// Password: "",               // redis密码，没有则留空
		// DB:       0,                // 默认数据库，默认是0
	})

	defer func(*redis.Client) {
		rdb.Close()
	}(rdb)

	// err := rdb.Set(ctx, "cy", "123", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "cy").Result()
	// if err != redis.Nil {
	// 	panic(err)
	// }
	// log.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	log.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	log.Println("key2", val2)
	// }

	val3, err := rdb.GetSet(ctx, "cy", "321").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("new cy : ", val3)
}
