package scanner

import (
	"GoScan/internal/utils"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Function responsible for scanning provided URL for directories included in provided wordlist
func ScanForDirectories(baseURL string, wordlist string, statusCode int, threads int) {
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

			if resp.StatusCode == statusCode {
				fmt.Println(payload)
				fmt.Println(resp.StatusCode)
			}

			time.Sleep(delay)
		}(payload)
	}

	wg.Wait()
}
