package main

import (
	"ena/parser"
	"os"
)

func main() {
	args := os.Args[1:]
	var config_path string
	var destination string = "."
	for i, arg := range args {
		if arg == "-c" {
			config_path = args[i+1]
		}
		if arg == "-d" {
			destination = args[i+1]
		}
	}

	if config_path == "" {
		panic("Config file is not specified")
	}

	parser.Parse(config_path, destination)
}
