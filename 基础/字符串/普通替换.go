package main

import (
	"fmt"
	"strings"
)


func main() {
	s := "  需要替换的  字符串  "

	//如果n<0会替换所有old子串
	s = strings.Replace(s, "需要替换的字符串", "新词", -1)
	fmt.Println(s)

	//去掉前后\r\n\t 等空行
	s = strings.TrimSpace(s)
	fmt.Println(s)

	// 去掉前缀的 字符
	s = strings.TrimPrefix(s, "需要替换")
	fmt.Println(s)

	//去掉后缀尾巴的字符
	s = strings.TrimSuffix(s, "字符串")
	fmt.Println(s)
}
