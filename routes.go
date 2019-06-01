package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type props struct {
	Props []Prop
}

type Prop struct {
	Name  string
	Value string
}

func Index(w http.ResponseWriter, r *http.Request) {
	session, err := SessionsStore.Get(r, "cookie-name")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess := getSession(session)

	session.Values["uuid"] = sess
	session.Save(r, w)

	history, _ := db.GetAll(sess.Uuid)

	resultJSON, _ := json.Marshal(history)

	varmap := map[string]interface{}{
		"Searches": string(resultJSON),
	}

	t, _ := template.ParseFiles("app/dist/index.html")
	t.Execute(w, varmap)
}

func History(w http.ResponseWriter, r *http.Request) {
	// Pre-flight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	session, err := SessionsStore.Get(r, "cookie-name")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess := getSession(session)

	result, _ := db.GetAll(sess.Uuid)

	resultJson, _ := json.Marshal(result)

	w.Write(resultJson)
}

func Search(w http.ResponseWriter, r *http.Request) {

	// Pre-flight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	session, err := SessionsStore.Get(r, "cookie-name")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess := getSession(session)

	session.Values["uuid"] = sess
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	route := r.URL.Query().Get("r")
	stop := r.URL.Query().Get("s")

	predictions, _ := GetPredictions(route, stop)

	db.Create(sess.Uuid, route, stop)

	w.Write(predictions)
}

func getSession(s *sessions.Session) Session {
	val := s.Values["uuid"]

	var session = Session{}
	session, ok := val.(Session)

	if !ok {
		id, _ := uuid.NewUUID()
		fmt.Println("new uuid created " + id.String())

		return Session{Uuid: id.String()}
	}

	return session
}
