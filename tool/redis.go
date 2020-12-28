package tool

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	// 定义常量
	RedisClient    *redis.Pool
	REDIS_HOST     string
	REDIS_DB       int
	REDIS_PASSWORD string
)

func Redisinit() *redis.Pool {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("redis.db")
	REDIS_PASSWORD = beego.AppConfig.String("redis.password")

	// 建立连接池
	return &redis.Pool{
			// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 400),/// 最大空闲连接数
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 500),// 一个pool所能分配的最大的连接数目
		// 当设置成0的时候，该pool连接数没有限制
		IdleTimeout: 180 * time.Second, // 空闲连接超时时间，超过超时时间的空闲连s接会被关闭。
		// 如果设置成0，空闲连接将不会被关闭
		// 应该设置一个比redis服务端超时时间更短的时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)

			if err != nil {
				return nil, err
			}
			if REDIS_PASSWORD != "" {
				if _, err := c.Do("AUTH", REDIS_PASSWORD); err != nil {
					c.Close()
					return nil, err
				}
			}

			// 选择db
			//c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}
