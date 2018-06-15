package output

import (
	"fmt"

	"github.com/chriswalker/pd/pkg/pagerduty"
)

type StdOutputter struct{}

func NewStdOutputter() StdOutputter {
	return StdOutputter{}
}

func (s StdOutputter) Output(incidents []pagerduty.Incident) {
	fmt.Println("number of incidents =", len(incidents))
	for _, i := range incidents {
		fmt.Printf("[%d] %s, %s\n", i.IncidentNumber, i.Title, i.Status)
	}
}
