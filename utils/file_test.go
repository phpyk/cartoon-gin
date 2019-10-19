package utils

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"
)

func TestGetFileBufferFromLocal(t *testing.T) {
	path := "/Volumes/mydisk/yy-finish/544/218700/1.jpg"
	buffer := GetFileBufferFromLocal(path)
	fmt.Println(hex.EncodeToString(buffer))

	buffer = buffer[2:]
	fmt.Println(hex.EncodeToString(buffer))
	var f *os.File
	var filename = "/Users/kaiyue/Downloads/go.data"
	f, err := os.Create(filename)
	CheckError(err)
	_, err = f.Write(buffer)
	CheckError(err)

}
