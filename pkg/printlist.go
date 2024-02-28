package pkg

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func PrintList(list []FileInfo, file, dir string, hidden, index bool) {
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
	count := 0
	var nlist []FileInfo
	if hidden {
		count = len(list)
		nlist = list
	} else {
		for _, f := range list {
			if !strings.HasPrefix(f.Name, ".") {
				nlist = append(nlist, f)
			}
		}
		count = len(nlist)
	}
	size := int(math.Ceil(float64(MaxWidth) / float64(S_name)))
	lines := int(math.Ceil(float64(count) / float64(size)))
	for i := 0; i < lines; i++ {
		for j := 0; j < size; j++ {
			if i+j*lines > len(nlist)-1 {
				break
			}
			if file != "" { // there is specific file
				if file == nlist[i+j*lines].Name {
					if index { // print index
						fmt.Printf("%d ", nlist[i+j*lines].Index)
					}
					color := SelectColor(nlist[i+j*lines].Mode.String())
					name := color + nlist[i+j*lines].Name + "\033[0m"
					fmt.Println(name) // print file
				}
			} else { // there is no specific file
				if lines > 1 { // if there is more than one line we should arrange the print
					color := SelectColor(nlist[i+j*lines].Mode.String())
					name := color + nlist[i+j*lines].Name + "\033[0m"
					if index {
						fmt.Printf("%6d ", nlist[i+j*lines].Index)
					}
					temp := name + strings.Repeat(" ", S_name-len(nlist[i+j*lines].Name)+1)
					fmt.Print(temp)
				} else { // if there is only one line to be printed
					color := SelectColor(nlist[i+j*lines].Mode.String())
					name := color + nlist[i+j*lines].Name + "\033[0m"
					if index {
						fmt.Printf("%6d ", nlist[i+j*lines].Index)
					}
					temp := name + strings.Repeat(" ", 2)
					fmt.Print(temp)
				}
			}
		}
		if file == "" {
			fmt.Println()
		}
	}
}
