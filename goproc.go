package main

import (
	"fmt"
	"goproc/disk"
	"log"
)

func main() {

	s, err := disk.Partitions()
	if err != nil {
		// fatal: 致命的意思
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(err.Error())
	}

	fmt.Printf("The type is : %T\n", s)
	fmt.Println(s)
}
