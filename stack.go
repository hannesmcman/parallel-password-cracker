package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)
const pwd1 string = "password"
const pwd2 string = "passwor"

func main() {
	var a = stack.New()
	for i := 97; i < 123; i++ {
		a.Push(string(i))
	}


	fmt.Println(pwd1 == pwd2)
	fmt.Println(pwd1 == pwd2)
	
	var length = a.Len()
	for i := 0; i<length; i++ {
		fmt.Println(a.Pop())
	}

}