package payload

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Data struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

func ConvertTextToJson(nameFile string) {
	file, err := os.Open(nameFile + ".txt")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := []Data{}
	for scanner.Scan() {
		payload := scanner.Text()
		arr = append(arr, Data{Name: nameFile, Payload: payload})
	}
	CreateFile(arr, "payload/"+nameFile)
}
func Create() {
	mysql := Read("payload/mysql")
	mssql := Read("payload/mssql")
	oracle := Read("payload/oracle")
	postgres := Read("payload/postgresql")
	arr := []Data{}
	i, j, k, l := 0, 0, 0, 0
	for i < len(mysql) || j < len(mssql) || k < len(oracle) || l < len(postgres) {
		if i < len(mysql) {
			arr = append(arr, mysql[i])
			i++
		}
		if j < len(mssql) {
			arr = append(arr, mssql[j])
			j++
		}
		if k < len(oracle) {
			arr = append(arr, oracle[k])
			k++
		}
		if l < len(postgres) {
			arr = append(arr, postgres[l])
			l++
		}
	}
	CreateFile(arr, "payload/shake")
}

func CreateFile(data []Data, nameFile string) {
	file, err := os.Create(nameFile + ".json")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	fmt.Printf("JSON written to %v successfully.\n", nameFile+".json")
}

func Push(newData Data) {
	// Open the people.json file (for both reading and writing)
	file, err := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file content into a byte slice
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Declare a slice to hold the existing people
	var dataPayload []Data

	// If the file is not empty, attempt to unmarshal the existing data
	if len(data) > 0 {
		// If the file content is not empty, unmarshal the data
		err = json.Unmarshal(data, &dataPayload)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
	}

	// Append the new person to the slice
	dataPayload = append(dataPayload, newData)

	// Truncate the file content (clear the file) before writing the updated data
	err = file.Truncate(0)
	if err != nil {
		log.Fatalf("Error truncating file: %v", err)
	}

	// Write the updated data back to the file
	// Move the file pointer back to the beginning of the file for writing
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalf("Error seeking to the start of the file: %v", err)
	}

	// Encode the updated data and write it back to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: pretty-print the JSON
	err = encoder.Encode(dataPayload)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	fmt.Println("person4 appended to people.json successfully.")

}

func Read(nameFile string) []Data {
	file, err := os.OpenFile(nameFile+".json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var dataPayload []Data

	if len(data) > 0 {
		err = json.Unmarshal(data, &dataPayload)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
	}
	return dataPayload
}
