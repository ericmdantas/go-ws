package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.Handle("/ws", websocket.Handler(func(client *websocket.Conn) {
		ticker := time.Tick(33 * time.Millisecond)
		rep := 0
		max := 10000

		for {
			select {
			case <-ticker:
				websocket.JSON.Send(client, map[string]string{"msg": "yo!"})
				rep++
			}

			if rep == max {
				break
			}
		}

		fmt.Println("aaaaaaand, free!")

	}))

	http.ListenAndServe(":3333", nil)
}
