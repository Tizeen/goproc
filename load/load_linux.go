package load

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetLoad() (*Loads, error) {

	var loads *Loads

	// loadavg文件内容很短，直接全部读取，返回[]byte切片
	data, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		log.Fatal("open /proc/loadavg error")
	}

	loadSlice, err := parseLoad(string(data))
	if err != nil {
		log.Fatal("parse data error")
	}

	loads = &Loads{
		Load1:  loadSlice[0],
		Load5:  loadSlice[1],
		Load15: loadSlice[2],
	}

	return loads, err
}

func parseLoad(data string) (loads []float64, err error) {

	loads = make([]float64, 3)
	// 注意Fields和Split的区别， Fields会去除空格，Split如果根据空格来分割，会保留原先存在的空格
	parts := strings.Fields(data)
	// 文件内容小于3,则有问题
	if len(parts) < 3 {
		return nil, fmt.Errorf("unexcepted content in /proc/loadavg")
	}

	// 遍历和赋值
	for i, load := range parts[0:3] {
		loads[i], err = strconv.ParseFloat(load, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse load '%s': %s", load, err)
		}
	}
	return loads, err
}
