package main

import "net/http"

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/chat.html")
	})
	http.HandleFunc("/chat", ToServer)
	http.ListenAndServe(":8080", nil)
}
