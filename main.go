package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Vents struct {
	Vents []Vent `json:"vents"`
}

type Vent struct {
	Date int64 `json:"date"`
	Content string `json:"content"`
}

func VentsHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := os.ReadFile("vents")
	var vents Vents
	json.Unmarshal(b, &vents)

	switch r.Method {
	case "GET": {

		for i := 0; i < len(vents.Vents); i++ {
			time := time.Unix(vents.Vents[i].Date, 0);

			fmt.Fprintf(w, "%02d/%02d/%02d %02d:%02d ", time.Year(), time.Month(), time.Day(),
				time.Hour(), time.Minute());
			fmt.Fprintf(w, "%s\n", vents.Vents[i].Content);
		}
	}
	case "POST": {
		body, _ := io.ReadAll(r.Body)
		new := &Vent {
			Date: time.Now().Unix(),
			Content: string(body),
		}

		vents.Vents = append(vents.Vents, *new)

		json, _ := json.Marshal(vents);

		os.WriteFile("vents", json, 0644);
	}
	default:
		http.Error(w, "Request Invalid", http.StatusMethodNotAllowed);
	}
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
