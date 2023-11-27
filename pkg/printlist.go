package pkg

import (
	"fmt"
	"sort"
	"strings"
)

func PrintList(list []FileInfo, file string, hidden, index bool) {
	// fileCount := 0
	// ////
	// for i, entry := range list {
	// 	if file != "" { // there is specific file
	// 		if file == entry.Name {
	// 			if index { // print index
	// 				fmt.Printf("%6d ", entry.Index)
	// 			}
	// 			color := SelectColor(entry.Mode.String())
	// 			name := color + entry.Name + "\033[0m"
	// 			fmt.Printf("%s", name) // print file
	// 			fileCount++
	// 		}
	// 	} else { // there is no specific file
	// 		color := SelectColor(entry.Mode.String())
	// 		name := color + entry.Name + "\033[0m"
	// 		if !strings.HasPrefix(entry.Name, ".") { // print all file without hidden file
	// 			if index {
	// 				fmt.Printf("%6d ", entry.Index)
	// 			}
	// 			fmt.Printf("%v", name)
	// 			// if i != len(list)-1 {
	// 			// 	fmt.Print("  ")
	// 			// }
	// 			fileCount++
	// 			if fileCount%4 == 0 && i != len(list)-1 {
	// 				fmt.Println()
	// 				fileCount = 0
	// 			} else {
	// 				fmt.Print(" ")
	// 			}
	// 		} else if hidden { // print all file
	// 			if index {
	// 				fmt.Printf("%d ", entry.Index)
	// 			}
	// 			fmt.Printf("%v", name)
	// 			// if i != len(list)-1 {
	// 			// 	fmt.Print("  ")
	// 			// }
	// 			if fileCount%4 == 0 && i != len(list)-1 {
	// 				fmt.Println()
	// 				fileCount = 0
	// 			} else {
	// 				fmt.Print(" ")
	// 			}
	// 		}
	// 	}
	// }

	// if fileCount%4 != 0 {
	// 	fmt.Println()
	// }
	// fmt.Println()
	fileNames := []string{}
	for _, entry := range list {
		if file != "" && file == entry.Name {
			if index {
				fileNames = append(fileNames, fmt.Sprintf("%6d %s", entry.Index, entry.Name))
			} else {
				fileNames = append(fileNames, entry.Name)
			}
		} else if file == "" {
			if !strings.HasPrefix(entry.Name, ".") || hidden {
				if index {
					fileNames = append(fileNames, fmt.Sprintf("%6d %s", entry.Index, entry.Name))
				} else {
					fileNames = append(fileNames, entry.Name)
				}
			}
		}
	}

	sort.Strings(fileNames)

	const numColumns = 4
	columnWidth := 20
	numFiles := len(fileNames)

	// for i := 0; i < numFiles; i += numColumns {
	// 	for j := i; j < i+numColumns && j < numFiles; j++ {
	// 		fmt.Printf("%-*s", columnWidth, fileNames[j])
	// 	}
	// 	fmt.Println()
	// }

	//filesPerColumn := (numFiles + numColumns - 1) / numColumns // Calculate files per column
	filesPerColumn := 4

	for i := 0; i < filesPerColumn; i++ {
		for j := i; j < numFiles; j += filesPerColumn {
			fmt.Printf("%-*s", columnWidth, fileNames[j])
		}
		if i < filesPerColumn-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}
