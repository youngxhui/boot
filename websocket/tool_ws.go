package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	mt, message, err := c.ReadMessage()
	for {

		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		id, err := strconv.Atoi(string(message))
		if err != nil {
			log.Println("err :", err)
		}
		log.Println("id", id)

		f := rand.Float32() * 100

		str := fmt.Sprintf("%.0f", f)
		err = c.WriteMessage(mt, []byte(str))
		//fmt.Println(runtime.)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}
