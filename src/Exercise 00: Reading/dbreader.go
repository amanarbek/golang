package main

import (
	"errors"
	"path/filepath"
)

type DBReader interface {
	Read(filename string) (*Recipes, error)
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
