package database

// import (
// 	// "database/sql"
// 	// "encoding/json"
// 	// "fmt"
// 	// "io/ioutil"
// 	// "net/http"

// 	_ "github.com/lib/pq"
// 	// "gopkg.in/yaml.v2"
// )

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

type User struct {
	To string `json:"user"`
}

// type psqlConfig struct {
// 	PSQLURI string `yaml:"psql_uri"`
// }

// var psql *sql.DB

// func ConnectToDB() {
// 	yamlData, err := ioutil.ReadFile("./config.yml")
// 	if err != nil {
// 		fmt.Println("unable to read config.yml file")
// 	}
// 	var config psqlConfig
// 	yaml.Unmarshal(yamlData, &config)

// 	psqlConn := config.PSQLURI
// 	var er error
// 	psql, er = sql.Open("postgres", psqlConn)
// 	if er != nil {
// 		fmt.Println(er)
// 	} else {
// 		fmt.Println("psql connected....")
// 	}

// }

// func GetUserMessages(w http.ResponseWriter, user User) []Message {
// 	fmt.Println(user, user.To)
// 	sqlQuery := fmt.Sprintf("SELECT * FROM messages WHERE to_user = '%v'", user.To)
// 	// fmt.Println(sqlQuery)
// 	data, er := psql.Query(sqlQuery)

// 	defer data.Close()
// 	delSqlQuery := fmt.Sprintf("DELETE FROM messages WHERE to_user='%v'", user.To)
// 	delData, errrr := psql.Query(delSqlQuery)
// 	defer delData.Close()
// 	if errrr != nil {
// 		fmt.Println(errrr)
// 	}
// 	if er != nil {
// 		fmt.Println("query error", er)
// 	}
// 	var manyMsgs []Message
// 	for data.Next() {
// 		var perMsg Message
// 		err := data.Scan(&perMsg.To, &perMsg.From, &perMsg.Message, &perMsg.Time)
// 		if err != nil {
// 			fmt.Println("Error Looping data", err)
// 		}
// 		manyMsgs = append(manyMsgs, perMsg)
// 	}
// 	// w.Header().Set("Content-Type","application/json")
// 	// json.NewEncoder(w).Encode(manyMsgs)
// 	return manyMsgs

// }

// func InsertMessage(msg Message) {

// 	sqlQuery := `INSERT INTO messages(to_user,from_user,message,time) VALUES($1,$2,$3,$4);`
// 	_, er := psql.Exec(sqlQuery, msg.To, msg.From, msg.Message, msg.Time)
// 	if er != nil {
// 		fmt.Println(er)
// 	}
// }
