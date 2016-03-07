package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	ws "github.com/gorilla/websocket"
)

func initWs() {
	fmt.Println("ws on")

	u := url.URL{Scheme: "ws", Host: "localhost:3333", Path: "/"}

	c, _, err := ws.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial error:", err)
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)

		for {
			_, message, err := c.ReadMessage()

			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Println("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(ws.TextMessage, []byte(t.String()))

			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}

func initHTTP() {
	fmt.Println("http on")

	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func main() {
	initWs()
	initHTTP()
}
