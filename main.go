package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func FetchBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []string
	books = append(books, "The Essential Rumi")
	books = append(books, "The book of Why")
	books = append(books, "When Breathe becomes air")
	books = append(books, "The Magic")
	books = append(books, "The Secret")


	response, _ := json.Marshal(books)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getbooks", FetchBooks).Methods("GET")
	log.Printf("Server Started");
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))
}
