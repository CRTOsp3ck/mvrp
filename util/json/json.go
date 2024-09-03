package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type IJsonUtil interface {
	PrintJson(v interface{})
	ParseJsonFile(filename string, v interface{}) (*interface{}, error)
}

type JsonUtil struct{}

func (j *JsonUtil) PrintJson(v interface{}) {
	// Convert the vendor struct to a JSON string with indents
	modelJson, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	// Print the formatted JSON string
	fmt.Printf("\n%s\n", strings.Repeat("-", 80))
	fmt.Printf("\n%T: %s\n", v, modelJson)
}

// ParseJsonFile reads a JSON file and unmarshals the data into the provided struct
// Returns a pointer to the struct
func (j *JsonUtil) ParseJsonFile(filename string, v interface{}) (interface{}, error) {
	// Open the JSON file
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into the Root struct
	err = json.Unmarshal(byteValue, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
