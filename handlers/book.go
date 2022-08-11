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

func (b *Book) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetBook()
	//convert from list to json
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "faild to marshal", http.StatusInternalServerError)
	}
	w.Write(d)
}
