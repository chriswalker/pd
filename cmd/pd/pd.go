package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chriswalker/pd/pkg/output"
	"github.com/chriswalker/pd/pkg/pagerduty"
)

var (
	token = flag.String("token", "", "(required) PagerDuty auth token")
)

func main() {
	flag.Parse()
	if *token == "" {
		fmt.Println("PagerDuty token required")
		os.Exit(1)
	}

	pd := pagerduty.NewClient(*token)
	incidents, err := pd.GetIncidents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output
	outputter := output.NewStdOutputter()
	outputter.Output(incidents)
}
