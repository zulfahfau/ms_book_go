package main

import (
	"log"
	"net/http"
	"os"
	"root/handlers"
)

func main() {
	//handlefunc -> takes func creates http - add to default server
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// })
	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Nice babay")
	// })

	http.ListenAndServe(":9090", sm)

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
