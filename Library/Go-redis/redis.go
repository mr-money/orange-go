package Go_redis

import "github.com/go-redis/redis/v8"

var Redis *redis.Client

//
// init
// @Description: 初始化链接
//
func init() {
	Redis = connectRedis()
}

//
// connectRedis
// @Description: 连接redis
// @return *redis.Client
//
func connectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
