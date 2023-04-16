package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

func Strem_Chat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()

	// 循环处理 WebSocket 消息
	for {
		// 读取消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		// 打印接收到的消息
		log.Printf("Received message: %s\n", msg)

		// 将接收到的消息原样发回客户端
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func Http_Chat(c *gin.Context) {

}
