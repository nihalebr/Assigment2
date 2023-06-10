package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/nihalebr/Assigment2/data"
	"github.com/nihalebr/Assigment2/utils"
)

// We call Helper function ReadJSON from utils sub package to read the JSON File.
// Looping through the pastries variable to find the donut which has name Old Fashioned.
// if true check for the Batter type Coffee already present.
// if false sort the the Batter by it's ID and find the last used ID value.
// Increment the Last id value by one and type as coffee append the Batter to Struct.
// Write the updated one to the file by using helper function WriteJSON.

func main() {
	var pastries []*data.Pastry
	filename := "Assigment2.json"
	pastries, err := utils.ReadJSON(filename)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range pastries {
		if value.Type == "donut" && value.Name == "Old Fashioned" {
			data.SortBatters(&value.Batters)
			if data.HasValueBatterType(&value.Batters, "Coffee") {
				fmt.Println("The Batter of type Coffee already Present in the JSON.")
				return
			}
			newID, err := strconv.Atoi(value.Batters.Batter[len(value.Batters.Batter)-1].ID)
			if err != nil {
				log.Fatal(err)
			}
			value.Batters.Batter = append(value.Batters.Batter, data.Batter{ID: strconv.Itoa(newID + 1), Type: "Coffee"})
			fmt.Printf("Successfully added batter type Coffee with id %v\n", newID+1)
		}
	}

	if err := utils.WriteJSON(filename, pastries); err != nil {
		log.Fatal(err)
	}
}
