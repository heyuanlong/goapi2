package mfomo

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

	return m
}

//-----------------------------------------------------------------------------------

// http://127.0.0.1:8080/mfomo/test?wallet_account=0x2d437Ec35E7d13a1AfF58250EeDc2808b92D9725&min_coins=0.5&chain_id=1
type subtest2Struct struct {
	WalletAccout string  `json:"wallet_account"` // 转币地址
	MinCoins     float64 `json:"min_coins"`      //最少投入eth数量
	ChainId      int     `json:"chain_id"`       //分红占比

}
type test2Struct struct {
	Status int            `json:"status"`
	Info   string         `json:"info"`
	Data   subtest2Struct `json:"data"`
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
	object := test2Struct{
		Status: kcode.SUCCESS_STATUS,
		Info:   kcode.GetCodeMsg(kcode.SUCCESS_STATUS),
		Data:   subObject,
	}
	kbase.ReturnData(c, object, callbackName)
}

//-----------------------------------------------------------------------------------

// http://127.0.0.1:8080/mfomo/test2
func (ts *Test) Test2(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("test2"))
}

//-----------------------------------------------------------------------------------
