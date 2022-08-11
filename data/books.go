package data

import (
	"encoding/json"
	"io"
	"time"
)

type Book struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

type Books []*Book

func (b *Books) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

func GetBooks() Books {
	return bookList
}

var bookList = []*Book{
	{ID: 1, Name: "Selena", Description: "Made by TereLiye forth", Price: 200000, SKU: "aaa12", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
	{ID: 2, Name: "BUMI", Description: "Made by TereLiye first", Price: 200000, SKU: "aaa13", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
}

func (b *Books) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}
