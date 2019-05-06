package redis

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"gitlab.meitu.com/mt-family/mtfamily-shop/context"
	"gitlab.meitu.com/mt-family/mtfamily-shop/errors"
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

// RLock redis分布式锁
func RLock(key string, notWaiting ...bool) error {
	redisConn := GetConn()
	defer redisConn.Close()
	i := 0
	for {
		ok, err := redisConn.SetNX(key, time.Now().Unix()+10)
		if err != nil {
			errObj := errors.ErrInternalError
			errObj.Log = "[Utils - RLock] 加锁失败 error:" + err.Error()
			return errObj
		}

		if !ok {
			if i == 6 {
				return errors.ErrServiceBusy
			}

			value, err := redisConn.Get(key)
			if err != nil {
				errObj := errors.ErrInternalError
				errObj.Log = "[Utils - RLock] 读取锁状态失败 error:" + err.Error()
				return errObj
			}
			expire, err := strconv.ParseInt(value, 10, 64)
			if time.Now().Unix() > expire {
				redisConn.Del(key)
			} else if len(notWaiting) > 0 && notWaiting[0] {
				return errors.ErrServiceBusy
			} else {
				time.Sleep(time.Duration(200) * time.Millisecond)
			}

			i++
			continue
		}

		break
	}
	return nil
}

// RUnLock 解除分布式锁
func RUnLock(key string) error {
	redisPool := context.GetRedisPool()
	defer redisPool.Close()
	_, err := redisPool.Del(key)
	return err
}
