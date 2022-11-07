package main

import (
	"fmt"
	"strings"
)

func main() {

	//将一系列字符串连接为一个字符串，之间用sep来分隔。
	s := []string{"我", "是", "谁", "呀"}
	fmt.Println(strings.Join(s, "///"))

	//使用 fmt.Sprintf
	host := "127.0.0.1"
	port := "27017"
	dbName := "dbName"
	uri := fmt.Sprintf("mongodb://%s:%s/%s", host, port, dbName)
	fmt.Println(uri)

}
