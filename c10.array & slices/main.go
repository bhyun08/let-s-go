package main

import "fmt"


func main() {
//	names1 := [5]string{"nico","lynn","dal"}  this is array
	names := []string{"nico","lynn","dal"} // this is slices
	names = append(names, "flynn") // append return new item
	fmt.Println(names)
}