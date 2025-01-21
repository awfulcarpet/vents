package main

import (
	"fmt"
	"log"
	"net/http"
)

func VentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "vents\n");
}

func Help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usage\n");
}

func Intro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Intro\n");
}


func main() {
	port := 8080;

	http.HandleFunc("/", Intro);
	http.HandleFunc("/help", Help);

	http.HandleFunc("/vent", VentsHandler);


	fmt.Printf("listening on port %d\n", port);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil));
}
