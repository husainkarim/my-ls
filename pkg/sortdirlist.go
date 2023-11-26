package pkg

// sort the directory list to print the files first then the directory
func SortDirList(list []Dir) []Dir {
	var new []Dir
	for _, v := range list {
		if !v.IsDir {
			new = append(new, v)
		}
	}
	for _, v := range list {
		if v.IsDir {
			new = append(new, v)
		}
	}
	return new
}
