package utils

import (
	"GoScan/pkg/config"
	"bufio"
	"fmt"
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

func CalculateLeft(wordlist string) {
	count := 0
	file, err := os.Open(wordlist)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
		fmt.Println(count)
	}	
}

func PrintHeaders() {
	config.Magenta.Printf("%-55s%-10s%-10s\n", "URL", "LENGTH", "STATUS_CODE\n-----------------------------------------------------------------------------")
}