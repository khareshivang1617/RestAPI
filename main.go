package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc`
// 	Content string `json:"content"`
// }

// type Articles []Article

// func allArticles(w http.ResponseWriter, r *http.Request) {
// 	articles := Articles{
// 		Article{Title: "Title 1", Desc: "Desc 1", Content: "Content 1"},
// 	}
// 	fmt.Println("Endpoint hit")
// 	json.NewEncoder(w).Encode(articles)
// }

// func postArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "POST Articles")
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Homepage endpoint")
// }

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	// myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/users", NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("GO ORM")
	InitialMigration()
	handleRequests()
}
