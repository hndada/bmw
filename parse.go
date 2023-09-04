package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readAllRecords(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the first line (headers) and discard
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Read the rest of the file
	raws, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return raws
}
