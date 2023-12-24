package pkg

import (
	"fmt"
	"strings"
)

func RecursiveList(list []FileInfo, dir, root, file string, time, reverse, hidden, index, longprint, recursive bool, f func(l []FileInfo, f string, h, i bool)) {
	fmt.Printf("%s:\n", root)                         // print the root of file
	f(list, file, hidden, index)                      // print the list depend on the formate requested
	dircontain, dirnames := AddDir(list, dir, hidden) // get the list of directory in the root
	maindir := MainRootDir(list)
	for i := range dircontain {
		l := GetFilesInfo(dircontain[i], maindir) // get new list for the new root
		newroot := root + "/" + dirnames[i]       // new printed root
		if strings.HasSuffix(root, "/") {
			newroot = root + dirnames[i]
		}
		fmt.Println()
		// put the new list into the recursive with new file and root
		Apply(l, dircontain[i], newroot, file, time, reverse, hidden, index, longprint, recursive)
	}
}

func MainRootDir(list []FileInfo) FileInfo {
	var maindir FileInfo
	for _, v := range list {
		if v.Name == "." {
			maindir = v
			maindir.Name = ".."
			break
		}
	}
	return maindir
}
