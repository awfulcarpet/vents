package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080;

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API First Steps\n");
	})


	fmt.Printf("listening on port %d\n", port);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil));
}
