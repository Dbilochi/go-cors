package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"
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
	handler := cors.Default().Handler(r)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST","PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Screen"},
	})
	handler = c.Handler(handler)
	r.HandleFunc("/getbooks", FetchBooks).Methods("GET")
	log.Printf("Server Started");
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), handler))
}
