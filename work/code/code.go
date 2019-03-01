package code

// 不可变参数

const (
	KEY    = "xxxxxxxxxxxxxxx"
	SECRET = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

//---------------------------------------------------------------------------------------------

var codeToMsg map[int]string
var codeToChnMsg map[int]string

const (
	SUCCESS_STATUS    = 100001
	OPERATION_WRONG   = 600001
	ACCESS_TOKEN_FAIL = 600011
	PARAM_WRONG       = 600040
)

func init() {
	codeToMsg = make(map[int]string)
	codeToMsg[SUCCESS_STATUS] = "success"
	codeToMsg[OPERATION_WRONG] = "operation is wrong"
	codeToMsg[ACCESS_TOKEN_FAIL] = "check access_token fail"
	codeToMsg[PARAM_WRONG] = "Parameter is wrong"

	codeToChnMsg = make(map[int]string)
	codeToChnMsg[SUCCESS_STATUS] = "成功"
	codeToChnMsg[OPERATION_WRONG] = "操作错误"
	codeToChnMsg[ACCESS_TOKEN_FAIL] = "check access_token fail"
	codeToChnMsg[PARAM_WRONG] = "参数有误"
}
func GetCodeMsg(code int) string {
	if msg, ok := codeToMsg[code]; ok {
		return msg
	}
	return ""
}
func GetCodeChnMsg(code int) string {
	if msg, ok := codeToChnMsg[code]; ok {
		return msg
	}
	return ""
}

//---------------------------------------------------------------------------------------------
