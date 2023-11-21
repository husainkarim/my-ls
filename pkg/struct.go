package pkg

import (
	"io/fs"
	"time"
)

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
}
