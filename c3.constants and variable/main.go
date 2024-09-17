package main

import "fmt"

func main() {
	const hello string = "hi" // this is constants
	name := "nico" // auto type check == type inference
	name = "lynn"
	fmt.Println(name)
}
