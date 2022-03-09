package main

import (
	db "chat/database"
	"chat/handlers"
	// "fmt"
	"os"

	"github.com/gorilla/mux"

	"net/http"
	"time"
)

func main() {
	Config := make(map[string]string)

	Config["url"] = "ec2-3-216-221-31.compute-1.amazonaws.com"
	Config["port"] = "5432"
	Config["user"] = "qzdiorjrtvqbyz"
	Config["password"] = "521a5ad9cc01b09f06b81b6c9eed3eb7a8d0b00e077453085e587aaa3e8cdf90"
	Config["db"] = "d91n7256lpp7sb"

	db.ConnectToDB(Config)

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/sendmsg", handlers.SendMsgHandler)
	router.HandleFunc("/getmsg", handlers.RecMsgHandler)

	port :=os.Getenv("PORT")
	// fmt.Println(port)
	// return
	if port==""{
		port = "8000"
	}
	// fmt.Println(port)
	server := &http.Server{
		Handler:      router,
		Addr:":"+port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.ListenAndServe()
}
