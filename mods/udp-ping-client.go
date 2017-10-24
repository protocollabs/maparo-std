package mods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type configClient struct {
	Port   int32  `json:"port"`
	Server string `json:"server"`
}

func NewConfigClient() configClient {
	return configClient{
		Port:   6666,
		Server: "::1",
	}
}

type modUdpPingClient struct {
	mode         string
	cli_args     map[string]string
	configClient configClient
}

// Default Values
func NewModUdpPingClient() modUdpPingClient {
	return modUdpPingClient{
		cli_args:     make(map[string]string),
		configClient: NewConfigClient(),
	}
}

func (m modUdpPingClient) handleConfigFile(filename string) error {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "config file is not readable %s\n", filename)
		}
		return err
	}

	var c []configClient
	json.Unmarshal(raw, &c)

	return nil
}

// first, check if a config file is given, if so parse and overwrite defaults,
// later, check for arguments, if available and valid, overwrite defaults and
// config. Arguments has the highest precedence
func (m modUdpPingClient) Init() error {
	// check if config file is available
	if filename, ok := m.cli_args["config"]; ok {
		err := m.handleConfigFile(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m modUdpPingClient) Start() error {
	return nil
}

func valid_option(s string) bool {
	return strings.Contains(s, "=")
}

func is_valid_mode(s string) bool {
	if s == "server" || s == "client" {
		return true
	}
	return false
}

// parse simple parse command line and safe for later use
// parse do not overwrite the defaults, this is done later
// in init, where sanity checks are done too
func (m modUdpPingClient) Parse() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("moaare args required")
	}

	// must be server or client
	mode := os.Args[2]
	if !is_valid_mode(mode) {
		return fmt.Errorf("not a valid mode: %s (server or client)", mode)
	}
	m.mode = mode

	// parse arguments
	for _, word := range os.Args[3:] {
		ok := valid_option(word)
		if !ok {
			return fmt.Errorf("not a valid option: %q, must be key=val", word)
		}
		kv := strings.Split(word, "=")
		m.cli_args[kv[0]] = kv[1]
	}

	return nil
}
