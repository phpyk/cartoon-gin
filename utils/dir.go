package utils

import (
	"os"
)

func CreateDirIfNotExists(dir string) bool {
	if !DirExists(dir) {
		DirCreate(dir)
	}
	return true
}

func DirExists(dir string) bool {
	d, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return d.IsDir()
}

func DirCreate(dir string) {
	e := os.Mkdir(dir, os.ModePerm)
	CheckError(e)
}
