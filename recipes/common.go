package common

import "runtime"

var ReplaceStr string

func SetUp() {
	if runtime.GOOS == "windows" {
		ReplaceStr = "\r\n"
	} else {
		ReplaceStr = "\n"
	}
}
