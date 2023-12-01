package main

import (
	"my-ls/pkg"
	"os"
	"sort"
	"testing"
)

var CUR string
var LIST []pkg.FileInfo

func init() {
	var temp pkg.FileInfo
	CUR, _ = os.Getwd()
	LIST = pkg.GetFilesInfo(CUR, temp)
}

func TestAddDir(t *testing.T) {
	a, b := pkg.AddDir(LIST, CUR, false)
	if len(a) != len(b) {
		t.Error("Error: error in generate directory list of the current directory!")
	}
}

func TestAddMainDir(t *testing.T) {
	_, err := pkg.AddMainDir(CUR, ".")
	if err != nil {
		t.Error("Error: error in generate file information of the current directory!")
	}
}

func TestCheckFileNameDir(t *testing.T) {
	file := "main.go"
	_, err := pkg.CheckFileNameDir(file, file, CUR)
	if err != nil {
		t.Error("Error: error in checking the directory name!")
	}
}

func TestCheckFlag(t *testing.T) {
	if !pkg.CheckFlag("tRlair") {
		t.Error("Error: error in checking the flag given!")
	}
}

func TestFillInfo(t *testing.T) {
	maindir, _ := os.Open(CUR)
	mainfile, _ := maindir.Stat()
	_, err := pkg.FillInfo(mainfile, "")
	if err != nil {
		t.Error("Error: error in generating File information!")
	}
}

func TestGetFilesInfo(t *testing.T) {
	if len(LIST) == 0 {
		t.Error("Error: no File information generated!")
	}
}

func TestSortList(t *testing.T) {
	pkg.ReverseList(LIST)
	temp := LIST
	sort.Slice(temp, func(i, j int) bool {
		return !(temp[i].Name > temp[j].Name)
	})
	pkg.SortList(LIST)
	if LIST[0] != temp[0] {
		t.Error("Error: in sorting!")
	}
}
