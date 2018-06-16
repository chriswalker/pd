package pagerduty

type Incidents struct {
	Incidents []Incident `json:"incidents"`
}

type Incident struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	IncidentNumber int    `json:"incident_number"`
	CreatedAt      string `json:"created_at"`
	Status         string `json:"status"`
	HTMLURL        string `json:"html_url"`
	//	Assignments      []Assignment      `json:"assignments"`
	Acknowledgements []Acknowledgement `json:"acknowledgements"`
}

type Assignment struct {
}

type Acknowledgement struct {
	At string       `json:"at"`
	By Acknowledger `json:"acknowledger"`
}

type Acknowledger struct {
	Name string `json:"summary"`
}
