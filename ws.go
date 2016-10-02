package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type msg struct {
	Type string `json:"type"`
	Data struct {
		Msg string `json:"msg"`
	} `json:"data"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.Handle("/echo", websocket.Handler(func(conn *websocket.Conn) {
		for {
			var b []byte
			var m msg

			websocket.Message.Receive(conn, &b)

			json.Unmarshal(b, &m)

			fmt.Println(m)

			bm, _ := json.Marshal(m)

			conn.Write(bm)
		}
	}))

	http.ListenAndServe(":1234", nil)
}
