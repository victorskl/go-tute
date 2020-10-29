package main

import (
	"net/http"

	"github.com/victorskl/go-tute/useless"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/useless", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(useless.Foobar()))
	}))
	http.ListenAndServe(":1234", nil)
}
