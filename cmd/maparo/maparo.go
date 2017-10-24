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
	Init() error
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
	m["mod-udp-ping-client"] = mods.NewModUdpPingClient()
	return m
}

func prepare_campaign_map() map[string]Mod {
	m := make(map[string]Mod)
	m["campaign-ping"] = mods.NewModUdpPingClient()
	return m
}

func usage(mm map[string]Mod, mc map[string]Mod) {
	fmt.Fprintf(os.Stderr, "invalid arg, must be mod- or campaign-\n")
}

func mode_mod(s map[string]Mod, name string) {
	mod, ok := s[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "not a valid module\n")
		os.Exit(1)
	}
	err := mod.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	err = mod.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func mode_campaign(s map[string]Mod, name string) {
	mod, ok := s[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "not a valid module\n")
		os.Exit(1)
	}
	err := mod.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Fprintf(os.Stderr, "mapago(c) - 2017\n")
	fmt.Fprintf(os.Stderr, "build version: %s\n", core.BuildVersion)
	fmt.Fprintf(os.Stderr, "build date:    %s\n", core.BuildDate)

	mm := prepare_mod_map()
	mc := prepare_campaign_map()

	name, err := parse_args()
	if err != nil {
		usage(mm, mc)
		os.Exit(1)
	}

	if strings.HasPrefix(name, "mod-") {
		mode_mod(mm, name)
	} else if strings.HasPrefix(name, "campaign-") {
		mode_campaign(mc, name)
	} else {
		usage(mm, mc)
		os.Exit(1)
	}

	os.Exit(0)
}
