package pkg

import (
	"fmt"
	"os"
	"syscall"
)

// add main and sub directory
func AddMainDir(dir, name string) FileInfo {
	maindir, err := os.Open(dir) // by the directory name open the file
	if err != nil {
		fmt.Println("Error reading main directory:", err)
		os.Exit(1)
	}
	mainfile, _ := maindir.Stat() // get file info
	var stat syscall.Stat_t
	err = syscall.Stat(dir, &stat) // get the stat of the file
	if err != nil {
		fmt.Println("Error getting stat for file:", dir)
		os.Exit(1)
	}
	temp := FillInfo(mainfile, &stat, name)
	return temp
}
