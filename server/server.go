package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("fan-al-harb is listening on port 8080")
	http.HandleFunc("/", PodName)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func PodName(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "The hostname for this pod is %s.", hostname)
}
