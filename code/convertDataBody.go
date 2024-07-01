package code

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func GetDataBodyAndSend(requestStringArr []string, payload string) io.ReadCloser {
	getBody := requestStringArr[len(requestStringArr)-1]
	getBody = strings.ReplaceAll(getBody, "*", payload)

	dataBody := []byte(getBody)
	return io.NopCloser(bytes.NewBuffer(dataBody))
}

func GetDataMultipartAndSend() {
	fmt.Print("this is data multipart")
}
