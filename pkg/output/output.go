package output

import (
	"fmt"
	"os"

	"github.com/chriswalker/pd/pkg/pagerduty"
	"github.com/fatih/color"
)

var (
	colours = map[string]func(format string, a ...interface{}) string{
		"triggered":    color.New(color.FgRed, color.Bold).SprintfFunc(),
		"acknowledged": color.New(color.FgYellow, color.Bold).SprintfFunc(),
		"resolved":     color.New(color.FgGreen, color.Bold).SprintfFunc(),
	}
)

type StdOutputter struct{}

func NewStdOutputter() StdOutputter {
	return StdOutputter{}
}

func (s StdOutputter) Output(incidents []pagerduty.Incident) {
	if len(incidents) == 0 {
		fmt.Println("No overnight incidents to report")
		os.Exit(0)
	}

	fmt.Printf("PagerDuty incidents since last night\n\n")

	for _, i := range incidents {
		fn := colours[i.Status]
		fmt.Printf("  %s %d: %s\n", fn("┃ %s", i.Status), i.IncidentNumber, i.Title)
		fmt.Printf("  %s Time: %s\n", fn("┃"), i.CreatedAt)
		if len(i.Acknowledgements) > 0 {
			fmt.Printf("  %s Acknowledged by %s, at %s\n", fn("┃"), i.Acknowledgements[0].By.Name, i.Acknowledgements[0].At)
		}
		fmt.Printf("  %s %s\n\n", fn("┃"), color.BlueString(i.HTMLURL))
	}
}
