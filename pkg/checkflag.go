package pkg

import "strings"

// check the flag givin
func CheckFlag(flag string) bool {
	flag = flag[1:]
	contain := "lRarti"
	for _, v := range flag {
		// check if the flag didn't contain any disallow flag
		if !strings.Contains(contain, string(v)) {
			return false
		}
	}
	return true
}
