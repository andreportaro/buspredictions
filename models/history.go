package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

type DB struct {
}

type SearchHistory struct {
	Uuid    string `json:"uuid" bson:"uuid"`
	RouteID string `json:"route_id" bson:"route_id"`
	StopID  string `json:"stop_id" bson:"stop_id"`
}

func (d *DB) Connect() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB("test")
}

func (d *DB) Create(uuid, route, stop string) {
	var s SearchHistory
	s = SearchHistory{Uuid: uuid, RouteID: route, StopID: stop}

	err := db.C("search_history").Insert(&s)

	if err != nil {
		log.Fatal(err)
	}
}

func (d *DB) GetAll(uuid string) ([]SearchHistory, error) {
	var s []SearchHistory
	err := db.C("search_history").Find(bson.M{"uuid": uuid}).All(&s)
	return s, err
}
