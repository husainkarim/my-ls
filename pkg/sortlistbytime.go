package pkg

// sort the list by time
func SortListByTime(list []FileInfo) {
	i := 0
	for len(list)-1 > i {
		if list[i+1].ModTime.After(list[i].ModTime) {
			list[i], list[i+1] = list[i+1], list[i]
			i = 0
		} else {
			i++
		}
	}
}
