package main

import "fmt"

func main() {
	w := func() {
		fmt.Println("Je suis une fonction anonyme sans return")
	}

	w()

	z := func() string {
		return "John"
	}()

	fmt.Println(z)
}
