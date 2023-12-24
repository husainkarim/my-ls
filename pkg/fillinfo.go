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
	curUser, _ := user.LookupId(fmt.Sprint(file.Sys().(*syscall.Stat_t).Uid))    // get user info
	group, _ := user.LookupGroupId(fmt.Sprint(file.Sys().(*syscall.Stat_t).Gid)) // get group info
	temp.Dir = file.IsDir()                                                      // is directory
	temp.Mode = file.Mode()                                                      // get mode
	temp.Nlinks = file.Sys().(*syscall.Stat_t).Nlink                             // number of files
	temp.User = curUser.Username                                                 // user
	temp.Group = group.Name                                                      // group
	temp.Size = file.Sys().(*syscall.Stat_t).Size                                // size
	temp.ModTime = file.ModTime()                                                // time
	temp.Index = int(file.Sys().(*syscall.Stat_t).Ino)                           // index
	temp.Block = int(file.Sys().(*syscall.Stat_t).Blocks)                        // block size
	if name != "" {                                                              // if there customize name
		temp.Name = name
	} else { // no specific name
		temp.Name = file.Name()
	}
	return temp, nil
}
