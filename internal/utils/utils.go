package utils

import (
	"fmt"
	"net/http"
)

func ScanForDirectories(baseURL string, wordlist string) {
	payload := constructPayload(baseURL, wordlist)

	resp, err := http.Get(payload)

	if err != nil {
		return
	}

	fmt.Println(resp.StatusCode)
}

func constructPayload(baseURL string, wordlist string) string {
	var payload string

	for _, word := range wordlist {
		payload = baseURL + "/" + string(word)
	}

	return payload
}