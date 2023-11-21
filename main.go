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
	if len(args) > 2 {  // if args length more than 2 error
		fmt.Println(args)
		fmt.Println("Error: argument length givin not supported it should be between 0 and 2 only!.")
		os.Exit(1)
	} else if len(args) != 0 { // if length not 0 get the flag & file
		if strings.HasPrefix(args[0], "-") && len(args[0]) > 1 { // '-' mean the argument is flag
			flag = args[0]
			if !pkg.CheckFlag(flag) { // check the flag is supported or not
				fmt.Println("Error: the flag not supported!")
				os.Exit(1)
			}
			args = args[1:] // remove the flag from list
		}
		if len(args) > 0 { // get the file name
			file = args[0]
			if strings.HasPrefix(file, "/") { // if the file stat with '/' this mean new directory
				err = os.Chdir(file)
				if err != nil {
					fmt.Println("can't fund this dir!")
					os.Exit(1)
				}
				currentDir, err = os.Getwd()
				if err != nil {
					fmt.Println("Error getting current directory:", err)
					return
				}
			} else { // check if it file or directory
				check, err := pkg.CheckFileNameDir(file, currentDir)
				if err != nil { // handle error
					fmt.Println(err)
					os.Exit(1)
				}
				if check { // if directory add it to root and empty the file
					currentDir += "/" + file
					file = ""
				}
			}
		}
	}
	// default value for the flags
	hidden := false
	index := false
	longprint := false
	recursive := false
	time := false
	reverse := false
	root := "."
	// get the list
	var empty pkg.FileInfo
	var listOfFile []pkg.FileInfo = pkg.GetFilesInfo(currentDir, empty)
	pkg.SortList(listOfFile)
	if flag == "" { // no flag normal print
		pkg.PrintList(listOfFile, file, hidden, index)
	} else { // if there is flag
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
		// put all the variable in the apply function
		pkg.Apply(listOfFile, currentDir, root, file, time, reverse, hidden, index, longprint, recursive)
	}
}
