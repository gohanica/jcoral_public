package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/stretchr/objx"
)

//メッセージのコンテンツをまとめる構造体です
type sendmessage struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	When      string `json:"date"`
	AvatarURL string `json:"profile"`
	ID        string `json:"id"`
}
type recievemessage struct {
	Name1    string `json:"name1"`
	Message1 string `json:"message1"`
}

func noteofajax(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	var recieve recievemessage
	json.Unmarshal(b, &recieve)
	fmt.Println(recieve)

	authCookie, err := r.Cookie("auth")
	if err != nil {
		panic(err)
	}

	var userData map[string]interface{}

	userData = objx.MustFromBase64(authCookie.Value)

	var send sendmessage

	send.Name = userData["name"].(string)
	send.Message = recieve.Message1
	send.When = time.Now().String()
	if avatarURL, ok := userData["avatar_url"]; ok {
		send.AvatarURL = avatarURL.(string)
	}
	send.ID = userData["openID"].(string)

	json.NewEncoder(w).Encode(send)

	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	no, err1 := db.Prepare("INSERT INTO message (name,message,hhh,avatarurl,id,url) VALUES(?,?,?,?,?,?)")
	if err1 != nil {
		panic(err1)
	}
	defer no.Close()
	_, err2 := no.Exec(send.Name, send.Message, send.When, send.AvatarURL, send.ID, r.URL.Path)
	if err2 != nil {
		panic(err2)

	}

}
