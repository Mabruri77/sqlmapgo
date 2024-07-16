package code

import (
	"net/http"
	"strings"
	"time"
)

func makeHeaderBody(requestStringArr []string, payload string, userAgent string) (*http.Response, float64, error) {
	isHttps := map[string]string{
		"HTTP/1.1": "http://",
		"HTTP/2":   "https://",
	}
	isHeaderRequest := map[string]bool{
		"Cookie":       true,
		"Referer":      true,
		"Content-Type": true,
	}

	makeUrl := strings.Split(requestStringArr[0], " ")
	link := isHttps[makeUrl[2]] + strings.Split(requestStringArr[1], " ")[1] + makeUrl[1]

	req, err := http.NewRequest(makeUrl[0], link, nil)

	if err != nil {
		return nil, 0, err
	}
	for i := 2; i < len(requestStringArr); i++ {
		if requestStringArr[i] == "" {
			break
		}
		reqStringSplit := strings.Split(requestStringArr[i], ":")
		if isHeaderRequest[reqStringSplit[0]] {
			reqStringSplit[1] = strings.ReplaceAll(reqStringSplit[1], "*", payload)
			req.Header.Set(reqStringSplit[0], reqStringSplit[1])
		}

	}

	contentType := req.Header.Get("Content-Type")
	if contentType == "multipart/form-data" {
		GetDataMultipartAndSend()
	} else if contentType != "" {
		req.Body = GetDataBodyAndSend(requestStringArr, payload)
		defer req.Body.Close()
	}
	req.Header.Set("User-Agent", userAgent)

	req.Header.Del("Accept-Encoding")

	client := &http.Client{}
	startTime := time.Now()
	resp, err := client.Do(req)
	elapsedTime := time.Since(startTime)

	if err != nil {
		return nil, 0, err
	}
	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, &elapsedTime, nil
	// }
	// fmt.Println(string(respBody))
	return resp, elapsedTime.Seconds(), nil
}
