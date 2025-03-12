package payload

import (
	"bufio"
	"fmt"
	"os"
)

func ReadPayloadsFromFile() (map[string]string, error) {
	// Open the file
	namePayloadArr := []string{"mysql", "postgresql", "oracle", "mssql"}
	payloads := make(map[string]string)

	for _, nameFile := range namePayloadArr {
		openNameFile := nameFile + ".txt"
		file, err := os.Open(openNameFile)
		if err != nil {
			return nil, err
		}
		count := 1
		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			payload := scanner.Text()
			payloads[fmt.Sprintf("%s-%d", nameFile, count)] = payload
			count++
		}
		file.Close()
	}
	return payloads, nil
}
