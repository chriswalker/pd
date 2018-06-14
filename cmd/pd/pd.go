package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	token        = flag.String("token", "", "(required) PagerDuty auth token")
	pagerDutyURL = "https://api.pagerduty.com/incidents"
)

func main() {
	flag.Parse()
	if *token == "" {
		fmt.Println("PagerDuty token required")
		os.Exit(1)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", pagerDutyURL, nil)
	if err != nil {
		// TODO
	}
	req.Header.Add("Authorization", "Token token="+*token)

	resp, err := client.Do(req)
	if err != nil {
		// TODO
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
