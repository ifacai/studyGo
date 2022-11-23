package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//只遍历 key
	//for k := range scoreMap

	//注意： 遍历map时的元素顺序与添加键值对的顺序无关。
}
