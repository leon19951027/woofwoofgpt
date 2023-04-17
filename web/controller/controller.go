package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/leon19951027/woofwoofgpt/openai"
)

var OpenaiApiToken string
var OpenaiApiUrlPrefix string

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

func Login(c *gin.Context) {

}

func Stream_Chat(c *gin.Context) {
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
		chatmsgs := &openai.ChatMessages{}
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal(msg, chatmsgs)
		if err != nil {
			log.Println(err)
			continue
		}
		doneChan := make(chan bool)
		msgsChan := make(chan string, 10)
		go openai.Chat(chatmsgs, OpenaiApiToken, OpenaiApiUrlPrefix, msgsChan, doneChan)

	LOOP:
		for {
			select {
			case resp := <-msgsChan:
				conn.WriteMessage(msgType, []byte(resp))
			case isDone := <-doneChan:
				if isDone {
					break LOOP
				}
			}
		}
		fmt.Println("ws done")

	}
}

func Chunk_Chat(c *gin.Context) {
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	// 设置响应内容类型
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//chatmsgs := &openai.ChatMessages{}
	b, _ := ioutil.ReadAll(c.Request.Body)
	chatmsgs := &openai.ChatMessages{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(b, chatmsgs)
	if err != nil {

		c.JSON(500, gin.H{"err": "反序列化错误,检查请求体"})

		return
	}
	//fmt.Println(chatmsgs)
	doneChan := make(chan bool)
	msgsChan := make(chan string, 10)
	go openai.Chat(chatmsgs, OpenaiApiToken, OpenaiApiUrlPrefix, msgsChan, doneChan)

LOOP:
	for {
		select {
		case resp := <-msgsChan:
			_, err := io.WriteString(c.Writer, resp)
			if err != nil {
				c.JSON(500, gin.H{"err": err})
				break LOOP
			}
			c.Writer.Flush()
		case isDone := <-doneChan:
			if isDone {
				break LOOP
			}
		}
	}
	fmt.Println("chunk stop")

}
