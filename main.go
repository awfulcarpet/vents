package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var secret_route string = "/aoutnhskmjqwvziahtnigcraeouthnsoagcyrstikoeutanzuaoekuthneouaghcrjktnuoamwvueoagchr"

type Vents struct {
	Vents []Vent `json:"vents"`
}

type Vent struct {
	Date int64 `json:"date"`
	Content string `json:"content"`
}

func GetJSON(path string) Vents {
	b, _ := os.ReadFile("vents")
	var vents Vents
	json.Unmarshal(b, &vents)

	return vents
}

func VentsHandler(w http.ResponseWriter, r *http.Request) {
	vents := GetJSON("vents")

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
	http.ServeFile(w, r, "text/help");
}

func Intro(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "text/intro");
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404);
	fmt.Fprintf(w, "404 Page not Found\n");
}

func Secret(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		f, _ := os.ReadFile("secret")

		date := time.Now().Unix()
		secret, _ := strconv.Atoi(string(f));

		if (date - int64(secret) < 10) {
			os.Remove("vents")
			fmt.Fprintf(w, "*deletes database*\n")
		} else {
			fmt.Fprintf(w, "you do not have permission to delete the database\n")
		}
		return
	}

	fmt.Fprintf(w, "Why hello there.  I see you read the code and saw all the routes. Good for you that you actually read the source code.\n")
	fmt.Fprintf(w, "As a special treat, I'll allow you to nuke the db\n")
	fmt.Fprintf(w, "All you have to do is POST this route within 10 seconds and the job will be done\n")
	f, _ := os.Create("secret")

	f.WriteString(fmt.Sprintf("%d", time.Now().Unix()))
}

func LatestHandler(w http.ResponseWriter, r *http.Request) {
	vents := GetJSON("vents")
	last := len(vents.Vents)

	if (last <= 0) {
		fmt.Fprintf(w, "there are not vents\n")
		return
	}


	newest := last - 1

	query := r.URL.Query()
	num, exists := query["n"]

	if exists == true {
		count, err := strconv.Atoi(num[0])

		if err == nil {
			newest = last - count
		}
	}

	if (newest < 0) {
		fmt.Fprintf(w, "supplied argument is larger than total amount of vents in database\n")
		return
	}

	for i := newest; i < last; i++ {
		time := time.Unix(vents.Vents[i].Date, 0);

		fmt.Fprintf(w, "%02d/%02d/%02d %02d:%02d ", time.Year(), time.Month(), time.Day(),
			time.Hour(), time.Minute());
		fmt.Fprintf(w, "%s\n", vents.Vents[i].Content);
	}
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "vents")
}


func main() {
	port := 8080;

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			NotFound(w, r);
			return;
		}
		Intro(w, r)
	});

	http.HandleFunc("/help", Help);

	http.HandleFunc("/vent", VentsHandler);
	http.HandleFunc("/vents", VentsHandler);
	http.HandleFunc("/latest", LatestHandler);
	http.HandleFunc("/json", JsonHandler);

	http.HandleFunc(secret_route, Secret);


	fmt.Printf("listening on port %d\n", port);

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil));
}
