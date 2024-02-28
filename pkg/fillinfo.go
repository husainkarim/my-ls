package pkg

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strings"
	"syscall"
)

// fill the information into the struct
func FillInfo(file fs.FileInfo, name, dir string) (FileInfo, error) {
	var temp FileInfo
	curUser, _ := user.LookupId(fmt.Sprint(file.Sys().(*syscall.Stat_t).Uid))    // get user info
	group, _ := user.LookupGroupId(fmt.Sprint(file.Sys().(*syscall.Stat_t).Gid)) // get group info
	temp.Dir = file.IsDir()                                                      // is directory
	temp.Mode = file.Mode()                                                      // get mode
	temp.Nlinks = int64(file.Sys().(*syscall.Stat_t).Nlink)                      // number of files
	if IntLen(temp.Nlinks) > S_link {
		S_link = IntLen(temp.Nlinks)
	}
	temp.User = curUser.Username // user
	if len(temp.User) > S_user {
		S_user = len(temp.User)
	}
	temp.Group = group.Name // group
	if len(temp.Group) > S_group {
		S_group = len(temp.Group)
	}
	temp.Size = file.Sys().(*syscall.Stat_t).Size // size
	Ssize := IntLen(temp.Size)
	// fmt.Println(temp.Mode.String())
	if strings.HasPrefix(temp.Mode.String()[1:], "c") {
		Ssize = 8
	}
	if Ssize > S_size {
		S_size = Ssize
	}
	temp.ModTime = file.ModTime()                      // time
	temp.Index = int(file.Sys().(*syscall.Stat_t).Ino) // index
	if IntLen(int64(temp.Index)) > S_index {
		S_index = IntLen(int64(temp.Index))
	}
	if name != "" { // if there customize name
		temp.Name = name
	} else { // no specific name
		temp.Name = file.Name()
	}
	if len(temp.Name) > S_name {
		S_name = len(temp.Name)
	}
	temp.Block = int(file.Sys().(*syscall.Stat_t).Blocks) // block size
	if strings.HasPrefix(temp.Mode.String(), "L") {
		currentDir, _ := os.Getwd()
		if dir != currentDir {
			temp.Link, _ = os.Readlink(dir + file.Name()) // get the link of the file
		} else {
			temp.Link, _ = os.Readlink(file.Name())
		}

	}
	return temp, nil
}

func IntLen(n int64) int {
	return len(fmt.Sprintf("%d", n))
}
