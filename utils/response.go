package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ##### 状态常量
const (
	SUCCESS = 200
	ERROR = 500
	// ### 无效登录凭证
	INVALID_AUTH_TOKEN = 5000
	// ### 没有权限操作
	INVALID_PERMISSION_OPERATE = 4001

	// ### 无效传参
	INVALID_REQUEST_PARAMS = 1000
)
var msgFlags = map[int]string{
	SUCCESS: "操作成功",
	ERROR: "操作出错",
	INVALID_AUTH_TOKEN: "无效登录凭证",
	INVALID_PERMISSION_OPERATE: "无权操作",
	INVALID_REQUEST_PARAMS: "参数传递错误",
}

func getMsg(code int)string{
	msg,ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[ERROR]
}

type Object struct {
	C *gin.Context
}
type response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
func (g *Object)Response(code int,data interface{},msg string){
	if msg == "" {
		msg = getMsg(code)
	}
	g.C.JSON(http.StatusOK,response{
		Code: code,
		Msg: msg,
		Data: data,
	})
}