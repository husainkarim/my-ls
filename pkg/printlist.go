package pkg

import (
	"fmt"
	"strings"
)

func PrintList(list []FileInfo, file string, hidden, index bool) {
	for i, entry := range list {
		if file != "" { // there is specific file
			if file == entry.Name {
				if index { // print index
					fmt.Printf("%d ", entry.Index)
				}
				fmt.Printf(entry.Name) // print file
			}
		} else { // there is no specific file
			name := entry.Name
			if entry.Dir { // if directory color blue
				name = "\033[1;34m" + entry.Name + "\033[0m"
			}
			if strings.HasSuffix(name, ".sample") && !entry.Dir { // if file end with '.sample' color green
				name = "\033[1;32m" + entry.Name + "\033[0m"
			}
			if !strings.HasPrefix(entry.Name, ".") { // print all file without hidden file
				if index {
					fmt.Printf("%d ", entry.Index)
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
