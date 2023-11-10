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

func (alert *Alert) log() {
	if alert.Status {
		if alert.Datatype == "cpu" {
			log.Warnf("[ALERT] [CPU_HIGH] [%v]\n", alert.Value)
		} else {
			log.Warnf("[ALERT] [BATTERY_LOW] [%v]\n", alert.Value)
		}
	} else {
		if alert.Datatype == "cpu" {
			log.Infof("[ALERT] [CPU_RESTORED] [%v]\n", alert.Value)
		} else {
			log.Infof("[ALERT] [BATTERY_RESTORED] [%v]\n", alert.Value)
		}
	}
}
