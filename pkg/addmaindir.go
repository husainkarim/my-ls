package pkg

import (
	"fmt"
	"os"
)

// add main and sub directory
func AddMainDir(dir, name string) (FileInfo, error) {
	maindir, err := os.Open(dir) // by the directory name open the file
	if err != nil {
		fmt.Println("Error reading main directory:", err)
		os.Exit(1)
	}
	mainfile, err := maindir.Stat() // get file info
	if err != nil {
		fmt.Println("Error getting stat for file:", dir)
		os.Exit(1)
	}
	temp, _ := FillInfo(mainfile, name, dir)
	return temp, nil
}
