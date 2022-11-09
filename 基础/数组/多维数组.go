package main

import (
	"fmt"
)

var arr0 [5][3]int
var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(arr0, arr1)
	fmt.Println(a, b)
	//输出结果：
	//
	//    [[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]] [[1 2 3] [7 8 9]]
	//    [[1 2 3] [4 5 6]] [[1 1] [2 2] [3 3]]

}
