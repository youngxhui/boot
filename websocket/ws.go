package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func main() {
	flag.Parse()
	log.SetFlags(0)
	//http.HandleFunc("/echo", echo)
	//http.HandleFunc("/", echo)
	http.HandleFunc("/", tool)
	http.HandleFunc("/notice", notice)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
