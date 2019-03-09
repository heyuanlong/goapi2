package main

import (
	"flag"
	_ "goapi2/initialize"
	ktest "goapi2/work/control/test"
	"strconv"
	//klog "github.com/heyuanlong/go-utils/common/log"
	kinit "goapi2/initialize"
	kroute "goapi2/route"
	kservice "goapi2/work/service"
)

func main() {
	types := flag.String("T", "", "启动类型，空：正常，race：刷比赛，bonus：刷分红，price：刷兑换比例")
	table_schema := flag.String("s", "test", "table_schema")
	table_name := flag.String("t", "test", "table_name")
	//begins := flag.Int64("B", 0, "开始时间戳，左闭右闭")
	//ends := flag.Int64("E", 0, "结束时间戳，左闭右闭")
	flag.Parse()

	kservice.InitSystem()
	if *types == "sql" {
		s := kservice.NewGenerateSqlStruct()
		s.Run(*table_schema, *table_name)
	}

	if *types == "api" {
		portStr, _ := kinit.Conf.GetString("server", "port")
		port, _ := strconv.Atoi(portStr)
		r := kroute.NewRouteStruct(port)
		r.SetMiddleware(kroute.SetCommonHeader)

		r.Load(ktest.NewTest())
		r.Run()
	}
	if *types == "settle" {
		kservice.Run()
	}

}
