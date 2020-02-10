package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Message 便宜上今だけ作成したメッセージ送受信用構造体
type Message struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

// Messages 送信用構造体の配列
type Messages []Message

var message Message
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

// ToServer クライアントからサーバーへメッセージ送信
func ToServer(w http.ResponseWriter, r *http.Request) {
	go ToClients()
	// アップグレード
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("アップグレードエラー", err)
	}
	defer websocket.Close()

	clients[websocket] = true

	for {
		// メッセージ受信
		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Fatal("メッセージ読み込みエラー")
			delete(clients, websocket)
			break
		}

		// メッセージ送信用へchan使用して送り出す
		broadcast <- message
		// fmt.Println(message)
	}
}

// ToClients サーバーからクライアントへメッセージ送信
func ToClients() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				fmt.Println("送信エラー")
				client.Close()
				delete(clients, client)
			}
		}
	}
}
