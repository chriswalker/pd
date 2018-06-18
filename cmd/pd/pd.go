package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chriswalker/pd/pkg/output"
	"github.com/chriswalker/pd/pkg/pagerduty"
	"github.com/chriswalker/spinner"
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

	incidents, err := getIncidents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output
	outputter := output.NewStdOutputter()
	outputter.Output(incidents)
}

func getIncidents() ([]pagerduty.Incident, error) {
	pd := pagerduty.NewClient(*token)

	spinner := spinner.NewSpinner()
	defer spinner.Stop()
	spinner.Prefix = "Getting incidents "
	spinner.Start()

	return pd.GetIncidents()
}
