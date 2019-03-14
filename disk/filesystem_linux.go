package disk

import (
	"bufio"
	"goproc/common"
	"log"
	"os"
	"strings"
)

func MountPointDetails() ([]fileSystemLabels, error) {
	file, err := os.Open("/proc/1/mounts")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var filesystems []fileSystemLabels
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		parts[1] = strings.Replace(parts[1], "\\040", " ", -1)
		parts[1] = strings.Replace(parts[1], "\\011", "\t", -1)

		if !strings.HasPrefix(parts[0], "/dev") {
			continue
		}

		filesystems = append(filesystems, fileSystemLabels{
			device:     parts[0],
			mountPoint: parts[1],
			fsType:     parts[2],
			options:    parts[3],
		})
	}

	return filesystems, scanner.Err()

}

func Partitions() ([]PartitionStat, error) {

	file, err := os.Open("/proc/self/mounts")
	if err != nil {
		log.Println("Error:", err.Error())
	}

	fs, err := getFileSystems()
	if err != nil {
		log.Println("Error:", err.Error())
	}
	//fmt.Println(fs)

	ret := make([]PartitionStat, 0, 5)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		d := PartitionStat{
			Device:     fields[0],
			MountPoint: fields[1],
			Fstype:     fields[2],
			Opts:       fields[3],
		}
		if d.Device == "none" || !common.StringsHas(fs, d.Fstype) {
			continue
		}
		ret = append(ret, d)
	}
	return ret, nil
}

func getFileSystems() ([]string, error) {
	file, err := os.Open("/proc/filesystems")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var ret []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "nodev") {
			ret = append(ret, strings.TrimSpace(line))
			continue
		}
		t := strings.Split(line, "\t")
		if len(t) != 2 || t[1] != "ext4" {
			continue
		}
		ret = append(ret, strings.TrimSpace(t[1]))
	}

	return ret, nil
}
