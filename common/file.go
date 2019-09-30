package common

import (
	"bytes"
	"io"
	"os"
)

func CreateFile(filename string, body []byte) {
	out, err := os.Create(filename)
	CheckError(err)
	_, e := io.Copy(out, bytes.NewReader(body))
	CheckError(e)
}

func GetFileBufferFromLocal(localPath string) ([]byte) {
	file,err := os.Open(localPath)
	CheckError(err)
	defer file.Close()

	fileinfo,err := file.Stat()
	CheckError(err)

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_,err = file.Read(buffer)
	CheckError(err)
	return buffer
}