package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[uint]*websocket.Conn)

func InitConnectionSocket(c *gin.Context) {
	idUser, _ := c.Get("userid")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		if idUser == nil {
			return false
		}
		return true
	}
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections

	clients[idUser.(uint)] = ws

	reader(ws)
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))
	}
}

func sender(idClient uint, message string) {
	if _, exits := clients[idClient]; !exits {
		fmt.Printf("THE CLIENT %d DON`T EXIST \n", idClient)
		return
	}
	err := clients[idClient].WriteMessage(websocket.TextMessage, []byte(message))
	fmt.Println()
	if err != nil {
		delete(clients, idClient)
		log.Printf("[WEBSOCKET] SEND A MESSAGE -> %v \n", err)
	}
}
