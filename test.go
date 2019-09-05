package main

import "crypto/md5"

func main() {
	pwd := []byte("123456")
	enPwd := md5.New()
	enPwd.Sum(pwd)

}
