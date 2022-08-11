package data

import "time"

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

var bookList = []*Book{
	{ID: 1, Name: "Selena", Description: "Made by TereLiye forth", Price: 200000, SKU: "aaa12", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
	{ID: 2, Name: "BUMI", Description: "Made by TereLiye first", Price: 200000, SKU: "aaa13", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
}

func GetBook() []*Book {
	return bookList
}

// func AddBook()[]*Book{

// }
