package eroticgo

import (
	"fmt"
	"strings"
)

const (
	BLACK   = COLOR("\033[1;30m%s\033[0m")
	RED     = COLOR("\033[1;31m%s\033[0m")
	GREEN   = COLOR("\033[1;32m%s\033[0m")
	YELLOW  = COLOR("\033[1;33m%s\033[0m")
	BLUE    = COLOR("\033[1;34m%s\033[0m")
	MAGENTA = COLOR("\033[1;35m%s\033[0m")
	CYAN    = COLOR("\033[1;36m%s\033[0m")
	WHITE   = COLOR("\033[1;37m%s\033[0m")
)

type COLOR string

func (R COLOR) Sprint(args ...interface{}) string {
	var messages = make([]string, 0, len(args))
	for _, v := range args {
		messages = append(messages, fmt.Sprint(v))
	}
	msg := strings.Join(messages, " ")
	sps := strings.Split(msg, "\n")

	newLines := make([]string, 0, len(sps))
	for _, sub := range sps {
		newLines = append(newLines, fmt.Sprintf(string(R), sub))
	}
	res := strings.Join(newLines, "\n")
	return res
}
