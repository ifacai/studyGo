package main

import (
	"fmt"
	"io"
	"os"
)


func checkFileExist(fileOrPath string) bool {
	//true 存在 false 不存在
	_, err := os.Stat(fileOrPath)
	if err != nil {
		return false
	}
	return true
}
func writeFile(path, fileName, content string) {
	// path like "./html/market/"
	var f *os.File
	fullFileName := path + fileName
	pathExist := checkFileExist(path)
	if !pathExist {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.Create(fullFileName)
	if err != nil {
		panic(fullFileName + " 创建失败")
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	if _, err = io.WriteString(f, content); err != nil {
		panic(fullFileName + " 写入错误")
	}
}
func main() {
	writeFile("/data/", "ip.txt", "hello world")
}
