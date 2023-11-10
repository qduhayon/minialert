package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Hello() {
	fmt.Println("Hello")
}

func SendMetric(url string, metric Metric) {
	// Send request
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(&metric)
	request, err := http.NewRequest("POST", url+"/SendDataMetric", buffer)
	client := &http.Client{}

	// Read response
	resp, err := client.Do(request)
	if err != nil {
		log.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		log.Warnf("StatusCode received: %v", resp.StatusCode)
	}
}

func ListAlert(url string) {
	// Send request
	request, err := http.NewRequest("GET", url+"/GetAlertHistory", nil)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Error(err)
		return
	}

	// Read response
	if resp.StatusCode != 200 {
		log.Warnf("StatusCode received: %v", resp.StatusCode)
	}
	var alertHistory []Alert
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&alertHistory)
	if err != nil {
		log.Error(err)
	} else {
		for _, alert := range alertHistory {
			alert.log()
		}
	}

}
