package main

import (
	"flag"
	"fmt"

	"github.com/tr41z/dirgo/internal/scanner"
	"github.com/tr41z/dirgo/internal/utils"
	"github.com/tr41z/dirgo/pkg/config"
)

func main() {
	config.Green.Println("================================================")
	config.Green.Println(`                      
	.--|  ||__|.----..-----..-----.
	|  _  ||  ||   _||  _  ||  _  |
	|_____||__||__|  |___  ||_____|
			 |_____|       
		`)
	config.Magenta.Println("© 2024 tr41z")

	config.Magenta.Println("\nLanguage used: Golang")
	config.Magenta.Print("Github Repo: https://github.com/tr41z/dirgo \n")

	config.Green.Print("================================================\n\n")
	
	var scanModeFlag = flag.String("m", "", "Scan mode (e.g. `dir` for directory scanning and `sub` for subdomain scanning)")
	var urlFlag = flag.String("u", "", "URL in format (e.g., 192.168.68.74:8080)")
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
	
	scanner.Scan(*urlFlag, *wordlistFlag, *statusFlag, *threadsFlag, *outputFlag, *scanModeFlag)
}
