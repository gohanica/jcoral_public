package main

import (
	"encoding/json"
	"fmt"
	"text/template"

	"io/ioutil"
	"net/http"
)

// Comment メッセージ用構造体
type Comment struct {
	Message  string `json:"message"`
	ID       string `json:"id"`
	Date     string `json:"date"`
	Profile  string `json:"profile"`
	Username string `json:"username"`
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
	http.HandleFunc("/book/templ", templ)

	server.ListenAndServe()

}

// テンプレート用ハンドラ
func templ(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/book.html")

	var tcomment Comment
	tcomment.Message = "おばけほーく"
	tcomment.Date = "2020/02/23/23:13"
	tcomment.ID = "19930130"
	tcomment.Profile = "../image/1.png"
	tcomment.Username = "千賀滉大"
	tcomments := []Comment{tcomment, tcomment}
	t.Execute(w, tcomments)
}

// 解析用ハンドラ
func process(w http.ResponseWriter, r *http.Request) {
	// メッセージ受信→json
	b, _ := ioutil.ReadAll(r.Body)
	front, _ := ioutil.ReadFile("tofront.json")
	var comment Comment

	json.Unmarshal(b, &comment)
	fmt.Println(comment)

	// json→メッセージ保存・送信
	output, _ := json.MarshalIndent(&comment, "", "\t\t")
	ioutil.WriteFile("messages.json", output, 0644)
	w.Header().Set("Content-Type", "application/json")
	w.Write(front)

}
