package main

import (
	"fmt"
)

func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		//for x:=0; x <=50; x++
		return fib(x-1) + fib(x-2)
	}
}
func main() {
	//var x uint
	fmt.Println(12)
}
