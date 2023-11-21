package pkg

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"syscall"
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
	currdir := AddMainDir(dir, ".") // main root
	maindir := mainD
	if mainD.Name == "" {
		maindir = AddMainDir(mdir, "..") // sub main root
	}
	List = append(List, currdir, maindir) // add both main and sub main
	for _, entry := range files {         // loop to add other file
		file, err := entry.Info()
		if err != nil {
			fmt.Println("Error getting Info for file:", entry.Name())
			continue
		}
		// get the stat for the file to use it in the FillInfo function
		var stat syscall.Stat_t
		err = syscall.Stat(dir+"/"+file.Name(), &stat)
		if err != nil {
			fmt.Println("Error getting stat for file:", file.Name())
			os.Exit(1)
		}
		temp := FillInfo(file, &stat, "") // add all element
		List = append(List, temp)         // add it to the list
	}
	return List
}
