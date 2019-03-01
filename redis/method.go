package redis

import (
	"errors"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var errForbidCommand error = errors.New("command was forbiden in shards")
var errWrongArguments error = errors.New("wrong number of arugments")
var errNoKey = errors.New("shards pool can't do command without key")

func isOKString(str string, err error) (bool, error) {
	if strings.ToUpper(str) == "OK" {
		return true, err
	}
	return false, err
}

func inStrArray(str string, array []string) bool {
	found := false
	for _, got := range array {
		if got == str {
			found = true
		}
	}
	return found
}

func (rm *redisMethod) Get(key string) (string, error) {
	return redis.String(rm.Do("GET", key))
}

func (rm *redisMethod) GetSet(key string, value interface{}) (string, error) {
	return redis.String(rm.Do("GETSET", key, value))
}

func (rm *redisMethod) Del(keys ...interface{}) (int, error) {
	return redis.Int(rm.Do("DEL", keys...))
}

func (rm *redisMethod) Exists(key ...interface{}) (int, error) {
	return redis.Int(rm.Do("EXISTS", key...))
}

func (rm *redisMethod) Expire(key string, expiration int64) (bool, error) {
	return redis.Bool(rm.Do("EXPIRE", key, expiration))
}

func (rm *redisMethod) ExpireAt(key string, tm int64) (bool, error) {
	return redis.Bool(rm.Do("EXPIREAT", key, tm))
}

func (rm *redisMethod) TTL(key string) (int, error) {
	return redis.Int(rm.Do("TTL", key))
}

func (rm *redisMethod) PTTL(key string) (int, error) {
	return redis.Int(rm.Do("PTTL", key))
}

func (rm *redisMethod) Type(key string) (string, error) {
	return redis.String(rm.Do("TYPE", key))
}

func (rm *redisMethod) Incr(key string) (int, error) {
	return rm.IncrBy(key, 1)
}

func (rm *redisMethod) IncrBy(key string, value int64) (int, error) {
	return redis.Int(rm.Do("INCRBY", key, value))
}

func (rm *redisMethod) IncrByFloat(key string, value float64) (float64, error) {
	return redis.Float64(rm.Do("INCRBYFLOAT", key, value))
}

func (rm *redisMethod) Decr(key string) (int, error) {
	return rm.DecrBy(key, 1)
}

func (rm *redisMethod) DecrBy(key string, value int64) (int, error) {
	return redis.Int(rm.Do("DECRBY", key, value))
}

func (rm *redisMethod) Append(key string, value string) (int, error) {
	return redis.Int(rm.Do("APPEND", key, value))
}

func (rm *redisMethod) GetRange(key string, start, end int64) (string, error) {
	return redis.String(rm.Do("GETRANGE", key, start, end))
}

func (rm *redisMethod) SetRange(key string, offset int64, value string) (int, error) {
	return redis.Int(rm.Do("SETRANGE", key, offset, value))
}

func (rm *redisMethod) MGet(keys ...interface{}) ([]string, error) {
	return redis.Strings(rm.Do("MGET", keys...))
}

func (rm *redisMethod) Set(key string, value interface{}) (bool, error) {
	return isOKString(redis.String(rm.Do("SET", key, value)))
}

func (rm *redisMethod) SetEX(key string, value interface{}, seconds int64) (bool, error) {
	return isOKString(redis.String(rm.Do("SET", key, value, "EX", seconds)))
}

func (rm *redisMethod) SetPX(key string, value interface{}, milliseconds int64) (bool, error) {
	return isOKString(redis.String(rm.Do("SET", key, value, "PX", milliseconds)))
}

func (rm *redisMethod) SetNX(key string, value interface{}) (bool, error) {
	return redis.Bool(rm.Do("SETNX", key, value))
}

func (rm *redisMethod) SetXX(key string, value interface{}) (bool, error) {
	return isOKString(redis.String(rm.Do("SET", key, value, "XX")))
}

func (rm *redisMethod) MSet(pairs ...interface{}) (bool, error) {
	return isOKString(redis.String(rm.Do("MSET", pairs...)))
}

func (rm *redisMethod) MSetNX(pairs ...interface{}) (bool, error) {
	return redis.Bool(rm.Do("MSETNX", pairs...))
}

func (rm *redisMethod) StrLen(key string) (int, error) {
	return redis.Int(rm.Do("STRLEN", key))
}

func (rm *redisMethod) GetBit(key string, offset int64) (int, error) {
	return redis.Int(rm.Do("GETBIT", key, offset))
}

func (rm *redisMethod) SetBit(key string, offset int64, value int) (int, error) {
	return redis.Int(rm.Do("SETBIT", key, offset, value))
}

func (rm *redisMethod) BitCount(key string, offsets ...int64) (int, error) {
	switch len(offsets) {
	case 0:
		return redis.Int(rm.Do("BITCOUNT", key))
	case 2:
		return redis.Int(rm.Do("BITCOUNT", key, offsets[0], offsets[1]))
	default:
		return 0, errWrongArguments
	}
}

func (rm *redisMethod) BitOpAnd(destKey string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, "AND", destKey)
	args = append(args, keys...)
	return redis.Int(rm.Do("BITOP", args...))
}

