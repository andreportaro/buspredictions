package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type props struct {
	Props []Prop
}

type Prop struct {
	Name  string
	Value string
}

func Index(w http.ResponseWriter, r *http.Request) {
	p := Prop{Name: "test", Value: "value"}

	t, _ := template.ParseFiles("app/dist/index.html")
	t.Execute(w, p)
}

func Search(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	// w.WriteHeader(http.StatusOK)

	route := r.URL.Query().Get("r")
	stop := r.URL.Query().Get("s")

	predictions, _ := GetPredictions(route, stop)

	fmt.Fprintln(w, predictions)
}
