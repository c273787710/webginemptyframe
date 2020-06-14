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

	// ### 无效
	INVALID_REQUEST_PARAMS = 1000
	INVALID_AUTH_NAME = 1001
	RECORD_IS_EXIT = 1002
	RECORD_NOT_EXIT = 1003
	CREATE_RECORD_ERR = 1004
	UPDATE_RECORD_ERR = 1005
	DELETE_RECORD_ERR = 1006

	ADMIN_LOCKING = 2002
	ADMIN_AUTH_ERR = 2003


)
var msgFlags = map[int]string{
	SUCCESS: "操作成功",
	ERROR: "操作出错",
	INVALID_AUTH_TOKEN: "无效登录凭证",
	INVALID_PERMISSION_OPERATE: "无权操作",
	INVALID_REQUEST_PARAMS: "参数传递错误",
	INVALID_AUTH_NAME: "用户名错误",
	ADMIN_LOCKING: "用户被锁定",
	ADMIN_AUTH_ERR: "用户登录失败",
	RECORD_IS_EXIT: "记录相关值已经存在",
	RECORD_NOT_EXIT: "记录不存在",
	CREATE_RECORD_ERR:"创建记录失败",
	UPDATE_RECORD_ERR: "更新记录失败",
	DELETE_RECORD_ERR: "删除记录失败",
}

func getMsg(code int)string{
	msg,ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[ERROR]
}

type object struct {
	C *gin.Context
}

func NewObject(c *gin.Context)object{
	return object{C:c}
}

type response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
func (g *object)Response(code int,data interface{},msg string){
	if msg == "" {
		msg = getMsg(code)
	}
	g.C.JSON(http.StatusOK,response{
		Code: code,
		Msg: msg,
		Data: data,
	})
}