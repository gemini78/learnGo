package main

import (
	"fmt"

	"github.com/gemini78/mymodule/footypes"
)

func main() {
	var foovar footypes.Foo
	foovar.TypeInt = 18
	foovar.TypeString = "Gemini"
	foovar.TypeBoolean = true

	fmt.Println(foovar)

}
