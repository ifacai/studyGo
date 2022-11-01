package main

import (
	"os"
)

func checkFileExist(fileFullPathAndName string) bool {
	//true 存在 false 不存在
	_, err := os.Stat(fileFullPathAndName)
	if err != nil {
		return false
	}
	return true
}
func main() {
	checkFileExist("/home/goProject/src/news/web/sitemapFiles/2.xml")
}
