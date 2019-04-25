package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	. "github.com/andreportaro/buspredictions/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var db = DB{}

type Session struct {
	Uuid string
}

func init() {
	InitConfig()

	gob.Register(Session{})

	db.Connect()
}

func main() {
	router := NewRouter()

	router.HandleFunc("/", Index)
	router.HandleFunc("/search", Search)
	router.HandleFunc("/history", History)

	fmt.Println("Listening on port " + GetPort())
	log.Fatal(http.ListenAndServe(":"+GetPort(), router))
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		return "9000"
	}
	return port
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Server CSS, JS & Images Statically.
	router.
		PathPrefix("/css/").
		Handler(http.StripPrefix("/css", http.FileServer(http.Dir("."+"/app/dist/css"))))

	router.
		PathPrefix("/js/").
		Handler(http.StripPrefix("/js", http.FileServer(http.Dir("."+"/app/dist/js"))))
	return router
}

func session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := SessionsStore.Get(r, "uuid")
		if err != nil {
			fmt.Println(err)
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

		fmt.Println(session.Values)

		// Save it before we write to the response/return from the handler.
		sessionError := session.Save(r, w)

		if sessionError != nil {
			log.Fatal(sessionError)
			// handle the error case
		}

		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}
