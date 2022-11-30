package main

import (
	"ena/parser"
)

func main() {
	p := parser.InitParser("test.yaml")
	p.ParseFolders()
}
