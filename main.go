package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello, I'm Golab")

	args := os.Args
	fmt.Println("Args List:")
	for index, arg := range args {
		if index != 0 {
			fmt.Println(arg)
		}
	}
}
