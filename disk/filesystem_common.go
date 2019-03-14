package disk

import "encoding/json"

type fileSystemLabels struct {
	device, mountPoint, fsType, options string
}

type PartitionStat struct {
	Device     string `json:"device"`
	MountPoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Opts       string `json:"options"`
}

func (p PartitionStat) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}
