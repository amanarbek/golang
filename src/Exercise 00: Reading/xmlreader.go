package main

import (
	"encoding/xml"
	"os"
)

type XMLReader struct{}

func (r XMLReader) Read(filename string) (*Recipes, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var recipes Recipes
	err = xml.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return &recipes, nil
}
