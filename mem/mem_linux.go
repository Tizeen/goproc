package mem

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetMemInfo() (*Mem, error) {

	// 打开文件
	file, err := os.Open("/proc/meminfo")

	if err != nil {
		log.Fatal(err)
	}

	// 函数返回时关闭打开的文件
	defer file.Close()

	// 调用parseMemInfo函数解析内存
	memInfo, err := parseMemInfo(file)

	mem := &Mem{
		Total:     memInfo["MemTotal_bytes"],
		Available: memInfo["Available_bytes"],
		Free:      memInfo["Free_bytes"],
		Buffer:    memInfo["Buffers_bytes"],
		Cached:    memInfo["Cached_bytes"],
	}

	return mem, err
}

func parseMemInfo(r io.Reader) (map[string]float64, error) {

	var (
		memInfo = map[string]float64{}
		// 返回一个scanner扫描器
		scanner = bufio.NewScanner(r)
		re      = regexp.MustCompile(`\((.*)\)`)
	)

	for scanner.Scan() {
		// 返回当前行： MemTotal:        8085152 kB
		line := scanner.Text()
		// 根据空格拆分字符串，返回string切片
		parts := strings.Fields(line)

		fv, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value in meminfo: %s", err)
		}

		// 将冒号去掉
		key := parts[0][:len(parts[0])-1]
		// 将括号部分去掉
		key = re.ReplaceAllString(key, "_{$1}")

		// 根据长度判断是否存在单位
		switch len(parts) {
		case 2:
		case 3:
			fv *= 1024
			key = key + "_bytes"
		default:
			return nil, fmt.Errorf("invalid line in meminfo: %s", line)
		}

		memInfo[key] = fv
	}

	return memInfo, scanner.Err()
}
