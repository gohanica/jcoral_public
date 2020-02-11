package main

import (
	"time"
)

func Crawler() {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	layout := time.RFC1123Z
	for {
		select {
		case <-t.C:
		now:=time.Now()
		date:=date_from_db()
		var format_date:=[]time.Time
		for i range date{
		format_date
		}
		if
		

		}
	}
}
