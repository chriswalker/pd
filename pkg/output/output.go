package output

import (
	"fmt"

	"github.com/chriswalker/pd/pkg/pagerduty"
	"github.com/fatih/color"
)

type StdOutputter struct{}

func NewStdOutputter() StdOutputter {
	return StdOutputter{}
}

func (s StdOutputter) Output(incidents []pagerduty.Incident) {
	for _, i := range incidents {
		fn := getColourFunc(i)
		fmt.Printf("%s\n", fn("[%d] %s, %s", i.IncidentNumber, i.Title, i.Status))
		fmt.Printf("(%s)\n", i.HTMLURL)
	}
}

func getColourFunc(incident pagerduty.Incident) func(format string, a ...interface{}) string {
	if incident.Status == "triggered" {
		return color.New(color.FgRed).SprintfFunc()
	}
	if incident.Status == "acknowledged" {
		return color.New(color.FgYellow).SprintfFunc()
	}
	return color.New(color.FgGreen).SprintfFunc()
}
