package main

import "fmt"

func main() {
	//使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
	//    delete(map, key)
	//map:表示要删除键值对的map
	//key:表示要删除的键值对的键
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	delete(scoreMap, "小明") //将小明:100从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
