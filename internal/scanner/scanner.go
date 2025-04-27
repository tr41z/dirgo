package scanner

import (
	"net/http"
	"sync"
	"time"
	"fmt"
	"log"
	"os"
	
	"github.com/tr41z/dirgo/internal/utils"
	"github.com/tr41z/dirgo/pkg/config"
)

// Function responsible for scanning provided URL for directories included in provided wordlist
func Scan(baseURL string, wordlist string, statusCode int, threads int, filePath string, scanMode string) {
	delay := 10 * time.Millisecond
	payloads := utils.ConstructPayload(baseURL, wordlist, scanMode)

	var wg sync.WaitGroup
	sem := make(chan struct{}, threads)

	if err := os.Truncate(filePath, 0); err != nil {
		// Ignore the error
	}	

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
				config.Blue.Printf("%-40s%-10d%-10d\n", payload, int(resp.ContentLength), int(resp.StatusCode))

				// Creatring results file with append mode
				f, ferr := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if ferr != nil {
					log.Fatalln(ferr)
					return
				}

				// Writing results to file
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