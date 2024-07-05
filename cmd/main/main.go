package main

import (
	"GoScan/internal/scanner"
	"GoScan/internal/utils"
	"flag"
	"fmt"
)

func main() {
	var urlFlag = flag.String("u", "", "URL in format (e.g., http://192.168.68.74:8080)")
	var statusFlag = flag.Int("s", 200, "HTTP status code (e.g., 200 for OK)")
	var wordlistFlag = flag.String("w", "", "Path to the wordlist file")
	var threadsFlag = flag.Int("t", 5, "Number of concurrent threads")
	var outputFlag = flag.String("o", "./results.txt", "Path for saving the results file (e.g., results.txt)")

	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("No URL flag provided!")
		return
	}

	if *wordlistFlag == "" {
		fmt.Println("No wordlist provided!")
		return
	}

	utils.PrintHeaders()
	scanner.ScanForDirectories(*urlFlag, *wordlistFlag, *statusFlag, *threadsFlag, *outputFlag)
}
