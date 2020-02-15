package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

//　クライアントデータの大元になっています。
type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

type message struct {
	Name      string
	Message   string
	When      string
	AvatarURL string
}

type authHandler struct {
	next http.Handler
}

//チャットの部屋のデータです
type room struct {
	forward chan *message
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

type templateHandlertochat struct {
	filenametochat string
	temptochat     *template.Template
}
type templateHandlertothread struct {
	filenametothread string
	temptothread     *template.Template
}
type templateHandlertologin struct {
	filenametologin string
	temptologin     *template.Template
}
type receivecontent struct {
	As string
	Bs string
	Cs string
	Ds string
	Fs string
}
type receivethreads struct {
	At string
	Bt string
	Ct string
	Dt int
	Ft string
}
