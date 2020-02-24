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
	// Heading  string
}

// Onepage 一ページに表示される内容
type Onepage struct {
	Comment []Comment
	Heading string
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
	tcomment.ID = "ID:19930130"
	tcomment.Profile = "../image/0.png"
	tcomment.Username = "千賀滉大"
	// tcomment.Heading = "2020/02/23"

	var tcomment2 Comment
	tcomment2.Message = "おばけほーく２"
	tcomment2.Date = "2020/02/23/23:13"
	tcomment2.ID = "ID:19930130"
	tcomment2.Profile = "../image/1.png"
	tcomment2.Username = "千賀滉大２"
	// tcomment2.Heading = "2020/02/24"

	tcomments := []Comment{tcomment, tcomment2, tcomment2}
	tcomments2 := []Comment{tcomment, tcomment2, tcomment, tcomment2}

	var onepage1 Onepage
	onepage1.Comment = tcomments
	onepage1.Heading = "2020/02/23"

	var onepage2 Onepage
	onepage2.Comment = tcomments2
	onepage2.Heading = "2020/02/24"
	tcommentss := []Onepage{onepage1, onepage2}
	// tcommentss := [][]Comment{tcomments, tcomments2}
	t.Execute(w, tcommentss)
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
