package goversion

import (
	"fmt"
	"os"
)

var version string = "dev"

func init() {
	if len(os.Args) == 2 && os.Args[1] == "-version" {
		fmt.Println(version)
		os.Exit(0)
	}
}

func Version() string {

	return version
}
