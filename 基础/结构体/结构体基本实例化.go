package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "小名"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"小名", city:"北京", age:18}
}
