package pkg

import "strings"

// filter to get the name of the file only without the expansion
func FilterFileName(name string) string {
	if name == "." || name == ".." {
		return name
	}
	if strings.HasPrefix(name, ".") {
		return "." + FilterFileName(strings.TrimPrefix(name, "."))
	}
	list := strings.Split(name, ".")
	return list[0]
}
