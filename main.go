package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var counter = 0
var lastWords []string

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	hostname, _ := os.Hostname()
	lastWord := r.URL.Path[1:]

	if lastWord == "clear" {
		lastWords = []string{}
	} else if lastWord != "" {
		lastWords = append(lastWords, lastWord)
	}

	fmt.Fprintf(w, "Request: %v\nHostname: %v\nWords:%v", counter, hostname, lastWords)
}

func main() {
	port := "8080"
	log.Println("Listening in :" + port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
