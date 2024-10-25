package Cache

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/go-redis/redis/v8"
	"github.com/shockerli/cvt"
	"orange-go/Config"
	"runtime"
	"strings"
	"time"
)

var Redis *redis.Client
var Cxt = context.Background()

// init
// @Description: 初始化链接
func init() {
	Redis = connectRedis()

	if cvt.String(Redis.Ping(Cxt)) != "ping: PONG" {
		log.FATAL.Panicln(Redis.Ping(Cxt))
	} else {
		log.INFO.Println("Redis [" + Redis.Options().Addr + "]: Connect Success!")
	}
}

// connectRedis
// @Description: 连接redis
// @return *redis.Client
func connectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		//连接信息
		Network: "tcp", //网络类型，tcp or unix，默认tcp
		Addr: fmt.Sprintf("%s:%s", //主机名+冒号+端口，默认localhost:6379
			Config.Configs.Web.Redis.Host,
			Config.Configs.Web.Redis.Port,
		),
		Password: Config.GetFieldByName(Config.Configs.Web.Redis, "Pwd"),         //密码
		DB:       cvt.Int(Config.GetFieldByName(Config.Configs.Web.Redis, "Db")), // redis数据库index

		//连接池容量及闲置连接数量
		PoolSize:     4 * runtime.GOMAXPROCS(0), // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: 10,                        //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		//闲置连接检查包括IdleTimeout，MaxConnAge
		IdleCheckFrequency: 60 * time.Second, //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        5 * time.Minute,  //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:         0 * time.Second,  //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

		//命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
	})

	return rdb
}

// SetKey
// @Description: 生成缓存key
// @param key 嵌套key名 如 key1,key2,key3...
// @return string
func SetKey(key ...string) string {
	var keys []string
	keys = append(keys, key...)

	return strings.Join(keys, ":")
}

// RememberString
// @Description:不存在则写入缓存数据后返回 字符串
// @param key 缓存key
// @param value 缓存数据
// @param expiration 过期时间 0为不过期
// @return string
func RememberString(key string, value func() string, expiration time.Duration) string {
	//获取缓存数据
	data, err := Redis.Get(Cxt, key).Result()
	if err != nil {
		//缓存为空 返回传入数据
		if err.Error() == "redis: nil" || data == "" {
			if value() != "" {
				//写入缓存
				Redis.Set(
					Cxt,
					key, value(),
					expiration,
				)
			}

			return value()
		}

		log.ERROR.Println(err)

		return ""
	}

	return data
}

// RememberZScore
// @Description: 不存在则写入缓存数据后返回 有序集合
// @param key 缓存key
// @param member 成员名
// @param value 缓存数据
// @param expiration 过期时间 <=0为不过期
// @return float64
func RememberZScore(key, member string, value func() float64, expiration time.Duration) float64 {
	//获取缓存数据
	data, err := Redis.ZScore(Cxt, key, member).Result()
	if err != nil {
		//缓存为空 返回传入数据
		if err.Error() == "redis: nil" || data == 0 {

			if value() > 0 {
				//写入缓存
				zData := redis.Z{
					Member: member,
					Score:  value(),
				}

				Redis.ZAdd(
					Cxt,
					key, &zData,
				)
				if expiration > 0 {
					Redis.Expire(Cxt, key, expiration)
				}
			}

			return value()
		}

		log.ERROR.Println(err)

		return 0
	}

	return data
}

// RememberHash
// @Description: 不存在则写入缓存数据后返回 哈希
// @param key 缓存key
// @param field 字段名
// @param value 缓存值
// @param expiration 过期时间 <=0为不过期
// @return string
func RememberHash(key, field string, value func() string, expiration time.Duration) string {
	//获取缓存数据
	data, err := Redis.HGet(Cxt, key, field).Result()
	if err != nil {
		//缓存为空 返回传入数据
		if err.Error() == "redis: nil" || data == "" {
			if value() != "" {
				//写入缓存
				Redis.HSet(
					Cxt,
					key, field, value(),
				)
				if expiration > 0 {
					Redis.Expire(Cxt, key, expiration)
				}
			}

			return value()
		}

		log.ERROR.Println(err)

		return ""
	}

	return data
}
