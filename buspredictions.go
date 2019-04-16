package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("app/dist"))
	http.Handle("/", fs)

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.WriteHeader(http.StatusOK)

		route := r.URL.Query().Get("r")
		stop := r.URL.Query().Get("s")

		predictions, _ := GetPredictions(route, stop)

		fmt.Fprintln(w, predictions)
	})

	fmt.Println("Listening on port" + GetPort())

	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "9000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
