package initialize

import (
	"runtime"
	kconf "github.com/heyuanlong/go-utils/common/conf"
)

const CONFIG_PATH = "conf/config.cfg"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())	//多核设置
	kconf.SetFile(CONFIG_PATH)

}
