package pkg

import (
	"fmt"
	"os"
	"strings"
)

func PrintList(list []FileInfo, file string, hidden, index bool) {
	if file != "" {
		link, _ := os.Readlink(file)
		if link != "" {
			if strings.HasSuffix(link, "/") {
				var empty FileInfo
				list = GetFilesInfo(link, empty)
				file = ""
			}
		}
	}
	for i, entry := range list {
		if file != "" { // there is specific file
			if file == entry.Name {
				if index { // print index
					fmt.Printf("%6d ", entry.Index)
				}
				color := SelectColor(entry.Mode.String())
				name := color + entry.Name + "\033[0m"
				fmt.Printf("%s", name) // print file
			}
		} else { // there is no specific file
			color := SelectColor(entry.Mode.String())
			name := color + entry.Name + "\033[0m"
			if !strings.HasPrefix(entry.Name, ".") { // print all file without hidden file
				if index {
					fmt.Printf("%6d ", entry.Index)
				}
				fmt.Printf("%v", name)
				if i != len(list)-1 {
					fmt.Print("  ")
				}
			} else if hidden { // print all file
				if index {
					fmt.Printf("%d ", entry.Index)
				}
				fmt.Printf("%v", name)
				if i != len(list)-1 {
					fmt.Print("  ")
				}
			}
		}
	}
	fmt.Println()
}

// func GetFile(list []FileInfo, name string) (FileInfo, error) {
// 	for _, v := range list {
// 		if v.Name == name {
// 			return v, nil
// 		}
// 	}
// 	var empty FileInfo
// 	return empty, errors.New("Error")
// }
