package scanner

import (
	"GoScan/internal/utils"
	"GoScan/pkg/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// Function responsible for scanning provided URL for directories included in provided wordlist
func ScanForDirectories(baseURL string, wordlist string, statusCode int, threads int, filePath string) {
	delay := 10 * time.Millisecond
	payloads := utils.ConstructPayload(baseURL, wordlist)

	var wg sync.WaitGroup
	sem := make(chan struct{}, threads)

	for _, payload := range payloads {
		wg.Add(1)
		go func(payload string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <- sem }()
			
			resp, err := http.Get(payload)
			if err != nil {
				fmt.Printf("Request to %s failed: %v\n", payload, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == statusCode && statusCode != 0 {
				config.Blue.Printf("%-55s%-10d%-10d\n", payload, int(resp.ContentLength), int(resp.StatusCode))

				// creatring results file with append mode
				f, ferr := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if ferr != nil {
					log.Fatalln(ferr)
					return
				}

				// writing results to file
				_, werr := f.WriteString(payload +"\n")
				if werr != nil {
					log.Fatalln(werr)
					f.Close()
					return
				}
			}

			time.Sleep(delay)
		}(payload)
	}

	wg.Wait()
}
