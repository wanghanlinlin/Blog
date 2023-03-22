package res

type ErrorCode int

var (
	SystemError   ErrorCode = 1001 //系统错误
	ArgumentError ErrorCode = 1002 //文件上传错误
)

var ErrorCodeMap = map[ErrorCode]string{
	1001: "系统错误",
	1002: "参数错误",
}
