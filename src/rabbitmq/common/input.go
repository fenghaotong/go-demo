package common

import (
	"os"
	"strings"
)

func BodyForm(args []string) string {
	var s string
	if len(args) < 2 || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], "")
	}
	return s
}

func SeverityFrom(args []string) string  {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}