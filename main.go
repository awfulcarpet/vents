package main

import (
	"fmt"
	"log"
	"net/http"
)

func VentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			fmt.Fprintf(w, "here are some vents\n");
		default:
			http.Error(w, "Request Invalid", http.StatusMethodNotAllowed);
	}

	fmt.Fprintf(w, "vents\n");
}

func Help(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "help");
}

func Intro(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "intro");
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404);
	fmt.Fprintf(w, "404 Page not Found\n");
}


func main() {
	port := 8080;

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			NotFound(w, r);
			return;
		}
	});

	http.HandleFunc("/help", Help);

	http.HandleFunc("/vent", VentsHandler);
	http.HandleFunc("/vents", VentsHandler);


	fmt.Printf("listening on port %d\n", port);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil));
}
