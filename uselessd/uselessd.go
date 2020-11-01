package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/victorskl/go-tute/useless"
	uselessUtils "github.com/victorskl/go-tute/useless/utils"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/useless", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(useless.Foobar())) // using useless package from useless module
	}))

	// using utils package from useless module
	log.Println(fmt.Sprintf("uselessUtils.Add(1, 1) is: %d", uselessUtils.Add(1, 1)))

	log.Println("Listing for requests at ws://localhost:1234/useless")
	http.ListenAndServe(":1234", nil)
}
