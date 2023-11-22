package pkg

import (
	"fmt"
	"io/fs"
	"os/user"
	"syscall"
)

// fill the information into the struct
func FillInfo(file fs.FileInfo, name string) (FileInfo, error) {
	var temp FileInfo
	var err error
	curr, err := user.LookupId(fmt.Sprint(file.Sys().(*syscall.Stat_t).Uid))
	if err != nil {
		fmt.Println("Error getting user info for file:", file.Name())
		return temp, err
	}
	group, _ := user.LookupGroupId(curr.Gid)
	temp.Dir = file.IsDir()                            // is directory
	temp.Mode = file.Mode()                            // get mode
	temp.Nlinks = file.Sys().(*syscall.Stat_t).Nlink   // number of files
	temp.User = curr.Username                          // user
	temp.Group = group.Name                            // group
	temp.Size = file.Sys().(*syscall.Stat_t).Size      // size
	temp.ModTime = file.ModTime()                      // time
	temp.Index = int(file.Sys().(*syscall.Stat_t).Ino) // index
	if name != "" {                                    // if there customize name
		temp.Name = name
	} else { // no specific name
		temp.Name = file.Name()
	}
	return temp, nil
}
