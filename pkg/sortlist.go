package pkg

func SortList(list []FileInfo) {
	i := 0
	for len(list)-1 > i {
		if list[i].Name > list[i+1].Name {
			list[i], list[i+1] = list[i+1], list[i]
			i = 0
		} else {
			i++
		}
	}
}
