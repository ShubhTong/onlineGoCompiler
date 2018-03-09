package main

import (
	"log"
	"net/http"
)

func ca(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Form.Get("body"))

	w.Write([]byte("here is bytee"))
}

func main() {
	http.HandleFunc("/compileFile", ca)
	http.Handle("/", http.FileServer(http.Dir("templates")))

	log.Println("Listening...")
	http.ListenAndServe(":3001", nil)
}
