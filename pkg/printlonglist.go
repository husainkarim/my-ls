package pkg

import (
	"fmt"
	"strings"
)

func PrintLongList(entry FileInfo) {
	// prepare the var to print
	mode := entry.Mode
	links := fmt.Sprintf("%3d", entry.Nlinks)
	user := entry.User
	group := entry.Group
	size := fmt.Sprintf("%6d", entry.Size)
	time := entry.ModTime.Format("Jan _2 15:04")
	name := entry.Name
	if entry.Dir {
		name = "\033[1;34m" + entry.Name + "\033[0m"
	}
	if strings.HasSuffix(name, ".sample") && !entry.Dir {
		name = "\033[1;32m" + entry.Name + "\033[0m"
	}
	// print the file with the correct format
	fmt.Println(mode, links, user, group, size, time, name)
}
