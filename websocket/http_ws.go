package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":11888", "http service address")

// http 的方式
//func __main() {
//	http.HandleFunc("/quote/list", httpSocketHandler)
//	log.Fatal(http.ListenAndServe(*addr, nil))
//}

func httpSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}
