package main

import (
	"fmt"

	"fortio.org/version"

	"github.com/fortio/semtest/v2/tlib"
)

func main() {
	_, _, fullV := version.FromBuildInfo()
	fmt.Print(fullV)
	tlib.DoSomething()
}
