package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
)

func getDiskUsage(path string) {

	var free, total, avail uint64

	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		fmt.Println("Error Fetching Disk Usage:", err)
		return
	}

	err = windows.GetDiskFreeSpaceEx(pathPtr, &avail, &total, &free)
	if err != nil {
		fmt.Println("Error Fetching Disk Usage:", err)
		return
	}

	used := total - free
	percentUsed := float64(used) / float64(total) * 100

	fmt.Printf("Disk usage of %s:\n", path)
	fmt.Printf("Total: %d GB\n", total/1e9)
	fmt.Printf("Used: %d GB (%.2f%%)\n", used/1e9, percentUsed)
	fmt.Printf("Free: %d GB\n", free/1e9)
	//fmt.Println("Free:", free, "Total:", total, "Available:", avail)
}

func main() {
	path := "C:\\"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Errors: '%s' Path doesn't exist.\n", path)
		return
	} else if err != nil {
		fmt.Printf("Error occurred whilre accessing path %s: %v \n", path, err)
		return
	}
	getDiskUsage(path)
}
