package common

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

const (
	AskAgain bool = true
	EndHere  bool = false
)

var ReplaceStr string

func SetUp() {
	if runtime.GOOS == "windows" {
		ReplaceStr = "\r\n"
	} else {
		ReplaceStr = "\n"
	}
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	return strings.Replace(text, ReplaceStr, "", -1)
}

func CheckInSlice(element int, list []int) bool {
	isInSlice := false
	for _, k := range list {
		if k == element {
			isInSlice = true
		}
	}
	return isInSlice
}
