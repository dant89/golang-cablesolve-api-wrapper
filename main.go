package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dant89/cablesolve-api-client-golang/gosoap"
	"log"
	"net/http"
	"time"
)

// GetIPLocationResponse will hold the Soap response
type SearchResponse struct {
	SearchResult string `xml:"SearchResult"`
}

// SearchResult will
type SearchResult struct {
}

var (
	r SearchResponse
)

func main() {
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}

	endpoint := "ComponentSearchServices"
	url := "https://username:password@cablesolve/cswebapi/" + endpoint + ".asmx?WSDL"

	fmt.Println(url)

	soap, err := gosoap.SoapClient(url, httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	// Use gosoap.ArrayParams to support fixed position params
	params := gosoap.Params{}

	res, err := soap.Call("Search", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

	res.Unmarshal(&r)

	// GetIpLocationResult will be a string. We need to parse it to XML
	result := SearchResult{}
	err = xml.Unmarshal([]byte(r.SearchResult), &result)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	log.Println(result)
}
