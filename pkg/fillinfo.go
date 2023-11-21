package pkg

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"syscall"
)

// fill the information into the struct
func FillInfo(file fs.FileInfo, stat *syscall.Stat_t, name string) FileInfo {
	var temp FileInfo
	var err error
	curr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user info for file:", file.Name())
		os.Exit(1)
	}
	group, _ := user.LookupGroupId(curr.Gid)
	temp.Dir = file.IsDir()       // is directory
	temp.Mode = file.Mode()       // get mode
	temp.Nlinks = stat.Nlink      // number of files
	temp.User = curr.Username     // user
	temp.Group = group.Name       // group
	temp.Size = file.Size()       // size
	temp.ModTime = file.ModTime() // time
	temp.Index = int(stat.Ino)    // index
	if name != "" {               // if there customize name
		temp.Name = name
	} else { // no specific name
		temp.Name = file.Name()
	}
	return temp
}
