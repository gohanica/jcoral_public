package main

import (
	"log"
	"net/http"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {

	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		google.New("32005940021-dookshlkvj9jpdjlnjl0qjjsjbk4mkuc.apps.googleusercontent.com", "6PJmUe_4jxcitmExnlHs2Z_D", "http://localhost:8080/auth/callback/google"),
	)
	threadreturn()
	http.HandleFunc("/note/make", threadmake)
	//http.HandleFunc("/note/ajax", noteofajax)
	http.Handle("/note/", MustAuth(&templateHandlertonote{filenametonote: "note.html"}))
	http.Handle("/notelist/", MustAuth(&templateHandlertonotelist{filenametonotelist: "notelist.html"}))
	http.Handle("/login/", &templateHandlertologin{filenametologin: "login.html"})
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
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("LIstenAndServe", err)
	}

}
