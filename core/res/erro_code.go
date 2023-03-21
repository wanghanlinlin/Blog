package res

type ErrorCode int

var (
	SystemError ErrorCode = 1001 //系统错误
)

var ErrorCodeMap = map[ErrorCode]string{
	1001: "系统错误",
}
