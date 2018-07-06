package pagerduty

// Incidents holds a list of individual Incident items.
type Incidents struct {
	Incidents []Incident `json:"incidents"`
}

// Incident contains all the details for a specific PagerDuty incident.
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

// Assignment details who the incident is assigned to.
type Assignment struct {
}

// Acknowledgement details who acknowledged the incident.
type Acknowledgement struct {
	At string       `json:"at"`
	By Acknowledger `json:"acknowledger"`
}

// Acknowledger holds the name of the person who acknowledged the incident.
type Acknowledger struct {
	Name string `json:"summary"`
}
