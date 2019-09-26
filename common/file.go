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
