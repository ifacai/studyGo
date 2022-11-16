package main

import (
	"fmt"
)

//string底层就是一个byte的数组，因此，也可以进行切片操作。

func main() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)
}
