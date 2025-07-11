package reader

import (
	"encoding/xml"
	"os"
	"project/models"
)

type XMLReader struct{}

func (r XMLReader) Read(filename string) (*models.Recipes, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var recipes models.Recipes
	err = xml.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return &recipes, nil
}
