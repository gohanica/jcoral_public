package main

import (
	"net/http"
)

func POST(w http.ResponseWriter r *http.Request){
	if r.Method == "GET" {
		http.NotFound(w, r)
		return
	}
//	c, err := r.Cookie(CookieName)
//	if err != nil {
//		fmt.Fprintln(w, "cannot get user info")
//	}
//	userInfo := Cookie_to_UserInfo(c)

//	if userInfo.Id == 0 {
//		userInfo.Name = "plz enter your name in the box"
//		cookie := UserInfo_to_Cookie(*userInfo)
//		http.SetCookie(w, cookie)
//		http.Redirect(w, r, HTTP_localhost+Port+Forum_URL, 301)
//		return
//	}
	date := time.Now()
	contentType := r.Header["Content-Type"]
//	contentdata := ContentData{}
	contype := strings.Split(contentType[0], ";")
	switch contype[0] {
//	case "application/x-www-form-urlencoded":
//		content := r.PostFormValue("")
//		contentdata = ContentData{
//			Type:        "text",
//			Date:        date,
//			Content:     content,
//			Contributer: *userInfo,
//			IsImage:     false,
//		}
//	case "multipart/form-data":
//		local_filename := StoreContent(r, date, userInfo)
//		contentdata = ContentData{
//			Type:        "image",
//			Date:        date,
//			Content:     local_filename, //file are stored at StoreContent()
//			Contributer: *userInfo,
//			IsImage:     true,
//		}
		case "application/json":
		
		

	}

	Insert_Content(&contentdata)

	//リダイレクト
	http.Redirect(w, r, HTTP_localhost+Port+Forum_URL, 301)
}
