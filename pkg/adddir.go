package pkg

import "strings"

// create two lis for directory root and directory name
func AddDir(list []FileInfo, dir string, hidden bool) ([]string, []string) {
	var rootList []string
	var dirname []string
	for _, v := range list {
		if v.Name == "." || v.Name == ".." {
			continue
		}
		if !hidden { // if hidden file not needed avoid it
			if !strings.HasPrefix(v.Name, ".") { // hidden directory always start with '.'
				if v.Dir { // if the file is directory go on
					root := dir + "/" + v.Name + "/"  //root name
					dirname = append(dirname, v.Name) // add name
					rootList = append(rootList, root) // add root
				}
			}
		} else { // if all file needed
			if v.Dir { // if the file is directory go on
				root := dir + "/" + v.Name + "/"
				dirname = append(dirname, v.Name) // add name
				rootList = append(rootList, root) // add root
			}
		}
	}
	return rootList, dirname
}
