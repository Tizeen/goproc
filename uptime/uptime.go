package uptime

import "encoding/json"

type UpTime struct {
	Time float64 `json:"uptime"`
}

func (u UpTime) String() string {
	s, _ := json.Marshal(u)
	return string(s)
}
