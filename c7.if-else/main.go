package main

import (
	"fmt"
)

func canIdrink(age int) bool {

	if 	koreanAge := age + 2; koreanAge < 18 { // only use if block variable
		return false
	}
	return true
	
}

func main() {
	fmt.Println((canIdrink(16)))
}