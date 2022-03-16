package database

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

type User struct {
	To string `json:"user"`
}
