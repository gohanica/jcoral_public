package main

import (
	"log"
	"net/http"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	// 認証系に必要なデータです
	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		google.New("32005940021-dookshlkvj9jpdjlnjl0qjjsjbk4mkuc.apps.googleusercontent.com", "6PJmUe_4jxcitmExnlHs2Z_D", "http://localhost:8080/auth/callback/google"),
	)
	//サーバーが落ちた時、ウェブソケット通信の部屋の再起動をしてくれます
	threadreturn()
	//ハンドラたちです
	//スレッド選択のテンプレートハンドラです
	http.Handle("/thread/", MustAuth(&templateHandlertothread{filenametothread: "thread.html"}))
	//スレッドの作成をつかさどりマス
	http.HandleFunc("/thread/make", threadmake)
	//チャットページのテンプレートハンドラです
	http.Handle("/chat/", MustAuth(&templateHandlertochat{filenametochat: "chat.html"}))
	//ログインページ用のテンプレートハンドラです
	http.Handle("/login/", &templateHandlertologin{filenametologin: "login.html"})
	//ログインのロジックをつかさどります
	http.HandleFunc("/auth/", loginHandler)
	//ログアウトのためのハンドラです
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	//おなじみこれでサーバーを建てます
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("LIstenAndServe", err)
	}
}
