package pkg

// revrse the list
func ReverseList(list []FileInfo) {
	base := len(list) - 1
	for i := 0; i <= len(list)/2; i++ {
		list[i], list[base-i] = list[base-i], list[i]
	}
}
