package utils

import (
	"os"
	"time"

	"github.com/dchest/captcha"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 4
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
	// Standard width and height of a captcha image.
	StdWidth  = 240
	StdHeight = 80
)

var (
	GlobalStore = captcha.NewMemoryStore(CollectNum, Expiration)
)

func GetNewCaptcha() (capid string, imgData string) {
	captcha.SetCustomStore(GlobalStore)
	capid = captcha.NewLen(DefaultLen)

	fname := os.Getenv("GOPATH") + "/tmp/" + capid + ".png"
	f, err := os.Create(fname)
	CheckError(err)
	defer f.Close()

	err = captcha.WriteImage(f, capid, StdWidth, StdHeight)
	CheckError(err)

	return capid, ImageBase64String(fname)
}

func GetCaptchaVal(id string) []byte {
	captcha.SetCustomStore(GlobalStore)
	return GlobalStore.Get(id, false)
}

func CheckCaptcha(capid string, value []byte) bool {
	captcha.SetCustomStore(GlobalStore)
	return captcha.Verify(capid, value)
}
