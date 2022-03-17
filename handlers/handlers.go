package handlers

import (
	db "chat/database"
	"encoding/json"
	"fmt"

	"time"

	"net/http"
)

var messages map[string][]db.Message = make(map[string][]db.Message)

var usersOnline map[string]int = make(map[string]int)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./static/html/homepage.html")

}

func SendMsgHandler(w http.ResponseWriter, r *http.Request) {

	var msg db.Message
	er := json.NewDecoder(r.Body).Decode(&msg)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
	}

	usersOnline[msg.From] += 1
	messages[msg.To] = append(messages[msg.To], msg)
	fmt.Println(messages)

}

func RecMsgHandler(w http.ResponseWriter, r *http.Request) {

	var user db.User
	er := json.NewDecoder(r.Body).Decode(&user)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
	}

	usersOnline[user.To] += 1
	for i := 0; i < 10; i++ {

		msgs := messages[user.To]
		if len(msgs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msgs)
			delete(messages, user.To)
			return
		}
		time.Sleep(time.Second)

	}

}

func UsersOnlineHandler(w http.ResponseWriter, r *http.Request) {
	checkOffline()
	users, _ := json.Marshal(usersOnline)
	resData := string(users)
	w.Write([]byte(resData))
}

func checkOffline() {

	for k, v := range usersOnline {
		if v > 10 {
			usersOnline[k] = 10
		} else if v < 0 {
			delete(usersOnline, k)
		} else{
			usersOnline[k] -= 1
		}
	}
	// time.Sleep(3 * time.Second)

}
