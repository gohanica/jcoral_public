package main

import "time"

// Contentdata 投稿のデータ
type Contentdata struct {
	Time    time.Time
	Content string
	User    Userdata
}

// Userdata ユーザーの情報
type Userdata struct {
	Name      string
	ID        int
	AvatarURL string
}

// Thread スレッドの情報
type Thread struct {
	Data Contentdata
	Tags []string
}

// Comment コメント一つのデータ
type Comment struct {
	Data Contentdata
	ID   int
	Good int
	Bad  int
}