func (rm *redisMethod) BitOpOr(destKey string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, "OR", destKey)
	args = append(args, keys...)
	return redis.Int(rm.Do("BITOP", args...))
}

func (rm *redisMethod) BitOpXor(destKey string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, "XOR", destKey)
	args = append(args, keys...)
	return redis.Int(rm.Do("BITOP", args...))
}

func (rm *redisMethod) BitOpNot(destKey, key string) (int, error) {
	return redis.Int(rm.Do("BITOP", "NOT", destKey, key))
}

func (rm *redisMethod) BitPos(key string, bit int64, offsets ...int64) (int, error) {
	switch len(offsets) {
	case 0:
		return redis.Int(rm.Do("BITPOS", key, bit))
	case 1:
		return redis.Int(rm.Do("BITPOS", key, bit, offsets[0]))
	case 2:
		return redis.Int(rm.Do("BITPOS", key, bit, offsets[0], offsets[1]))
	default:
		return 0, errWrongArguments
	}
}

func (rm *redisMethod) HDel(key string, fields ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, fields...)
	return redis.Int(rm.Do("HDEL", args...))
}

func (rm *redisMethod) HExists(key, field string) (bool, error) {
	return redis.Bool(rm.Do("HEXISTS", key, field))
}

func (rm *redisMethod) HGet(key, field string) (string, error) {
	return redis.String(rm.Do("HGET", key, field))
}

func (rm *redisMethod) HGetAll(key string) (map[string]string, error) {
	return redis.StringMap(rm.Do("HGETALL", key))
}

func (rm *redisMethod) HIncrBy(key, field string, value int) (int, error) {
	return redis.Int(rm.Do("HINCRBY", key, field, value))
}

func (rm *redisMethod) HIncrByFloat(key, field string, value float64) (float64, error) {
	return redis.Float64(rm.Do("HINCRBYFLOAT", key, field, value))
}

func (rm *redisMethod) HMGet(key string, fields ...interface{}) ([]interface{}, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, fields...)
	return redis.Values(rm.Do("HMGET", args...))
}

func (rm *redisMethod) HMSet(key string, pairs ...interface{}) (bool, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, pairs...)
	return isOKString(redis.String(rm.Do("HMSET", args...)))
}

func (rm *redisMethod) HSet(key string, field string, value interface{}) (bool, error) {
	return redis.Bool(rm.Do("HSET", key, field, value))
}

func (rm *redisMethod) HSetNX(key string, field string, value interface{}) (bool, error) {
	return redis.Bool(rm.Do("HSETNX", key, field, value))
}

func (rm *redisMethod) HVals(key string) ([]string, error) {
	return redis.Strings(rm.Do("HVALS", key))
}

func (rm *redisMethod) HLen(key string) (int, error) {
	return redis.Int(rm.Do("HLEN", key))
}

func (rm *redisMethod) BRPopLPush(source, destination string, timeout uint64) (string, error) {
	return redis.String(rm.Do("BRPOPLPUSH", source, destination, timeout))
}

func (rm *redisMethod) LIndex(key string, index int64) (string, error) {
	return redis.String(rm.Do("LINDEX", key, index))
}

func (rm *redisMethod) LInsert(key string, op string, pivot, value interface{}) (int, error) {
	return redis.Int(rm.Do("LINSERT", key, op, pivot, value))
}

func (rm *redisMethod) LInsertBefore(key string, pivot, value interface{}) (int, error) {
	return redis.Int(rm.Do("LINSERT", key, "BEFORE", pivot, value))
}

func (rm *redisMethod) LInsertAfter(key string, pivot, value interface{}) (int, error) {
	return redis.Int(rm.Do("LINSERT", key, "AFTER", pivot, value))
}

