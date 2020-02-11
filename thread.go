package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func threadmake(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	t := now.Nanosecond()
	s := r.PostFormValue("text")
	room := newRoom()
	URI1 := strconv.Itoa(t)
	fmt.Println(s)
	w.Header().Set("Location", "/chat/"+URI1)
	http.Handle("/chat/"+URI1+"/room", room)
	w.WriteHeader(http.StatusTemporaryRedirect)

	go room.run()
}
