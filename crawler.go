package main

import (
	"time"
)

func Crawler() {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			now := time.Now()
			dates := date_from_db()
			for _, date := range dates {
				duration := date.AddDate(0, 0, 5).Sub(now)
				if duration < 0 {
					delete_comments(date)
				}
			}
		}
	}
}
