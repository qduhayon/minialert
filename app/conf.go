package app

import (
	"encoding/json"
	"os"
)

type Conf struct {
	CpuThreshold     int
	BatteryThreshold int
}

func loadConf() (conf Conf, err error) {
	file, err := os.Open("miniconf.json")
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&conf)

	file.Close()

	return
}
