package pkg

import (
	"fmt"
	"strings"
)

func LongListFormat(list []FileInfo, file string, hidden, index bool) {
	if file == "" { // to print the total check the number of files
		count := 0
		for _, v := range list {
			if !hidden { // if hidden false ignore hidden files
				if !strings.HasPrefix(v.Name, ".") {
					count += v.Block
				}
			} else {
				count += v.Block
			}
		}
		// print total
		fmt.Printf("total %d\n", (count / 2))
	}
	// print files
	for _, entry := range list {
		if file != "" { // if there is file name print that file only
			if entry.Name == file {
				PrintLongList(entry)
			}
		} else { // no specific file name print the files
			if !hidden { // ignore hidden file
				if !strings.HasPrefix(entry.Name, ".") {
					if index {
						fmt.Printf("%6d ", entry.Index)
					}
					PrintLongList(entry)
				}
			} else { // print all file
				if index {
					fmt.Printf("%6d ", entry.Index)
				}
				PrintLongList(entry)
			}
		}
	}
}
