package initialize

import (
	"log"
	"runtime"

	conf "github.com/heyuanlong/go-utils/common/conf"
	klog "github.com/heyuanlong/go-utils/common/log"
	kgorm "github.com/heyuanlong/go-utils/db/gorm"
	kredis "github.com/heyuanlong/go-utils/db/redis"
	"github.com/jinzhu/gorm"
)

const CONFIG_PATH = "conf/config.cfg"

var Conf *conf.Kconf

var LogError *klog.Llog
var LogWarn *klog.Llog
var LogInfo *klog.Llog
var LogDebug *klog.Llog

var Gorm *gorm.DB
var RedisPool *kredis.RedisPool

func init() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核设置
	Conf = conf.NewKconf(CONFIG_PATH)

	LogError = klog.Error
	LogWarn = klog.Warn
	LogInfo = klog.Info
	LogDebug = klog.Debug

	mysql_user, _ := Conf.GetString("mysql", "user")
	mysql_password, _ := Conf.GetString("mysql", "password")
	mysql_ip, _ := Conf.GetString("mysql", "ip")
	mysql_port, _ := Conf.GetString("mysql", "port")
	mysql_mysqldb, _ := Conf.GetString("mysql", "mysqldb")
	Gorm, err = kgorm.NewGorm(mysql_user, mysql_password, mysql_ip, mysql_port, mysql_mysqldb)
	if err != nil {
		log.Println(err)
	}

	redis_host, _ := Conf.GetString("redis", "host")
	redis_port, _ := Conf.GetString("redis", "port")
	redis_auth, _ := Conf.GetString("redis", "auth")
	RedisPool, err = kredis.NewRedisPool(redis_host, redis_port, redis_auth)
	if err != nil {
		log.Println(err)
	}

}
