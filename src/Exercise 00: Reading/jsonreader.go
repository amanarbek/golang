package main

import (
	"encoding/json"
	"os"
)

type JSONReader struct{}

func (r JSONReader) Read(filename string) (*Recipes, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var recipes Recipes
	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return &recipes, nil
}
