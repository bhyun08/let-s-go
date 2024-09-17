package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) { // go's function can have many return types 
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) { // ... is able to recive many arguments
	fmt.Println(words)
}

func main() {
	totalLenght, upperName := lenAndUpper("nico")
	// totalLenght, _ := lenAndUpper("nico") can ignore return by use _
	fmt.Println(totalLenght, upperName)
	repeatMe("nico","lynn","dal","marl","flynn")
}
