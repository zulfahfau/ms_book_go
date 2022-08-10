package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// struct
type Hello struct {
	//variable log tipe nya log.logger
	log *log.Logger
}

// func new hello bakal make type hello dan going to return hello handler
func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Nice throw")
	//so ya didepan funtion itu nama kelas gitu
	// io.ReadAll(r.Body)
	d, err := io.ReadAll(r.Body)
	fmt.Fprintf(w, "Hello %s", d)
	if err != nil {
		http.Error(w, ("oops"), http.StatusInternalServerError)
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("oops"))
		return
	}
	// log.Println("data :", d)
}
