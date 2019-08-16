package redis

import (
	"github.com/gomodule/redigo/redis"
)

var redisInstance *redis.Pool

func Init(address, password string) {
	redisInstance = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 0,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

type redisMethod struct {
	redis.Conn
}

func GetConn() *redisMethod {
	return &redisMethod{
		redisInstance.Get(),
	}
}
