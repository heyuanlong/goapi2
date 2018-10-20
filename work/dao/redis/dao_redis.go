package redis

import (
	"github.com/garyburd/redigo/redis"

	klog "github.com/heyuanlong/go-utils/common/log"
	kredis "github.com/heyuanlong/go-utils/db/redis"
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
