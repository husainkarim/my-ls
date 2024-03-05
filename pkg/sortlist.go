package pkg

import "strings"

// sort the file information list by name
func SortList(list []FileInfo) {
	i := 0
	for len(list)-1 > i {
		name0 := FilterFileName(list[i].Name)
		name1 := FilterFileName(list[i+1].Name)
		if name0 != "." && name0 != ".." {
			name0 = strings.TrimPrefix(name0, ".")
		}
		if name1 != "." && name1 != ".." {
			name1 = strings.TrimPrefix(name1, ".")
		}
		if strings.ToLower(name0) > strings.ToLower(name1) {
			list[i], list[i+1] = list[i+1], list[i]
			i = 0
		} else {
			i++
		}
	}
}
