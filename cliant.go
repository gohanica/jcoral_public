package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

func (g *client) read() {

	for {
		var msg *message
		if err := g.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = g.userData["name"].(string)
			if avatarURL, ok := g.userData["avatar_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}

			g.room.forward <- msg
		} else {
			break
		}

	}
	g.socket.Close()

}

func (g *client) write(r *http.Request) {
	fmt.Println(r.URL.Path)
	for msg := range g.send {
		if err := g.socket.WriteJSON(msg); err != nil {
			break
		}
		fmt.Println(msg)
	}
	g.socket.Close()
}
