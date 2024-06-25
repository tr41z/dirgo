package main

import (
	"GoScan/internal/scanner"
	"flag"
	"fmt"
)

func main() {
	var urlFlag = flag.String("u", "", "URL in format <http(s)://IP:PORT>")
	var statusFlag = flag.Int("s", 0, "HTTP status code (e.g., 200 for OK)")

	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("No URL flag provided!")
		return
	}

	if *statusFlag == 0 {
		fmt.Println("No status code provided, using default: 200")
		*statusFlag = 200
	} 

	scanner.ScanForDirectories("http://127.0.0.1:8080", "../../wordlists/common.txt")
}