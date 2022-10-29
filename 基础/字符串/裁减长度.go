package main

import (
	"strings"
	"unicode/utf8"
)

func cutString(content string, maxLength int) string {
	length := utf8.RuneCountInString(content)
	if length > maxLength {
		content = mbSubstr(content, 0, maxLength)
	}
	return strings.TrimSpace(content)
}

func mbSubstr(str string, s, e int) string {
	strRune := []rune(str)
	strLen := len(strRune)
	if s >= e || e > strLen {
		panic("长度有误")
	}
	return string(strRune[s:e])
}
func main() {

}
