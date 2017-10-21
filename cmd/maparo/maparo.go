package main

import (
	"fmt"

	"github.com/protocollabs/maparo/core"
)

func main() {
	fmt.Printf("maparo\n")
	fmt.Printf("build version: %#v\n", core.BuildVersion)
	fmt.Printf("build date: %#v\n", core.BuildDate)
}
