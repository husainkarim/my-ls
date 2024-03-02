package pkg

import (
	"errors"
	"os"
)

// check all file name givin in the argument
func CheckFileNameDir(file, dir string) (bool, bool, error) {
	file = RemoveExtraSlashes(file)
	files, err := os.Stat(dir + "/" + file)
	if err != nil {
		return false, false, errors.New("./my-ls: cannot access '" + file + "': no such file or directory")
	}
	if files.IsDir() {
		link, _ := os.Readlink(file)
		if link != "" {
			return false, true, nil
		}
		return true, false, nil
	}
	if files.Mode().IsRegular() {
		return false, true, nil
	}
	return false, false, errors.New("./my-ls: cannot access '" + file + "': no such file or directory")
}
