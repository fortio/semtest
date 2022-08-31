package main

import (
	"fmt"

	"fortio.org/fortio/version"
)

func main() {
	_, _, fullV := version.FromBuildInfo()
	fmt.Print(fullV)
}
