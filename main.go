package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
	filename string
	temp     *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.temp = template.Must(template.ParseFiles(filepath.Join("htmls", t.filename)))

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	err1 := t.temp.Execute(w, data)
	if err1 != nil {
		panic(err1)
	}

}

func main() {

	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		google.New("32005940021-dookshlkvj9jpdjlnjl0qjjsjbk4mkuc.apps.googleusercontent.com", "6PJmUe_4jxcitmExnlHs2Z_D", "http://localhost:8080/auth/callback/google"),
	)

	//	r := newRoom()
	http.Handle("/thread/", MustAuth(&templateHandler{filename: "thread.html"}))
	http.HandleFunc("/thread/make", threadmake)
	http.Handle("/chat/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login/", &templateHandler{filename: "login.html"})
	//	http.Handle("/chat/room", r)
	http.HandleFunc("/auth/", loginHandler)
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
	//	go r.run()
	go 
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("LIstenAndServe", err)
	}
}
