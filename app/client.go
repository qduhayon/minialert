package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"minialert/logger"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Hello() {
	fmt.Println("Hello")
}

func SendMetric(url string, metric Metric) {
	// Set logger
	log.SetFormatter(&logger.ClientFormatter{})

	// Send request
	buffer := bytes.Buffer{}
	json.NewEncoder(&buffer).Encode(&metric)
	request, err := http.NewRequest("POST", url+"/SendDataMetric", &buffer)
	client := &http.Client{}

	// Read response
	resp, err := client.Do(request)
	if err != nil {
		log.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		log.Warnf("StatusCode received: %v", resp.StatusCode)
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Warn(string(b))
		}
	}
}

func ListAlert(url string) {
	// Set logger
	log.SetFormatter(&logger.ClientFormatter{})

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
		// Display alerts from most recent to oldest
		for i := len(alertHistory) - 1; i >= 0; i-- {
			alertHistory[i].display()
		}
	}

}
