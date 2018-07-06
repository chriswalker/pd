package pagerduty

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var (
	pagerDutyURL = "https://api.pagerduty.com/incidents"
)

// Client represents a PagerDuty client, and contains the http.Client
// used for interactions, and the user's PagerDuty token.
type Client struct {
	httpClient *http.Client
	token      string
}

// NewClient returns an initialised PagerDuty client.
func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

// GetIncidents gets a list of incidents that happened between 1700 hrs the
// previous day, and now.
func (p *Client) GetIncidents() ([]Incident, error) {
	req, err := http.NewRequest("GET", pagerDutyURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create 'GET incidents' request")
	}
	req.Header.Add("Authorization", "Token token="+p.token)
	req.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")

	t := getTimeYesterday()

	q := req.URL.Query()
	q.Add("since", t.Format(time.RFC3339))
	req.URL.RawQuery = q.Encode()

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve incidents")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	incidents := Incidents{}
	err = json.Unmarshal(body, &incidents)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal incidents")
	}

	return incidents.Incidents, nil
}

// getTimeYesterday calculates the time.Time representing 1700 hrs of the previous
// day.
func getTimeYesterday() time.Time {
	t := time.Now()

	hrs := t.Hour() + 7 // Takes number of hours from now, until 5pm (approx) previous day
	prevTime := t.Add(time.Duration(-hrs) * time.Hour)

	return prevTime
}
