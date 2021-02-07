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
