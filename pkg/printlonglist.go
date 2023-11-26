package pkg

import (
	"fmt"
	"strings"
)

func PrintLongList(entry FileInfo) {
	// prepare the var to print
	mode := strings.ToLower(entry.Mode.String())
	if len(mode) == 11 {
		mode = mode[1:]
	}
	links := fmt.Sprintf("%3d", entry.Nlinks)
	user := entry.User
	group := entry.Group
	size := fmt.Sprintf("%7d", entry.Size)
	time := entry.ModTime.Format("Jan _2 15:04")
	color := SelectColor(entry.Mode.String())
	name := color + entry.Name + "\033[0m"
	// print the file with the correct format
	fmt.Println(mode, links, user, group, size, time, name)
}
