package main

import "net/http"
import "log"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalf("error serving: %v", err)
	}
}
