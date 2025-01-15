package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"unicode/utf16"

	human "github.com/dustin/go-humanize"
)

const formatter = "%-5s %7s %7s %7s %7s\n"

func getDisks() ([]string, error) {
	n, e := windows.GetLogicalDriveStrings(0, nil)
	if e != nil {
		return nil, e
	}
	a := make([]uint16, n)

	windows.GetLogicalDriveStrings(n, &a[0])
	s := string(utf16.Decode(a))
	return strings.Split(strings.TrimRight(s, "\x00"), "\x00"), nil
}

func getDiskUsage(path string) (string, error) {

	var free, total, avail uint64

	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return "", fmt.Errorf("error Fetching Disk Usage: %s", err)
	}

	err = windows.GetDiskFreeSpaceEx(pathPtr, &avail, &total, &free)
	if err != nil {
		return "", fmt.Errorf("error Fetching Disk Usage: %s", err)
	}

	used := total - free
	percentUsed := fmt.Sprintf("%2.2f%%", float64(used)/float64(total)*100)

	return fmt.Sprintf(formatter, path, human.Bytes(total), human.Bytes(used), human.Bytes(free), percentUsed), nil
}

func main() {
	// if the disk name is specified by a command line argument
	if len(os.Args) > 1 {
		path := os.Args[1]

		diskInfo, err := getDiskUsage(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(formatter, "Disk", "Size", "Used", "Avail", "Use%")
		fmt.Print(diskInfo)
		return
	}

	// otherwise, getting information on all disks
	disksList, err := getDisks()
	if err != nil {
		panic(err)
	}

	fmt.Printf(formatter, "Disk", "Size", "Used", "Avail", "Use%")

	for _, s := range disksList {
		diskInfo, _ := getDiskUsage(s)
		fmt.Print(diskInfo)
	}

}
