package main

import (
	c "crypto/md5"
	"encoding/json"
	"net/http"

	"golang.org/x/net/websocket"
)

type msg struct {
	Type string `json:"type"`
	Hash []byte `json:"hash"`
	Data struct {
		CreatedAt int64  `json:"createdAt"`
		Msg       string `json:"msg"`
	} `json:"data"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.Handle("/echo", websocket.Handler(func(conn *websocket.Conn) {
		for {
			var b []byte
			var m msg

			websocket.Message.Receive(conn, &b)

			hasher := c.New()
			hasher.Write(b)

			json.Unmarshal(b, &m)

			m.Hash = hasher.Sum(nil)

			bm, _ := json.Marshal(m)

			conn.Write(bm)
		}
	}))

	http.ListenAndServe(":1234", nil)
}
