package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"root/data"
)

type Book struct {
	log *log.Logger
}

func NewBook(log *log.Logger) *Book {
	return &Book{log}
}

func (b *Book) getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("ini get")
	// GET PRODUCT
	lp := data.GetBook()
	//convert from list to json
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "faild to marshal", http.StatusInternalServerError)
	}
	w.Write(d)
}

func (b *Book) addBook(w http.ResponseWriter, r *http.Request) {
	// POST book
	log.Println("ini post")
}

func (b *Book) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.getBook(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.addBook(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
