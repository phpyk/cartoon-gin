package common

import (
	"github.com/dchest/captcha"
	"os"
	"time"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 6
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
	// Standard width and height of a captcha image.
	StdWidth  = 240
	StdHeight = 80
)

func GetNewCaptcha() string {
	var capid string
	capid = captcha.New()

	fname := os.Getenv("GOPATH")+"/tmp/"+capid+".png"
	f,err := os.Create(fname)
	CheckError(err)
	defer f.Close()

	err = captcha.WriteImage(f,capid,StdWidth,StdHeight)
	CheckError(err)
	return capid
}

func CheckCaptcha(capid string,value []byte) bool {
	return captcha.Verify(capid,value)
}


