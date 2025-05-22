package eroticgo

import (
	"fmt"
	"strings"
)

const (
	BLACK   = COLOR("\033[1;30m%s\033[0m") // Bold Black // 粗体黑色
	RED     = COLOR("\033[1;31m%s\033[0m") // Bold Red // 粗体红色
	GREEN   = COLOR("\033[1;32m%s\033[0m") // Bold Green // 粗体绿色
	YELLOW  = COLOR("\033[1;33m%s\033[0m") // Bold Yellow // 粗体黄色
	BLUE    = COLOR("\033[1;34m%s\033[0m") // Bold Blue // 粗体蓝色
	MAGENTA = COLOR("\033[1;35m%s\033[0m") // Bold Magenta // 粗体洋红
	CYAN    = COLOR("\033[1;36m%s\033[0m") // Bold Cyan // 粗体青色
	WHITE   = COLOR("\033[1;37m%s\033[0m") // Bold White // 粗体白色

	GRAY    = COLOR("\033[1;90m%s\033[0m") // Bold Bright Black // 粗体亮黑色 (灰色)
	SCARLET = COLOR("\033[1;91m%s\033[0m") // Bold Bright Red // 粗体鲜红色
	LIME    = COLOR("\033[1;92m%s\033[0m") // Bold Bright Green // 粗体柠檬绿色
	AMBER   = COLOR("\033[1;93m%s\033[0m") // Bold Bright Yellow // 粗体琥珀色
	AZURE   = COLOR("\033[1;94m%s\033[0m") // Bold Bright Blue // 粗体天蓝色
	PINK    = COLOR("\033[1;95m%s\033[0m") // Bold Bright Magenta // 粗体粉红色
	AQUA    = COLOR("\033[1;96m%s\033[0m") // Bold Bright Cyan // 粗体水青色
	IVORY   = COLOR("\033[1;97m%s\033[0m") // Bold Bright White // 粗体雪白
)

type COLOR string

func NewCOLOR(format string) COLOR {
	return COLOR(format)
}

func (co COLOR) Sprint(args ...interface{}) string {
	var messages = make([]string, 0, len(args))
	for _, v := range args {
		messages = append(messages, fmt.Sprint(v))
	}
	msg := strings.Join(messages, " ")
	sps := strings.Split(msg, "\n")

	newLines := make([]string, 0, len(sps))
	for _, sub := range sps {
		newLines = append(newLines, fmt.Sprintf(string(co), sub))
	}
	res := strings.Join(newLines, "\n")
	return res
}

func (co COLOR) ShowMessage(msgs ...any) {
	fmt.Println(co.Sprint("----------------------------------------"))
	fmt.Println(co.Sprint("----------------------------------------"))
	fmt.Println(co.Sprint(msgs...))
	fmt.Println(co.Sprint("----------------------------------------"))
	fmt.Println(co.Sprint("----------------------------------------"))
}
