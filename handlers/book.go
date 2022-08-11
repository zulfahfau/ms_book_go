package handlers

import (
	"log"
	"net/http"
	"root/data"
)

type Books struct {
	log *log.Logger
}

func NewBooks(log *log.Logger) *Books {
	return &Books{log}
}

func (b *Books) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.getBooks(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.addBook(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (b *Books) getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("ini get")
	// GET PRODUCT
	lp := data.GetBooks()
	//convert from list to json
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (b *Books) addBook(w http.ResponseWriter, r *http.Request) {
	// POST book
	log.Println("ini post")
	book := &data.Book{}
	err := book.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "failed to unmarshal", http.StatusBadRequest)
	}
	log.Printf("Book :%#v", book)
	data.AddBook(book)
}
