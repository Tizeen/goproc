package net

import (
	"io/ioutil"
	"log"
	"strings"
)

func GetNetCardInfo() (*NetDevices, error) {
	// 获取所有网卡名字
	netCardNames, err := getNetCardsName()
	if err != nil {
		log.Fatal(err)
	}

	var (
		// 保存网卡数据
		deviceStat []DeviceStat
		// 保存网卡状态：启动、停止或其他状态
		netCardStatus []string
	)

	for _, name := range netCardNames {
		name, err := getNetCardStatus(name)
		if err != nil {
			log.Println("error:", err)
		}
		netCardStatus = append(netCardStatus, name)
	}

	for i := 0; i < len(netCardNames); i++ {
		deviceStat = append(deviceStat, DeviceStat{Name: netCardNames[i], Operstate: netCardStatus[i]})
	}
	devices := &NetDevices{
		Devices: deviceStat,
	}

	return devices, nil

}

func getNetCardsName() ([]string, error) {

	files, err := ioutil.ReadDir("/sys/class/net")
	if err != nil {
		return nil, err
	}

	var netCards []string

	for _, file := range files {
		netCards = append(netCards, file.Name())
	}

	return netCards, nil
}

func getNetCardStatus(name string) (string, error) {

	statFile := "/sys/class/net/" + name + "/operstate"

	data, err := ioutil.ReadFile(statFile)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(data), "\n"), nil
}
