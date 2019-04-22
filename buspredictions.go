package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func main() {
	InitConfig()

	http.Handle("/search", session(http.HandlerFunc(Search)))

	http.Handle("/", session(http.HandlerFunc(Index)))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/dist"))))

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

func session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := SessionsStore.Get(r, "uuid")
		if err != nil {
			fmt.Println(err)
			// Handle the error
		}

		if val, ok := session.Values["uuid"].(string); ok {
			fmt.Println("Session exists")
			fmt.Println(val)
		} else {
			fmt.Println("uuid does not exist")
			id, _ := uuid.NewUUID()

			session.Values["uuid"] = id.String()

			fmt.Println(session.Values)
		}

		if val, ok := session.Values["uuid"].(string); ok {
			w.Header().Set("tests", val)
		}

		// Pre-flight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		fmt.Println(session.Values)

		// Save it before we write to the response/return from the handler.
		sessionError := session.Save(r, w)
		if sessionError != nil {
			fmt.Println("error saving session")
			fmt.Println(sessionError)
			// handle the error case
		}

		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}
