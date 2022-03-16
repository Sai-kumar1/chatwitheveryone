package handlers

import (
	db "chat/database"
	"encoding/json"
	"fmt"

	// "fmt"
	// "sync"
	"time"

	// "fmt"
	"net/http"
)

// var requestMap map[string]*http.Request = make(map[string]*http.Request)
// var responseWriterMap map[string]http.ResponseWriter = make(map[string]http.ResponseWriter)

// var mutex sync.Mutex

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
	// if msg.To is not having an active request then insert in db
	// else respond to active request
	// if requestMap[msg.To] != nil {
		// RecMsgHandler(responseWriterMap[msg.To], requestMap[msg.To])
		// mutex.Lock()
		// delete(requestMap, msg.To)
		// delete(responseWriterMap,msg.To)
		// mutex.Unlock()
		// w.WriteHeader(http.StatusOK)
	// } else {
		// db.InsertMessage(msg)
	// }
	// fmt.Println(requestMap)
	// db.InsertMessage(msg)
	messages[msg.To] = append(messages[msg.To],msg)
		fmt.Println(messages)

}

func RecMsgHandler(w http.ResponseWriter, r *http.Request) {

	var user db.User
	er := json.NewDecoder(r.Body).Decode(&user)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
	}

	// msgs := db.GetUserMessages(w, user)
	// fmt.Println(msgs)
	for i:=0;i<10;i++{
		// msgs := db.GetUserMessages(w, user)
		msgs := messages[user.To]
		if len(msgs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msgs)
			delete(messages,user.To)
			return
		}
		time.Sleep(time.Second)
		
	}
	// if len(msgs) > 0 {
	// 	// fmt.Println(requestMap,responseWriterMap)
	// 	// fmt.Println(w,"\n",r)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(msgs)
	// } 
	// else {
		// mutex.Lock()
		// requestMap[user.To] = r
		// responseWriterMap[user.To] = w
		// mutex.Unlock()
		// w.Header().Set("Connection","Keep-Alive")
	// }
	// fmt.Println(requestMap)
	// fmt.Println(r,"\n",w)
}
