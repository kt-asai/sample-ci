package main

import (
	"fmt"

	"github.com/kt-asai/sample-ci/echo"
)

func main() {
	msg := echo.Echo("shota")
	fmt.Printf("%s", msg)
}