func (rm *redisMethod) LLen(key string) (int, error) {
	return redis.Int(rm.Do("LLEN", key))
}

func (rm *redisMethod) LPop(key string) (string, error) {
	return redis.String(rm.Do("LPOP", key))
}

func (rm *redisMethod) LPush(key string, values ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, values...)
	return redis.Int(rm.Do("LPUSH", args...))
}

func (rm *redisMethod) LPushX(key string, value interface{}) (int, error) {
	return redis.Int(rm.Do("LPUSHX", key, value))
}

func (rm *redisMethod) LRange(key string, start, stop int64) ([]string, error) {
	return redis.Strings(rm.Do("LRANGE", key, start, stop))
}

func (rm *redisMethod) LRem(key string, count int64, value interface{}) (int, error) {
	return redis.Int(rm.Do("LREM", key, count, value))
}

func (rm *redisMethod) LSet(key string, index int64, value interface{}) (bool, error) {
	return isOKString(redis.String(rm.Do("LSET", key, index, value)))
}

func (rm *redisMethod) LTrim(key string, start, stop int64) (bool, error) {
	return isOKString(redis.String(rm.Do("LTRIM", key, start, stop)))
}

func (rm *redisMethod) RPop(key string) (string, error) {
	return redis.String(rm.Do("RPOP", key))
}

func (rm *redisMethod) RPopLPush(source, destination string) (string, error) {
	return redis.String(rm.Do("RPOPLPUSH", source, destination))
}

func (rm *redisMethod) RPush(key string, values ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, values...)
	return redis.Int(rm.Do("RPUSH", args...))
}

func (rm *redisMethod) RPushX(key string, value interface{}) (int, error) {
	return redis.Int(rm.Do("RPUSHX", key, value))
}

func (rm *redisMethod) SAdd(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, members...)
	return redis.Int(rm.Do("SADD", args...))
}

func (rm *redisMethod) SCard(key string) (int, error) {
	return redis.Int(rm.Do("SCARD", key))
}

func (rm *redisMethod) SDiff(keys ...interface{}) ([]string, error) {
	return redis.Strings(rm.Do("SDIFF", keys...))
}

func (rm *redisMethod) SDiffStore(destination string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, keys...)
	return redis.Int(rm.Do("SDIFFSTORE", args...))
}

func (rm *redisMethod) SInter(keys ...interface{}) ([]string, error) {
	return redis.Strings(rm.Do("SINTER", keys...))
}

func (rm *redisMethod) SInterStore(destination string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, keys...)
	return redis.Int(rm.Do("SINTERSTORE", args...))
}

func (rm *redisMethod) SIsMember(key string, member interface{}) (bool, error) {
	return redis.Bool(rm.Do("SISMEMBER", key, member))
}

func (rm *redisMethod) SMove(source, destination string, member interface{}) (bool, error) {
	return redis.Bool(rm.Do("SMOVE", source, destination, member))
}

func (rm *redisMethod) SMembers(key string) ([]string, error) {
	return redis.Strings(rm.Do("SMEMBERS", key))
}

func (rm *redisMethod) SPop(key string) (string, error) {
	return redis.String(rm.Do("SPOP", key))
}

func (rm *redisMethod) SPopN(key string, count int64) ([]string, error) {
	return redis.Strings(rm.Do("SPOP", key, count))
}

func (rm *redisMethod) SRandMember(key string) (string, error) {
	return redis.String(rm.Do("SRANDMEMBER", key))
}

func (rm *redisMethod) SRandMemberN(key string, count int64) ([]string, error) {
	return redis.Strings(rm.Do("SRANDMEMBER", key, count))
}

func (rm *redisMethod) SRem(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, members...)
	return redis.Int(rm.Do("SREM", args...))
}

func (rm *redisMethod) SUnion(keys ...interface{}) ([]string, error) {
	return redis.Strings(rm.Do("SUNION", keys...))
}

func (rm *redisMethod) SUnionStore(destionation string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destionation)
	args = append(args, keys...)
	return redis.Int(rm.Do("SUNIONSTORE", args...))
}

func (rm *redisMethod) ZAdd(key string, pairs ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, pairs...)
	return redis.Int(rm.Do("ZADD", args...))
}

func (rm *redisMethod) ZCard(key string) (int, error) {
	return redis.Int(rm.Do("ZCARD", key))
}

func (rm *redisMethod) ZCount(key string, min, max interface{}) (int, error) {
	return redis.Int(rm.Do("ZCOUNT", key, min, max))
}

