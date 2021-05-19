package common

import (
	"os"
	"strings"
)

func BodyForm(args []string) string {
	var s string
	if (len(args) < 2 || os.Args[1] == ""){
		s = "hello"
	} else {
		s = strings.Join(args[1:], "")
	}
	return s
}
