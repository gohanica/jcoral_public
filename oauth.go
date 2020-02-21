package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

type authHandler struct {
	next http.Handler
}

//認証の利用をつかさどります
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
			"openID":     user.IDForProvider("google"),
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  authCookieValue,
			Path:   "/",
			MaxAge: 10000})

		w.Header().Set("Location", "/notelist")
		w.WriteHeader(http.StatusTemporaryRedirect)

	//サポートしてない操作及びプロバイダーの場合
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
