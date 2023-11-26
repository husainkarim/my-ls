package pkg

import (
	"strings"
)

func SelectColor(mode string) string {
	mode = strings.ToLower(mode)
	if len(mode) == 11 {
		mode = mode[1:]
	}
	if strings.HasPrefix(mode, "d") { // blue
		return "\033[1;34m"
	}
	if strings.HasPrefix(mode, "l") { // cyan
		return "\033[1;36m"
	}
	if strings.HasPrefix(mode, "c") { // yellow
		return "\033[1;33m"
	}
	if mode == "-rwxr-xr-x" { // Green
		return "\033[1;32m"
	}
	if strings.HasPrefix(mode, "u") { // white with red highlight
		return "\x1b[37;41m"
	}
	if strings.HasPrefix(mode, "g") { // black with yellow highlight
		return "\x1b[30;43m"
	}
	return "\033[0m"
}
