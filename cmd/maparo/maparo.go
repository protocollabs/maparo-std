package main

import (
	"fmt"
	"os"

	"github.com/protocollabs/maparo/core"
	"github.com/protocollabs/maparo/mods"
)


type Mod interface {
	Parse() (error)
	Init(n core.ModConf) (error)
	Start() (error)
}

func parse_args() (string, error) {
    if len(os.Args) > 1 {
		return os.Args[1], nil
	}

	return "", fmt.Errorf("list or count subcommand is required")
}


func main() {
	fmt.Fprintf(os.Stderr, "build version: %s\n", core.BuildVersion)
	fmt.Fprintf(os.Stderr, "build date:    %s\n", core.BuildDate)

	mod_name, err := parse_args()
	if err != nil {
		fmt.Fprintf(os.Stderr, "no valid mod gievn\n")
		os.Exit(1)
	}

	s := make(map[string]Mod)
	s["udp-ping"] = mods.ModUdpPing{}

	mod, ok := s[mod_name]
	if !ok {
		fmt.Fprintf(os.Stderr, "not a valid module\n")
		os.Exit(1)
	}
	mod.Parse()

}
