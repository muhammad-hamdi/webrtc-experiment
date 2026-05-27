package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var socketRooms map[string]*Room = make(map[string]*Room)

var mu = &sync.Mutex{} // Protect clients map

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}

	// init/handshake/connection start flow
	// 1. Check if user sent user_id
	// 2. if user_id find user reference in users
	// 3. if not create user object and append in users
	// 4. Check if user sent canvas_id
	// 5. if canvas_id find canvas ref in canvases
	// 6. find room associated with canvas
	// 7. if room, create client and subscribe to room
	// 8. if not, create room and client and subscribe client to newly created room
	// 6. if no canvas_id, create canvas, room, client and do appropriate appends and subscription

	//// for #1 we'll do a pre-read, first user message is either a NewCanvas or ConnectToCanvas
	//// ConnectToCanvas will be treated differently if the connection is active?
	///// Check said canvas and find if an exisitng room is there, otherwise create the room for the canvas

	_, message, err := conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error: %v", err)
		}
	}

	fmt.Println(message)

	// var initMessage Message
	// json.Unmarshal(message, &initMessage)

	// Of course the lookups will be replaced with db queries later
	// but we're testing with in-memory data for now

	var room *Room
	// lookup the room
	if room == nil {
		room = newRoom()
		go room.run()
		mu.Lock()
		socketRooms[room.Id] = room
		mu.Unlock()
	}

	client := &Client{room: room, id: uuid.Must(uuid.NewV4()).String(), conn: conn, send: make(chan []byte, 256)}
	room.register <- client

	go client.readPump()
	go client.writePump()
}
