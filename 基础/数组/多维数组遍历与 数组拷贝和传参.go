package main

import "fmt"

func main() {
	//多维数组遍历：

	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
	// 数组拷贝和传参
	var arr3 [5]int
	printArr(&arr3)
	fmt.Println(arr3)
	arr4 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr4)
	fmt.Println(arr4)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
