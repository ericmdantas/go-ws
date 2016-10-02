package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.Handle("/echo", websocket.Handler(func(conn *websocket.Conn) {
		conn.Write([]byte("yo!"))
	}))

	http.ListenAndServe(":1234", nil)
}
