package app

import (
	log "github.com/sirupsen/logrus"
)

type Alert struct {
	Datatype  string
	Value     int
	Threshold int
	Status    bool
	Date      string
}

// Function used for logging in the server
func (alert *Alert) log() {
	if alert.Status {
		if alert.Datatype == "cpu" {
			log.Warnf("[ALERT] [CPU_HIGH] [%v]", alert.Value)
		} else {
			log.Warnf("[ALERT] [BATTERY_LOW] [%v]", alert.Value)
		}
	} else {
		if alert.Datatype == "cpu" {
			log.Infof("[ALERT] [CPU_RESTORED] [%v]", alert.Value)
		} else {
			log.Infof("[ALERT] [BATTERY_RESTORED] [%v]", alert.Value)
		}
	}
}

// Function used to diplay alerts in the client
func (alert *Alert) display() {
	if alert.Status {
		if alert.Datatype == "cpu" {
			log.Warnf("[%s] [ALERT] [CPU_HIGH] [%v]", alert.Date, alert.Value)
		} else {
			log.Warnf("[%s] [ALERT] [BATTERY_LOW] [%v]", alert.Date, alert.Value)
		}
	} else {
		if alert.Datatype == "cpu" {
			log.Infof("[%s] [ALERT] [CPU_RESTORED] [%v]", alert.Date, alert.Value)
		} else {
			log.Infof("[%s] [ALERT] [BATTERY_RESTORED] [%v]", alert.Date, alert.Value)
		}
	}
}
