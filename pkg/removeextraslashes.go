package pkg

import "strings"

func RemoveExtraSlashes(path string) string {
	for {
		newPath := strings.ReplaceAll(path, "//", "/")
		if newPath == path {
			return newPath
		}
		path = newPath
	}
}
