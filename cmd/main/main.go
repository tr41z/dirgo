package main

import (
	"GoScan/internal/scanner"
	"flag"
	"fmt"
)

func main() {
	var urlFlag = flag.String("u", "", "URL in format (e.g., http://192.168.68.74:8080)")
	var statusFlag = flag.Int("s", 200, "HTTP status code (e.g., 200 for OK)")
	var wordlistFlag = flag.String("w", "", "Path to the wordlist file")
	var threadsFlag = flag.Int("t", 10, "Number of concurrent threads")

	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("No URL flag provided!")
		return
	}

	if *wordlistFlag == "" {
		fmt.Println("No wordlist provided!")
		return
	}

	scanner.ScanForDirectories(*urlFlag, *wordlistFlag, *statusFlag, *threadsFlag)
}
