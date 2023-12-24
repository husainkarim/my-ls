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
	linke := entry.Link
	link := ""
	if linke != "" {
		if strings.HasSuffix(linke, "/") {
			link = "-> " + "\033[1;34m" + linke + "\033[0m"
		} else {
			link = "-> " + linke
		}
	}
	// print the file with the correct format
	fmt.Println(mode, links, user, group, size, time, name, link)
}
