package utils

import (
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha() (string, string, error) {
	driver := &base64Captcha.DriverString{
		Length:          4,
		Height:          40,
		Width:           80,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		NoiseCount:      0,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
	}

	c := base64Captcha.NewCaptcha(driver, store)
	return c.Generate()
}

func VerifyCaptcha(id, VerifyValue string) bool {
	return store.Verify(id, VerifyValue, true)
}
