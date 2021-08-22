package main

import "fmt"

// global var
// var name string = "Brad"

func main() {
	// using var
	// var name string = "Brad"
	var age int32 = 37
	var isCool = true
	var size float64 = 1.3

	// shorthand
	// name := "HaiXin"
	// email := "Haixin.zhang.777@gmail.com"
	name, email := "HaiXin", "Haixin.zhang.777@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)

}
