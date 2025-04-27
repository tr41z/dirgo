package utils

import (
	"github.com/tr41z/dirgo/pkg/config"
	"bufio"
	"log"
	"os"
)

// Function responsible for creating list of payloads constructed from combination of URL and each word from wordlist
func ConstructPayload(baseURL string, wordlist string, scanMode string) []string {
	var payloads []string

	file, err := os.Open(wordlist)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if (scanMode == "dir") {
			payload := baseURL + "/" + scanner.Text()
			payloads = append(payloads, payload)
		} else if (scanMode == "sub") {
			payload := scanner.Text() + "." + baseURL
			payloads = append(payloads, payload)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return payloads
} 

func PrintHeaders() {
	config.Magenta.Printf("%-40s%-10s%-10s\n", "URL", "LENGTH", "STATUS_CODE\n-------------------------------------------------------------")
}