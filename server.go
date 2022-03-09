package main

import (
	db "chat/database"
	"chat/handlers"
	"github.com/gorilla/mux"

	"net/http"
	"time"
)

func main() {
	Config := make(map[string]string)

	Config["url"] = "127.0.0.1"
	Config["port"] = "5432"
	Config["user"] = "postgres"
	Config["password"] = "saikumar"
	Config["db"] = "chat"

	db.ConnectToDB(Config)

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/sendmsg", handlers.SendMsgHandler)
	router.HandleFunc("/getmsg", handlers.RecMsgHandler)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.ListenAndServe()
}
