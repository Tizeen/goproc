package cpu

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetCpu() (*Cpu, error) {
	cpuInfo, err := getCpuInfo()
	if err != nil {
		log.Fatalf(err.Error())
	}

	cpu := &Cpu{
		User:    cpuInfo[0],
		Nice:    cpuInfo[1],
		System:  cpuInfo[2],
		Idle:    cpuInfo[3],
		Iowait:  cpuInfo[4],
		Irq:     cpuInfo[5],
		Softirq: cpuInfo[6],
	}

	return cpu, nil
}

func getCpuInfo() ([]float64, error) {
	// 打开文件
	file, err := os.Open("/proc/stat")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var (
		cpuInfo []float64
		// 创建一个buffer的scanner
		scanner = bufio.NewScanner(file)
	)

	for scanner.Scan() {
		// 获取第一行内容，也是目标内容
		line := scanner.Text()
		// 根据空格分割内容，得到切片
		fields := strings.Fields(line)
		// 从第1个位置开始遍历切片，第0个位置是“cpu”，不需要
		for i := 1; i < len(fields); i++ {
			// 转换成float64类型
			val, err := strconv.ParseFloat(fields[i], 64)
			if err != nil {
				log.Fatal(err)
			}
			cpuInfo = append(cpuInfo, val)
		}
	}
	return cpuInfo, scanner.Err()
}
