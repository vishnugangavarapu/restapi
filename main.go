package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BOOK struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
}

var books []BOOK

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&BOOK{})
}
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book BOOK
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
			var book BOOK
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000))
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return

		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	r := mux.NewRouter()

	books = append(books, BOOK{ID: "1", Isbn: "44543", Title: "Book one", Author: &Author{Firstname: "john", Lastname: "Doe"}})
	books = append(books, BOOK{ID: "2", Isbn: "76555", Title: "Book two", Author: &Author{Firstname: "vishnu", Lastname: "gangavarapu"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("CREATE")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("UPDATE")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
