package main

import (
	"flag"
	"fmt"
)

var (
	host = flag.String("host", "http://localhost:8000", "http host of the server")
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("A command is required")
		flag.Usage()
		return
	}

	switch flag.Arg(0) {
	case "auth":
		authCommands(flag.Args()[1:])
	case "todos":
		todoCommands(flag.Args()[1:])
	default:
		fmt.Println("Invalid command:", flag.Arg(0))
		flag.Usage()
	}
}
