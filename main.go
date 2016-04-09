package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type message struct {
	Message string `json:"message"`
}

type clManager struct {
	clients []*websocket.Conn
}

func (cl *clManager) add(client *websocket.Conn) {
	if len(cl.clients) == 0 {
		cl.clients = append(cl.clients, client)
		return
	}

	for _, v := range cl.clients {
		if client != v {
			cl.clients = append(cl.clients, client)
		}
	}
}

func (cl *clManager) broadcast(msg message) {
	for _, v := range cl.clients {
		websocket.JSON.Send(v, msg)
	}
}

func (cl *clManager) remove(client *websocket.Conn) {
	for i, v := range cl.clients {
		if v == client {
			cl.clients = append(cl.clients[:i], cl.clients[i+1:]...)
		}
	}
}

var cl = new(clManager)

func socketFn(client *websocket.Conn) {
	cl.add(client)

	for {
		var m message

		if err := websocket.JSON.Receive(client, &m); err != nil {
			log.Println(err)

			if err == io.EOF {
				cl.remove(client)
			}

			break
		}

		cl.broadcast(message{m.Message})
	}
}

func main() {
	port := ":3333"

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/socket", websocket.Handler(socketFn))

	fmt.Println(port)

	http.ListenAndServe(port, nil)
}
