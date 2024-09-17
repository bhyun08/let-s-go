package main

import (
	"fmt"
)

func superAdd(numbers ...int) int{
	total := 0
	for _, number := range numbers { // index ignore
		total += number
	}

	return total
}

func main() {
	total := superAdd(1,2,3,3,44,5,6,7) // array
	fmt.Println("total is : ",total)
}