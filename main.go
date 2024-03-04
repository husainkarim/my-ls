package main

import (
	"fmt"
	"my-ls/pkg"
	"os"
	"strings"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	args := os.Args[1:] // get the argument
	flag := ""          // default
	file := ""          // default
	var DirList []pkg.Dir
	var dir pkg.Dir
	wrong := false
	if len(args) != 0 { // if length not 0 get the flag & file
		for _, s := range args {
			if strings.HasPrefix(s, "-") && len(s) > 1 { // '-' mean the argument is flag
				flag += s
				text, check := pkg.CheckFlag(flag)
				if !check { // check the flag is supported or not
					fmt.Println("./my-ls: invalid option '" + text + "'")
					os.Exit(1)
				}
			} else { // get the file name
				file = s
				if strings.HasPrefix(file, "/") { // if the file start with '/' this mean new directory
					if !strings.HasSuffix(file, "/") {
						file += "/"
					}
					dir.Name = file
					dir.Path = strings.ReplaceAll(file, "//", "/")
					dir.File = ""
					dir.IsDir = true
					DirList = append(DirList, dir)
				} else { // check if it file or directory
					isdir, isfile, err := pkg.CheckFileNameDir(file, currentDir)
					if err != nil { // handle error
						fmt.Println(err)
						wrong = true
					} else {
						if isdir { // if directory add it to root and empty the file
							dir.Name = file
							dir.Path = currentDir + "/" + file
							dir.File = ""
							dir.IsDir = true
							DirList = append(DirList, dir)
						}
						if isfile {
							if strings.Contains(file, "/") { // if file contain '/' this mean it is in another directory
								dir.Name = strings.Split(file, "/")[0]
								dir.Path = currentDir + "/" + strings.Split(file, "/")[:len(strings.Split(file, "/"))-1][0]
								dir.File = strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
								dir.IsDir = false
								DirList = append(DirList, dir)
							} else { // if file is in the current directory
								dir.Name = ""
								dir.Path = currentDir
								dir.File = file
								dir.IsDir = false
								DirList = append(DirList, dir)
							}
						}
					}
				}
			}
		}
	} else {
		dir.Name = ""
		dir.Path = currentDir
		dir.File = ""
		dir.IsDir = false
		DirList = append(DirList, dir)
	}
	if flag != "" && len(DirList) == 0 && !wrong {
		dir.Name = ""
		dir.Path = currentDir
		dir.File = ""
		dir.IsDir = false
		DirList = append(DirList, dir)
	}
	sortedList := pkg.SortDirList(DirList)
	hidden := false
	index := false
	longprint := false
	recursive := false
	time := false
	reverse := false
	if strings.Contains(flag, "t") { // check if time sort required
		time = true
	}
	if strings.Contains(flag, "r") { // check if reverse sort required
		reverse = true
	}
	if strings.Contains(flag, "a") { // print all file with hidden file
		hidden = true
	}
	if strings.Contains(flag, "l") { // use long print list
		longprint = true
	}
	if strings.Contains(flag, "R") { // use recursive print
		recursive = true
	}
	if strings.Contains(flag, "i") { // if index for the file required
		index = true
	}
	for c, v := range sortedList {
		path := v.Path
		file = v.File
		// default value for the flags
		root := "."
		if v.Name != "" {
			root = v.Name
		}
		// get the list
		var empty pkg.FileInfo
		var listOfFile []pkg.FileInfo = pkg.GetFilesInfo(path, empty)
		pkg.SortList(listOfFile)
		// put all the variable in the apply function
		if !recursive && v.IsDir && len(sortedList) > 1 {
			fmt.Println(v.Name + ":")
		}
		pkg.Apply(listOfFile, path, root, file, time, reverse, hidden, index, longprint, recursive)
		if c != len(sortedList)-1 {
			fmt.Println()
		}
	}
}
