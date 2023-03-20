package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Sucess = 200
	Erro   = 500
)

// 响应结构体
type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// 总返回方法
func Result(code int, data any, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

func Ok(data any, message string, c *gin.Context) {
	Result(Sucess, data, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Sucess, data, "查询成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Sucess, map[string]interface{}{}, message, c)
}

func Error(data any, message string, c *gin.Context) {
	Result(Erro, data, message, c)
}

func ErrorWithMessage(message string, c *gin.Context) {
	Result(Erro, map[string]any{}, message, c)
}

func ErrorWithData(data any, c *gin.Context) {
	Result(Erro, data, "查询失败", c)
}

func ErrorWithCode(code ErrorCode, c *gin.Context) {
	message, ok := ErrorCodeMap[code]
	if ok {
		Result(int(code), map[string]any{}, message, c)
		return
	}
	Result(Erro, map[string]any{}, "未知错误", c)
}
