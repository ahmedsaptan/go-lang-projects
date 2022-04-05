package main

import (
	"fmt"
	"log"
	"net/http"
)

func formFunc(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "pasedForm() error: %v", err)
		return 
	}

	fmt.Fprintf(w, "Post request successfully");
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name= %s\n", name);
	fmt.Fprintf(w, "address= %s\n", address)
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}
func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formFunc)
	http.HandleFunc("/hello", helloFunc)

	fmt.Println("server is satrt on 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
