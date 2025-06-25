package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/m-mfp/skyrim-alchemy-scrapper/webscrapper"
)

func getEffects(ingredient string) ([]string, error) {

	reader, file, err := webscrapper.ReadCSV()

	if err != nil {
		return nil, fmt.Errorf("something went wrong with readCSV: %w", err)
	}
	defer file.Close()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		if strings.ToLower(record[0]) == ingredient {
			return record[1:], nil
		}
	}
	return nil, nil
}
