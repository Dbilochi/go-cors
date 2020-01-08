package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FetchBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("hi")
	books := make(map[string]string)
	books["Programming Ruby"] = "Dave Thomas"
	books["VueJs up & Running"] = "Callum Macrae"
	books["Angular in action"] = "Jeremy"
	response,_ := json.Marshal(books)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getbooks", FetchBooks).Methods("GET")
	log.Printf("Server Started at Port:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
