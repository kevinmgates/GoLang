package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":3000"
	fmt.Println("Starting server...") // println writes to stdout

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!") // Fprint uses the io writer
	})
	fmt.Println("Listening on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
