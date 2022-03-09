package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

type User struct {
	To string `json:"user"`
}

var psql *sql.DB

func ConnectToDB(config map[string]string){
	host := config["url"]
	port := config["port"]
	user := config["user"]
	password := config["password"]
	dbname := config["db"]

	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var er error
	psql, er = sql.Open("postgres", psqlConn)
	if er!=nil{
		fmt.Println(er)
	}
	
}

func GetUserMessages(w http.ResponseWriter,user User) {
	fmt.Println(user,user.To)
	sqlQuery := fmt.Sprintf("SELECT * FROM messages WHERE to_user = '%v'", user.To)
	fmt.Println(sqlQuery)
	data, er := psql.Query(sqlQuery)
	delSqlQuery := fmt.Sprintf("DELETE FROM messages WHERE to_user='%v'",user.To)
	_, errrr := psql.Query(delSqlQuery)
	if errrr!=nil{
		fmt.Println(errrr)
	}
	if er != nil {
		fmt.Println("query error",er)
	}
	var manyMsgs []Message
	for data.Next() {
		var perMsg Message
		err := data.Scan(&perMsg.To, &perMsg.From, &perMsg.Message, &perMsg.Time)
		if err != nil {
			fmt.Println("Error Looping data", err)
		}
		manyMsgs = append(manyMsgs, perMsg)
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(manyMsgs)

}

func InsertMessage(msg Message) {

	sqlQuery := `INSERT INTO messages(to_user,from_user,message,time) VALUES($1,$2,$3,$4);`
	_, er := psql.Exec(sqlQuery, msg.To, msg.From, msg.Message, msg.Time)
	if er != nil {
		fmt.Println(er)
	}
}
