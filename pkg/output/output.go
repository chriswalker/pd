package output

import (
	"fmt"

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
	fmt.Printf("%v", incidents)
	for _, i := range incidents {
		fn := colours[i.Status]
		fmt.Printf("%s %d: %s\n", fn("┃ %s", i.Status), i.IncidentNumber, i.Title)
		if len(i.Acknowledgements) > 0 {
			fmt.Printf("%s Acknowledged by %s, at %s", fn("┃"), i.Acknowledgements[0].By.Name, i.Acknowledgements[0].At)
		}
		fmt.Printf("%s %s\n\n", fn("┃"), color.BlueString(i.HTMLURL))
	}
}
