package net

import "encoding/json"

type NetDevices struct {
	Devices []DeviceStat `json:"devices"`
}

type DeviceStat struct {
	Name      string `json:"name"`
	Operstate string `json:"state"`
}

func (n NetDevices) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}

func (n DeviceStat) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}
