package scanner

import (
	"fmt"
	"net/http"
	"GoScan/internal/utils"
)

// Function responsible for scanning provided url for directories included in provided wordlis
func ScanForDirectories(baseURL string, wordlist string) {
	payloads := utils.ConstructPayload(baseURL, wordlist)

	for _, payload := range(payloads) {
		resp, err := http.Get(payload)

		if err != nil {
			return
		}

		if resp.StatusCode == 200 {
			fmt.Println(payload)
			fmt.Println(resp.StatusCode)
		}
	}
}