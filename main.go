package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Entry point")
	listArgs()
}

func HelloWorld() string {
	return "Hello World!"
}

func listArgs() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
