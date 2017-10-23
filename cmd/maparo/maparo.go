package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/protocollabs/maparo/core"
	"github.com/protocollabs/maparo/mods"
)

type Mod interface {
	Parse() error
	Init(n core.ModConf) error
	Start() error
}

// campaign has to provide the same functions as
// Modules. Simple typedef then
type Campaign Mod

func parse_args() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}

	return "", fmt.Errorf("list or count subcommand is required")
}

func prepare_mod_map() map[string]Mod {
	m := make(map[string]Mod)
	m["udp-ping"] = mods.ModUdpPing{}
	return m
}

func prepare_campaign_map() map[string]Mod {
	m := make(map[string]Mod)
	m["campaign-ping"] = mods.ModUdpPing{}
	return m
}

func usage() {
	fmt.Fprintf(os.Stderr, "invalid arg, must be mod- or campaign-\n")
}

func mode_mod(name string) {
	s := prepare_mod_map()
	mod, ok := s[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "not a valid module\n")
		os.Exit(1)
	}
	mod.Parse()
}


func mode_campaign(name string) {
	s := prepare_campaign_map()
	mod, ok := s[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "not a valid module\n")
		os.Exit(1)
	}
	mod.Parse()
}

func main() {
	fmt.Fprintf(os.Stderr, "build version: %s\n", core.BuildVersion)
	fmt.Fprintf(os.Stderr, "build date:    %s\n", core.BuildDate)

	name, err := parse_args()
	if err != nil {
		usage()
		os.Exit(1)
	}

	if strings.HasPrefix(name, "mod-") {
		mode_mod(name)
	} else if strings.HasPrefix(name, "campaign-") {
		mode_campaign(name)
	} else {
		usage()
		os.Exit(1)
	}

	os.Exit(0)


}
