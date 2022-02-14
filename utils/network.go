package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendRequest(url string, vars map[string]string, output interface{}, verbose bool) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	queryString := req.URL.Query()
	for key, value := range vars {
		queryString.Add(key, value)
	}
	req.URL.RawQuery = queryString.Encode()
	if verbose {
		fmt.Printf("URL:\n%s\n\n", req.URL)
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode < 200 || res.StatusCode > 300 {
		log.Fatal("Request not successful.")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &output)
}
