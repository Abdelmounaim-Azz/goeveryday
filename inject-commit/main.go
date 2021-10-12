package main

import (
	"flag"
	"fmt"
)

// GitCommit is set at build-time
var GitCommit string

func main() {
	var version bool
	flag.BoolVar(&version, "version", false, "Print the version")
	flag.Parse()
	if version {
		fmt.Printf("Version: %s\n", GitCommit)
	}
	fmt.Printf("Hello user\n")
}
