package mem

import "encoding/json"

type Mem struct {
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
	Free      float64 `json:"free"`
	Buffer    float64 `json:"buffers"`
	Cached    float64 `json:"cached"`
}

func (m Mem) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
