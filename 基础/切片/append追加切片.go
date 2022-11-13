package main

import (
	"fmt"
)

func main() {

	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)
	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)

	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s1, s2)

	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]

	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。

}
