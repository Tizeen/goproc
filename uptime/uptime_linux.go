package uptime

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetUpTime() (*UpTime, error) {

	var upTime *UpTime

	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		log.Fatalf("Invalid content on file: /proc/uptime\n")
	}

	parts := strings.Fields(string(data))

	// 第一个值是系统启动到现在的时间（单位：秒）
	// 第二个值是系统空闲时间（单位：秒）
	time, err := strconv.ParseFloat(parts[0], 64)

	upTime = &UpTime{
		Time: time,
	}

	return upTime, err
}
