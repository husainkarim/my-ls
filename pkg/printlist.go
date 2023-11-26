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
