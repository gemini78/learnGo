package main

import "fmt"

type Animal interface {
	noise() string
	numberOfLegs() int
}

type Cat struct {
	name  string
	breed string
}

type Spider struct {
	name     string
	breed    string
	Venomous bool
}

func main() {
	cat := Cat{
		name:  "Kitty",
		breed: "Siamois",
	}
	printAnimalInfo(&cat)

	spider := Spider{
		name:     "Spiddy",
		breed:    "Veuve noire",
		Venomous: true,
	}

	printAnimalInfo(&spider)
}

func printAnimalInfo(a Animal) {
	fmt.Println("Le bruit de cet animal est", a.noise(), " et il possede ", a.numberOfLegs(), "pattes !")
}

func (c *Cat) noise() string {
	return "Miaou"
}

func (c *Cat) numberOfLegs() int {
	return 4
}

func (s *Spider) noise() string {
	return "Hiss"
}

func (s *Spider) numberOfLegs() int {
	return 8
}
