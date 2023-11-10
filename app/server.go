package app

import (
	"encoding/json"
	"fmt"
	"minialert/logger"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	battery            int  = 100
	cpu                int  = 5
	batteryThreshold   int  = 25
	cpuThreshold       int  = 50
	batteryAlarmStatus bool = false
	cpuAlarmStatus     bool = false
	alerts             []Alert
	mutex              sync.Mutex
)

func sendDataMetric(writer http.ResponseWriter, request *http.Request) {
	// Parse request body
	var metric Metric
	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&metric)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	// Update data values
	switch metric.Datatype {
	case "cpu":
		mutex.Lock()
		cpu = metric.Value
		mutex.Unlock()
		log.Info("New cpu value: ", cpu)
	case "battery":
		mutex.Lock()
		battery = metric.Value
		mutex.Unlock()
		log.Info("New battery value: ", battery)
	default:
		http.Error(writer, "Data type "+metric.Datatype+" not handled: use cpu or battery", http.StatusBadRequest)
	}
}

func getAlertHistory(writer http.ResponseWriter, request *http.Request) {
	err := json.NewEncoder(writer).Encode(alerts)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

// Monitoring loop
func monitorCpuBattery() {
	for {
		mutex.Lock()
		if cpu >= cpuThreshold && !cpuAlarmStatus {
			cpuAlarmStatus = true
			alert := Alert{"cpu", cpu, cpuThreshold, cpuAlarmStatus, time.Now().Format(time.DateTime)}
			alerts = append(alerts, alert)
			alert.log()
		}
		if cpu < cpuThreshold && cpuAlarmStatus {
			cpuAlarmStatus = false
			alert := Alert{"cpu", cpu, cpuThreshold, cpuAlarmStatus, time.Now().Format(time.DateTime)}
			alerts = append(alerts, alert)
			alert.log()
		}
		if battery <= batteryThreshold && !batteryAlarmStatus {
			batteryAlarmStatus = true
			alert := Alert{"battery", battery, batteryThreshold, batteryAlarmStatus, time.Now().Format(time.DateTime)}
			alerts = append(alerts, alert)
			alert.log()
		}
		if battery > batteryThreshold && batteryAlarmStatus {
			batteryAlarmStatus = false
			alert := Alert{"battery", battery, batteryThreshold, batteryAlarmStatus, time.Now().Format(time.DateTime)}
			alerts = append(alerts, alert)
			alert.log()
		}
		mutex.Unlock()

		// Sleep before next loop iteration
		time.Sleep(50 * time.Millisecond)
	}
}

func Serve(port int) {
	// Set logger
	log.SetFormatter(&logger.ServerFormatter{})

	// http routing
	http.HandleFunc("/SendDataMetric", sendDataMetric)
	http.HandleFunc("/GetAlertHistory", getAlertHistory)

	// Load alert configuration from json file
	conf, err := loadConf()
	if err != nil {
		log.Warn(err)
	} else {
		batteryThreshold = conf.BatteryThreshold
		cpuThreshold = conf.CpuThreshold
	}

	// monitor cpu and battery in a separate thread
	go monitorCpuBattery()

	// run server
	address := "localhost:" + fmt.Sprint(port)
	log.Infof("Listening on address %v\n", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
