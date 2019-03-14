package cpu

import "encoding/json"

type Cpu struct {
	User    float64 `json:"user"`
	Nice    float64 `json:"nice"`
	System  float64 `json:"system"`
	Idle    float64 `json:"idle"`
	Iowait  float64 `json:"iowait"`
	Irq     float64 `json:"irq"`
	Softirq float64 `json:"softirq"`
}

func (c Cpu) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}
