package main

import "fmt"

var x int // global

func main() {
	x = 5
	fmt.Println(x)
	f()
	showY()
}

func f() {
	x := 10 // local
	fmt.Println(x)
}
