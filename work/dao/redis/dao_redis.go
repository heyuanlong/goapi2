package redis

import (
	"errors"

	"github.com/garyburd/redigo/redis"

	klog "github.com/heyuanlong/go-utils/common/log"
	kredis "github.com/heyuanlong/go-utils/db/redis"
)

var (
	ErrSingleLockInvalidRedisConn = errors.New("SingleLock : Invalid redis conn")
	ErrSingleLockOperationFailed  = errors.New("SingleLock : Operation failed")
	ErrSingleLockNotLocked        = errors.New("SingleLock : Not locked")
	ErrSingleLockInvalidLockValue = errors.New("SingleLock : Invalid lock value")
	ErrSingleLockLockIsUnlocked   = errors.New("SingleLock : Lock is unlocked")
	ErrSingleLockInvalidRedisConf = errors.New("SingleLock : Invalid redis conf")
)

func init() {
	kredis.InitRedis()
}

func Set(key string, v interface{}, expire int64) error {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	if err := rc.Send("SET", key, v); err != nil {
		klog.Info.Println(err)
		return err
	}

	if expire > 0 {
		if err := rc.Send("EXPIRE", key, expire); err != nil {
			klog.Info.Println(err)
			return err
		}
	}
	if err := rc.Flush(); err != nil {
		klog.Info.Println(err)
		return err
	}
	return nil
}
func Incr(key string) error {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	if _, err := rc.Do("INCR", key); err != nil {
		klog.Info.Println(err)
		return err
	}
	return nil
}
func Decr(key string) error {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	if _, err := rc.Do("DECR", key); err != nil {
		klog.Info.Println(err)
		return err
	}
	return nil
}
func GetFloat64(key string) (float64, error) {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	v, err := redis.Float64(rc.Do("get", key))
	return v, err
}

func GetString(key string) (string, error) {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	v, err := redis.String(rc.Do("get", key))
	return v, err
}
func GetInt64(key string) (int64, error) {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	v, err := redis.Int64(rc.Do("get", key))
	return v, err
}

func GetIncr(key string) (int, error) {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	return redis.Int(rc.Do("incr", key))
}

func HgetAll(key string) (map[string]int64, error) {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	return redis.Int64Map(rc.Do("HGETALL", key))
}
func Hset(key string, subKey string, v interface{}) error {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	if err := rc.Send("HSET", key, subKey, v); err != nil {
		klog.Info.Println(err)
		return err
	}
	if err := rc.Flush(); err != nil {
		klog.Info.Println(err)
		return err
	}
	return nil
}
func Hdel(key string, subKey string) error {
	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)
	if err := rc.Send("HDEL", key, subKey); err != nil {
		klog.Info.Println(err)
		return err
	}
	if err := rc.Flush(); err != nil {
		klog.Info.Println(err)
		return err
	}
	return nil
}

func Lock(keyName string, value string, milliseconds int64) error {
	if len(keyName) == 0 ||
		len(value) == 0 {
		return ErrSingleLockNotLocked
	}

	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)

	//	try to lock
	rpl, err := redis.String(rc.Do("SET", keyName, value, "NX", "PX", milliseconds))
	if nil != err {
		return err
	}
	if rpl != "OK" {
		return ErrSingleLockOperationFailed
	}

	return nil
}

func Unlock(keyName string, value string) error {
	if len(keyName) == 0 ||
		len(value) == 0 {
		return ErrSingleLockNotLocked
	}

	rc := kredis.GetRedis()
	defer kredis.CloseRedis(rc)

	//	try to unlock
	//	avoid to unlock a lock not belongs to the locker
	lockValue, err := redis.String(rc.Do("GET", keyName))
	if nil != err {
		return err
	}
	if lockValue != value {
		return ErrSingleLockInvalidLockValue
	}

	rpl, err := redis.Int(rc.Do("DEL", keyName))
	if nil != err {
		return err
	}

	if rpl != 1 {
		return ErrSingleLockLockIsUnlocked
	}

	return nil
}
