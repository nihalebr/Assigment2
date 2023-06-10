package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/nihalebr/Assigment2/data"
)

// Takes filename in string and returns the slice of pointers in datatype Pastry.
// Reads the JSON from the filename and convert it into a go struct.
func ReadJSON(filename string) ([]*data.Pastry, error) {

	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully opened the JSON File.")
	defer jsonFile.Close()

	var pastries []*data.Pastry
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &pastries)
	if err != nil {
		return nil, err
	}
	return pastries, nil
}
