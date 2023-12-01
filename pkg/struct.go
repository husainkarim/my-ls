package pkg

import (
	"io/fs"
	"time"
)

// struct to store all information need for the files
type FileInfo struct {
	Name    string
	Dir     bool
	Mode    fs.FileMode
	Nlinks  uint64
	User    string
	Group   string
	Size    int64
	ModTime time.Time
	Index   int
	Block   int
}

// struct to store the file and the path for it with the necessary information
type Dir struct {
	Name  string
	Path  string
	File  string
	IsDir bool
}
