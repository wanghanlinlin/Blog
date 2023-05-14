package res

type ErrorCode int

var (
	SystemError          ErrorCode = 1001  //系统错误
	ArgumentError        ErrorCode = 1002  //参数错误
	WSCommunicationError ErrorCode = 1003  //WS系统通信错误
	WSReadMessageError   ErrorCode = 10031 //WS系统读取消息异常
	WSWriteMessageError  ErrorCode = 10032 //WS系统发送消息异常
)

var ErrorCodeMap = map[ErrorCode]string{
	1001:  "系统错误",
	1002:  "参数错误",
	1003:  "WS系统通信错误",
	10031: "WS系统读取消息异常",
	10032: "WS系统发送消息异常",
}
