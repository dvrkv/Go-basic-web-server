package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func slashHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	w.Write([]byte("Slash is a great way to check the server.\n\n"))
	w.Write([]byte("Web server works perfectly!\n"))
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Web server works perfectly!"
	if r.URL.Path != "/test" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("static/test.html")
	tmpl.Execute(w, msg)
	if err != nil {
		fmt.Fprintf(w, "ParseFiles() error: %v", err)
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		"Ben",
		"Milligan",
		27,
	}
	//Convert varible user of type User to JSON data
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(w, "Marshal() error: %v", err)
	}
	//Print
	w.Write([]byte("JSON data looks like this ↓\n\n"))
	w.Write([]byte(u))
}

func main() {
	//Call handlers
	http.HandleFunc("/", slashHandler)
	http.HandleFunc("/test", htmlHandler)
	http.HandleFunc("/json", jsonHandler)

	//Create connection
	port := ":8080"
	fmt.Print("Server listen on port", port, "\n")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe()", err)
	}
}
