package handlers

import (
	db "chat/database"
	"encoding/json"
	"fmt"
	"net/http"
)



func HomeHandler(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		http.ServeFile(w,r,"./static/html/homepage.html")
	}
	
}

func SendMsgHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST"{
		var msg db.Message
		er := json.NewDecoder(r.Body).Decode(&msg)
		if er!=nil{
			fmt.Println(er)
			return
		}
		db.InsertMessage(msg)
	}
}


func RecMsgHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST"{
		var user db.User
		er := json.NewDecoder(r.Body).Decode(&user)
		if er != nil{
			http.Error(w, er.Error(), http.StatusBadRequest)
        	return
		}
		db.GetUserMessages(w,user)
	}
}