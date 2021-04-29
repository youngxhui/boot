package main

import (
	"boot/protos"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/youngxhui/power/log"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"net/http"
	"time"
)

func tool(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connect")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info("upgrade:", err)
		return
	}

	defer c.Close()
	_, message, err := c.ReadMessage()
	wear := new(protos.Wear)
	if err = proto.Unmarshal(message, wear); err != nil {
		log.Waring(err.Error())
	}
	for {

		if err != nil {
			log.Waring("read:", err)
			break
		}
		f := rand.Float32()*10 + 20 + float32(time.Now().Second())

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
