package main

import (
	"fmt"
	"net/http"

	"github.com/matheus-alpe/go-gorilla-websockets/internal/websocket"
)

func main() {
	http.HandleFunc("/ws", websocket.Handler)

	fmt.Println("WebSocket server is running on http://localhost:8080/ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
