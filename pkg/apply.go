package pkg

func Apply(list []FileInfo, dir, root, file string, time, reverse, hidden, index, longprint, recursive bool) {
	// SortList(list)
	if time {
		SortListByTime(list)
	}
	if reverse {
		ReverseList(list)
	}
	if !longprint && !recursive { // normal print
		PrintList(list, file, hidden, index)
	}
	if longprint && !recursive { // long print
		LongListFormat(list, file, hidden, index)
	}
	if !longprint && recursive { // normal print + recursive
		if file != "" {
			PrintList(list, file, hidden, index)
		} else {
			RecursiveList(list, dir, root, file, time, reverse, hidden, index, longprint, recursive, PrintList)
		}
	}
	if longprint && recursive { // long print + recursive
		if file != "" {
			LongListFormat(list, file, hidden, index)
		} else {
			RecursiveList(list, dir, root, file, time, reverse, hidden, index, longprint, recursive, LongListFormat)
		}
	}
}
