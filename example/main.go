// run cmd:
// go build -ldflags "-X github.com/ryandeng/goversion.version=1.1" main.go
// ./main.go -version
package main

import (
	"fmt"
	"github.com/ryandeng/goversion"
)

func main() {
	fmt.Println("main's version is:", goversion.Version())
}
