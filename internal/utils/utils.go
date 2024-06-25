package utils

import (
	"bufio"
	"log"
	"os"
)

// Function responsible for creating list of payloads constructed from combination of URL and each word from wordlist
func ConstructPayload(baseURL string, wordlist string) []string {
	var payloads []string

	file, err := os.Open(wordlist)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payload := baseURL + "/" + scanner.Text()
		payloads = append(payloads, payload)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return payloads
}
