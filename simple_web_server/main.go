package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsFired   bool   `json:"is_fired"`
}

func usersPage(writer http.ResponseWriter, r *http.Request) {
	users := []User{
		{"John", "Smith", true},
		{"Sam", "Max", false},
	}

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(writer, users); err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
}

func mainPage(writer http.ResponseWriter, r *http.Request) {
	//user := User{"John", "Smith"}
	//js, _ := json.Marshal(user)
	//writer.Write(js)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(writer, nil); err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPage)

	port := ":9090"
	println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
