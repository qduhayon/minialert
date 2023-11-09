package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
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

// Struct used for json parsing and validation
type SendRequestBody struct {
	Datatype string
	Value    int
}

type Alert struct {
	Datatype  string
	value     int
	threshold int
	status    bool
	date      string
}

func sendDataMetric(writer http.ResponseWriter, request *http.Request) {
	// Parse request body
	var requestBody SendRequestBody
	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&requestBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	// Update data values
	switch requestBody.Datatype {
	case "cpu":
		mutex.Lock()
		cpu = requestBody.Value
		mutex.Unlock()
		fmt.Println("New cpu value: ", cpu)
	case "battery":
		mutex.Lock()
		battery = requestBody.Value
		mutex.Unlock()
		fmt.Println("New battery value: ", battery)
	default:
		http.Error(writer, "Data type "+requestBody.Datatype+" not handled", http.StatusBadRequest)
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
			alerts = append(alerts, Alert{"cpu", cpu, cpuThreshold, cpuAlarmStatus, time.Now().Format(time.RFC850)})
			fmt.Println(alerts[len(alerts)-1])
		}
		if cpu < cpuThreshold && cpuAlarmStatus {
			cpuAlarmStatus = false
			alerts = append(alerts, Alert{"cpu", cpu, cpuThreshold, cpuAlarmStatus, time.Now().Format(time.RFC850)})
			fmt.Println(alerts[len(alerts)-1])
		}
		if battery <= batteryThreshold && !batteryAlarmStatus {
			batteryAlarmStatus = true
			alerts = append(alerts, Alert{"battery", battery, batteryThreshold, batteryAlarmStatus, time.Now().Format(time.RFC850)})
			fmt.Println(alerts[len(alerts)-1])
		}
		if battery > batteryThreshold && batteryAlarmStatus {
			batteryAlarmStatus = false
			alerts = append(alerts, Alert{"battery", battery, batteryThreshold, batteryAlarmStatus, time.Now().Format(time.RFC850)})
			fmt.Println(alerts[len(alerts)-1])
		}
		mutex.Unlock()

		// Sleep before next loop iteration
		time.Sleep(50 * time.Millisecond)
	}
}

func Serve(port int) {
	// http routing
	http.HandleFunc("/SendDataMetric", sendDataMetric)
	http.HandleFunc("/GetAlertHistory", getAlertHistory)

	// monitor cpu and battery in a separate thread
	go monitorCpuBattery()

	// run server
	address := "localhost:" + fmt.Sprint(port)
	fmt.Printf("Listening on address %s\n", address)
	http.ListenAndServe(address, nil)
}
