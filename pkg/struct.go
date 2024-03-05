package pkg

import (
	"io/fs"
	"time"
)

// struct to store all information need for the files

var (
	MaxWidth int = 200
	S_link   int
	S_user   int
	S_group  int
	S_size   int
	S_name   int
	S_index  int
)

type FileInfo struct {
	Name    string
	Dir     bool
	Mode    fs.FileMode
	Nlinks  int64
	User    string
	Group   string
	Size    int64
	ModTime time.Time
	Index   int
	Block   int
	Link    string
}

// struct to store the file and the path for it with the necessary information
type Dir struct {
	Name  string
	Path  string
	File  string
	IsDir bool
}
