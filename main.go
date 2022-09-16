package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "heee", http.StatusNotFound)
		return

	}
	if r.Method != "GET" {
		http.Error(w, "hee wrong method", http.StatusMethodNotAllowed)
		return

	}
	fmt.Fprintf(w, "hello")
}
func FourmHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "hee", http.StatusNotFound)
		return

	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %v", name)
	fmt.Println(name)
	fmt.Fprintf(w, "address = %v", address)

}
func main() {
	StaticServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", StaticServer)
	http.HandleFunc("/hello", HelloHandler)

	http.HandleFunc("/form", FourmHandler)
	fmt.Println("starting server on 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)

	}

}
