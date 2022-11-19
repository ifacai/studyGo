package main

import "fmt"

func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)

	//new是一个内置的函数，它的函数签名如下：
	//    func new(Type) *Type
	//1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
	//2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
	c := new(int)
	d := new(bool)
	fmt.Printf("%T\n", c) // *int
	fmt.Printf("%T\n", d) // *bool
	fmt.Println(*c)       // 0
	fmt.Println(*d)       // false

	var e *int
	e = new(int)
	*e = 10
	fmt.Println(*e)

	//make
	//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：
	var f map[string]int
	f = make(map[string]int, 10)
	f["测试"] = 100
	fmt.Println(f)

	//new与make的区别
	//1.二者都是用来做内存分配的。
	//2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
