package pagerduty

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var (
	pagerDutyURL = "https://api.pagerduty.com/incidents"
)

type Client struct {
	httpClient *http.Client
	token      string
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (p *Client) GetIncidents() ([]Incident, error) {
	req, err := http.NewRequest("GET", pagerDutyURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create 'GET incidents' request")
	}
	req.Header.Add("Authorization", "Token token="+p.token)

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
