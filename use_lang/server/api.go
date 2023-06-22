package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"main-module/git"
)

type Task struct {
	Lang  string  `json:"langage"`
	Ratio float64 `json:"ratio"`
}

type Tasks []Task

func HandlerRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/artivles", returnArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	langages := git.User_lang()

	task := Tasks{}

	for langage, raito := range langages {
		task = append(task, Task{Lang: langage, Ratio: raito})
	}
	json.NewEncoder(w).Encode(task)
}
