package main

import "fmt"

//类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。

//类型定义

type NewInt int

//类型别名

type MyInt = int

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //  a main.NewInt
	fmt.Printf("type of b:%T\n", b) //  b int
}
