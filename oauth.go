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

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		fmt.Println(r.Header)
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
		fmt.Println(r.Header)
	}
}

//MustAuth aaa
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//type of path /auth/action/provider
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
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

	case "callback":

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			panic(err)
		}

		// get the credentials
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			panic(err)
		}

		// get the user
		user, err := provider.GetUser(creds)
		if err != nil {
			panic(err)

		}

		// save some data
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

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
