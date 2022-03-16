package handlers

import (
	db "chat/database"
	"encoding/json"
	"fmt"

	"time"

	"net/http"
)

var messages map[string][]db.Message = make(map[string][]db.Message)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./static/html/homepage.html")

}

func SendMsgHandler(w http.ResponseWriter, r *http.Request) {

	var msg db.Message
	er := json.NewDecoder(r.Body).Decode(&msg)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
	}
	
	messages[msg.To] = append(messages[msg.To],msg)
		fmt.Println(messages)

}

func RecMsgHandler(w http.ResponseWriter, r *http.Request) {

	var user db.User
	er := json.NewDecoder(r.Body).Decode(&user)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
	}

	for i:=0;i<10;i++{

		msgs := messages[user.To]
		if len(msgs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msgs)
			delete(messages,user.To)
			return
		}
		time.Sleep(time.Second)
		
	}
	
}
