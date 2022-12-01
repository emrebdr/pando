package main

import (
	"ena/parser"
	"fmt"
)

func main() {
	parser.Parse("test1.yaml")
	fmt.Println("Done")
}
