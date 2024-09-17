package main

import "fmt"

func main() {
	a := 2
	b := &a // b refer to a's memory adress
	*b = 202020 // b access a's memory
	fmt.Println(a)
}