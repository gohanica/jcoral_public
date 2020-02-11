package main

import (
	//	"fmt"
	"log"
	//	"time"
)

func CREATE_TABLE(str string) {
	_, err := DB.Query(`create table ` + DB_name + `.` + str + `(name varchar(255),id int);`)
	if err != nil {
		log.Fatal(err)
	}
}

//ret:none
//コンテンツのデータをデータベースに登録
func Insert_Content(data *ContentData) {
	ins, err := DB.Prepare("INSERT INTO " + Content_tbl_name + "(type,date,content,contributer_id) VALUES(?,?,?,?)")
	ShowErr(err)

	layout := "2006 JST Mon Jan 02 15:04:05"
	date := data.Date.Format(layout)
	_, err = ins.Exec(data.Type, date, data.Content, data.Contributer.Id)
	ShowErr(err)
}
