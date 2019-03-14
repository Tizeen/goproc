package hostname

import (
	"encoding/json"
	"os"
)

type HostName struct {
	Name string `json:"hostname"`
}

func (h HostName) String() string {
	s, _ := json.Marshal(h)
	return string(s)
}

func GetHostName() (*HostName, error) {
	name, err := os.Hostname()
	hostName := &HostName{
		Name: name,
	}
	return hostName, err
}
