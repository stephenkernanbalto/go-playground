package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrades protocol from http to tcp/websocket request
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func request_handler(w http.ResponseWriter, r *http.Request) {
	// 	convert connection to websocket connection
	connection, err := upgrader.Upgrade(w, r, nil)

	// if connection fails, log the error and return
	if err != nil {
		log.Println(err)
		return
	}

	for {
		// values returned are messageType, data, and error
		_, msg, err := connection.ReadMessage()

		// if the websocket itself has an error, log and kill connection
		if err != nil {
			log.Println()
			connection.Close()
			return
		}

		// otherwise, print the message received
		log.Println(string(msg))
	}
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
	}
}

