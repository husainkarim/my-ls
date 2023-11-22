package pkg

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// function to create the main list
func GetFilesInfo(dir string, mainD FileInfo) []FileInfo {
	var List []FileInfo
	var files []fs.DirEntry
	var err error
	files, err = os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}
	// split to remove main directory name
	l := strings.Split(dir, "/")
	// merge the name to get the name
	mdir := strings.Join(l[:len(l)-1], "/")
	currdir, _ := AddMainDir(dir, ".") // main root
	maindir := mainD
	if mainD.Name == "" {
		maindir, _ = AddMainDir(mdir, "..") // sub main root
	}
	List = append(List, currdir, maindir) // add both main and sub main
	for _, entry := range files {         // loop to add other file
		file, err := entry.Info()
		if err != nil {
			fmt.Println("Error getting Info for file:", entry.Name())
			os.Exit(1)
		}
		// get the stat for the file to use it in the FillInfo function
		temp, _ := FillInfo(file, "") // add all element
		List = append(List, temp)     // add it to the list
	}
	return List
}
