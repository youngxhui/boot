package main

import (
	"boot/protos"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func tool(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connect")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	_, message, err := c.ReadMessage()
	wear := new(protos.Wear)
	if err = proto.Unmarshal(message, wear); err != nil {
		log.Println(err.Error())
	}
	for {

		if err != nil {
			log.Println("read:", err)
			break
		}
		f := rand.Float32() * 100
		//str := fmt.Sprintf("%.0f", f)
		if f > 80 {
			go errorNotify(c, int(wear.Id))
		} else if f > 60 {
			go waringNotify(c, int(wear.Id))
		}

		wear := protos.Wear{
			Id:     wear.Id,
			ToolId: 1,
			Date:   time.Now().Format("15:04:05"),
			Loss:   f,
		}
		marshal, err := proto.Marshal(&wear)
		err = c.WriteMessage(websocket.BinaryMessage, marshal)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// 警告信息⚠
func errorNotify(c *websocket.Conn, id int) {

}

// 错误信息❌
func waringNotify(c *websocket.Conn, id int) {

}
