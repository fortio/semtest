package main

import (
	"fmt"

	"fortio.org/version"
)

func main() {
	_, _, fullV := version.FromBuildInfo()
	fmt.Print(fullV)
}
