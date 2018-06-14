package main

import (
	"flag"
	"fmt"
	"os"

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

	pd := pagerduty.NewPagerDutyClient(*token)
	incidents, err := pd.GetIncidents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output
	for _, i := range incidents {
		fmt.Printf("[%d] %s, %s\n", i.IncidentNumber, i.Title, i.Status)
	}

}
