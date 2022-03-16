package main

import (

	"chat/handlers"
	"fmt"

	"os"

	"github.com/gorilla/mux"

	"net/http"
	"time"
)

func main() {

	router := mux.NewRouter()
	
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/sendmsg", handlers.SendMsgHandler).Methods("POST")
	router.HandleFunc("/getmsg", handlers.RecMsgHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server started and listening at http://localhost:8000")
	server.ListenAndServe()
}
