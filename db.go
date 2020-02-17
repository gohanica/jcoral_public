package main

import (
	//	"fmt"
	"log"
	"time"
)

func CREATE_TABLE(str string) {
	_, err := DB.Query(`create table ` + DB_name + `.` + str + `(name varchar(255),id int);`)
	if err != nil {
		log.Fatal(err)
	}
}

func date_from_db() []time.Time {
	var dates []time.Time
	var date string
	rows, err := DB.Query(`select comparetime from ` + DB_name + `.threads;`)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&date)
		if err != nil {
			log.Fatal(err)
		}
		t, e := time.Parse(time.RFC1123Z, date)
		if e != nil {
			log.Fatal(e)
		}
		dates = append(dates, t)
	}
	return dates
}

func delete_comments(t time.Time) {
	var url string
	rows, err := DB.Query(`select url from `+DB_name+`.threads where comparetime=?;`, t)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&url)
		if err != nil {
			log.Fatal(err)
		}
		_, err = DB.Exec(`delete from `+DB_name+`.allchats where url= ?;`, url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func delete_comment_table(t time.Time) {
	var title string
	rows, err := DB.Query(`select title from `+DB_name+`.m_threads where create_time=?;`, t)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}
		_, err = DB.Exec(`drop table ` + DB_name + `.` + title)
		if err != nil {
			log.Fatal(err)
		}
	}
}
