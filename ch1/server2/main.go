package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var count int

func counter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mutex.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Fprintf(w, "URL.Path = %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
