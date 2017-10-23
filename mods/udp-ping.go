package mods

import "fmt"
import "flag"
import "os"
import "github.com/protocollabs/maparo/core"

type ModUdpPing struct {
	conf core.ModConf
}

func (m ModUdpPing) Init(conf core.ModConf) error {
	// save configuration
	m.conf = conf
	return nil
}

func (m ModUdpPing) Start() error {
	return nil
}

func (m ModUdpPing) Parse() error {
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	listTextPtr := listCommand.String("text", "", "Text to parse. (Required)")
	listMetricPtr := listCommand.String("metric", "chars", "Metric <chars|words|lines>. (Required)")
	listUniquePtr := listCommand.Bool("unique", false, "Measure unique values of a metric.")

	listCommand.Parse(os.Args[2:])

	if listCommand.Parsed() {
		if *listTextPtr == "" {
			fmt.Println("NO ARG GIBVEN")
			listCommand.PrintDefaults()
			return fmt.Errorf("FIXME")
		}

		metricChoices := map[string]bool{"chars": true, "words": true, "lines": true}
		if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
			listCommand.PrintDefaults()
			return fmt.Errorf("FIXME")
		}

		fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *listTextPtr, *listMetricPtr, *listUniquePtr)
	}

	return nil
}
