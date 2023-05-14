package echoapi

import (
	"AuroraPixel/core/res"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type EchoApi struct{}

// websocket通信升级器
var upgrade = websocket.Upgrader{
	//允许跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// controller
func (*EchoApi) Connect(c *gin.Context) {
	//升级http为ws连接
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.ErrorWithCodeData(err.Error(), res.WSCommunicationError, c)
	}

	//完成时候关闭连接
	defer ws.Close()

	//协程
	go func() {
		//事件完成
		<-c.Done()
		//事件完成处理
		logrus.Info("ws 失去连接")
	}()
	//无限循环监听消息
	for {
		//读取消息
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			res.ErrorAndContinue(err.Error(), res.WSReadMessageError, c)
			break
		}
		logrus.Infof("获取消息类型:%v,消息内容:%v", messageType, string(message))
		//回消息
		err2 := ws.WriteMessage(messageType, []byte("这是服务器返回消息"))
		if err2 != nil {
			res.ErrorAndContinue(err2.Error(), res.WSWriteMessageError, c)
			break
		}

	}

}

// 定义客户端
type Client struct {
	WS   *websocket.Conn //webSocket连接
	Send chan []byte     //管道消息
	Hub  *Hub            //消息中心
}

var (
	WriteTimeWait time.Duration = 2 * time.Minute
	ReadTimeWait  time.Duration = 4 * time.Minute
	PingPeriod    time.Duration = 2 * time.Second
	PongPeriod    time.Duration = 4 * time.Second
	ReadSizeLimit int64         = 65536
)

// 定义客户端发送消息
func (c *Client) WriteMessage() {
	//触发两秒的定时器
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		c.WS.Close()
	}()
	for {
		select {
		case Message, ok := <-c.Send:
			//获取客户端
			//设置写入超时时间
			c.WS.SetWriteDeadline(time.Now().Add(WriteTimeWait))
			if !ok {
				//关闭消息
				c.WS.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//文本编辑器
			wc, err := c.WS.NextWriter(websocket.TextMessage)
			if err != nil {
				logrus.Errorln("这是websocket.NextWriter错误信息:%v", err.Error())
				return
			}
			wc.Write(Message)
			if err := wc.Close(); err != nil {
				return
			}
		case <-ticker.C:
			//每两秒进行心跳传输
			err := c.WS.SetWriteDeadline(time.Now().Add(WriteTimeWait))
			if err != nil {
				return
			}
		}
	}
}

// 定义客户端读消息
func (c *Client) ReadMessage() {
	//注销进程关闭连接
	defer func() {
		c.Hub.UnRegister <- c
		c.WS.Close()
	}()
	//限制读取大小
	c.WS.SetReadLimit(ReadSizeLimit)
	//读取最大延迟时间
	c.WS.SetReadDeadline(time.Now().Add(ReadTimeWait))
	//在4s秒内没有获取Ping消息关闭连接
	c.WS.SetPongHandler(
		func(string) error { c.WS.SetReadDeadline(time.Now().Add(PongPeriod)); return nil })
	//循环读取消息
	for {
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("websocker读取消息失败:%v", err.Error())
			}
			break
		}
		//消息格式化
		//message = bytes.TrimSpace(bytes.Replace(message,newline,space,-1))
		//消息推送
		c.Hub.Broadcast <- message
	}
}

// 消息运行服务
func (e *EchoApi) EchoService(hub *Hub, c *gin.Context) {
	//升级http为ws连接
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.ErrorWithCodeData(err.Error(), res.WSCommunicationError, c)
	}
	//定义客户端
	client := &Client{
		Hub:  hub,
		WS:   ws,
		Send: make(chan []byte),
	}
	//注册客户端
	client.Hub.Register <- client
	//写消息
	go client.WriteMessage()
	//读消息
	go client.ReadMessage()
}

// 客户端管理中心
type Hub struct {
	Clients    map[*Client]bool //客户端集合
	Broadcast  chan []byte      //消息管理中心
	Register   chan *Client     //注册中心
	UnRegister chan *Client     //注销中心
}

// 新建管理中心
func NewHub() *Hub {
	return &Hub{
		make(map[*Client]bool),
		make(chan []byte),
		make(chan *Client),
		make(chan *Client),
	}
}

func (h *Hub) HubRun() {
	//无线循环监听
	for {
		select {
		//客户端注册
		case Client := <-h.Register:
			h.Clients[Client] = true
		//客户端注销
		case Client := <-h.UnRegister:
			//从map中拿到客户判断是否存在，存在旧删除
			if _, ok := h.Clients[Client]; ok {
				//删除客户端
				delete(h.Clients, Client)
				//关闭管道
				close(Client.Send)
			}
			//客户端发送消息
		case Message := <-h.Broadcast:
			//给每个客户端推送消息
			for Client := range h.Clients {
				select {
				//进行消息推送给客户端
				case Client.Send <- Message:
				default:
					//关闭客户端
					close(Client.Send)
					delete(h.Clients, Client)
				}
			}
		}

	}
}
