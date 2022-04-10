package main

import (
	"errors"
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Bonjour à tous, je m'appelle %v. \n", name)
}

func sayBye(name string) (string, error) {
	if name == "" {
		return "", errors.New("\bErreur: vous avez oublié de spécifier un nom")
	}
	byeMessage := fmt.Sprintf("%v s'en va! Bonne soirée.", name)
	return byeMessage, nil
}

func main() {
	fmt.Println(sayBye(""))
}
