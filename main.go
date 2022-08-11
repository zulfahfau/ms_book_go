package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"root/handlers"
	"time"
)

func main() {
	//handlefunc -> takes func creates http - add to default server
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	bh := handlers.NewBooks(l)

	sm := http.NewServeMux()
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/book", bh)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// })
	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Nice babay")
	// })

	//make a server
	s := &http.Server{
		Addr: ":9090", Handler: sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	// s.ListenAndServe()
	a := make(chan os.Signal, 1)
	signal.Notify(a, os.Interrupt)
	signal.Notify(a, os.Kill)

	sig := <-a
	log.Println("Recieved terminate, shutdown", sig)
	//wait till complete then shutdown the server
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

//MATERI VARIABLES

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// //var can be categorized

// var (
// 	actorName  = "actor"
// 	seasonName = "sea;son"
// )
// var (
// 	counter int = 0
// )

// //shadow variables : if u declare var outside func main, and
// //declare some var inside the func, the outside var will be shadowed by the inside.

// func main() {
// 	//3 types of variables
// 	// var i int
// 	// i=42
// 	// var j int = 12
// 	//	var j float32 = 12 u can change it to another type
// 	// k := 42
// 	// fmt.Printf("%v, %T", j, j)

// 	//u can change or convert the type of var only like this :
// 	var a int = 1
// 	var b string
// 	b = strconv.Itoa(a)
// 	//print convert variables
// 	fmt.Printf(b)
// }

//----------------------------------------------------------------

//MATERI PRIMITIVES
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	n := 100 == 100
// 	m := 100 == 101
// 	fmt.Printf("%v, %T\n", n, n)
// 	fmt.Printf("%v, %T", m, m)

// }
