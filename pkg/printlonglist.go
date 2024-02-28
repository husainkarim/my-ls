package pkg

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
)

func PrintLongList(entry FileInfo, dir string) {
	// prepare the var to print
	mode := strings.ToLower(entry.Mode.String())
	if len(mode) == 11 {
		mode = mode[1:]
	}
	links := strings.Repeat(" ", S_link-len(fmt.Sprintf("%d", entry.Nlinks))) + fmt.Sprintf("%d", entry.Nlinks)
	user := entry.User + strings.Repeat(" ", S_user-len(entry.User))
	group := entry.Group + strings.Repeat(" ", S_group-len(entry.Group))
	size := strings.Repeat(" ", S_size-len(fmt.Sprintf("%d", entry.Size))) + fmt.Sprintf("%d", entry.Size)
	if strings.HasPrefix(mode, "c") {
		fileInfo, _ := os.Stat(dir + "/" + entry.Name)
		if fileInfo.Mode()&os.ModeDevice == os.ModeDevice && fileInfo.Mode()&os.ModeCharDevice == os.ModeCharDevice {
			stat, ok := fileInfo.Sys().(*syscall.Stat_t) // Cast to syscall.Stat_t for device numbers
			if !ok {
				fmt.Println("Error accessing device information:", entry)
			}
			size = fmt.Sprintf("%d, %d", stat.Rdev>>8, stat.Rdev&0xff) // Extract major and minor numbers
		}
	}
	currenttime := time.Now()
	sixMonthsAgo := currenttime.AddDate(0, -6, 0)
	time := entry.ModTime.Format("Jan _2 15:04")
	if entry.ModTime.Before(sixMonthsAgo) {
		time = entry.ModTime.Format("Jan _2  2006")
	}
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
