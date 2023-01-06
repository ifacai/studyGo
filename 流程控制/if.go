package main

import "fmt"

func main() {
	var a int = 10
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else if a < 30{
		fmt.Println("...")
	} else{
		fmt.Println("...")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}
