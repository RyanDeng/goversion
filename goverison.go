package goversion

import (
	"flag"
	"fmt"
	"os"
)

var version string = "dev"

func init() {
	b := flag.Bool("version", false, "application verison")
	flag.Parse()
	if *b {
		fmt.Println(version)
		os.Exit(0)
	}
}

func Version() string {
	return version
}
