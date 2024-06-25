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
 
    Scanner := bufio.NewScanner(file)
    Scanner.Split(bufio.ScanWords)
 
    for Scanner.Scan() {
        payload := baseURL + "/" + string(Scanner.Text())
		payloads = append(payloads, payload)
    }
    if err := Scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return payloads
}