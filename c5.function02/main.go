package main

import (
	"fmt"
	"strings"
)

func lenAndIpper(name string) (length int, uppercase string) { // naked return
	defer fmt.Println("I'm done") // when function ends, defer run
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(lenAndIpper("nico"))
}