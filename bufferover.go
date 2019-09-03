package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Bufferover struct {
	Meta struct {
		Runtime string `json:"Runtime"`
		Errors interface{} `json:"Errors"`
		Message string `json:"Message"`
		FileNames []string `json:"FileNames"`
		TOS string `json:"TOS"`
	} `json:"Meta"`
	FDNSA []string `json:"FDNS_A"`
	RDNS []string `json:"RDNS"`
}
func fetchBufferover(domain string) ([]string,error) {
	var s []string
	fetchURL := fmt.Sprintf("https://dns.bufferover.run/dns?q=.%s", domain)
	resp, err := http.Get(fetchURL)
	if err != nil {
		return s, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	b := Bufferover{}
	json.Unmarshal([]byte(data), &b)
	for _,c := range b.FDNSA{
		d := strings.Split(c,",")
		s = append(s, d[1])
	}
	for _,j := range b.RDNS{
		h := strings.Split(j,",")
		s = append(s, h[1])
	}
   return s, nil
}