func (rm *redisMethod) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return redis.Float64(rm.Do("ZINCRBY", key, increment, member))
}

func (rm *redisMethod) ZInterStore(destination string, nkeys int, params ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, nkeys)
	args = append(args, params...)
	return redis.Int(rm.Do("ZINTERSTORE", args...))
}

func (rm *redisMethod) ZLexCount(key, min, max string) (int, error) {
	return redis.Int(rm.Do("ZLEXCOUNT", key, min, max))
}

func (rm *redisMethod) ZRange(key string, start, stop int64) ([]string, error) {
	return redis.Strings(rm.Do("ZRANGE", key, start, stop))
}

func (rm *redisMethod) ZRangeWithScores(key string, start, stop int64) (map[string]string, error) {
	return redis.StringMap(rm.Do("ZRANGE", key, start, stop, "WITHSCORES"))
}

func (rm *redisMethod) ZRangeByLex(key, min, max string, offset, count int64) ([]string, error) {
	return redis.Strings(rm.Do("ZRANGEBYLEX", key, min, max, "LIMIT", offset, count))
}

func (rm *redisMethod) ZRangeByScore(key, min, max interface{}, offset, count int64) ([]string, error) {
	return redis.Strings(rm.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count))
}

func (rm *redisMethod) ZRank(key, member string) (int, error) {
	return redis.Int(rm.Do("ZRANK", key, member))
}

func (rm *redisMethod) ZRem(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, members...)
	return redis.Int(rm.Do("ZREM", args...))
}

func (rm *redisMethod) ZRemRangeByLex(key, min, max string) (int, error) {
	return redis.Int(rm.Do("ZREMRANGEBYLEX", key, min, max))
}

func (rm *redisMethod) ZRemRangeByRank(key string, start, stop int64) (int, error) {
	return redis.Int(rm.Do("ZREMRANGEByRANK", key, start, stop))
}

func (rm *redisMethod) ZRemRangeByScore(key, min, max interface{}) (int, error) {
	return redis.Int(rm.Do("ZREMRANGEBYSCORE", key, min, max))
}

func (rm *redisMethod) ZRevRange(key string, start, stop int64) ([]string, error) {
	return redis.Strings(rm.Do("ZREVRANGE", key, start, stop))
}

func (rm *redisMethod) ZRevRangeWithScores(key string, start, stop int64) (map[string]string, error) {
	return redis.StringMap(rm.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
}

func (rm *redisMethod) ZRevRangeByLex(key, max, min string, offset, count int64) ([]string, error) {
	return redis.Strings(rm.Do("ZREVRANGEBYLEX", key, max, min, "LIMIT", offset, count))
}

func (rm *redisMethod) ZRevRangeByScore(key, max, min interface{}, offset, count int64) ([]string, error) {
	return redis.Strings(rm.Do("ZREVRANGEBYSCORE", key, max, min, "LIMIT", offset, count))
}

func (rm *redisMethod) ZRevRangeByScoreWithScores(key, max, min interface{}, offset, count int64) (map[string]string, error) {
	return redis.StringMap(rm.Do("ZREVRANGEBYSCORE", key, max, min, "WITHSCORES", "LIMIT", offset, count))
}

func (rm *redisMethod) ZRevRank(key, member string) (int, error) {
	return redis.Int(rm.Do("ZREVRANK", key, member))
}

func (rm *redisMethod) ZScore(key, member string) (float64, error) {
	return redis.Float64(rm.Do("ZSCORE", key, member))
}

func (rm *redisMethod) ZUnionStore(destination string, nkeys int, params ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, nkeys)
	args = append(args, params...)
	return redis.Int(rm.Do("ZUNIONSTORE", args...))
}

func (rm *redisMethod) PFAdd(key string, els ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, els...)
	return redis.Int(rm.Do("PFADD", args...))
}

func (rm *redisMethod) PFCount(keys ...interface{}) (int, error) {
	return redis.Int(rm.Do("PFCOUNT", keys...))
}

func (rm *redisMethod) PFMerge(dest string, keys ...interface{}) (bool, error) {
	args := make([]interface{}, 0)
	args = append(args, dest)
	args = append(args, keys...)
	return isOKString(redis.String(rm.Do("PFMERGE", args...)))
}

func (rm *redisMethod) Publish(channel, msg string) (int, error) {
	return redis.Int(rm.Do("PUBLISH", channel, msg))
}

func (rm *redisMethod) RawCommand(command string, args ...interface{}) (interface{}, error) {
	return rm.Do(command, args...)
}
