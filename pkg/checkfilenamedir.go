package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// check all file name givin in the argument
func CheckFileNameDir(file, dir string) (bool, error) {
	// read all file in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return false, err
	}
	filename := strings.TrimSuffix(file, "/")
	for _, v := range files {
		name := v.Name()
		if filename == name { // if name == givin file name
			if v.IsDir() { // if the file is directory return true
				return true, nil
			} else { // if the file is file name return false
				return false, nil
			}
		}
	}
	// if file not found in the directory
	return false, errors.New("'" + file + "' no such file or directory")
}
