package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type templateHandlertologin struct {
	filenametologin string
	temptologin     *template.Template
}

type templateHandlertonotelist struct {
	filenametonotelist string
	temptonotelist     *template.Template
}

type templateHandlertonote struct {
	filenametonote string
	temptonote     *template.Template
}

type receivecontent struct {
	Name      string
	Message   string
	When      string
	AvatarURL string
	ID        string
	URL       string
}

type receivethreads struct {
	URL         string
	Title       string
	Tags        string
	Time        int
	Comparetime string
	Name        string
}

func (t *templateHandlertologin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.temptologin = template.Must(template.ParseFiles(filepath.Join("htmls", t.filenametologin)))

	data := map[string]interface{}{
		"Host": r.Host,
	}

	err1 := t.temptologin.Execute(w, data)
	if err1 != nil {
		panic(err1)
	}

}

func (t *templateHandlertonotelist) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.temptonotelist = template.Must(template.ParseFiles(filepath.Join("htmls", t.filenametonotelist)))

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

	err1 := t.temptonotelist.Execute(w, result)
	if err1 != nil {
		panic(err1)
	}

}

func (t *templateHandlertonote) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.temptonote = template.Must(template.ParseFiles(filepath.Join("htmls", t.filenametonote)))

	db, err := sql.Open("mysql", "root:Nyantech0604@/db_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err9 := db.Query("SELECT * FROM message WHERE url = ?;", r.URL.Path+"/ajax")
	if err9 != nil {
		panic(err9)
	}
	defer rows.Close()
	result := []receivecontent{}

	for rows.Next() {
		var data receivecontent

		err := rows.Scan(&data.Name, &data.Message, &data.When, &data.AvatarURL, &data.ID, &data.URL)
		if err != nil {
			panic(err)
		}

		result = append(result, data)

	}

	err1 := t.temptonote.Execute(w, result)
	if err1 != nil {
		panic(err1)
	}

}
