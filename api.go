package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type body struct {
	Directions []direction `xml:"predictions>direction"`
}

type prediction struct {
	Vehicle string `xml:"vehicle,attr"`
	Seconds string `xml:"seconds,attr"`
	Minutes string `xml:"minutes,attr"`
}

type direction struct {
	Predictions []prediction `xml:"prediction"`
}

// Calls the nextbus publicXMLfeed endpoint with the command=prediction
func GetPredictions(route string, stopID string) (string, error) {
	if xmlBytes, err := getXML("http://webservices.nextbus.com/service/publicXMLFeed?command=predictions&a=ttc&r=" + route + "&stopId=" + stopID); err != nil {
		log.Printf("Failed to get XML: %v", err)

		return "Failed to get XML", err
	} else {
		var result body
		xml.Unmarshal(xmlBytes, &result)

		jsonData, _ := json.Marshal(result)
		fmt.Println(string(jsonData))

		return string(jsonData), nil
	}
}

// getXML
func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
