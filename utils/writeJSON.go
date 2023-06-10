package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/nihalebr/Assigment2/data"
)

// Takes filename in string and pastries in slice of pointer in datatype Pastry.
// converts the slice into slice of bytes and write the Bytes to the specified filename.
func WriteJSON(filename string, pastries []*data.Pastry) error {
	jsonBytes, err := json.MarshalIndent(pastries, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonBytes, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Successfully written the JSON into file.")
	return nil
}
