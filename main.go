package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Score struct {
	Name  string `json:"name,omitempty"`
	Score int    `json:"score,omitempty"`
}

var scores []Score

func main() {
	//Init router
	router := mux.NewRouter()

	//Create some data
	scores = append(scores, Score{Name: "Frank", Score: 100})
	scores = append(scores, Score{Name: "Joe", Score: 300})
	scores = append(scores, Score{Name: "Pete", Score: 999})

	//Define routes
	router.HandleFunc("/score", GetScore).Methods("GET")
	router.HandleFunc("/score", AddScore).Methods("POST")

	//Start http server
	log.Fatal(http.ListenAndServe(":8000", router))
}

//Returns current scores
func GetScore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetScore requested")
	json.NewEncoder(w).Encode(scores)
}

//Adds new score
func AddScore(w http.ResponseWriter, r *http.Request) {
	var score Score
	_ = json.NewDecoder(r.Body).Decode(&score)
	if score.Name != "" {
		//check for duplicates
		var found bool = false
		for index, s := range scores {
			if s.Name == score.Name {
				scores[index] = score
				found = true
			}
		}
		if !found {
			scores = append(scores, score)
		}
		json.NewEncoder(w).Encode(scores)
	} else {
		fmt.Fprintf(w, "Name was empty")
	}
}
