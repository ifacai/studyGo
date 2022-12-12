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
	fmt.Printf("p1=%v\n", p1)  
	fmt.Printf("p1=%#v\n", p1) 
}
