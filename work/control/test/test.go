package test

import (
	"net/http"
	"strconv"

	kinit "goapi2/initialize"
	kcode "goapi2/work/code"
	kbase "goapi2/work/control/base"

	"github.com/gin-gonic/gin"
)

type Test struct {
}

func NewTest() *Test {
	return &Test{}
}
func (ts *Test) Load() []kbase.RouteWrapStruct {
	m := make([]kbase.RouteWrapStruct, 0)
	m = append(m, kbase.Wrap("GET", "/mfomo/test1", ts.Test1, 0))
	m = append(m, kbase.Wrap("GET", "/mfomo/test2", ts.Test2, 0))
	m = append(m, kbase.Wrap("GET", "/mfomo/test3", ts.test3, 0))

	return m
}

//-----------------------------------------------------------------------------------

// http://127.0.0.1:8080/mfomo/test1?wallet_account=0x2d437Ec35E7d13a1AfF58250EeDc2808b92D9725&min_coins=0.5&chain_id=1
type subtest2Struct struct {
	WalletAccout string  `json:"wallet_account"` // 转币地址
	MinCoins     float64 `json:"min_coins"`      //最少投入eth数量
	ChainId      int     `json:"chain_id"`       //分红占比

}

func (ts *Test) Test1(c *gin.Context) {
	callbackName := kbase.GetParam(c, "callback")
	wallet_account := kbase.GetParam(c, "wallet_account")
	kinit.LogWarn.Println("wallet_account:", wallet_account)

	min_coins, err := strconv.ParseFloat(kbase.GetParam(c, "min_coins"), 0)
	if err != nil {
		kinit.LogWarn.Println(min_coins, err)
		kbase.SendErrorJsonStr(c, kcode.PARAM_WRONG, callbackName)
		return
	}
	chain_id, err := strconv.Atoi(kbase.GetParam(c, "chain_id"))
	if err != nil {
		kbase.SendErrorJsonStr(c, kcode.PARAM_WRONG, callbackName)
		return
	}

	subObject := subtest2Struct{
		WalletAccout: wallet_account,
		MinCoins:     min_coins,
		ChainId:      chain_id,
	}
	kbase.ReturnDataI(c, kcode.SUCCESS_STATUS, subObject, callbackName)
}

//-----------------------------------------------------------------------------------

// http://127.0.0.1:8080/mfomo/test2
func (ts *Test) Test2(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("test2"))
}

//-----------------------------------------------------------------------------------
// http://127.0.0.1:8080/mfomo/test3
func (ts *Test) test3(c *gin.Context) {
	m := make(map[string]interface{})
	m["status"] = 1
	m["amount"] = 1.333
	m["Str"] = "xxxxxxxxxxxxxx"

	obj := struct {
		WalletAccout string  `json:"wallet_account"` // 转币地址
		MinCoins     float64 `json:"min_coins"`      //最少投入eth数量
		maxCoins     float64 `json:"min_coins"`      //最少投入eth数量
	}{
		WalletAccout: "wallet_account",
		MinCoins:     9.08,
		maxCoins:     8888,
	}

	mx := make(map[string]interface{})
	mx["status"] = 3
	mx["amount"] = 3.333

	m["struct"] = obj
	m["Map"] = mx
	kbase.ReturnDataI(c, kcode.SUCCESS_STATUS, m, "")
}
