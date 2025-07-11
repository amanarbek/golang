package reading

import (
	"encoding/json"
	"os"
	"project/models"
)

type JSONReader struct{}

func (r JSONReader) Read(filename string) (*models.Recipes, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var recipes models.Recipes
	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return &recipes, nil
}
