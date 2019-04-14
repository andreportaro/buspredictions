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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.WriteHeader(http.StatusOK)

		predictions, _ := getPredictions("80", "11777")

		fmt.Fprintln(w, predictions)
	})

	fmt.Println("Now listening on port :9000")
	log.Fatal(http.ListenAndServe(":9000", nil))

}

func getPredictions(route string, stopID string) (string, error) {
	if xmlBytes, err := GetXML("http://webservices.nextbus.com/service/publicXMLFeed?command=predictions&a=ttc&r=" + route + "&stopId=" + stopID); err != nil {
		log.Printf("Failed to get XML: %v", err)
		return "Failed to get XML", err

	} else {
		// fmt.Println(string(xmlBytes))
		var result body
		xml.Unmarshal(xmlBytes, &result)

		jsonData, _ := json.Marshal(result)
		fmt.Println(string(jsonData))
		return string(jsonData), nil
	}
}

// GetXML
func GetXML(url string) ([]byte, error) {
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
