package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Star メッセージ用構造体
type Star struct {
	Message string `json:"message"`
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/book/ajax", process)
	http.Handle("/book/", http.StripPrefix("/book/", http.FileServer(http.Dir("templates/"))))
	server.ListenAndServe()

}

// 解析
func process(w http.ResponseWriter, r *http.Request) {
	// メッセージ受信→json
	b, _ := ioutil.ReadAll(r.Body)
	var star Star
	json.Unmarshal(b, &star)
	fmt.Println(star)

	// json→メッセージ保存・送信
	output, _ := json.MarshalIndent(&star, "", "\t\t")
	ioutil.WriteFile("messages.json", output, 0644)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}
