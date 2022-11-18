package main

import "fmt"

func main() {

	//当一个指针被定义后没有分配到任何变量时，它的值为 nil
	//空指针的判断
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}
