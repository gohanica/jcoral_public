package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/objx"
)

func threadmake(w http.ResponseWriter, r *http.Request) {
	now1 := time.Now()
	compare := now1.Format(time.RFC1123Z)

	now := time.Now()
	t := now.Nanosecond()
	s := r.PostFormValue("text")
	tag := r.PostFormValue("tag")

	URI1 := strconv.Itoa(t)
	u := string("/note/" + URI1)

	http.HandleFunc("/note/"+URI1+"/ajax", noteofajax)

	authCookie, err := r.Cookie("auth")
	if err != nil {
		panic(err)
	}

	var userData map[string]interface{}

	userData = objx.MustFromBase64(authCookie.Value)

	name0 := userData["name"].(string)
	name := "`" + name0 + "`"

	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	no, err1 := db.Prepare("INSERT INTO notelist (url,title,tags,time,comparetime,name) VALUES(?,?,?,?,?,?)")
	if err1 != nil {
		panic(err1)
	}
	defer no.Close()
	_, err2 := no.Exec(u, s, tag, t, compare, name)
	if err2 != nil {
		panic(err2)

	}

	w.Header().Set("Location", "/note/"+URI1)
	w.WriteHeader(http.StatusTemporaryRedirect)

}

func threadreturn() {
	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err9 := db.Query("SELECT * FROM notelist;")
	if err9 != nil {
		panic(err9)
	}
	defer rows.Close()
	result := []receivethreads{}

	for rows.Next() {
		var data receivethreads

		err := rows.Scan(&data.URL, &data.Title, &data.Tags, &data.Time, &data.Comparetime, &data.Name)
		if err != nil {
			panic(err)
		}

		result = append(result, data)

	}
	for _, d := range result {
		URI1 := d.URL
		http.HandleFunc(URI1+"/ajax", noteofajax)

	}

}
