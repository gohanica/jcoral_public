package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

func (g *client) read(r *http.Request) {

	for {
		var msg *message
		if err := g.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().String()
			msg.Name = g.userData["name"].(string)
			if avatarURL, ok := g.userData["avatar_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}

			g.room.forward <- msg
			//データベースへの保存です主にチャットデータになります。
			db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
			if err != nil {
				panic(err)
			}
			defer db.Close()
			no, err1 := db.Prepare("INSERT INTO allchats (url,avatar,name,content,time) VALUES(?,?,?,?,?)")
			if err1 != nil {
				panic(err1)
			}
			defer no.Close()
			_, err2 := no.Exec(r.URL.Path, msg.AvatarURL, msg.Name, msg.Message, msg.When)
			if err2 != nil {
				panic(err2)

			}
		} else {
			break
		}

	}
	g.socket.Close()

}

// クライアントへのレスポンスをおくるメゾットです
func (g *client) write() {
	for msg := range g.send {
		if err := g.socket.WriteJSON(msg); err != nil {
			break
		}

	}
	g.socket.Close()
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		fmt.Println(r.Header)
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

//MustAuth はハンドラのラップ用の認証系の関数です
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//ログインのロジックをつかさどるハンドラ関数です
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	//ログイン時
	case "login":

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			panic(err)
		}

		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		//ログイン後
	case "callback":

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			panic(err)
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			panic(err)
		}

		user, err := provider.GetUser(creds)
		if err != nil {
			panic(err)

		}

		authCookieValue := objx.New(map[string]interface{}{
			"name":       user.Name(),
			"avatar_url": user.AvatarURL(),
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  authCookieValue,
			Path:   "/",
			MaxAge: 10000})

		w.Header().Set("Location", "/thread")
		w.WriteHeader(http.StatusTemporaryRedirect)

	//サポートしてない操作及びプロバイダーの場合
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

//ただの初期化関数です
func newRoom() *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

//チャットの部屋のメゾットです
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// クライアントの入室時
			r.clients[client] = true
		case client := <-r.leave:
			// クライアントの退出時
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			//  メッセージを送るとき、来た時
			for client := range r.clients {
				client.send <- msg
			}
		}

	}
}

//ただの定数定義です
const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

//アップグレイダーの初期化です
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

//部屋のロジックのハンドラです
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	authCookie, err := req.Cookie("auth")
	if err != nil {
		panic(err)

	}
	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read(req)

}

func (t *templateHandlertochat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.temptochat = template.Must(template.ParseFiles(filepath.Join("htmls", t.filenametochat)))
	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err9 := db.Query("SELECT * FROM allchats WHERE url = ?;", r.URL.Path+"/room")
	if err9 != nil {
		panic(err9)
	}
	defer rows.Close()
	result := []receivecontent{}

	for rows.Next() {
		var data receivecontent

		err := rows.Scan(&data.As, &data.Bs, &data.Cs, &data.Ds, &data.Fs)
		if err != nil {
			panic(err)
		}

		result = append(result, data)

	}

	err1 := t.temptochat.Execute(w, result)
	if err1 != nil {
		panic(err1)
	}

}

func (t *templateHandlertothread) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.temptothread = template.Must(template.ParseFiles(filepath.Join("htmls", t.filenametothread)))

	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err9 := db.Query("SELECT * FROM threads;")
	if err9 != nil {
		panic(err9)
	}
	defer rows.Close()
	result := []receivethreads{}

	for rows.Next() {
		var data receivethreads

		err := rows.Scan(&data.At, &data.Bt, &data.Ct, &data.Dt, &data.Ft)
		if err != nil {
			panic(err)
		}

		result = append(result, data)

	}

	err1 := t.temptothread.Execute(w, result)
	if err1 != nil {
		panic(err1)
	}

}

func (t *templateHandlertologin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.temptologin = template.Must(template.ParseFiles(filepath.Join("htmls", t.filenametologin)))

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	err1 := t.temptologin.Execute(w, data)
	if err1 != nil {
		panic(err1)
	}

}

//スレッド作成をつかさどるハンドラです
func threadmake(w http.ResponseWriter, r *http.Request) {
	now1 := time.Now()
	compare := now1.Format(time.RFC1123Z)

	now := time.Now()
	room := newRoom()
	t := now.Nanosecond()
	s := r.PostFormValue("text")
	tag := r.PostFormValue("tag")

	URI1 := strconv.Itoa(t)
	http.Handle("/chat/"+URI1+"/room", room)
	u := string("/chat/" + URI1)

	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	no, err1 := db.Prepare("INSERT INTO threads (url,title,tags,time,comparetime) VALUES(?,?,?,?,?)")
	if err1 != nil {
		panic(err1)
	}
	defer no.Close()
	_, err2 := no.Exec(u, s, tag, t, compare)
	if err2 != nil {
		panic(err2)

	}

	w.Header().Set("Location", "/chat/"+URI1)
	w.WriteHeader(http.StatusTemporaryRedirect)
	go room.run()

}

//スレッドの立て直しです
func threadreturn() {
	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err9 := db.Query("SELECT * FROM threads;")
	if err9 != nil {
		panic(err9)
	}
	defer rows.Close()
	result := []receivethreads{}

	for rows.Next() {
		var data receivethreads

		err := rows.Scan(&data.At, &data.Bt, &data.Ct, &data.Dt, &data.Ft)
		if err != nil {
			panic(err)
		}

		result = append(result, data)

	}
	for _, d := range result {
		URI1 := d.At

		room := newRoom()
		http.Handle(URI1+"/room", room)
		go room.run()
	}

}
