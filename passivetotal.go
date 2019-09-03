package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"os"
)

type Jsons struct {
	Subdomains []string `json:"subdomains"`
	Success bool `json:"success"`
	QueryValue string `json:"queryValue"`
	PrimaryDomain string `json:"primaryDomain"`
}

func fetchPassivetotal(domain string) ([]string, error) {
	apiEmail := os.Getenv("RISK_EMAIL")
	apiKey := os.Getenv("RISK_KEY")
	if apiKey == "" && apiEmail == ""{
		return []string{}, nil
	}
	client := &http.Client{}
	fetchURL := fmt.Sprintf("https://api.passivetotal.org/v2/enrichment/subdomains?query=%s", domain)
	req, err := http.NewRequest("GET", fetchURL, nil)
	req.SetBasicAuth(apiEmail, apiKey)
	if err != nil {
		return []string{}, nil
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	s := Jsons{}
	var j []string
	json.Unmarshal([]byte(data), &s)
	fmt.Println(s.Subdomains)
	for _,c := range s.Subdomains{
		d := c + "." + domain
		j = append(j,d)
	}
	return j, err
}
