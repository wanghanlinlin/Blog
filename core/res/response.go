package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	Success = 200
	Erro    = 500
)

// 响应结构体
type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// 总返回方法
func Result(code int, data any, message string, c *gin.Context) {
	logrus.Errorf("Res,错误代码:%v,错误数据:%v,错误消息:%v", code, data, message)
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

func Ok(data any, message string, c *gin.Context) {
	Result(Success, data, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "查询成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, message, c)
}

func Error(data any, message string, c *gin.Context) {
	Result(Erro, data, message, c)
}

func ErrorWithMessage(message string, c *gin.Context) {
	Result(Erro, map[string]any{}, message, c)
	logrus.Panicf(message)
}

func ErrorWithData(data any, c *gin.Context) {
	Result(Erro, data, "查询失败", c)
	logrus.Panicf("查询失败")
}

func ErrorWithCode(code ErrorCode, c *gin.Context) {
	message, ok := ErrorCodeMap[code]
	if ok {
		Result(int(code), map[string]any{}, message, c)
		logrus.Panicf(message)
	}
	Result(Erro, map[string]any{}, "未知错误", c)
	logrus.Panicf("未知错误")
}

func ErrorWithCodeData(data any, code ErrorCode, c *gin.Context) {
	message, ok := ErrorCodeMap[code]
	if ok {
		Result(int(code), data, message, c)
		logrus.Panicf(message)
	}
	Result(Erro, data, "未知错误", c)
	logrus.Panicf("未知错误")
}

func ErrorAndContinue(data any, code ErrorCode, c *gin.Context) {
	message, ok := ErrorCodeMap[code]
	if ok {
		Result(int(code), data, message, c)
		return
	}
	Result(Erro, data, "未知错误", c)
}
