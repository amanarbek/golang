package reading

import (
	"errors"
	"path/filepath"
	"project/models"
)

type DBReader interface {
	Read(filename string) (*models.Recipes, error)
}

func GetReader(filename string) (DBReader, error) {
	switch filepath.Ext(filename) {
	case ".json":
		return JSONReader{}, nil
	case ".xml":
		return XMLReader{}, nil
	default:
		return nil, errors.New("неподдерживаемый тип файла")
	}
}
