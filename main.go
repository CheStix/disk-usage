package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"

	human "github.com/dustin/go-humanize"
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
	percentUsed := fmt.Sprintf("%2.2f%%", float64(used)/float64(total)*100)

	formatter := "%-5s %7s %7s %7s %7s\n"
	fmt.Printf(formatter, "Disk", "Size", "Used", "Avail", "Use%")
	fmt.Printf(formatter, path, human.Bytes(total), human.Bytes(used), human.Bytes(free), percentUsed)
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
